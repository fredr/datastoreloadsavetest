package appenginetest

import (
	"strings"
	"testing"
	"time"

	"google.golang.org/appengine/datastore"
)

func TestLoadSaveStruct(t *testing.T) {
	v := TestEntity{
		NormalTime:   time.Now(),
		EmbeddedTime: EmbeddedTime{Time: time.Now()},
	}

	props, err := datastore.SaveStruct(&v)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(props)
	for _, prop := range props {
		if strings.HasSuffix(prop.Name, ".") {
			t.Error("invalid prop", prop.Name)
		}
	}
}
