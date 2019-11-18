package messages

import (
	"html/template"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"go.tmthrgd.dev/booked/web"
)

func Mount(r chi.Router) {
	r.Get("/messages", web.TrailingSlashRedirect)
	r.Get("/messages/*", handler())
}

func handler() http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		path := chi.URLParam(r, "*")
		if path == "" {
			return threadsHandler(w, r)
		}

		if t, ok := threads[path]; ok {
			return threadHandler(w, r, t)
		}

		return os.ErrNotExist
	})
}

var threadsTmpl = web.NewTemplate("threads.tmpl", nil)

var threadsHandlerData struct {
	once                      sync.Once
	Archived, Inbox, Requests []string
}

func threadsHandler(w http.ResponseWriter, r *http.Request) error {
	data := &threadsHandlerData
	data.once.Do(func() {
		for name := range threads {
			switch {
			case strings.HasPrefix(name, "archived_threads/"):
				data.Archived = append(data.Archived, name)
			case strings.HasPrefix(name, "inbox/"):
				data.Inbox = append(data.Inbox, name)
			case strings.HasPrefix(name, "message_requests/"):
				data.Requests = append(data.Requests, name)
			}
		}

		sort.Strings(data.Archived)
		sort.Strings(data.Inbox)
		sort.Strings(data.Requests)
	})

	return web.WriteTemplateResponse(w, threadsTmpl, data)
}

var messageTmpl = web.NewTemplate("message.tmpl", template.FuncMap{
	"sub": func(a, b int) int { return a - b },
})

func threadHandler(w http.ResponseWriter, r *http.Request, t *threadJSON) error {
	senderCount := make(map[string]int)
	for _, p := range t.Participants {
		senderCount[p.Name] = 0
	}

	var (
		photos, videos, audio, gifs, files     int
		stickers, plans, shares, calls, missed int
	)
	stickerCount := make(map[string]int)
	for _, msg := range t.Messages {
		senderCount[msg.SenderName]++

		photos += len(msg.Photos)
		videos += len(msg.Videos)
		audio += len(msg.AudioFiles)
		gifs += len(msg.GIFs)
		files += len(msg.Files)

		if msg.Sticker.URI != "" {
			stickers++
			stickerCount[msg.Sticker.URI]++
		}

		switch msg.Type {
		case messagePlan:
			plans++
		case messageShare:
			shares++
		case messageCall:
			calls++

			if msg.Missed {
				missed++
			}
		}
	}

	senders := make([]senderValue, 0, len(senderCount))
	for sender, count := range senderCount {
		senders = append(senders, senderValue{sender, count})
	}
	sort.Slice(senders, func(i, j int) bool {
		if senders[i].MessageCount != senders[j].MessageCount {
			return senders[i].MessageCount > senders[j].MessageCount
		}

		return senders[i].SenderName < senders[j].SenderName
	})

	sticker := make([]stickerValue, 0, len(stickerCount))
	for uri, count := range stickerCount {
		sticker = append(sticker, stickerValue{uri, count})
	}
	sort.Slice(sticker, func(i, j int) bool {
		if sticker[i].Count != sticker[j].Count {
			return sticker[i].Count > sticker[j].Count
		}

		return sticker[i].URI < sticker[j].URI // Stable sort.
	})

	return web.WriteTemplateResponse(w, messageTmpl, &struct {
		*threadJSON

		Senders  []senderValue
		Stickers []stickerValue

		PhotosCount, VideosCount, AudioCount, GIFsCount, FilesCount          int
		StickersCount, PlansCount, SharesCount, CallsCount, MissedCallsCount int
	}{
		t,

		senders,
		sticker,

		photos, videos, audio, gifs, files,
		stickers, plans, shares, calls, missed,
	})
}

type senderValue struct {
	SenderName   string
	MessageCount int
}

type stickerValue struct {
	URI   string
	Count int
}
