package other

type AudioPlayer interface {
	Play(mp3 Mp3) string
}

type Mp3 struct {
	Data []byte
}

func (mp3 Mp3) PlayAudio() string {
	return "Playing MP3 with data " + string(mp3.Data)
}
