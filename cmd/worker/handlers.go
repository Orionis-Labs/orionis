package worker

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/relevant-community/oracle/x/atom/types"
	"github.com/orionis-labs/orionis/x/orionis/types"
	"github.com/relevant-community/oracle/x/oracle/client/cli"
	oracletypes "github.com/relevant-community/oracle/x/oracle/types"
	rpctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/spf13/cobra"
)

type OnrInrData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// HandleTx is our custom tx handler
func HandleTx(cmd *cobra.Command, txEvent rpctypes.ResultEvent) error {
	return nil
}

// HandleBlock is our custom block handler
func HandleBlock(cmd *cobra.Command, blockEvent rpctypes.ResultEvent) error {
	helper, err := cli.NewWorkerHelper(cmd, blockEvent)
	if err != nil {
		return err
	}

	// for each claim type in the array, run the approriate handler
	for _, param := range helper.OracleParams.ClaimParams {
		switch param.ClaimType {
		case types.OnrInrClaim:
			getAtomPrice(cmd, helper, param)
		}
	}

	return nil
}

func getAtomPrice(cmd *cobra.Command, helper *cli.WorkerHelper, param oracletypes.ClaimParams) error {

	// use VotePeriod to check if its time to submit a claim
	if helper.IsRoundStart(param.ClaimType) == false {
		return nil
	}

	// fetch ONR/INR pair price from Binance
	var onrInrData = OnrInrData{}
	r, err := http.Get("https://api.binance.com/api/v3/ticker/")
	if err != nil {
		fmt.Println("Error fetching ONR confirmation")
		return err
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&onrInrData)

	// create a Claim about OnrINR price
	atomUsd := types.NewMsgOnrInr("", 10, uint64(helper.BlockHeight), "result", "txId")

	// run process for the given claimType
	if atomUsd == nil {
		fmt.Println("Error creating claim")
		return nil
	}

	// submit our claim
	helper.SubmitWorkerTx(atomUsd)
	return nil
}
