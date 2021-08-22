package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

var _ sdk.Msg = &MsgOnrInr{}

func NewMsgOnrInr(creator string, amount uint64, blockHeight uint64, result string, txnId string) *MsgOnrInr {
	return &MsgOnrInr{
		Creator:     creator,
		Amount:      amount,
		BlockHeight: blockHeight,
		Result:      result,
		TxnId:       txnId,
	}
}

func (msg *MsgOnrInr) Route() string {
	return RouterKey
}

func (msg *MsgOnrInr) Type() string {
	return OnrInrClaim
}

func (msg *MsgOnrInr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOnrInr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOnrInr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// Hash returns the hash of an Claim Content.
func (c *MsgOnrInr) Hash() tmbytes.HexBytes {
	bz, err := c.Marshal()
	if err != nil {
		panic(err)
	}
	return tmhash.Sum(bz)
}

// GetRoundID returns the block height for when the data was used.
func (c *MsgOnrInr) GetRoundID() uint64 {
	return uint64(c.BlockHeight)
}

// GetConcensusKey returns a key the oracle will use of vote consensus
// for deterministic results it should be the same as the hash of the content
// for nondeterministic content it should be a constant
func (c *MsgOnrInr) GetConcensusKey() string {
	return ""
}
