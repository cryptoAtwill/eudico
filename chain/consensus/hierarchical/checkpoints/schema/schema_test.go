package schema_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/consensus/hierarchical"
	"github.com/filecoin-project/lotus/chain/consensus/hierarchical/checkpoints/schema"
	checkTypes "github.com/filecoin-project/lotus/chain/consensus/hierarchical/checkpoints/types"
	"github.com/filecoin-project/lotus/chain/consensus/hierarchical/checkpoints/utils"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/stretchr/testify/require"
)

func TestMarshalCheckpoint(t *testing.T) {
	c1, _ := schema.Linkproto.Sum([]byte("a"))
	epoch := abi.ChainEpoch(1000)
	ch := schema.NewRawCheckpoint(hierarchical.RootSubnet, epoch)
	ch.SetPrevious(c1)

	// Add child checkpoints
	ch.AddListChilds(utils.GenRandChecks(3))

	// Marshal
	var buf bytes.Buffer
	err := ch.MarshalCBOR(&buf)
	require.NoError(t, err)

	// Unmarshal and check equal
	ch2 := &schema.Checkpoint{}
	err = ch2.UnmarshalCBOR(&buf)
	require.NoError(t, err)
	eq, err := ch.Equals(ch2)
	require.NoError(t, err)
	require.True(t, eq)

	// Same for marshal binary
	b, err := ch.MarshalBinary()
	require.NoError(t, err)

	// Unmarshal and check equal
	ch2 = &schema.Checkpoint{}
	err = ch2.UnmarshalBinary(b)
	require.NoError(t, err)
	eq, err = ch.Equals(ch2)
	require.NoError(t, err)
	require.True(t, eq)

	// Check that Equals works.
	c1, _ = schema.Linkproto.Sum([]byte("b"))
	epoch = abi.ChainEpoch(1001)
	ch = schema.NewRawCheckpoint(hierarchical.RootSubnet, epoch)
	ch.SetPrevious(c1)
	eq, err = ch.Equals(ch2)
	require.NoError(t, err)
	require.False(t, eq)

}

func TestMarshalEmptyPrevious(t *testing.T) {
	epoch := abi.ChainEpoch(1000)
	ch := schema.NewRawCheckpoint(hierarchical.RootSubnet, epoch)
	pr, _ := ch.PreviousCheck()
	require.Equal(t, pr, schema.NoPreviousCheck)

	// Add child checkpoints
	ch.AddListChilds(utils.GenRandChecks(3))

	// Marshal
	var buf bytes.Buffer
	err := ch.MarshalCBOR(&buf)
	require.NoError(t, err)

	// Unmarshal and check equal
	ch2 := &schema.Checkpoint{}
	err = ch2.UnmarshalCBOR(&buf)
	require.NoError(t, err)
	eq, err := ch.Equals(ch2)
	require.NoError(t, err)
	require.True(t, eq)

	// Same for marshal binary
	b, err := ch.MarshalBinary()
	require.NoError(t, err)

	// Unmarshal and check equal
	ch2 = &schema.Checkpoint{}
	err = ch2.UnmarshalBinary(b)
	require.NoError(t, err)
	eq, err = ch.Equals(ch2)
	require.NoError(t, err)
	require.True(t, eq)
}

func TestSignature(t *testing.T) {
	ctx := context.Background()
	w, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		t.Fatal(err)
	}
	addr, err := w.WalletNew(ctx, types.KTSecp256k1)
	require.NoError(t, err)
	env := &schema.SingleSignEnvelope{addr.String(), []byte("test")}
	sig, err := schema.NewSignature(env, checkTypes.SingleSignature)
	require.NoError(t, err)
	b, err := sig.MarshalBinary()
	require.NoError(t, err)
	sig2 := &schema.Signature{}
	err = sig2.UnmarshalBinary(b)
	require.NoError(t, err)
	require.True(t, sig.Equal(*sig2))
	sig3 := &schema.Signature{}
	require.False(t, sig.Equal(*sig3))
}

func TestEncodeDecodeSignature(t *testing.T) {
	origsig := schema.Signature{
		SignatureID: 3,
		Sig:         []byte("test-data"),
	}
	sigBytes, err := origsig.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	if len(sigBytes) == 0 {
		t.Fatal("did not encode sig")
	}

	var sig schema.Signature
	if err := sig.UnmarshalBinary(sigBytes); err != nil {
		t.Fatal(err)
	}
	if sig.SignatureID != origsig.SignatureID {
		t.Fatal("got wrong protocol ID")
	}
	if !bytes.Equal(sig.Sig, origsig.Sig) {
		t.Fatal("did not get expected data")
	}
	if !sig.Equal(origsig) {
		t.Fatal("sig no equal after decode")
	}

	// Zero the bytes and ensure the decoded struct still works.
	// This will fail if UnmarshalBinary did not copy the inner data bytes.
	copy(sigBytes, make([]byte, 1024))
	if !sig.Equal(origsig) {
		t.Fatal("sig no equal after buffer zeroing")
	}

	sig.SignatureID = origsig.SignatureID + 1
	if sig.Equal(origsig) {
		t.Fatal("sig should not be equal")
	}
}
