package internal

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

func HttpDownloader(url, dir string) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	dest := path.Join(dir, path.Base(r.URL.Path))

	f, err := os.Create(dest)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	n, err := io.Copy(f, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	s := strconv.Itoa(int(n))
	log.Println(s)
}
