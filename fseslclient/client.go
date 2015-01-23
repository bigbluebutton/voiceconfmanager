package fseslclient

import (
	"errors"
	"fmt"
	"github.com/fiorix/go-eventsocket/eventsocket"
	"regexp"
	"strconv"
)

func getMemberIdFromEvent(e *eventsocket.Event) string {
	return e.Get("Member-ID")
}

func getCallerIdNumFromEvent(e *eventsocket.Event) string {
	return e.Get("Caller-Caller-ID-Number")
}

func getCallerIdNameFromEvent(e *eventsocket.Event) string {
	return e.Get("Caller-Caller-ID-Name")
}

func getRecordFilenameFromEvent(e *eventsocket.Event) string {
	return e.Get("Path")
}

func getSpeakFromEvent(e *eventsocket.Event) string {
	return e.Get("Speak")
}

func isUserMuted(s string) bool {
	speak, err := strconv.ParseBool(s)
	muted := false
	if err != nil {
		fmt.Printf("Error parsing bool %s\n", err)
		return false
	} else {
		if speak {
			muted = false
			fmt.Printf("speak %s\n", speak)
		} else {
			muted = true
			fmt.Printf("speak %s\n", speak)
		}
	}

	return muted
}

func isUserTalking(e *eventsocket.Event) bool {
	talk, err := strconv.ParseBool(e.Get("Talking"))
	talking := false
	if err != nil {
		if talk {
			talking = false
		} else {
			talking = true
		}
	}

	return talking
}

func isUserJoinedThoughGlobalAudio(s string) bool {
	re := regexp.MustCompile("(GLOBAL_AUDIO)_(.*)$")
	result_slice := re.FindAllStringSubmatch(s, -1)
	if result_slice == nil {
		return false
	} else {
		return true
	}
}

type UserType struct {
	UserId   string
	Username string
}

func getUser(s string) (UserType, error) {
	re := regexp.MustCompile("(?P<userid>.*)-bbbID-(?P<username>.*)$")
	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch(s, -1)

	if r2 != nil {
		md := map[string]string{}
		for i, n := range r2[0] {
			md[n1[i]] = n
		}
		return UserType{md["userid"], md["username"]}, nil
	} else {
		return UserType{}, errors.New("Failed to match regexp.")
	}
}

func handleUseJoinedEvent(e *eventsocket.Event) VoiceUserJoinedEvent {
	memberId := getMemberIdFromEvent(e)
	callerIdNum := getCallerIdNumFromEvent(e)
	callerIdName := getCallerIdNameFromEvent(e)
	speak := getSpeakFromEvent(e)
	muted := isUserMuted(speak)
	talking := isUserTalking(e)
	confName := e.Get("Conference-Name")
	//	confSize := e.Get("Conference-Size")

	//   voiceUserId := callerIdName
	globalAudio := isUserJoinedThoughGlobalAudio(callerIdName)

	voiceUser := VoiceUserJoinedEvent{}

	if globalAudio != true {
		user, err := getUser(callerIdName)
		if err != nil {
			voiceUser := VoiceUserJoinedEvent{}
			voiceUser.ConferenceId = confName
			voiceUser.VoiceUserId = user.UserId
			voiceUser.CallerIdNum = callerIdNum
			voiceUser.CallerIdName = user.Username
			voiceUser.Muted = muted
			voiceUser.Talking = talking
			voiceUser.Locked = false
			voiceUser.UserId = memberId

		}

	}
	return voiceUser
}
