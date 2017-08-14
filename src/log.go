package main

import (
	"github.com/inconshreveable/log15"
	"path/filepath"
)


func InitLogger() {
	// set log15
	logPath := "log"
	logFile := "gloria.log"

	log = log15.New()
	if err := DirExistedOrCreate(logPath); err != nil {
		log.Warn("log path create failed")
		return
	}
	handler, _ := log15.FileHandler(filepath.Join(logPath, logFile), log15.TerminalFormat())
	log.SetHandler(handler)
	log.Info("start logging", "path", filepath.Join(logPath, logFile))
}
