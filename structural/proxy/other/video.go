package other

import (
	"fmt"
	"time"
)

type ThirdPartyYoutubeLib interface {
	ListVideo() []string
}

type ThirdPartyYoutubeImpl struct{}

func (t ThirdPartyYoutubeImpl) ListVideo() []string {
	<-time.After(200 * time.Millisecond)
	return []string{
		"https://youtube.com/watch/1",
		"https://youtube.com/watch/2",
	}
}

type Video struct {
	url string
}

type CacheYoutube struct {
	videos  map[string]*Video
	Service ThirdPartyYoutubeLib
}

func (c *CacheYoutube) GetVideo(url string) *Video {
	if c.videos == nil {
		c.videos = make(map[string]*Video)
	}
	video, ok := c.videos[url]
	if !ok {
		fmt.Println("Fetching video from service...")
		video = &Video{url: url}
		c.videos[url] = video
	} else {
		fmt.Println("Fetching video from cache...")
	}
	return video
}
