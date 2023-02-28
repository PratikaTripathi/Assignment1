package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pratika/VideoPlatform/structt"
)



var Videos = structt.GetVideos()
var Video=structt.Video{}

func GetViews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	Videos.Wait.Add(1)
	ch := make(chan int)
	go getV(ch)
	fmt.Println(<-ch)
	Videos.Wait.Wait()
	s := "Number of views for Video with id " + Videos.VideoId + " is " + strconv.Itoa(Videos.VideoCount)
	w.Write([]byte(s))
}

func getV(ch chan<- int) {
	defer Videos.Wait.Done()
	Videos.Lock.Lock()
	defer Videos.Lock.Unlock()
	ch <- Videos.VideoCount
}

func IncrementViewCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	for i := 0; i < 10; i++ {
		Videos.Wait.Add(1)
		go inc()
		Videos.Wait.Wait()
	}
	w.Write([]byte("View Count increased by 10"))
}

func inc() {
	defer Videos.Wait.Done()
	Videos.Lock.Lock()
	defer Videos.Lock.Unlock()
	Videos.VideoCount++
}
