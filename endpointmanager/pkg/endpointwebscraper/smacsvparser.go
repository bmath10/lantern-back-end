package endpointwebscraper

import (
	"io"
	"strings"

	"github.com/onc-healthit/lantern-back-end/endpointmanager/pkg/helpers"

	"os"
	log "github.com/sirupsen/logrus"
)

func SMACSVParser(vendorURL string, fileToWriteTo string) {
	var lanternEntryList []LanternEntry
	var endpointEntryList EndpointList

	csvFilePath := "./SMAEndpointDirectory.csv"

	csvReader, file, err := helpers.QueryAndOpenCSV(vendorURL, csvFilePath, true)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var entry1 LanternEntry
		var entry2 LanternEntry

		organizationName := strings.TrimSpace(rec[0])
		URL1 := strings.TrimSpace(rec[1])
		URL2 := strings.TrimSpace(rec[2])

		if len(URL1) > 0 {
			entry1.OrganizationName = organizationName
			entry1.URL = URL1
			lanternEntryList = append(lanternEntryList, entry1)
		}

		if len(URL2) > 0 {
			entry2.OrganizationName = organizationName
			entry2.URL = URL2
			lanternEntryList = append(lanternEntryList, entry2)
		}
	}

	endpointEntryList.Endpoints = lanternEntryList

	err = WriteEndpointListFile(endpointEntryList, fileToWriteTo)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
