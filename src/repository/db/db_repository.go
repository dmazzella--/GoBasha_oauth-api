package db

import (
	"github.com/dmazzella--/GoBasha_oauth-api/src/clients/cassandra"
	"github.com/dmazzella--/GoBasha_oauth-api/src/domain/access_token"
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
	"github.com/gocql/gocql"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires from access_tokens where access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) values (?,?,?,?) ;"
	queryUpdateExpires     = "UPDATE  access_tokens set expires = ? where access_token = ? ;"
)

type dbRepository struct {
}

func New() DbRepository {

	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()

	defer session.Close()
	var result access_token.AccessToken
	err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires)
	if err != nil {
		if err.Error() == gocql.ErrNotFound.Error() {
			return nil, errors.NewNotFoundError("not found")
		}
		return nil, errors.NewBadRequestError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()
	defer session.Close()

	if err := session.Query(queryCreateAccessToken, token.AccessToken, token.UserId, token.ClientId, token.Expires).Scan(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()
	defer session.Close()

	if err := session.Query(queryUpdateExpires, token.Expires, token.Expires).Scan(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
