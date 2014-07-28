package core

import (      
    "log"   
	"github.com/couchbaselabs/go-couchbase"
)

var b *couchbase.Bucket

type Coords struct {
    //~ X int `json:"x"`
    //~ Y int `json:"y"`
    //~ Z int `json:"z"`
    Contains string `json:"c"`
}

type WorldParams struct {
    TimeFactor int `json:"timefactor"`
    MaxX int `json:"maxx"`
    MaxY int `json:"maxy"`
    MaxZ int `json:"maxz"`
}

func GetWorldParams() WorldParams {
	var wpr WorldParams
	b.Get("worldparams", &wpr)
	return wpr
}

func AddToBucket(v interface{}, id string) (added bool, err error){
	added, err = b.Add(id, 0, v)
	if err != nil {
        log.Fatalf("Failed to store %s (%s)\n", id, err)
    }
    if added {
		log.Printf("stored %s", id)
	} else {
		log.Printf("%s already exists", id)
	}
	return added, err
}

func SetToBucket(v interface{}, id string) error{
	err := b.Set(id, 0, v)
	if err != nil {
        log.Fatalf("Failed to store %s (%s)\n", id, err)
    }
	return err
}

func GetBucket() *couchbase.Bucket{
	if b == nil {
		log.Println("Creating couchbase connection")
		var err error
		b, err = couchbase.GetBucket("http://localhost:8091/", "default", "default")
		if err != nil {
			log.Fatalf("Failed to connect to default bucket (%s)\n", err)
		}
	}
    return b
}
