package main

import (
    "net/http"
    "fmt"
    //~ "encoding/json"
    "github.com/mazecrown/arthos/core"
)

func lookHandler(w http.ResponseWriter, r *http.Request) {	
	//get central coords from params
	//build list of coords around that
	//add those coords to a map of 'local' coords starting at 0,0,0
	//get bulk of coords
	//for each result get add the contents of the coords to a new map
	//get bulk using the new map to get the contents of the coords
	//add the contents results to the corresponding 'local' coord
	//return the local coords and contents to the client 
	
	//get stuff
	//loop grasslist	
	b := core.GetBucket()
	grasslist := []string{}
	b.Get("grass:list", &grasslist)
	
	grasses, _ := b.GetBulk(grasslist)
	
	response := []string{}
	i := 0
	for k := range grasses {
		//get the body of the item and
		s := string(grasses[k].Body)
		
		if i < len(grasses)-1 {
			s = s + ","
		}
		response = append(response, s)
		i ++
	}
	
	//~ bytes, _ := json.Marshal(grasses)
	
	//write the response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, response)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
//~ func InitWeb() {
    http.HandleFunc("/api/look/", lookHandler)
    http.HandleFunc("/", staticHandler)
    http.ListenAndServe("localhost:7777", nil)
}
