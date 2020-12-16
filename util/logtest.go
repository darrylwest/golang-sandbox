package main

import (
	"fmt"
	"github.com/darrylwest/cassava-logger/logger"
	"time"
)

func main() {
	fmt.Printf("Version: %s\n", logger.Version)

	handler, _ := logger.NewTimeRotatingFileHandler("./logger", logger.WhenDay, 1)

	logr := logger.NewLogger(handler)

	for i := 0; i < 10; i++ {
		fmt.Printf("logging %d\n", i)

		logr.Debug("my debug log statement %d", i)
		logr.Info("my info log statement %d", i)
		logr.Warn("my warn log statement %d", i)
		// logr.Error("my error log statement %d", i)

		time.Sleep(50 * time.Millisecond)
	}

}
