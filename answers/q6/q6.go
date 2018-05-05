package q6

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Deactivete_serveys() {
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	for {
		duration := time.Now()
		collection := session.DB("simplesurveys").C("survey")
		find := bson.M{"status": true, "expiry": bson.M{"$lt": duration}}
		change := bson.M{"$set": bson.M{"status": false}}
		collection.Update(find, change)
		time.Sleep(10 * time.Second)
	}
}

//this isn sponed in main() of the service
