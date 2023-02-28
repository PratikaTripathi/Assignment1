package structt

import "sync"

type Video struct {
	VideoId    string
	VideoCount int
	Lock       sync.Mutex
	Wait       sync.WaitGroup
}

var Videos Video


func GetVideos() Video{
	Videos.VideoId = "ndwg218y"
	Videos.VideoCount = 0
	return Videos
}
