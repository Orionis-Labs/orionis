package types

type BankKeeper interface {
	// Methods imported from bank should be defined here
}

// type (
// 	// OracleKeeper interface
// 	OracleKeeper interface {
// 		FinalizeRound(ctx sdk.Context, claimType string, roundID uint64)
// 		GetPendingRounds(ctx sdk.Context, roundType string) (rounds []uint64)
// 		TallyVotes(ctx sdk.Context, claimType string, roundID uint64) *oracletypes.RoundResult
// 		GetClaim(ctx sdk.Context, hash tmbytes.HexBytes) oracle.Claim
// 	}
// )
