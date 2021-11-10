package controllers

import (
	"github.com/backend-service/api/v1"
	"github.com/backend-service/dataservices"
)

// ControllerDescriber -
type ControllerDescriber interface {
	Connect(connectionString string) (setupError error)
	Ping() error
	Close() error

	BeginTransaction() (txn dataservices.BackendServiceDBInterface, err *api.APIError)
	CommitTransaction() (err *api.APIError)
	RollbackTransaction() (err *api.APIError)
}
