package handler

import (
	"ecdownloader/internal"
	"log"
	"os"
	"sync"
)

func ArgHandler(args internal.Args) {
	url, urls := args.Url, args.Urls.Get()
	dir := args.Directory
	status, err := isWritable(dir)
	if !status {
		log.Println(err)
		return
	}

	// parse other argument cases
	downloadUrls := []string{}
	if len(url) != 0 {
		downloadUrls = append(downloadUrls, url)
	}
	for _, url := range urls {
		if len(url) != 0 {
			downloadUrls = append(downloadUrls, url)
		}
	}
	// Expect all urls and directory is valid
	DownloadHandler(downloadUrls, dir)
}

func DownloadHandler(urls []string, dir string) {
	downloaders := []internal.HttpFileDownloader{}
	for _, u := range urls {
		downloaders = append(downloaders, internal.InitHttpFileDownloader(u, dir))
	}
	wg := sync.WaitGroup{}
	wg.Add(len(downloaders))
	for _, obj := range downloaders {
		go start(obj, &wg)
	}
	wg.Wait()
}

func start(downloader internal.HttpFileDownloader, wg *sync.WaitGroup) {
	defer wg.Done()
	downloader.Preprocess()
	downloader.Run()
}

func isWritable(path string) (bool, error) {
	file, err := os.CreateTemp(path, "tmpfile")
	if err != nil {
		return false, err
	}

	defer os.Remove(file.Name())
	defer file.Close()

	return true, nil
}
