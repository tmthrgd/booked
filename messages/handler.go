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
	var (
		photos, videos, audio, gifs, files     int
		stickers, plans, shares, calls, missed int
	)
	for _, msg := range t.Messages {
		photos += len(msg.Photos)
		videos += len(msg.Videos)
		audio += len(msg.AudioFiles)
		gifs += len(msg.GIFs)
		files += len(msg.Files)

		if msg.Sticker.URI != "" {
			stickers++
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

	return web.WriteTemplateResponse(w, messageTmpl, &struct {
		*threadJSON

		PhotosCount, VideosCount, AudioCount, GIFsCount, FilesCount          int
		StickersCount, PlansCount, SharesCount, CallsCount, MissedCallsCount int
	}{
		t,

		photos, videos, audio, gifs, files,
		stickers, plans, shares, calls, missed,
	})
}
