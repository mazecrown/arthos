package core

import (      
    "log"   
	"github.com/couchbaselabs/go-couchbase"
	"strconv"
)

var b *couchbase.Bucket

type Coords struct {
    //~ X int `json:"x"`
    //~ Y int `json:"y"`
    //~ Z int `json:"z"`
    Contains string `json:"c"`
}

var lookradius = 4
var looksize = (lookradius * 2) + 1
var looklength = looksize*looksize*looksize
func look(fromX,fromY,fromZ int) []Coords {
	//get list of coords centred on this one
	//wrap the x and y coords. ceiling and floor the z.
	
	//when you know the size of the slice use make for more efficiency
	keys := make([]string, looklength)
	
	//get the list of keys for bulk get
	wpr := GetWorldParams()
	for x := fromX - lookradius; x <= fromX + lookradius; x++  {	
		realX := wrapToMax(x, wpr.MaxX)
		
		for y := 0; y < looksize; y++  {
			realY := wrapToMax(y, wpr.MaxY)
			
			for z := 0; z < looksize; z++  {
				if z < wpr.MaxZ && z >= 0 {//otherwise we know it will be an empty coord
					//add key to bulk list!
					keys = append(keys, "coords:"+strconv.Itoa(realX)+","+strconv.Itoa(realY)+","+strconv.Itoa(z))
				}
			}
		}
	}
	
	//do the bulk get
	//map[string]*gomemcached.MCResponse, error
	//~ response, err := b.GetBulk(keys)
	
	//FUTURE cull the list of visible objects using rays?
	
	//~ coords := make([][][]interface{},looksize)
	//~ for x := range coords {
        //~ coords[x] = make([][]interface{}, looksize) /* again the type? */
        //~ for y := range coords[x] {
            //~ coords[x][y] = make([]interface{}, looksize) /* again the type? */
			//~ for z := range coords[x][y] {
				//~ coords[x][y][z] = //get the data
			//~ }
        //~ }
    //~ }
	
	//add those coords to a map of 'local' coords starting at 0,0,0
	//get bulk of coords
	//for each result get add the contents of the coords to a new map
	//get bulk using the new map to get the contents of the coords
	burk := make([]Coords, 10)
	return burk
}

func wrapToMax(value, max int) int {
	if value >= max {//if is greater than Max then wrap
		value = value - max
	} else if value < 0 {//if is negative, we've gone off the edge and need to wrap
		value = max + value
	} 
	return value
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
