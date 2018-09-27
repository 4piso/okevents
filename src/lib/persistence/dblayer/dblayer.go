package dblayer

import (
	"github.com/4piso/okevents/src/lib/persistence"
	"github.com/4piso/okevents/src/lib/persistence/mongolayer"
)

// DBType define the db type
type DBType string

// global constant
const (
	MONGODB DBType = "mongodb"
)

// NewPersistenceLayer (DBType, connection) => persistence.DatabaseHandler, error
func NewPersistenceLayer(options DBType, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}

	return nil, nil
}
