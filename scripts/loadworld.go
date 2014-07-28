package main

import (      
    "log"   
	"strconv"
	"github.com/mazecrown/arthos/core"
	"github.com/mazecrown/arthos/nature"
)

func main() {
	//clear existing world
	//create world params object
	timefactor := 60000
	maxX, maxY, maxZ := 24, 24, 2
	worldparams := core.WorldParams{timefactor, maxX, maxY, maxZ}
	
	b := core.GetBucket()
    
	core.AddToBucket(worldparams, "worldparams")	
	
	var wpr core.WorldParams = core.GetWorldParams()
	log.Printf("got world params=%+v", wpr)
	
	//create world coord objects
	count := 0	
	c := ""
	for z := 0; z < maxZ; z++ {
		for x := 0; x < maxX; x++  {	
			for y := 0; y < maxY; y++  {
				count ++
				coords := core.Coords{c}
				log.Printf("adding coords=%+v", coords)
				core.AddToBucket(coords, "coords:"+strconv.Itoa(x)+","+strconv.Itoa(y)+","+strconv.Itoa(z))	
			}
		}
	}
	log.Printf("loaded %d coords", count)	
	
	//create grass objects in first layer of world
	//first we create a master "list" type of object which is simply an array which will allow us to keep track of all our grass keys	
	grasslist := []string{}
	core.AddToBucket(grasslist, "grass:list")	
	
	//create the grass objects themselves
	count = 0
	for x := 0; x < maxX; x++  {	
		for y := 0; y < maxY; y++  {
			count ++
			//z is always 0 as we only want 1 layer for now
			z:= 0
			
			//create grass object
			grass := nature.CreateGrass()
			
			//add to bucket
			grassid := "grass:"+strconv.Itoa(count)
			added, _ := core.AddToBucket(grass, grassid)	
			
			//add to grasslist document
			b.Get("grass:list", &grasslist)			
			grasslist = append(grasslist, grassid)
			core.SetToBucket(grasslist, "grass:list")
			
			//if successful then also update the correct coord with a reference to the grass object
			if added {
				coords := core.Coords{}
				coordsid := "coords:"+strconv.Itoa(x)+","+strconv.Itoa(y)+","+strconv.Itoa(z)
				b.Get(coordsid, &coords)
				coords.Contains = grassid
				core.SetToBucket(coords, coordsid)
			}
		}
	}
}
