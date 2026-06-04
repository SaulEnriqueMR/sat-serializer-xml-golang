package comprobante

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/helpers"
)

type Comprobante40 struct {
	Version           string                `xml:"Version,attr" bson:"Version" json:"Version"`
	Serie             *string               `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie,omitempty"`
	Folio             *string               `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	Fecha             time.Time             `bson:"Fecha" json:"Fecha"`
	Sello             string                `xml:"Sello,attr" bson:"Sello" json:"Sello"`
	FormaPago         *string               `xml:"FormaPago,attr" bson:"FormaPago,omitempty" json:"FormaPago,omitempty"`
	NoCertificado     string                `xml:"NoCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	Certificado       string                `xml:"Certificado,attr" bson:"Certificado" json:"Certificado"`
	CondicionesPago   *string               `xml:"CondicionesDePago,attr" bson:"CondicionesPago,omitempty" json:"CondicionesPago,omitempty"`
	Subtotal          float64               `xml:"SubTotal,attr" bson:"Subtotal" json:"Subtotal"`
	Descuento         *float64              `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento,omitempty"`
	Moneda            string                `xml:"Moneda,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio        *float64              `xml:"TipoCambio,attr" bson:"TipoCambio,omitempty" json:"TipoCambio,omitempty"`
	Total             float64               `xml:"Total,attr" bson:"Total" json:"Total"`
	TipoComprobante   string                `xml:"TipoDeComprobante,attr" bson:"TipoComprobante" json:"TipoComprobante"`
	Exportacion       string                `xml:"Exportacion,attr" bson:"Exportacion" json:"Exportacion"`
	MetodoPago        *string               `xml:"MetodoPago,attr" bson:"MetodoPago,omitempty" json:"MetodoPago,omitempty"`
	LugarExpedicion   string                `xml:"LugarExpedicion,attr" bson:"LugarExpedicion" json:"LugarExpedicion"`
	Confirmacion      *string               `xml:"Confirmacion,attr" bson:"Confirmacion,omitempty" json:"Confirmacion,omitempty"`
	InformacionGlobal *InformacionGlobal40  `xml:"InformacionGlobal" bson:"InformacionGlobal,omitempty" json:"InformacionGlobal,omitempty"`
	CfdisRelacionados *[]CfdiRelacionados40 `xml:"CfdiRelacionados" bson:"CfdisRelacionados,omitempty" json:"CfdisRelacionados,omitempty"`
	Emisor            Emisor40              `xml:"Emisor" bson:"Emisor" json:"Emisor"`
	Receptor          Receptor40            `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Conceptos         []Concepto40          `xml:"Conceptos>Concepto" bson:"Conceptos" json:"Conceptos"`
	Impuestos         *Impuestos40          `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
	// Complemento       Complemento            `xml:"Complemento" bson:"Complemento" json:"Complemento"`
	// Addenda           *Addenda               `xml:"Addenda" bson:"Addenda,omitempty" json:"Addenda,omitempty"`
}

type InformacionGlobal40 struct {
	Periodicidad string `xml:"Periodicidad,attr" bson:"Periodicidad" json:"Periodicidad"`
	Meses        string `xml:"Meses,attr" bson:"Meses" json:"Meses"`
	Anio         string `xml:"Año,attr" bson:"Anio" json:"Anio"`
}

type CfdiRelacionados40 struct {
	TipoRelacion      string              `xml:"TipoRelacion,attr" bson:"TipoRelacion" json:"TipoRelacion"`
	UuidsRelacionados []UuidRelacionado40 `xml:"CfdiRelacionado" bson:"UuidsRelacionados" json:"UuidsRelacionados"`
}

type UuidRelacionado40 struct {
	Uuid string `xml:"UUID,attr" bson:"Uuid" json:"Uuid"`
}

type Contribuyente40 struct {
	Rfc           string `json:"Rfc" bson:"Rfc"`
	Nombre        string `json:"Nombre" bson:"Nombre"`
	RegimenFiscal string `json:"RegimenFiscal" bson:"RegimenFiscal"`
}

type Emisor40 struct {
	Contribuyente40  `json:",inline" bson:",inline"`
	FacAtrAdquirente *string `xml:"FacAtrAdquirente,attr" bson:"FacAtrAdquirente,omitempty" json:"FacAtrAdquirente,omitempty"`
}

func (e *Emisor40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
		Nombre           string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
		RegimenFiscal    string  `xml:"RegimenFiscal,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
		FacAtrAdquirente *string `xml:"FacAtrAdquirente,attr" bson:"FacAtrAdquirente,omitempty" json:"FacAtrAdquirente,omitempty"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	e.Rfc = aux.Rfc
	e.Nombre = aux.Nombre
	e.RegimenFiscal = aux.RegimenFiscal
	e.FacAtrAdquirente = aux.FacAtrAdquirente
	return nil
}

type Receptor40 struct {
	Contribuyente40  `json:",inline" bson:",inline"`
	DomicilioFiscal  string  `xml:"DomicilioFiscalReceptor,attr" bson:"DomicilioFiscal" json:"DomicilioFiscal"`
	ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCFDI" json:"UsoCFDI"`
}

