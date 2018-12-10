package mongo

import (
	// "github.com/astaxie/beego"
	"gopkg.in/mgo.v2" // change to "github.com/mongodb/mongo-go-driver"
	log "github.com/sirupsen/logrus"
)

type Session struct {
	session *mgo.Session
}

func NewMongoDbSession(mongodb_url string) (*Session, error) {
	session, err := mgo.Dial(mongodb_url)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	log.Info("Successfully connected to mongodb driver at ", mongodb_url)

	return &Session{session}, err
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
		log.Info("Mongodb session is closed")
	}
}

func (s *Session) Copy() *mgo.Session {
	log.Info("Session: ", s.session.Copy())
	return s.session.Copy()
}

func (s *Session) New() *mgo.Session {
	return s.session
}

func (s *Session) Ping() error {
	return s.session.Ping()
}