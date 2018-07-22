package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JKolios/whatsflyingoverhead/conf"
	"github.com/JKolios/whatsflyingoverhead/dump1090-fa"
	"github.com/fsnotify/fsnotify"
)

func main() {
	var config *conf.Config
	var err error
	if config, err = conf.LoadConfig(); err != nil {
		log.Fatalf("Error loading configuration: %v", err.Error())
	}
	log.Printf("%+v", *config)

	var fileWatcher *fsnotify.Watcher
	if fileWatcher, err = fsnotify.NewWatcher(); err != nil {
		log.Fatalf("Error creating filesystem watcher: %v", err.Error())
	}

	aircraftFilePath := config.JSONFileDir + "/aircraft.json"

	if err = fileWatcher.Add(aircraftFilePath); err != nil {
		log.Fatalf("Error watching aircraft file: %v", err.Error())
	}
	defer fileWatcher.Close()

	waitChan := make(chan int)

	go func() {

		for {
			select {
			case <-fileWatcher.Events:
				log.Println("FS event")
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

				for _, aircraft := range visibleAircraftFile.Aircraft {
					if aircraft.HasCoordinates() {
						log.Printf("Flight: %v, Distance: %v\n", aircraft.Flight, aircraft.Distance(config.ReceiverLat, config.ReceiverLon, config.ReceiverHeight))
					}
				}

				log.Printf("%+v\n", visibleAircraftFile)

			}
		}
	}()

	<-waitChan

}
