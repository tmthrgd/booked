package main

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"
)

var threads = make(map[string]*messageJSON)

func isMessageJSON(f *zip.File) bool {
	return strings.HasPrefix(path.Base(f.Name), "message_") &&
		path.Ext(f.Name) == ".json"
}

func parseMessageJSON(f *zip.File) error {
	fail := func(err error) error {
		return fmt.Errorf("booked: failed to parse %s: %w", f.Name, err)
	}

	if !isMessageJSON(f) {
		panic("booked: internal error: not message.json")
	}

	rc, err := f.Open()
	if err != nil {
		return fail(err)
	}
	defer rc.Close()

	dec := json.NewDecoder(rc)
	dec.DisallowUnknownFields()

	var msg messageJSON
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

	if t, dup := threads[msg.ThreadPath]; dup {
		t.Messages = append(t.Messages, msg.Messages...)
	} else {
		threads[msg.ThreadPath] = &msg
	}

	return nil
}

func sortMessages() {
	for _, t := range threads {
		sort.Slice(t.Messages, func(i, j int) bool {
			return t.Messages[i].TimestampMS >
				t.Messages[j].TimestampMS
		})
	}
}

type messageJSON struct {
	Participants []struct {
		Name string
	}
	Messages []struct {
		SenderName  string `json:"sender_name"`
		TimestampMS uint64 `json:"timestamp_ms"`
		Content     string
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
		GIFS []struct {
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
			Reaction string
			Actor    string
		}
		Plan struct {
			Title     string
			Location  string
			Timestamp uint64
		}
		Share struct {
			Link      string
			ShareText string `json:"share_text"`
		}
		CallDuration uint64 `json:"call_duration"`
		Type         messageType
		Missed       bool
		Users        []struct {
			Name string
		}
	}
	Title              string
	IsStillParticipant bool       `json:"is_still_participant"`
	ThreadType         threadType `json:"thread_type"`
	ThreadPath         string     `json:"thread_path"`
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
