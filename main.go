//
// This program consumes a CSV file of devices that need to be added
// to a given collection in Horde and adds them to Horde.  The program
// expects a CSV file that conforms to RFC 4180 and which contains the
// following fields in order:
//
//   1. Device name
//   2. IMEI
//   3. IMSI
//
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	nbiot "github.com/telenordigital/nbiot-go"
)

var csvFilename = flag.String("csv", "", "Name of CSV file")
var collectionID = flag.String("collection", "", "Collection ID of collection we add devices to")

func main() {
	flag.Parse()

	// Check that we got both a CSV file and a Collection ID
	if *csvFilename == "" || *collectionID == "" {
		log.Fatalf("Please specify both CSV file and Collection ID")
	}

	// Open the CSV file and make it ready for being parsed as CSV
	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalf("Unable to open %s: %v", *csvFilename, err)
	}
	defer csvFile.Close()
	r := csv.NewReader(bufio.NewReader(csvFile))

	// Create a Horde client
	client, err := nbiot.New()
	if err != nil {
		log.Fatal("Error creating client:", err)
	}

	// Process the CSV file and add each entry to Horde.  If an entry
	// fails we keep going.  One reason for failure might be that the
	// device already exists.  You should always check the output to
	// ensure that everything went fine.
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file %s: %v", *csvFilename, err)
		}

		name := line[0]
		imei := line[1]
		imsi := line[2]

		createdDevice, err := client.CreateDevice(*collectionID, nbiot.Device{
			CollectionID: collectionID,
			IMEI:         &imei,
			IMSI:         &imsi,
			Tags: map[string]string{
				"name": name,
			},
		})

		if err != nil {
			log.Printf("Error creating device: %v", err)
		}
		log.Printf("Created device:  %+v", createdDevice)
	}

}
