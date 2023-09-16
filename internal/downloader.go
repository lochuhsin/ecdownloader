package internal

import (
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

type HttpFileDownloader struct {
	url           string
	dir           string
	acceptRanges  bool
	contentType   []string
	contentLength int
	splitCount    int
	client        RequestClient
}

func InitHttpFileDownloader(url, dir string) HttpFileDownloader {
	return HttpFileDownloader{url: url, dir: dir, acceptRanges: false, client: RequestClient{}}
}

func (h *HttpFileDownloader) Preprocess() {
	// Using head method to get basic download information
	resp := h.client.Head(h.url)

	headers := resp.Header
	if val, ok := headers["Content-Type"]; ok {
		h.contentType = val
	}
	if val, ok := headers["Accept-Ranges"]; ok {
		for _, str := range val {
			if str == "bytes" {
				h.acceptRanges = true
			}
		}
	}
	if val, ok := headers["Content-Length"]; ok {
		intval, err := strconv.Atoi(val[0])
		if err == nil {
			h.contentLength = intval
		}
	}
}

func (h *HttpFileDownloader) Run() {
	if h.acceptRanges {
		h.concurrentRun(h.url, h.dir)
	} else {
		h.run(h.url, h.dir)
	}

}

func (h *HttpFileDownloader) run(url, dir string) {
	r := h.client.Get(url)
	if r == nil {
		return
	}
	if r.Body == nil {
		log.Println("empty response body from", url)
		return
	}
	defer r.Body.Close()
	dest := path.Join(dir, path.Base(url))

	f, err := os.Create(dest)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *HttpFileDownloader) concurrentRun(url, dir string) {}
