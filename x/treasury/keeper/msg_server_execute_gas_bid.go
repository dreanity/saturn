package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

var (
	tenScaledBy18 = sdk.NewDec(10).Power(18)
	tenScaledBy6  = sdk.NewDec(10).Power(6)
)

func (k msgServer) ExecuteGasBid(goCtx context.Context, msg *types.MsgExecuteGasBid) (*types.MsgExecuteGasBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	gasMeter := ctx.GasMeter()

	treasurer := k.GetTreasurer(ctx)
	if treasurer.Address != msg.Creator {
		return nil, types.ErrCreatorIsNotTreasurer
	}

	defer func() {
		gasConsumed := gasMeter.GasConsumed()
		gasMeter.RefundGas(gasConsumed, "")
	}()

	if len(strings.Trim(msg.Chain, " ")) == 0 {
		return nil, types.ErrInvalidChain
	}

	if len(strings.Trim(msg.TokenAddress, " ")) == 0 {
		return nil, types.ErrInvalidTokenAddress
	}

	paid, ok := sdk.NewIntFromString(msg.PaidAmount)
	if !ok || paid.IsZero() || paid.IsNegative() {
		return nil, types.ErrInvalidPaidAmount
	}

	recepient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, types.ErrInvalidRecipient
	}

	gasBid, found := k.GetGasBid(ctx, msg.Chain)
	if !found && msg.BidNumber == 0 {
		gasBid.Chain = msg.Chain
		gasBid.Number = 0
	} else if gasBid.Number+1 != msg.BidNumber {
		return nil, types.ErrInvalidBidNumber
	}

	gasPrice, found := k.GetGasPrice(ctx, msg.Chain, msg.TokenAddress)
	if !found {
		return nil, types.ErrNotFoundGasPrice
	}

	gasPriceValue, _ := sdk.NewDecFromStr(gasPrice.Value)

	paidDec := sdk.NewDecFromInt(paid).Quo(tenScaledBy18)
	gasPriceDec := gasPriceValue.Quo(tenScaledBy6)
	amountIssued := paidDec.Mul(gasPriceDec).Mul(tenScaledBy6).Ceil().TruncateInt()

	amountIssuedCoin := sdk.NewCoin("uhydrogen", amountIssued)
	amountIssuedCoins := sdk.NewCoins(amountIssuedCoin)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amountIssuedCoins); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recepient, amountIssuedCoins); err != nil {
		return nil, err
	}

	gasBid.Number = msg.BidNumber
	k.SetGasBid(ctx, gasBid)

	ctx.EventManager().EmitTypedEvent(&types.GasBidExecuted{
		BidNumber:    msg.BidNumber,
		Recepient:    msg.Recipient,
		MintedAmount: amountIssued.String(),
		PaidAmount:   msg.PaidAmount,
		Chain:        msg.Chain,
		TokenAddress: msg.TokenAddress,
		GasPrice:     gasPrice.Value,
	})

	return &types.MsgExecuteGasBidResponse{}, nil
}
