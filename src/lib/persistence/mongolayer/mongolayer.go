package mongolayer

import (
	"github.com/4piso/okevents/src/lib/persistence"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Global contanst of the db
const (
	DB     = "okevents"
	USERS  = "users"
	EVENTS = "events"
)

// MongoDBLayer struct definition
type MongoDBLayer struct {
	session *mgo.Session
}

/* defining the contruction function for initialization of the code */

// NewMongoDBLayer contructor method
func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	// connection to the db
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{
		session: s,
	}, err
}

// AddEvent method
func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	// get a fresh mongo session
	s := mgoLayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	// check if the id for the location is also good
	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

// FindEvent  id => event, error
func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	// get fresh connection
	s := mgoLayer.getFreshSession()
	defer s.Close()

	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)

	return e, err
}

// FindEventByName string => event, error
func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	// get a fresh connection from the pool
	s := mgoLayer.getFreshSession()
	defer s.Close()

	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)

	return e, err
}

// FindAllAvailableEvents () => events, error
func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	// get fresh session
	s := mgoLayer.getFreshSession()
	defer s.Close()

	e := []persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(nil).All(&e)

	return e, err
}

// getFreshSession this is a helper function to add fresh database connection from the connection pool
func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
