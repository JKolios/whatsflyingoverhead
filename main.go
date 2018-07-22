package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

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
			case event := <-fileWatcher.Events:
				log.Printf("FS event: %+v\n", event)

				if event.Op != fsnotify.Write {
					continue
				}

				aircraftFileBytes, err := ioutil.ReadFile(config.JSONFileDir + "/aircraft.json")
				if err != nil {
					log.Printf("Error reading the aircraft file: %v\n", err.Error())
					waitChan <- 0
					return
				}

				var visibleAircraftFile dump1090fa.AircraftFile
				if err = json.Unmarshal(aircraftFileBytes, &visibleAircraftFile); err != nil {
					log.Printf("Error unmarshaling the aircraft file: %v\n", err.Error())
					waitChan <- 0
					return
				}

				for _, aircraft := range visibleAircraftFile.Aircraft {
					if aircraft.HasCoordinates() {
						log.Printf("Flight: %v, Distance: %v\n", aircraft.Flight, aircraft.Distance(config.ReceiverLat, config.ReceiverLon, config.ReceiverHeight))
					}
				}
			}
		}
	}()

	go func() {
		var err error
		for {
			select {
			case err = <-fileWatcher.Errors:
				log.Printf("File Watcher error: %v", err.Error())
			}
		}
	}()

	<-waitChan

}
