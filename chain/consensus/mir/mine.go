package mir

import (
	"context"
	"fmt"
	"time"

	"github.com/filecoin-project/go-address"
	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/lotus/chain/consensus/platform/logging"
	"github.com/filecoin-project/lotus/chain/types"
	ltypes "github.com/filecoin-project/lotus/chain/types"
)

// Mine handles block "mining" using Mir framework.
//
// Mine implements the following algorithm:
// 1. Retrieve messages and cross-messages from mempool and create Mir requests from them.
//    Note, that messages can be added into mempool via the libp2p mechanism and the CLI.
// 2. Store the requests in cache, that is, at present, a key-value storage in memory.
//    Each key is the request hash.
// 3. Send those hashes to the Mir node in FIFO mode for each client.
// 4. Receive ordered hashes from the Mir node and retrieve the corresponding requests from the storage.
// 5. Create the next Filecoin block.
//    Note, only a leader Eudico node, chosen by round-robin election, creates a block.
// 6. Sunk this block without sending it over the libp2p network.
//
// There are two ways how mining with Mir can be started:
// 1) Environment variables: validators ID and their network addresses are passed
//    via EUDICO_MIR_VALIDATORS variable.
//    This approach can be used to run Mir in the root network and for simple demos.
// 2) Hierarchical consensus framework: validators IDs and their network addresses
//    are received via state, after each validator joins the subnet.
//    This is used to run Mir in a subnet.
func Mine(ctx context.Context, addr address.Address, api v1api.FullNode) error {
	log.With("addr", addr).Infof("Mir miner started")
	defer log.With("addr", addr).Infof("Mir miner completed")

	m, err := NewManager(ctx, addr, api)
	if err != nil {
		return fmt.Errorf("unable to create a manager: %w", err)
	}
	log := logging.FromContext(ctx, log).With("miner", m.ID())

	log.Infof("Miner info:\n\twallet - %s\n\tnetwork - %s\n\tsubnet - %s\n\tMir ID - %s\n\tvalidators - %v",
		m.Addr, m.NetName, m.SubnetID, m.MirID, m.Validators)

	mirErrors := m.Start(ctx)
	mirHead := m.App.ChainNotify

	submit := time.NewTicker(SubmitInterval)
	defer submit.Stop()

	for {
		base, err := api.ChainHead(ctx)
		if err != nil {
			log.Errorw("failed to get chain head", "error", err)
			continue
		}

		// Miner (leader) for an epoch is assigned deterministically using round-robin.
		// All other validators use the same Miner in the block.
		epochMiner := m.Validators[int(base.Height())%len(m.Validators)].Addr
		nextEpoch := base.Height() + 1

		select {
		case <-ctx.Done():
			log.With("epoch", nextEpoch).Debug("Mir miner: context closed")
			return nil
		case err := <-mirErrors:
			return fmt.Errorf("miner consensus error: %w", err)
		case batch := <-mirHead:
			msgs, crossMsgs := m.GetMessages(batch)
			log.With("epoch", nextEpoch).
				Infof("try to create a block: msgs - %d, crossMsgs - %d", len(msgs), len(crossMsgs))

			bh, err := api.MinerCreateBlock(ctx, &lapi.BlockTemplate{
				Miner:            epochMiner,
				Parents:          base.Key(),
				BeaconValues:     nil,
				Ticket:           &ltypes.Ticket{VRFProof: nil},
				Epoch:            base.Height() + 1,
				Timestamp:        uint64(time.Now().Unix()),
				WinningPoStProof: nil,
				Messages:         msgs,
				CrossMessages:    crossMsgs,
			})
			if err != nil {
				log.With("epoch", nextEpoch).
					Errorw("creating a block failed", "error", err)
				continue
			}
			if bh == nil {
				log.With("epoch", nextEpoch).
					Debug("created a nil block")
				continue
			}

			err = api.SyncBlock(ctx, &types.BlockMsg{
				Header:        bh.Header,
				BlsMessages:   bh.BlsMessages,
				SecpkMessages: bh.SecpkMessages,
				CrossMessages: bh.CrossMessages,
			})
			if err != nil {
				log.With("epoch", nextEpoch).
					Errorw("unable to sync a block", "error", err)
				continue
			}

			log.With("epoch", nextEpoch).
				Infof("%s mined a block %v", epochMiner, bh.Cid())
		default:
			msgs, err := api.MpoolSelect(ctx, base.Key(), 1)
			if err != nil {
				log.With("epoch", nextEpoch).
					Errorw("unable to select messages from mempool", "error", err)
			}

			crossMsgs, err := api.GetUnverifiedCrossMsgsPool(ctx, m.SubnetID, base.Height()+1)
			if err != nil {
				log.With("epoch", nextEpoch).
					Errorw("unable to select cross-messages from mempool", "error", err)
			}

			m.SubmitRequests(ctx, m.GetRequests(msgs, crossMsgs))
		}
	}
}
