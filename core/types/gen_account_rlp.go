// Code generated by rlpgen. DO NOT EDIT.

package types

import "github.com/ur-os/go-ethereum/rlp"
import "io"

func (obj *StateAccount) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteUint64(obj.Nonce)
	if obj.Balance == nil {
		w.Write(rlp.EmptyString)
	} else {
		if obj.Balance.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(obj.Balance)
	}
	w.WriteBytes(obj.Root[:])
	w.WriteBytes(obj.CodeHash)
	w.ListEnd(_tmp0)
	return w.Flush()
}
