package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Statistics struct {
	Views []View `xml:"views>view"`
}

type View struct {
	Name  string `xml:"name,attr" json:"name"`
	Zones []Zone `xml:"zones>zone" json:"zones"`
}

type Zone struct {
	Name string `xml:"name,attr" json:"name"`
	Type string `xml:"type" json:"type"`
}

func getStatistics(server string) (Statistics, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/xml/v3/zones", server))
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var data Statistics
	err = xml.Unmarshal(body, &data)

	return data, err
}

func GetZones(server string, zoneType string) ([]byte, error) {
	data, err := getStatistics(server)

	if err != nil {
		return nil, err
	}

	var zones []Zone

	for _, view := range data.Views {
		zones = append(zones, view.Zones...)
	}
	if zoneType == "" {

		jsonrep, err := json.Marshal(zones)

		return jsonrep, err

	} else {
		var filteredZones []Zone
		for _, zone := range zones {
			if zoneType == zone.Type {
				filteredZones = append(filteredZones, zone)
			}
		}
		jsonrep, err := json.Marshal(filteredZones)

		return jsonrep, err
	}
}
