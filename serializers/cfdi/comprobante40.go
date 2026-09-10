package cfdi

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/helpers"
	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/serializers"
)

type Comprobante40 struct {
	Version           string                  `xml:"Version,attr" bson:"Version" json:"Version"`
	Serie             *string                 `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie,omitempty"`
	Folio             *string                 `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	Fecha             time.Time               `xml:"Fecha,attr" bson:"Fecha" json:"Fecha"`
	Sello             string                  `xml:"Sello,attr" bson:"Sello" json:"Sello"`
	FormaPago         *string                 `xml:"FormaPago,attr" bson:"FormaPago,omitempty" json:"FormaPago,omitempty"`
	NoCertificado     string                  `xml:"NoCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado       string                  `xml:"Certificado,attr" bson:"Certificado" json:"Certificado"`
	CondicionesDePago *string                 `xml:"CondicionesDePago,attr" bson:"CondicionesDePago,omitempty" json:"CondicionesDePago,omitempty"`
	SubTotal          float64                 `xml:"SubTotal,attr" bson:"SubTotal" json:"SubTotal"`
	Descuento         *float64                `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento,omitempty"`
	Moneda            string                  `xml:"Moneda,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio        *float64                `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty" json:"TipoCambio,omitempty"`
	Total             float64                 `xml:"Total,attr" bson:"Total" json:"Total"`
	TipoDeComprobante string                  `xml:"TipoDeComprobante,attr" bson:"TipoDeComprobante" json:"TipoDeComprobante"`
	Exportacion       string                  `xml:"Exportacion,attr" bson:"Exportacion" json:"Exportacion"`
	MetodoPago        *string                 `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty" json:"MetodoPago,omitempty"`
	LugarExpedicion   string                  `xml:"LugarExpedicion,attr" bson:"LugarExpedicion" json:"LugarExpedicion"`
	Confirmacion      *string                 `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty" json:"Confirmacion,omitempty"`
	InformacionGlobal *InformacionGlobal40    `xml:"InformacionGlobal" bson:"InformacionGlobal,omitempty" json:"InformacionGlobal,omitempty"`
	CfdiRelacionados  *[]CfdiRelacionados40   `xml:"CfdiRelacionados" bson:"CfdiRelacionados,omitempty" json:"CfdiRelacionados,omitempty"`
	Emisor            Emisor40                `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor          Receptor40              `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	RfcProvCertif     string                  `bson:"RfcProvCertif" json:"RfcProvCertif"`
	Conceptos         []Concepto40            `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
	Impuestos         *Impuestos40            `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
	Complemento       serializers.Complemento `xml:"Complemento" bson:"Complemento" json:"Complemento"`
	Addenda           *serializers.Addenda    `xml:"Addenda" bson:"Addenda,omitempty" json:"Addenda,omitempty"`
}

func (c *Comprobante40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Fecha string `xml:"Fecha,attr"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	fecha, err := helpers.ParseDatetime(aux.Fecha)
	if err != nil {
		return err
	}
	c.Fecha = fecha
	return nil
}

type InformacionGlobal40 struct {
	Periodicidad string `xml:"Periodicidad,attr" bson:"Periodicidad" json:"Periodicidad"`
	Meses        string `xml:"Meses,attr" bson:"Meses" json:"Meses"`
	Anio         string `xml:"Año,attr" bson:"Anio" json:"Anio"`
}

type CfdiRelacionados40 struct {
	TipoRelacion    string              `xml:"TipoRelacion,attr" bson:"TipoRelacion" json:"TipoRelacion"`
	CfdiRelacionado []CfdiRelacionado40 `xml:"CfdiRelacionado" bson:"CfdiRelacionado" json:"CfdiRelacionado"`
}

type CfdiRelacionado40 struct {
	UUID string `xml:"UUID,attr" bson:"UUID" json:"UUID"`
}

func (cr *CfdiRelacionado40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		UUID string `xml:"UUID,attr"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	cr.UUID = strings.ToUpper(aux.UUID)
	return nil
}

type Emisor40 struct {
	Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre           string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	RegimenFiscal    string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
	FacAtrAdquirente *string `xml:"FacAtrAdquirente,attr" bson:"FacAtrAdquirente,omitempty" json:"FacAtrAdquirente,omitempty"`
}

