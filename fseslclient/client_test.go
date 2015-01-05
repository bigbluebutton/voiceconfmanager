package fseslclient

import (
	"testing"
	"fmt"
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

func TestIsUserMutedTrue(t *testing.T) {
	speak := "true"
	fmt.Printf("speak=%s\n", speak)
    muted := isUserMuted(speak)
    if muted != false {
        t.Error("Expected false, got ", muted)
    }
}

func TestIsUserMutedFalse(t *testing.T) {
	speak := "false"
    muted := isUserMuted(speak)
    if muted != true {
        t.Error("Expected true, got ", muted)
    }
}
