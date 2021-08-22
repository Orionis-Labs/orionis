package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/orionis-labs/orionis/x/orionis/types"
)

var _ = strconv.Itoa(0)

func CmdOnrInr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "onr-inr [amount] [blockHeight] [result] [txnId]",
		Short: "Broadcast message onrInr",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAmount, _ := strconv.ParseUint(args[0], 10, 64)
			argsBlockHeight, _ := strconv.ParseUint(args[1], 10, 64)
			argsResult := string(args[2])
			argsTxnId := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOnrInr(clientCtx.GetFromAddress().String(), uint64(argsAmount), uint64(argsBlockHeight), string(argsResult), string(argsTxnId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
