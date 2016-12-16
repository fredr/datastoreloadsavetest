package appenginetest

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/gorilla/mux"
)

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/do", do)
	http.Handle("/", router)
}

type EmbeddedTime struct {
	time.Time
}

type TestEntity struct {
	NormalTime   time.Time
	EmbeddedTime EmbeddedTime
}

func do(res http.ResponseWriter, req *http.Request) {
	te1 := TestEntity{
		NormalTime:   time.Now(),
		EmbeddedTime: EmbeddedTime{Time: time.Now()},
	}

	te2 := TestEntity{}

	props, err := datastore.SaveStruct(&te1)
	if err != nil {
		panic(err)
	}

	c := appengine.NewContext(req)
	log.Infof(c, "%+v", props)

	err = datastore.LoadStruct(&te2, props)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	res.Write([]byte(fmt.Sprintf("te1\n%+v\n\nte2\n%+v", te1, te2)))
}
