package other

import "fmt"

type VideoFlyweight struct {
	Title     string
	Thumbnail string
}

var flyweights map[string]*VideoFlyweight

func init() {
	// Inisialisasi daftar flyweight
	flyweights = make(map[string]*VideoFlyweight)

	// Menambahkan flyweight untuk video1
	flyweights["https://youtube.com/watch/1"] = &VideoFlyweight{
		Title:     "Video 1",
		Thumbnail: "https://thumbnail.com/1",
	}

	// Menambahkan flyweight untuk video2
	flyweights["https://youtube.com/watch/2"] = &VideoFlyweight{
		Title:     "Video 2",
		Thumbnail: "https://thumbnail.com/2",
	}
}

type ThirdPartyYoutubeLib interface {
	ListVideo() []string
}

type ThirdPartyYoutubeImpl struct{}

func (t ThirdPartyYoutubeImpl) ListVideo() []string {
	return []string{
		"https://youtube.com/watch/1",
		"https://youtube.com/watch/2",
	}
}

type Video struct {
	url       string
	Flyweight *VideoFlyweight
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
		// Jika video belum ada di cache, buat objek baru dengan flyweight yang sesuai
		fmt.Println("Fetching video from service...")
		flyweight, ok := flyweights[url]
		if !ok {
			// Jika flyweight tidak ditemukan, buat yang baru
			flyweight = &VideoFlyweight{
				Title:     fmt.Sprintf("Video %v", len(flyweights)+1),
				Thumbnail: fmt.Sprintf("https://thumbnail.com/%v", len(flyweights)+1),
			}
			flyweights[url] = flyweight
		}
		video = &Video{
			url:       url,
			Flyweight: flyweight,
		}
		c.videos[url] = video
	} else {
		fmt.Println("Fetching video from cache...")
	}
	return video
}
