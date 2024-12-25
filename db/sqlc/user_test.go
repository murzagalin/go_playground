package db

import (
	"context"
	"testing"

	"github.com/murzagalin/go_playground/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: util.RandomString(10),
		FullName:       util.RandomString(6),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUsr := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), createdUsr.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUsr.Username, user.Username)
	require.Equal(t, createdUsr.FullName, user.FullName)
	require.Equal(t, createdUsr.HashedPassword, user.HashedPassword)
	require.Equal(t, createdUsr.Email, user.Email)
	require.Equal(t, createdUsr.CreatedAt, user.CreatedAt)
	require.Equal(t, createdUsr.PasswordChangedAt, user.PasswordChangedAt)
}
