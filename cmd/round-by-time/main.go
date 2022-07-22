package main

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	round := roundAt(1658478253)
	fmt.Println("Round #", round)
}

func roundAt(time uint64) uint64 {
	genesisTime := uint64(1595431050)
	preiod := uint64(30)
	round := uint64(1)

	if time < genesisTime {
		round = 1
	} else {

		timeDiff := sdk.NewDecFromBigInt(sdk.NewIntFromUint64(time - genesisTime).BigInt())
		roundForTime := timeDiff.QuoRoundUp(sdk.NewDecFromInt(sdk.NewIntFromUint64(preiod)))
		nextRound := roundForTime.Ceil().TruncateInt().Uint64() + 1

		println("TimeDiff uint64   : ", time-genesisTime)
		println("TimeDiff sdk.Dec  : ", sdk.NewDecFromBigInt(sdk.NewIntFromUint64(time-genesisTime).BigInt()).String())
		println("TimeDiff / Period : ", timeDiff.QuoRoundUp(sdk.NewDecFromInt(sdk.NewIntFromUint64(preiod))).String())
		println("Next round        : ", roundForTime.Ceil().TruncateInt().Uint64()+1)

		round = nextRound
	}

	return round
}