type Receptor40 struct {
	Rfc                     string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre                  string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	DomicilioFiscalReceptor string  `xml:"DomicilioFiscalReceptor,attr" bson:"DomicilioFiscalReceptor" json:"DomicilioFiscalReceptor"`
	ResidenciaFiscal        *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	NumRegIdTrib            *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	RegimenFiscalReceptor   string  `xml:"RegimenFiscalReceptor,attr" bson:"RegimenFiscalReceptor" json:"RegimenFiscalReceptor"`
	UsoCFDI                 string  `xml:"UsoCFDI,attr" bson:"UsoCFDI" json:"UsoCFDI"`
}

type Concepto40 struct {
	ClaveProdServ       string                           `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProdServ"`
	NoIdentificacion    *string                          `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Cantidad            float64                          `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	ClaveUnidad         string                           `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad              *string                          `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Descripcion         string                           `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       float64                          `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe             float64                          `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Descuento           *float64                         `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento,omitempty"`
	ObjetoImp           string                           `xml:"ObjetoImp,attr" bson:"ObjetoImp" json:"ObjetoImp"`
	Impuestos           *ImpuestosConcepto40             `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
	ACuentaTerceros     *ACuentaTerceros40               `xml:"ACuentaTerceros" bson:"ACuentaTerceros,omitempty" json:"ACuentaTerceros,omitempty"`
	InformacionAduanera *[]InformacionAduanera40         `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
	CuentaPredial       *[]CuentaPredial40               `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial,omitempty"`
	ComplementoConcepto *serializers.ComplementoConcepto `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty" json:"ComplementoConcepto,omitempty"`
	Parte               *[]Parte40                       `xml:"Parte" bson:"Parte,omitempty" json:"Parte,omitempty"`
}

type ImpuestosConcepto40 struct {
	Traslados   *[]TrasladoConcepto40  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
	Retenciones *[]RetencionConcepto40 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
}

type TrasladoConcepto40 struct {
	Base       float64  `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
}

type RetencionConcepto40 struct {
	Base       float64 `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type ACuentaTerceros40 struct {
	RfcACuentaTerceros             string `xml:"RfcACuentaTerceros,attr" bson:"RfcACuentaTerceros" json:"RfcACuentaTerceros"`
	NombreACuentaTerceros          string `xml:"NombreACuentaTerceros,attr" bson:"NombreACuentaTerceros" json:"NombreACuentaTerceros"`
	RegimenFiscalACuentaTerceros   string `xml:"RegimenFiscalACuentaTerceros,attr" bson:"RegimenFiscalACuentaTerceros" json:"RegimenFiscalACuentaTerceros"`
	DomicilioFiscalACuentaTerceros string `xml:"DomicilioFiscalACuentaTerceros,attr" bson:"DomicilioFiscalACuentaTerceros" json:"DomicilioFiscalACuentaTerceros"`
}

type InformacionAduanera40 struct {
	NumeroPedimento string `xml:"NumeroPedimento,attr" bson:"NumeroPedimento" json:"NumeroPedimento"`
}

type CuentaPredial40 struct {
	Numero string `xml:"Numero,attr" bson:"Numero" json:"Numero"`
}

type Parte40 struct {
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr" bson:"ClaveProdServ" json:"ClaveProdServ"`
	NoIdentificacion    *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Cantidad            float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Descripcion         string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       *float64                 `xml:"ValorUnitario,attr" bson:"ValorUnitario,omitempty" json:"ValorUnitario,omitempty"`
	Importe             *float64                 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduanera40 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
}

type Impuestos40 struct {
	TotalImpuestosRetenidos   *float64                `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos,omitempty"`
	TotalImpuestosTrasladados *float64                `xml:"TotalImpuestosTrasladados,attr" bson:"TotalImpuestosTrasladados,omitempty" json:"TotalImpuestosTrasladados,omitempty"`
	Retenciones               *[]RetencionImpuestos40 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
	Traslados                 *[]TrasladoImpuestos40  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
}

type RetencionImpuestos40 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type TrasladoImpuestos40 struct {
	Base       float64  `xml:"Base,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"Importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
}
