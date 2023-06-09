package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/1AMTEDDY/Backend/util"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountsParams{
		Owner:    util.RandomOwner(), // randomly generated
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func CreateRandomAccounts(t *testing.T) Account { // create a random account to use for unit testing
	arg := CreateAccountsParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccounts(t *testing.T) {
	CreateRandomAccounts(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccounts(t)

	account2, err := testQueries.GetAccounts(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

//func TestUpdateAccount(t *testing.T) {
//	account1 := CreateRandomAccounts(t) // create a random account to use for unit testing
//	arg := UpdateAccountsParams{
//		ID:       account1.ID, // randomly generated
//		Balance:  util.RandomMoney(),
//		Currency: util.RandomCurrency(),
//		Owner:    util.RandomOwner(),
//	}

//	account2, err := testQueries.UpdateAccounts(context.Background(), arg)
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)

//	require.Equal(t, account1.ID, account2.ID)
//	require.Equal(t, account1.Owner, account2.Owner)
//	require.Equal(t, arg.Balance, account2.Balance)
//	require.Equal(t, arg.Currency, account2.Currency)
//	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
//}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccounts(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccounts(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccounts(t)
	}

	arg := ListAccountsParams{
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
