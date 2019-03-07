// genlifx is a tool used to parse the machine readable version of the
// LIFX LAN Protocol, and generates constants, data structures and
// methods that meet the lan.Payloader interface.
package main

import (
	"flag"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

var fURL = flag.String(
	"url",
	"https://raw.githubusercontent.com/LIFX/public-protocol/master/protocol.yml",
	"url of LIFX machine-readable protocol in YAML format",
)

var data struct {
	Enums   map[string]enum
	Fields  map[string]field
	Packets map[string]map[string]packet
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("genlifx: ")

	resp, err := http.Get(*fURL)
	if err != nil {
		log.Fatalf("Unable to GET %s:\n\t%s", *fURL, err)
	}
	defer resp.Body.Close()

	if err := yaml.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Unable to decode product YAML:\n\t%s", err)
	}

	genEnums(data.Enums)
	genFields(data.Fields)
	genPackets(data.Packets)
}
