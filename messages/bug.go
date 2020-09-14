package messages

import (
	"archive/zip"
	"context"
	"fmt"
	"net/url"
	"path"
	"reflect"
	"runtime"
	"strings"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

var (
	couldFix   uint64 // atomic
	couldntFix uint64 // atomic
)

func FixupURIs(rcs ...*zip.ReadCloser) error {
	eg, ctx := errgroup.WithContext(context.Background())
	sem := make(chan struct{}, runtime.NumCPU())

	for _, thread := range threads {
		for _, msg := range thread.Messages {
			select {
			case sem <- struct{}{}:
			case <-ctx.Done():
				return eg.Wait()
			}

			mv := reflect.ValueOf(msg).Elem()
			eg.Go(func() error {
				defer func() { <-sem }()
				return fixBrokenURIs(mv, rcs)
			})
		}
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	if couldFix > 0 {
		fmt.Printf("fixed %d broken interncache-frc.fbcdn.net URIs\n", couldFix)
	}
	if couldntFix > 0 {
		fmt.Printf("failed to fix %d broken interncache-frc.fbcdn.net URIs\n", couldntFix)
	}

	return nil
}

func fixBrokenURIs(v reflect.Value, rcs []*zip.ReadCloser) error {
	for i := 0; i < v.NumField(); i++ {
		ft := v.Type().Field(i)
		fv := v.Field(i)

		switch ft.Type.Kind() {
		case reflect.Struct:
			if err := fixBrokenURIs(fv, rcs); err != nil {
				return err
			}
			continue
		case reflect.Array, reflect.Slice:
			if ft.Type.Elem().Kind() != reflect.Struct {
				continue
			}

			for i := 0; i < fv.Len(); i++ {
				if err := fixBrokenURIs(fv.Index(i), rcs); err != nil {
					return err
				}
			}
			continue
		case reflect.String:
		default:
			continue
		}

		if err := fixBrokenURI(fv, rcs); err != nil {
			return err
		}
	}

	return nil
}

func fixBrokenURI(v reflect.Value, rcs []*zip.ReadCloser) error {
	// Can Facebook please stop breaking the data download thanks. Do they even
	// test it, like at all? It's ridiculous.

	bu := v.String()
	if !strings.HasPrefix(bu, "https://interncache-frc.fbcdn.net/") {
		return nil
	}

	u, err := url.Parse(bu)
	if err != nil {
		return err
	}

	name := path.Base(u.Path)
	for _, rc := range rcs {
		for _, file := range rc.File {
			if !matchNames(name, path.Base(file.Name)) {
				continue
			}

			v.SetString(file.Name)

			atomic.AddUint64(&couldFix, 1)
			return nil
		}
	}

	atomic.AddUint64(&couldntFix, 1)
	return nil
}

func matchNames(a, b string) bool {
	// Facebook. What even is this shit?

	extA, extB := path.Ext(a), path.Ext(b)
	if extA != extB {
		return false
	}

	aNoExt := strings.TrimSuffix(a, extA)
	bNoExt := strings.TrimSuffix(b, extB)
	return strings.HasPrefix(bNoExt, aNoExt)
}
