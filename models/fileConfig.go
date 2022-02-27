package models

import "encoding/xml"

type FileConfig struct {
	XMLName      xml.Name `xml:"file_config"`
	BaseLocation string   `xml:"base_location"`
	MainFile     string   `xml:"main_file,attr"`
	BatFile      string   `xml:"bat_file,attr"`
}
