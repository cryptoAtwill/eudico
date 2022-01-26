// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package sca

import (
	"fmt"
	"io"
	"math"
	"sort"

	address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	schema "github.com/filecoin-project/lotus/chain/consensus/hierarchical/checkpoints/schema"
	types "github.com/filecoin-project/lotus/chain/types"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufConstructorParams = []byte{130}

func (t *ConstructorParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufConstructorParams); err != nil {
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

	// t.CheckpointPeriod (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CheckpointPeriod)); err != nil {
		return err
	}

	return nil
}

func (t *ConstructorParams) UnmarshalCBOR(r io.Reader) error {
	*t = ConstructorParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
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
	// t.CheckpointPeriod (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CheckpointPeriod = uint64(extra)

	}
	return nil
}

var lengthBufCheckpointParams = []byte{129}

func (t *CheckpointParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCheckpointParams); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Checkpoint ([]uint8) (slice)
	if len(t.Checkpoint) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Checkpoint was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Checkpoint))); err != nil {
		return err
	}

	if _, err := w.Write(t.Checkpoint[:]); err != nil {
		return err
	}
	return nil
}

func (t *CheckpointParams) UnmarshalCBOR(r io.Reader) error {
	*t = CheckpointParams{}

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

	// t.Checkpoint ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Checkpoint: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Checkpoint = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Checkpoint[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufSCAState = []byte{140}

func (t *SCAState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSCAState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.NetworkName (address.SubnetID) (string)
	if len(t.NetworkName) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.NetworkName was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.NetworkName))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.NetworkName)); err != nil {
		return err
	}

	// t.TotalSubnets (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TotalSubnets)); err != nil {
		return err
	}

	// t.MinStake (big.Int) (struct)
	if err := t.MinStake.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Subnets (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Subnets); err != nil {
		return xerrors.Errorf("failed to write cid field t.Subnets: %w", err)
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

	// t.CheckMsgsRegistry (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.CheckMsgsRegistry); err != nil {
		return xerrors.Errorf("failed to write cid field t.CheckMsgsRegistry: %w", err)
	}

	// t.Nonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	// t.BottomUpNonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.BottomUpNonce)); err != nil {
		return err
	}

	// t.BottomUpMsgsMeta (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.BottomUpMsgsMeta); err != nil {
		return xerrors.Errorf("failed to write cid field t.BottomUpMsgsMeta: %w", err)
	}

	// t.AppliedBottomUpNonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.AppliedBottomUpNonce)); err != nil {
		return err
	}

	// t.AppliedTopDownNonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.AppliedTopDownNonce)); err != nil {
		return err
	}

	return nil
}

func (t *SCAState) UnmarshalCBOR(r io.Reader) error {
	*t = SCAState{}

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

	// t.NetworkName (address.SubnetID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.NetworkName = address.SubnetID(sval)
	}
	// t.TotalSubnets (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TotalSubnets = uint64(extra)

	}
	// t.MinStake (big.Int) (struct)

	{

		if err := t.MinStake.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinStake: %w", err)
		}

	}
	// t.Subnets (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Subnets: %w", err)
		}

		t.Subnets = c

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
	// t.CheckMsgsRegistry (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.CheckMsgsRegistry: %w", err)
		}

		t.CheckMsgsRegistry = c

	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	// t.BottomUpNonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.BottomUpNonce = uint64(extra)

	}
	// t.BottomUpMsgsMeta (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.BottomUpMsgsMeta: %w", err)
		}

		t.BottomUpMsgsMeta = c

	}
	// t.AppliedBottomUpNonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.AppliedBottomUpNonce = uint64(extra)

	}
	// t.AppliedTopDownNonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.AppliedTopDownNonce = uint64(extra)

	}
	return nil
}

var lengthBufSubnet = []byte{137}

