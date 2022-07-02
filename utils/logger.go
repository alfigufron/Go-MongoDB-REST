package utils

import (
	"time"

	"github.com/ik5/rotatefilehook"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logLevel := logrus.InfoLevel
	logFormatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02 Jan 2006 15:04:05",
	}
	filename := "logs/" + time.Now().Format("January_2006") + ".log"

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    50,
		MaxBackups: 3,
		Level:      logLevel,
		Formatter:  logFormatter,
	})

	if err != nil {
		logrus.Fatal("Failed Setup RotateFileHook %v", err)
	}

	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(logFormatter)

	logrus.AddHook(rotateFileHook)
}
