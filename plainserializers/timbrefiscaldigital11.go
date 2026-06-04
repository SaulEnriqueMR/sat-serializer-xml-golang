package plainserializers

type TimbreFiscalDigital11 struct {
	Version          string  `xml:"Version,attr" bson:"Version" json:"Version"`
	UUID             string  `xml:"UUID,attr" bson:"UUID" json:"UUID"`
	FechaTimbrado    string  `xml:"FechaTimbrado,attr" json:"FechaTimbrado" bson:"FechaTimbrado"`
	RfcProvCertif    string  `xml:"RfcProvCertif,attr" bson:"RfcProvCertif" json:"RfcProvCertif"`
	Leyenda          *string `xml:"Leyenda,attr" bson:"Leyenda,omitempty" json:"Leyenda,omitempty"`
	SelloCFD         string  `xml:"SelloCFD,attr" bson:"SelloCFD" json:"SelloCFD"`
	NoCertificadoSAT string  `xml:"NoCertificadoSAT,attr" bson:"NoCertificadoSAT" json:"NoCertificadoSAT"`
	SelloSAT         string  `xml:"SelloSAT,attr" bson:"SelloSAT" json:"SelloSAT"`
}
