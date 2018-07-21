package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/JKolios/whatsflyingoverhead/conf"
	"github.com/JKolios/whatsflyingoverhead/dump1090-fa"
)

func main() {
	var config *conf.Config
	var err error
	if config, err = conf.LoadConfig(); err != nil {
		log.Fatalf("Error loading configuration: %v", err.Error())
	}
	log.Printf("%+v", *config)

	scanTicker := time.NewTicker(config.ScanPeriod)
	waitChan := make(chan int)

	go func() {

		for {
			select {
			case <-scanTicker.C:
				log.Println("Tick")
				aircraftFile, err := os.Open(config.JSONFileDir + "/aircraft.json")
				defer aircraftFile.Close()
				if err != nil {
					log.Printf("Error reading the aircraft file: %v\n", err.Error())
					waitChan <- 0
				}
				jsonDecoder := json.NewDecoder(aircraftFile)
				var visibleAircraftFile dump1090fa.AircraftFile
				if err = jsonDecoder.Decode(&visibleAircraftFile); err != nil {
					log.Printf("Error unmarshaling the aircraft file: %v\n", err.Error())
					waitChan <- 0
					return
				}

				log.Printf("%+v\n", visibleAircraftFile)

			}
		}
	}()

	<-waitChan

}
