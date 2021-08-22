package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/orionis-labs/orionis/app"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {

	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

// import (
// 	"os"

// 	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
// 	"github.com/relevant-community/oracle/app"
// 	"github.com/relevant-community/oracle/cmd/oracled/cmd"
// )

// func main() {
// 	rootCmd, _ := cmd.NewRootCmd()
// 	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
// 		os.Exit(1)
// 	}
// }
