package internal

import "flag"

type ArrayStrArgs struct {
	urls []string
}

func (arr *ArrayStrArgs) String() string {
	return "TODO: handle this string representation"
}

func (arr *ArrayStrArgs) Set(flag string) error {
	arr.urls = append(arr.urls, flag)
	return nil
}

func (arr *ArrayStrArgs) Get() []string {
	return arr.urls
}

type Args struct {
	Url       string
	Urls      ArrayStrArgs
	Help      string
	Mode      string
	Directory string
}

var Argument Args = Args{}

func init() {
	flag.StringVar(&Argument.Url, "u", "", "Download url")
	flag.StringVar(&Argument.Url, "url", "", "Download url")
	flag.StringVar(&Argument.Help, "h", "", "Help")
	flag.StringVar(&Argument.Help, "help", "", "Help")
	flag.StringVar(&Argument.Mode, "m", "", "Program mode")
	flag.StringVar(&Argument.Mode, "mode", "", "Program mode")
	flag.StringVar(&Argument.Directory, "d", "./", "Download directory")
	flag.StringVar(&Argument.Directory, "dir", "./", "Download directory")
	flag.Var(&Argument.Urls, "urls", "input list of urls to download")
}
