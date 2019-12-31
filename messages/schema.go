package messages

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var threads = make(map[string]*threadJSON)

func IsMessageJSON(f *zip.File) bool {
	return strings.HasPrefix(path.Base(f.Name), "message_") &&
		path.Ext(f.Name) == ".json"
}

func ParseMessageJSON(f *zip.File) error {
	fail := func(err error) error {
		return fmt.Errorf("booked: failed to parse %s: %w", f.Name, err)
	}

	if !IsMessageJSON(f) {
		panic("booked: internal error: not message.json")
	}

	rc, err := f.Open()
	if err != nil {
		return fail(err)
	}
	defer rc.Close()

	dec := json.NewDecoder(rc)
	dec.DisallowUnknownFields()

	var msg threadJSON
	if err := dec.Decode(&msg); err != nil {
		return fail(err)
	}

	if dec.More() {
		return fail(errors.New("trailing json data"))
	}

	if err := rc.Close(); err != nil {
		return fail(err)
	}

	switch msg.ThreadType {
	case threadRegular:
	case threadRegularGroup:
	case threadPending:
	default:
		return fail(fmt.Errorf("unknown thread_type %q", msg.ThreadType))
	}

	for _, m := range msg.Messages {
		switch m.Type {
		case messageGeneric:
		case messageShare:
		case messagePlan:
		case messageCall:
		case messageSubscribe:
		case messageUnsubscribe:
		default:
			return fail(fmt.Errorf("unknown message type %q", m.Type))
		}
	}

	var n int
	for _, m := range msg.Messages {
		if !isSpamMessage(m) {
			msg.Messages[n] = m
			n++
		}
	}
	if n == 0 {
		return nil
	}
	for i := range msg.Messages[n:] {
		msg.Messages[n+i] = nil
	}
	msg.Messages = msg.Messages[:n]

	if t, dup := threads[msg.ThreadPath]; dup {
		t.Messages = append(t.Messages, msg.Messages...)
		sortMessages(t)
	} else {
		sortMessages(&msg)
		threads[msg.ThreadPath] = &msg
	}

	return nil
}

func sortMessages(t *threadJSON) {
	sort.Slice(t.Messages, func(i, j int) bool {
		return t.Messages[i].TimestampMS > t.Messages[j].TimestampMS
	})
}

func isSpamMessage(m *messageJSON) bool {
	switch m.Type {
	case messageGeneric:
		switch {
		case m.Content == "You can now call each other and see information such as Active Status and when you've read messages.":
			return true
		case strings.HasPrefix(string(m.Content), "Say hi to your new Facebook friend, ") && strings.HasSuffix(string(m.Content), "."):
			return true
		case m.Content == m.SenderName+" just joined Messenger! Be the first to send a welcome message or sticker.":
			return true
		case m.Content == m.SenderName+" is on Messenger.":
			return true
		}
	case messageSubscribe:
		return true
	case messageUnsubscribe:
		return true
	}

	return false
}

type utf8String string

var latin1Encoder = charmap.ISO8859_1.NewEncoder()

func (s *utf8String) UnmarshalJSON(data []byte) error {
	var sv string
	err := json.Unmarshal(data, &sv)
	if err != nil {
		return err
	}

	*(*string)(s), err = latin1Encoder.String(sv)
	return err
}

type threadJSON struct {
	Participants []struct {
		Name utf8String
	}
	Messages           []*messageJSON
	Title              utf8String
	IsStillParticipant bool       `json:"is_still_participant"`
	ThreadType         threadType `json:"thread_type"`
	ThreadPath         string     `json:"thread_path"`
}

type messageJSON struct {
	SenderName  utf8String `json:"sender_name"`
	TimestampMS uint64     `json:"timestamp_ms"`
	Content     utf8String
	Photos      []struct {
		URI               string
		CreationTimestamp uint64 `json:"creation_timestamp"`
	}
	Videos []struct {
		URI               string
		CreationTimestamp uint64 `json:"creation_timestamp"`
		Thumbnail         struct {
			URI string
		}
	}
	AudioFiles []struct {
		URI               string
		CreationTimestamp uint64 `json:"creation_timestamp"`
	} `json:"audio_files"`
	GIFs []struct {
		URI string
	}
	Files []struct {
		URI               string
		CreationTimestamp uint64 `json:"creation_timestamp"`
	}
	Sticker struct {
		URI string
	}
	Reactions []struct {
		Reaction utf8String
		Actor    utf8String
	}
	Plan struct {
		Title     utf8String
		Location  utf8String
		Timestamp uint64
	}
	Share struct {
		Link      string
		ShareText utf8String `json:"share_text"`
	}
	CallDuration uint64 `json:"call_duration"`
	Type         messageType
	Missed       bool
	Users        []struct {
		Name utf8String
	}
}

type messageType string

const (
	messageGeneric     messageType = "Generic"
	messageShare       messageType = "Share"
	messagePlan        messageType = "Plan"
	messageCall        messageType = "Call"
	messageSubscribe   messageType = "Subscribe"
	messageUnsubscribe messageType = "Unsubscribe"
)

type threadType string

const (
	threadRegular      threadType = "Regular"
	threadRegularGroup threadType = "RegularGroup"
	threadPending      threadType = "Pending"
)
