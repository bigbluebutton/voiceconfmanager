package fseslclient

type FsEslEvent struct {
	name string
}

func (event FsEslEvent) Name() string {
	return event.name
}

type VoiceUserJoinedEvent struct {
	FsEslEvent

	ConferenceId string
	VoiceUserId  string
	CallerIdNum  string
	CallerIdName string
	Muted        bool
	Talking      bool
	Locked       bool
	UserId       string
}

type VoiceUserLeftEvent struct {
	ConferenceId string
	VoiceUserId  string
}

type VoiceUserMutedEvent struct {
	ConferenceId string
	VoiceUserId  string
	Muted        bool
}

type VoiceUserTalkingEvent struct {
	ConferenceId string
	VoiceUserId  string
	Talking      bool
}

type VoiceStartRecordingEvent struct {
	ConferenceId string
	Timestamp    string
	Filename     string
	Recording    bool
}
