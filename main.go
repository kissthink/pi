package main

import (
	"flag"
	"log"
	"os"
	"github.com/smhouse/pi/db"
	"github.com/smhouse/pi/mqtt"
	"github.com/smhouse/pi/ghttp"
)

var mqttPort, httpPort int
var dbFile, logFile string

func init() {
	flag.IntVar(&mqttPort, "q", 1883, "mqtt server port")
	flag.IntVar(&httpPort, "h", 80, "http server port")
	flag.StringVar(&dbFile, "d", "pi.db", "path to database file")
	flag.StringVar(&logFile, "l", "", "path to log file")
	flag.Parse()
}

func main() {
	if logFile != "" {
		logWriter, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		log.SetOutput(logWriter)
	}

	err := db.OpenDatabase(dbFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.CreateAdmin()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		err = mqtt.StartServer(mqttPort)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	ghttp.StartHTTP(httpPort)
}
