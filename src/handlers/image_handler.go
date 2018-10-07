package handlers

import (
	"bytes"
	"fmt"
	"loaders"
	"log"
	"mime"
	"net/http"
)

var _contentType string
var _content string
var i int16 = 0

// ImageHandler: handler images
func ImageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Starting image handler proccess")
		normalizedURL := loaders.NormalizeURL(r.RequestURI)
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		if _contentType != "" && i > 2 {
			loadFromCache(w)
			return
		}
		// don't worry about errors
		response, e := http.Get(normalizedURL)
		if e != nil {
			log.Fatal(e)
		}
		contentType, _, _ := mime.ParseMediaType(response.Header.Get("Content-Type"))
		// if resp.ContentLength != 0 && !validContentType(p.ContentTypes, contentType) {
		if response.ContentLength == 0 {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			msg := fmt.Sprintf("forbidden content-type: %q", contentType)
			log.Print(msg)
			http.Error(w, msg, http.StatusForbidden)
			return
		}

		i = i + 1
		log.Print(contentType)
		w.Header().Set("Content-Type", contentType)
		_contentType = contentType

		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		_content = buf.String()
		fmt.Fprint(w, _content)
		// fmt.Fprint(w, response.Body)
		// io.Copy(w, response.Body)

	})
}

func loadFromCache(w http.ResponseWriter) {
	log.Print("Loaded from cache with content type " + _contentType)
	w.Header().Set("Content-Type", _contentType)
	fmt.Fprint(w, _content)
}
