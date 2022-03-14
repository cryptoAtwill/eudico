// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package checkpointing

import (
	"fmt"
	"io"
	"math"
	"sort"

	address "github.com/filecoin-project/go-address"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufResolveMsg = []byte{131}

func (t *ResolveMsg) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufResolveMsg); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.From (address.SubnetID) (string)
	// if len(t.From) > cbg.MaxLength {
	// 	return xerrors.Errorf("Value in field t.From was too long")
	// }

	// if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.From))); err != nil {
	// 	return err
	// }
	// if _, err := io.WriteString(w, string(t.From)); err != nil {
	// 	return err
	// }

	// t.Type (resolver.MsgType) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return err
	}

	// t.Cid (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Cid); err != nil {
		return xerrors.Errorf("failed to write cid field t.Cid: %w", err)
	}

	// t.CrossMsgs (sca.CrossMsgs) (struct)
	if err := t.CrossMsgs.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *ResolveMsg) UnmarshalCBOR(r io.Reader) error {
	*t = ResolveMsg{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.From (address.SubnetID) (string)

	// {
	// 	sval, err := cbg.ReadStringBuf(br, scratch)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	t.From = address.SubnetID(sval)
	// }
	// t.Type (resolver.MsgType) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Type = MsgType(extra)

	}
	// t.Cid (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Cid: %w", err)
		}

		t.Cid = c

	}
	// t.CrossMsgs (sca.CrossMsgs) (struct)

	{

		if err := t.CrossMsgs.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.CrossMsgs: %w", err)
		}

	}
	return nil
}
