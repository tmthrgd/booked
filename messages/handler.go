package messages

import (
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"go.tmthrgd.dev/booked/web"
)

func Mount(r chi.Router) {
	r.Get("/messages", web.TrailingSlashRedirect)
	r.Get("/messages/", handler())
}

func handler() http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return threadsHandler(w, r)
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