func (t *Subnet) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSubnet); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.ID (address.SubnetID) (string)
	if len(t.ID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ID was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.ID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ID)); err != nil {
		return err
	}

	// t.ParentID (address.SubnetID) (string)
	if len(t.ParentID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ParentID was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.ParentID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ParentID)); err != nil {
		return err
	}

	// t.Stake (big.Int) (struct)
	if err := t.Stake.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Funds (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Funds); err != nil {
		return xerrors.Errorf("failed to write cid field t.Funds: %w", err)
	}

	// t.TopDownMsgs (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.TopDownMsgs); err != nil {
		return xerrors.Errorf("failed to write cid field t.TopDownMsgs: %w", err)
	}

	// t.Nonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	// t.CircSupply (big.Int) (struct)
	if err := t.CircSupply.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (sca.Status) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.PrevCheckpoint (schema.Checkpoint) (struct)
	if err := t.PrevCheckpoint.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Subnet) UnmarshalCBOR(r io.Reader) error {
	*t = Subnet{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 9 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ID (address.SubnetID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.ID = address.SubnetID(sval)
	}
	// t.ParentID (address.SubnetID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.ParentID = address.SubnetID(sval)
	}
	// t.Stake (big.Int) (struct)

	{

		if err := t.Stake.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Stake: %w", err)
		}

	}
	// t.Funds (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Funds: %w", err)
		}

		t.Funds = c

	}
	// t.TopDownMsgs (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.TopDownMsgs: %w", err)
		}

		t.TopDownMsgs = c

	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	// t.CircSupply (big.Int) (struct)

	{

		if err := t.CircSupply.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.CircSupply: %w", err)
		}

	}
	// t.Status (sca.Status) (uint64)

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
	// t.PrevCheckpoint (schema.Checkpoint) (struct)

	{

		if err := t.PrevCheckpoint.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PrevCheckpoint: %w", err)
		}

	}
	return nil
}

var lengthBufFundParams = []byte{129}

func (t *FundParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufFundParams); err != nil {
		return err
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *FundParams) UnmarshalCBOR(r io.Reader) error {
	*t = FundParams{}

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

	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Value: %w", err)
		}

	}
	return nil
}

var lengthBufSubnetIDParam = []byte{129}

func (t *SubnetIDParam) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSubnetIDParam); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.ID (string) (string)
	if len(t.ID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ID was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.ID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ID)); err != nil {
		return err
	}
	return nil
}

func (t *SubnetIDParam) UnmarshalCBOR(r io.Reader) error {
	*t = SubnetIDParam{}

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

	// t.ID (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.ID = string(sval)
	}
	return nil
}

var lengthBufCrossMsgs = []byte{130}

func (t *CrossMsgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCrossMsgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Msgs ([]types.Message) (slice)
	if len(t.Msgs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Msgs was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Msgs))); err != nil {
		return err
	}
	for _, v := range t.Msgs {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Metas ([]schema.CrossMsgMeta) (slice)
	if len(t.Metas) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Metas was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Metas))); err != nil {
		return err
	}
	for _, v := range t.Metas {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *CrossMsgs) UnmarshalCBOR(r io.Reader) error {
	*t = CrossMsgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Msgs ([]types.Message) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Msgs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Msgs = make([]types.Message, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v types.Message
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Msgs[i] = v
	}

	// t.Metas ([]schema.CrossMsgMeta) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Metas: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Metas = make([]schema.CrossMsgMeta, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v schema.CrossMsgMeta
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Metas[i] = v
	}

	return nil
}

var lengthBufMetaTag = []byte{130}

func (t *MetaTag) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufMetaTag); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.MsgsCid (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.MsgsCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.MsgsCid: %w", err)
	}

	// t.MetasCid (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.MetasCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.MetasCid: %w", err)
	}

	return nil
}

func (t *MetaTag) UnmarshalCBOR(r io.Reader) error {
	*t = MetaTag{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.MsgsCid (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.MsgsCid: %w", err)
		}

		t.MsgsCid = c

	}
	// t.MetasCid (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.MetasCid: %w", err)
		}

		t.MetasCid = c

	}
	return nil
}

var lengthBufMsgParams = []byte{129}

func (t *MsgParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufMsgParams); err != nil {
		return err
	}

	// t.Msg (types.Message) (struct)
	if err := t.Msg.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *MsgParams) UnmarshalCBOR(r io.Reader) error {
	*t = MsgParams{}

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

	// t.Msg (types.Message) (struct)

	{

		if err := t.Msg.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Msg: %w", err)
		}

	}
	return nil
}
