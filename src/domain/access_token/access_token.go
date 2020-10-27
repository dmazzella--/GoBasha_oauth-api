package access_token

import (
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
	"strings"
	"time"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	expirationTime := time.Unix(token.Expires, 0)
	return expirationTime.Before(time.Now())
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access Token Id")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("Invalid User Id")
	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid Client Id")
	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("Invalid Expires")

	}

	return nil
}
