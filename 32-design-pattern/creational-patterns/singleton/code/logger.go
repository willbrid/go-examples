package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	file *os.File
}

var lock = &sync.Mutex{}
var loggerInstance *Logger

func getInstance() *Logger {
	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if loggerInstance == nil {
			logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

			if err != nil {
				log.Fatalf("Erreur d'ouverture du fichier de log : %v", err)
			}

			fmt.Println("Creating Logger instance now.")
			loggerInstance = &Logger{file: logFile}
		} else {
			fmt.Println("Logger instance already created.")
		}
	} else {
		fmt.Println("Logger instance already created.")
	}

	return loggerInstance
}
