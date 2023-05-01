package flyweight

import (
	"design-patterns/structural/flyweight/other"
	"fmt"
	"testing"
)

func TestVideo(t *testing.T) {
	cache := other.CacheYoutube{
		Service: other.ThirdPartyYoutubeImpl{},
	}

	fmt.Println(cache.GetVideo("https://youtube.com/watch/1"))
	fmt.Println(cache.GetVideo("https://youtube.com/watch/2"))
	fmt.Println(cache.GetVideo("https://youtube.com/watch/1"))
}
