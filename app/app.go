package appenginetest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/appengine/datastore"
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
	EmbeddedTime EmbeddedTime
	NormalTime   time.Time
}

func do(res http.ResponseWriter, req *http.Request) {
	t1 := TestEntity{
		NormalTime:   time.Now(),
		EmbeddedTime: EmbeddedTime{Time: time.Now()},
	}

	t2 := TestEntity{}

	props, err := datastore.SaveStruct(&t1)
	if err != nil {
		panic(err)
	}

	err = datastore.LoadStruct(&t2, props)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	res.Write([]byte(fmt.Sprintf("props\n%+v\n\nt1\n%+v\n\nt2\n%+v\n", props, t1, t2)))
}
