package q7

import (
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/servicehandlers"

	"gopkg.in/mgo.v2/bson"
)

func Authenticate(f func(h http.Handler, sema chan struct{})) func(h http.Handler, sema chan struct{}) {
	var r http.Request
	var w http.ResponseWriter
	return func(h http.Handler, sema chan struct{}) {
		session := dao.MgoSession.Clone()
		defer session.Close()
		var response dao.Session

		sessionClctn := session.DB("simplesurveys").C("session")
		query := sessionClctn.Find(bson.M{"token": r.Header["token"][0]})
		err := query.One(&response)
		if err != nil {
			servicehandlers.UnauthorizedAccess("Unauthorised").RenderResponse(w)
		} else {
			f(h, sema)
		}
	}
}

//coundnot consume this
