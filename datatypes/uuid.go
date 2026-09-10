package datatypes

import (
	"encoding/xml"
	"strings"
)

type Uuid struct {
	T string
}

func (u *Uuid) UnmarshalXMLAttr(attr xml.Attr) error {
	u.T = strings.ToUpper(attr.Value)
	return nil
}
