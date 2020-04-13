package patients_db

import (
	mgo "gopkg.in/mgo.v2"
	"log"
)

var (
	mgoSession *mgo.Session
	Client     *mgo.Database
)

const (
	url = "localhost"
)

func init() {
	GetSession().Ping()
	defer GetSession().Close()
	log.Println("database successfully configured!")
}

func GetSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(url)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}