package fseslclient

import (
	"testing"
)

func TestIsUserJoinedThoughGlobalAudioTrue(t *testing.T) {
	s := "GLOBAL_AUDIO_abc1234-foo bar"
    global := isUserJoinedThoughGlobalAudio(s)
    if global != true {
        t.Error("Expected true, got ", global)
    }
}

func TestIsUserJoinedThoughGlobalAudioFalse(t *testing.T) {
	s := "GLOBAL_AUDIOFAIL_abc1234-foo bar"
    global := isUserJoinedThoughGlobalAudio(s)
    if global != false {
        t.Error("Expected false, got ", global)
    }
}