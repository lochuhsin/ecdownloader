package handler

import (
	"ecdownloader/internal"
	"log"
	"os"
)

func ArgHandler(args internal.Args) {
	url := args.Url
	urls := args.Urls.Get()
	dir := args.Directory

	status, err := isWritable(dir)
	if !status {
		log.Println(err)
		return
	}

	if len(url) != 0 {
		internal.HttpDownloader(url, dir)
	}

	if len(urls) != 0 {
		for _, u := range urls {
			internal.HttpDownloader(u, dir)
		}
	}

}

func isWritable(path string) (bool, error) {
	tmpFile := "tmpfile"

	file, err := os.CreateTemp(path, tmpFile)
	if err != nil {
		return false, err
	}

	defer os.Remove(file.Name())
	defer file.Close()

	return true, nil
}
