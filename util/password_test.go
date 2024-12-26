package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)

	hashedPassword, err := HashPassowrd(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestPasswordFailure(t *testing.T) {
	password := RandomString(10)

	hashedPassword, err := HashPassowrd(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(fmt.Sprintf("salt%s", password), hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestDuplicateHashPassword(t *testing.T) {
	password := RandomString(10)
	hashedPassword1, err := HashPassowrd(password)
	require.NoError(t, err)
	hashedPassword2, err := HashPassowrd(password)
	require.NoError(t, err)

	require.NotEqual(t, hashedPassword1, hashedPassword2)
	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)
	err = CheckPassword(password, hashedPassword2)
	require.NoError(t, err)
}
