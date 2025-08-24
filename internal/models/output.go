package models

import (
	"encoding/xml"
)

//WmosMessage representa el mensaje de salida en formate WMOS XML
type WMOSMessage struct {
	XMLName xml.Name `xml:"WmosMessage"`
	Header Header `xml:"Header"`
	Body Body `xml:"Body"`
}