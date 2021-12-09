// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package subnet

import (
	"fmt"
	"io"
	"math"
	"sort"

	address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	hierarchical "github.com/filecoin-project/lotus/chain/consensus/hierarchical"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufSubnetState = []byte{140}

func (t *SubnetState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSubnetState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Name (string) (string)
	if len(t.Name) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Name was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Name))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Name)); err != nil {
		return err
	}

	// t.ParentID (hierarchical.SubnetID) (string)
	if len(t.ParentID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ParentID was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.ParentID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ParentID)); err != nil {
		return err
	}

	// t.Consensus (hierarchical.ConsensusType) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Consensus)); err != nil {
		return err
	}

	// t.MinMinerStake (big.Int) (struct)
	if err := t.MinMinerStake.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Miners ([]address.Address) (slice)
	if len(t.Miners) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Miners was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Miners))); err != nil {
		return err
	}
	for _, v := range t.Miners {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.TotalStake (big.Int) (struct)
	if err := t.TotalStake.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Stake (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Stake); err != nil {
		return xerrors.Errorf("failed to write cid field t.Stake: %w", err)
	}

	// t.Status (subnet.Status) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.Genesis ([]uint8) (slice)
	if len(t.Genesis) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Genesis was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Genesis))); err != nil {
		return err
	}

	if _, err := w.Write(t.Genesis[:]); err != nil {
		return err
	}

	// t.CheckPeriod (abi.ChainEpoch) (int64)
	if t.CheckPeriod >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CheckPeriod)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.CheckPeriod-1)); err != nil {
			return err
		}
	}

	// t.Checkpoints (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Checkpoints); err != nil {
		return xerrors.Errorf("failed to write cid field t.Checkpoints: %w", err)
	}

	// t.WindowChecks (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.WindowChecks); err != nil {
		return xerrors.Errorf("failed to write cid field t.WindowChecks: %w", err)
	}

	return nil
}

func (t *SubnetState) UnmarshalCBOR(r io.Reader) error {
	*t = SubnetState{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 12 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Name (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Name = string(sval)
	}
	// t.ParentID (hierarchical.SubnetID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.ParentID = hierarchical.SubnetID(sval)
	}
	// t.Consensus (hierarchical.ConsensusType) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Consensus = hierarchical.ConsensusType(extra)

	}
	// t.MinMinerStake (big.Int) (struct)

	{

		if err := t.MinMinerStake.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinMinerStake: %w", err)
		}

	}
	// t.Miners ([]address.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Miners: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Miners = make([]address.Address, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v address.Address
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Miners[i] = v
	}

	// t.TotalStake (big.Int) (struct)

	{

		if err := t.TotalStake.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalStake: %w", err)
		}

	}
	// t.Stake (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Stake: %w", err)
		}

		t.Stake = c

	}
	// t.Status (subnet.Status) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = Status(extra)

	}
	// t.Genesis ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Genesis: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Genesis = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Genesis[:]); err != nil {
		return err
	}
	// t.CheckPeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.CheckPeriod = abi.ChainEpoch(extraI)
	}
	// t.Checkpoints (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Checkpoints: %w", err)
		}

		t.Checkpoints = c

	}
	// t.WindowChecks (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.WindowChecks: %w", err)
		}

		t.WindowChecks = c

	}
	return nil
}

var lengthBufConstructParams = []byte{134}

func (t *ConstructParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufConstructParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.NetworkName (string) (string)
	if len(t.NetworkName) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.NetworkName was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.NetworkName))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.NetworkName)); err != nil {
		return err
	}

	// t.Name (string) (string)
	if len(t.Name) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Name was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Name))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Name)); err != nil {
		return err
	}

	// t.Consensus (hierarchical.ConsensusType) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Consensus)); err != nil {
		return err
	}

	// t.MinMinerStake (big.Int) (struct)
	if err := t.MinMinerStake.MarshalCBOR(w); err != nil {
		return err
	}

	// t.DelegMiner (address.Address) (struct)
	if err := t.DelegMiner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.CheckPeriod (abi.ChainEpoch) (int64)
	if t.CheckPeriod >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CheckPeriod)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.CheckPeriod-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *ConstructParams) UnmarshalCBOR(r io.Reader) error {
	*t = ConstructParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.NetworkName (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.NetworkName = string(sval)
	}
	// t.Name (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Name = string(sval)
	}
	// t.Consensus (hierarchical.ConsensusType) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Consensus = hierarchical.ConsensusType(extra)

	}
	// t.MinMinerStake (big.Int) (struct)

	{

		if err := t.MinMinerStake.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinMinerStake: %w", err)
		}

	}
	// t.DelegMiner (address.Address) (struct)

	{

		if err := t.DelegMiner.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DelegMiner: %w", err)
		}

	}
	// t.CheckPeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.CheckPeriod = abi.ChainEpoch(extraI)
	}
	return nil
}

var lengthBufCheckVotes = []byte{129}

func (t *CheckVotes) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCheckVotes); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Miners ([]address.Address) (slice)
	if len(t.Miners) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Miners was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Miners))); err != nil {
		return err
	}
	for _, v := range t.Miners {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *CheckVotes) UnmarshalCBOR(r io.Reader) error {
	*t = CheckVotes{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miners ([]address.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Miners: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Miners = make([]address.Address, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v address.Address
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Miners[i] = v
	}

	return nil
}
