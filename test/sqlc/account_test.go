package testSqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ndbac/go-log/src/sqlc"
	"github.com/ndbac/go-log/test"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) sqlc.Account {
	arg := sqlc.CreateAccountParams{
		Owner:    test.RandomOwner(),
		Balance:  test.RandomMoney(),
		Currency: test.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	savedAccount := createRandomAccount(t)
	resAccount, err := testQueries.GetAccount(context.Background(), savedAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, resAccount)

	require.Equal(t, savedAccount.ID, resAccount.ID)
	require.Equal(t, savedAccount.Owner, resAccount.Owner)
	require.Equal(t, savedAccount.Balance, resAccount.Balance)
	require.Equal(t, savedAccount.Currency, resAccount.Currency)
	require.WithinDuration(t, savedAccount.CreatedAt, resAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	savedAccount := createRandomAccount(t)

	arg := sqlc.UpdateAccountParams{
		ID:      savedAccount.ID,
		Balance: test.RandomMoney(),
	}

	resAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, resAccount)

	require.Equal(t, savedAccount.ID, resAccount.ID)
	require.Equal(t, savedAccount.Owner, resAccount.Owner)
	require.Equal(t, arg.Balance, resAccount.Balance)
	require.Equal(t, savedAccount.Currency, resAccount.Currency)
	require.WithinDuration(t, savedAccount.CreatedAt, resAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	savedAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), savedAccount.ID)
	require.NoError(t, err)

	resAccount, err := testQueries.GetAccount(context.Background(), savedAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, resAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := sqlc.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
