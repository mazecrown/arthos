package nature

import (
	"math/rand"
	"time"	
	"log"
	"strconv"	
	"github.com/mazecrown/arthos/core"
)

type Grass struct {
    Name string `json:"name"`
    Height float64 `json:"amount"`
}

func random(min, max float64) float64 {
  return rand.Float64() * (max - min) + min
}

func CreateGrass() Grass{
	//create random number between 0.75 and 1.25 for our 'amount' of grass to give some variation
	r := random(0.75, 1.25)
	//create grass object
	grass := Grass{"Grass", r}
	return grass
}

func GrassProc(grassid string, timefactor int) {
	//do grassy things here
	ticker := time.NewTicker(time.Millisecond * time.Duration(timefactor))	
	
	go func() {
		log.Println("spawning process for " + grassid)
        for _ = range ticker.C {
			b := core.GetBucket()
			grass := Grass{}
			b.Get(grassid, &grass)
			
            //add a random number to the grasses existing height up to a maxheight
            r := random(0.5, 1.5)
            grass.Height = grass.Height + r
            var maxHeight float64 = 7
            if grass.Height > maxHeight {
				grass.Height = maxHeight
			}
            
            //write the grass back to the bucket
            core.SetToBucket(grass, grassid)
            
			log.Println(grassid + " grew by " + strconv.FormatFloat(r, 'f', 6, 64))
        }
    }()
}
