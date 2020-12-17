package main

import (
	"github.com/ofabry/go-callvis/examples/main/webcrawpkg"
	"github.com/ofabry/go-callvis/examples/main/include"
)

func main() {
	rootCmd()
	initializeSetup()
	
}

func rootCmd(){
	var c WebCrawSchedule
	c.urlExecution()
	c.urlResistry()
	
}
func initializeSetup() {
	include.ConfSettings()
}

type WebCrawSchedule struct{}

func (WebCrawSchedule) urlExecution() {
	webcrawpkg.RunTask()
}

func (WebCrawSchedule) urlResistry() {
	webcrawpkg.T.Identifier()
}

