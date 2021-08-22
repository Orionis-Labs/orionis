package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/orionis-labs/orionis/x/nameservice/types"
	"github.com/stretchr/testify/assert"
)

func createNWhois(keeper *Keeper, ctx sdk.Context, n int) []types.Whois {
	items := make([]types.Whois, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendWhois(ctx, items[i])
	}
	return items
}

func TestWhoisGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetWhois(ctx, item.Id))
	}
}

func TestWhoisExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasWhois(ctx, item.Id))
	}
}

func TestWhoisRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhois(ctx, item.Id)
		assert.False(t, keeper.HasWhois(ctx, item.Id))
	}
}

func TestWhoisGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllWhois(ctx))
}

func TestWhoisCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNWhois(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetWhoisCount(ctx))
}
