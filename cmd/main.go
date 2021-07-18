package main

import (
	"fmt"
	"github.com/asppj/droneDeploy/conf"
	http2 "github.com/asppj/droneDeploy/internal/httpServer"

	"github.com/spf13/pflag"
)

func main() {
	versionInfo := pflag.BoolP("version", "v", false, "show version info.")
	pflag.Parse()
	fmt.Printf("%s\n\n", conf.BuildVersion())
	if versionInfo != nil && *versionInfo {
		return
	}
	http2.Init()
}
