package datatypes

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/helpers"
)

type ISODateTime struct {
	T time.Time
}

func (d *ISODateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	t, err := helpers.ParseDatetime(attr.Value)
	if err != nil {
		return err
	}
	d.T = t
	return nil
}
