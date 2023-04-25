package other

type Mp3ToKasetAdapter struct {
	MediaPlayer Walkman
}

func (m Mp3ToKasetAdapter) Play(mp3 Mp3) string {
	kaset := Kaset{
		PitaMusik: string(mp3.Data),
	}

	return m.MediaPlayer.Play(kaset)
}
