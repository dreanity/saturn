package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dreanity/saturn/x/treasury/types"
)

func (k msgServer) ChangeGasPrices(goCtx context.Context, msg *types.MsgChangeGasPrices) (*types.MsgChangeGasPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	gasMeter := ctx.GasMeter()

	treasurer := k.GetTreasurer(ctx)
	if msg.Creator != treasurer.Address {
		return nil, types.ErrCreatorIsNotTreasurer
	}

	isValid := true
	for _, gasPrice := range msg.GasPrices {
		if len(strings.Trim(gasPrice.Currency, " ")) == 0 {
			isValid = false
			break
		}

		if value, ok := sdk.NewIntFromString(gasPrice.Value); !ok || value.IsNegative() {
			isValid = false
			break
		}
	}

	if !isValid {
		gasConsumed := gasMeter.GasConsumed()
		gasMeter.RefundGas(gasConsumed, "")

		return nil, types.ErrInvalidGasPrice
	}

	for _, gasPrice := range msg.GasPrices {
		currentGasPrice, _ := k.GetGasPrice(ctx, gasPrice.Currency)

		currentGasPrice.Currency = gasPrice.Currency
		currentGasPrice.Value = gasPrice.Value
		k.SetGasPrice(ctx, currentGasPrice)
	}

	gasConsumed := gasMeter.GasConsumed()
	gasMeter.RefundGas(gasConsumed, "")

	ctx.EventManager().EmitTypedEvent(&types.GasPricesChanged{
		GasPrices: msg.GasPrices,
	})

	return &types.MsgChangeGasPricesResponse{}, nil
}
