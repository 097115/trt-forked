package main

import (
	"fmt"
	"strings"

	"github.com/Murtaza-Udaipurwala/trt/cli"
	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/Murtaza-Udaipurwala/trt/tui"
)

const AppVersion = "0.0.4"

func main() {
	version, username, password, url, port := cli.ParseArgs()

	if version {
		fmt.Println("TRT v." + AppVersion)
		return
	}

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	session := core.Session{}
	session.URL = fmt.Sprintf("%s:%d/transmission/rpc", url, port)
	session.Username = username
	session.Password = password
	session.CompileRegex()
	session.NewSessionID()

	tui.Run(&session)
}
