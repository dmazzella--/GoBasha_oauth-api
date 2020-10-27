package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, at.AccessToken, "", "brant new access token should not have AccessToken defined")
	assert.NotEqual(t, 0, at.UserId, "UserId should not be defined")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "brand new access token should be expired")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Should not be expired")

}

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}
