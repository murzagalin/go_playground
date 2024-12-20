package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/murzagalin/go_playground/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Balance:  util.RandomInt(0, 1000),
		Currency: util.RandomCurrency(),
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
	createdAcc := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), createdAcc.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAcc.Owner, account.Owner)
	require.Equal(t, createdAcc.Balance, account.Balance)
	require.Equal(t, createdAcc.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAcc.ID,
		Balance: createdAcc.Balance + 10,
	}

	updatedAcc, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAcc)

	require.Equal(t, createdAcc.Owner, updatedAcc.Owner)
	require.Equal(t, arg.Balance, updatedAcc.Balance)
	require.Equal(t, createdAcc.Currency, updatedAcc.Currency)
	require.Equal(t, createdAcc.ID, updatedAcc.ID)
}

func TestDeleteAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createdAcc.ID)

	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), createdAcc.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
