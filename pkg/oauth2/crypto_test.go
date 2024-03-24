package oauth2_test

import (
	"testing"

	"github.com/maordavidov/oauth2c/pkg/oauth2"
	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	r1 := oauth2.RandomString(10)
	r2 := oauth2.RandomString(10)

	require.NotEqual(t, r1, r2)
}
