package messages

import (
	"html/template"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

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
	dateCount := make(map[string]int)
	if len(t.Messages) > 2 {
		newest := timeFromMS(t.Messages[0].TimestampMS)
		oldest := timeFromMS(t.Messages[len(t.Messages)-1].TimestampMS)

		for d := truncateDay(oldest); d.Before(newest); d = d.AddDate(0, 0, 1) {
			dateCount[d.Format("2006-01-02")] = 0
		}
	}

	senderCount := make(map[string]int)
	for _, p := range t.Participants {
		senderCount[string(p.Name)] = 0
	}

	var (
		photos, videos, audio, gifs, files     int
		stickers, plans, shares, calls, missed int
	)
	stickerCount := make(map[string]int)
	for _, msg := range t.Messages {
		d := timeFromMS(msg.TimestampMS)
		dateCount[d.Format("2006-01-02")]++

		senderCount[string(msg.SenderName)]++

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

	heatmap := make([]heatmapValue, 0, len(dateCount))
	for date, count := range dateCount {
		heatmap = append(heatmap, heatmapValue{date, count})
	}
	sort.Slice(heatmap, func(i, j int) bool {
		return heatmap[i].Date < heatmap[j].Date
	})

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

		Heatmap  []heatmapValue
		Senders  []senderValue
		Stickers []stickerValue

		PhotosCount, VideosCount, AudioCount, GIFsCount, FilesCount          int
		StickersCount, PlansCount, SharesCount, CallsCount, MissedCallsCount int
	}{
		t,

		heatmap,
		senders,
		sticker,

		photos, videos, audio, gifs, files,
		stickers, plans, shares, calls, missed,
	})
}

func timeFromMS(ms uint64) time.Time {
	return time.Unix(int64(ms/1000), int64(ms%1000*1000))
}

func truncateDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

type heatmapValue struct {
	Date         string
	MessageCount int
}

type senderValue struct {
	SenderName   string
	MessageCount int
}

type stickerValue struct {
	URI   string
	Count int
}
