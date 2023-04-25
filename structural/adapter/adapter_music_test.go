package adapter

import (
	"design-patterns/structural/adapter/other"
	"fmt"
	"testing"
)

func TestAdapterMusic(t *testing.T) {
	mp3 := other.Mp3{
		Data: []byte(`This is Taylor Swift song`),
	}

	walkman := other.Walkman{}

	adapter := other.Mp3ToKasetAdapter{MediaPlayer: walkman}
	fmt.Println(adapter.Play(mp3))
}
