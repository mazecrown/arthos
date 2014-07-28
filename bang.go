package main

import (      
    "log"   
	"github.com/mazecrown/arthos/core"
	"github.com/mazecrown/arthos/nature"
)

func main() {
	//loop grasslist	
	b := core.GetBucket()
	grasslist := []string{}
	b.Get("grass:list", &grasslist)
	for i := range grasslist {
		//for each grass spawn a grassproc
		grassid := grasslist[i]
		log.Printf("grass=%+v", grassid)
		
		//get world constants
		var wpr core.WorldParams = core.GetWorldParams()
		
		//spawn grass process
		nature.GrassProc(grassid, wpr.TimeFactor)
	}	
	
	//block forever until the program is killed
	select{}
}
