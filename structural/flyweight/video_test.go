package flyweight

import (
	"design-patterns/structural/flyweight/other"
	"fmt"
	"testing"
)

func TestVideo(t *testing.T) {
	youtube := other.CacheYoutube{
		Service: other.ThirdPartyYoutubeImpl{},
	}

	// Minta video pertama
	video1 := youtube.GetVideo("https://youtube.com/watch/1")
	fmt.Println("Title:", video1.Flyweight.Title)
	fmt.Println("Thumbnail:", video1.Flyweight.Thumbnail)

	// Minta video kedua
	video2 := youtube.GetVideo("https://youtube.com/watch/2")
	fmt.Println("Title:", video2.Flyweight.Title)
	fmt.Println("Thumbnail:", video2.Flyweight.Thumbnail)

	// Minta video pertama lagi
	video3 := youtube.GetVideo("https://youtube.com/watch/1")
	fmt.Println("Title:", video3.Flyweight.Title)
	fmt.Println("Thumbnail:", video3.Flyweight.Thumbnail)
}