func (r *Receptor40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Rfc              string  `xml:"Rfc,attr" bson:"Rfc" json:"Rfc"`
		Nombre           string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
		DomicilioFiscal  string  `xml:"DomicilioFiscalReceptor,attr" bson:"DomicilioFiscal" json:"DomicilioFiscal"`
		ResidenciaFiscal *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
		NumRegIdTrib     *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
		RegimenFiscal    string  `xml:"RegimenFiscalReceptor,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
		UsoCFDI          string  `xml:"UsoCFDI,attr" bson:"UsoCFDI" json:"UsoCFDI"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	r.Rfc = aux.Rfc
	r.Nombre = aux.Nombre
	r.RegimenFiscal = aux.RegimenFiscal
	r.DomicilioFiscal = aux.DomicilioFiscal
	r.ResidenciaFiscal = aux.ResidenciaFiscal
	r.NumRegIdTrib = aux.NumRegIdTrib
	r.RegimenFiscal = aux.RegimenFiscal
	r.UsoCFDI = aux.UsoCFDI
	return nil
}

type Concepto40 struct {
	ClaveProductoServicio string                   `xml:"ClaveProdServ,attr" bson:"ClaveProductoServicio" json:"ClaveProductoServicio"`
	NoIdentificacion      *string                  `xml:"NoIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Cantidad              float64                  `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	ClaveUnidad           string                   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad                *string                  `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Descripcion           string                   `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario         float64                  `xml:"ValorUnitario,attr" bson:"ValorUnitario" json:"ValorUnitario"`
	Importe               float64                  `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	Descuento             *float64                 `xml:"Descuento,attr" bson:"Descuento,omitempty" json:"Descuento,omitempty"`
	ObjetoImpuesto        string                   `xml:"ObjetoImp,attr" bson:"ObjetoImpuesto" json:"ObjetoImpuesto"`
	Impuestos             *ImpuestosConcepto40     `xml:"Impuestos" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
	ACuentaTerceros       *ACuentaTerceros40       `xml:"ACuentaTerceros" bson:"ACuentaTerceros,omitempty" json:"ACuentaTerceros,omitempty"`
	InformacionAduanera   *[]InformacionAduanera40 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
	CuentaPredial         *[]CuentaPredial40       `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial,omitempty"`
	// ComplementoConcepto *ComplementoConcepto     `xml:"ComplementoConcepto" bson:"ComplementoConcepto,omitempty" json:"ComplementoConcepto,omitempty"`
	Parte *[]Parte40 `xml:"Parte" bson:"Parte,omitempty" json:"Parte,omitempty"`
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
	Rfc             string `xml:"RfcACuentaTerceros,attr" bson:"Rfc" json:"Rfc"`
	Nombre          string `xml:"NombreACuentaTerceros,attr" bson:"Nombre" json:"Nombre"`
	RegimenFiscal   string `xml:"RegimenFiscalACuentaTerceros,attr" bson:"RegimenFiscal" json:"RegimenFiscal"`
	DomicilioFiscal string `xml:"DomicilioFiscalACuentaTerceros,attr" bson:"DomicilioFiscal" json:"DomicilioFiscal"`
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

func (c *Comprobante40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Version           string                `xml:"Version,attr"`
		Serie             *string               `xml:"Serie,attr"`
		FechaRaw          string                `xml:"Fecha,attr"`
		Sello             string                `xml:"Sello,attr"`
		FormaPago         *string               `xml:"FormaPago,attr"`
		NoCertificado     string                `xml:"NoCertificado,attr"`
		Certificado       string                `xml:"Certificado,attr"`
		CondicionesPago   *string               `xml:"CondicionesDePago,attr"`
		Subtotal          float64               `xml:"SubTotal,attr"`
		Descuento         *float64              `xml:"Descuento,attr"`
		Moneda            string                `xml:"Moneda,attr"`
		TipoCambio        *float64              `xml:"TipoCambio,attr"`
		Total             float64               `xml:"Total,attr"`
		TipoComprobante   string                `xml:"TipoDeComprobante,attr"`
		Exportacion       string                `xml:"Exportacion,attr"`
		MetodoPago        *string               `xml:"MetodoPago,attr"`
		LugarExpedicion   string                `xml:"LugarExpedicion,attr"`
		Confirmacion      *string               `xml:"Confirmacion,attr"`
		InformacionGlobal *InformacionGlobal40  `xml:"InformacionGlobal"`
		CfdisRelacionados *[]CfdiRelacionados40 `xml:"CfdiRelacionados"`
		Emisor            Emisor40              `xml:"Emisor"`
		Receptor          Receptor40            `xml:"Receptor"`
		Conceptos         []Concepto40          `xml:"Conceptos>Concepto"`
		Impuestos         *Impuestos40          `xml:"Impuestos"`
		// Complemento       Complemento            `xml:"Complemento" `
		// Addenda           *Addenda               `xml:"Addenda" bson:"Addenda,`

	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	fechaEmision, err := helpers.ParseDatetime(aux.FechaRaw)
	if err != nil {
		return err
	}
	c.Fecha = fechaEmision
	return nil
}

func (ur *UuidRelacionado40) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		Uuid string `xml:"UUID,attr"`
	}

	var aux tmp

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	ur.Uuid = strings.ToUpper(aux.Uuid)
	return nil
}
