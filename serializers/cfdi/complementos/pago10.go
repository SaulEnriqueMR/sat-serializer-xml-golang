package complementos

import (
	"encoding/xml"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/datatypes"
)

type Pagos10 struct {
	XMLName xml.Name `xml:"http://www.sat.gob.mx/Pagos Pagos"`
	Version string   `xml:"Version,attr" bson:"Version" json:"Version"`
}

type Pago10 struct {
	FechaPago          datatypes.ISODateTime `xml:"FechaPago,attr" bson:"FechaPago" json:"FechaPago"`
	FormaDePagoP       string                `xml:"FormaDePagoP,attr" bson:"FormaDePagoP" json:"FormaDePagoP"`
	MonedaP            string                `xml:"MonedaP,attr" bson:"MonedaP" json:"MonedaP"`
	TipoCambioP        *float64              `xml:"TipoCambioP,attr" bson:"TipoCambioP,omitempty" json:"TipoCambioP,omitempty"`
	Monto              float64               `xml:"Monto,attr" bson:"Monto" json:"Monto"`
	NumOperacion       *string               `xml:"NumOperacion,attr" bson:"NumOperacion,omitempty" json:"NumOperacion,omitempty"`
	RfcEmisorCtaOrd    *string               `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCtaOrd,omitempty" json:"RfcEmisorCtaOrd,omitempty"`
	NomBancoOrdExt     *string               `xml:"NomBancoOrdExt,attr" bson:"NomBancoOrdExt,omitempty" json:"NomBancoOrdExt,omitempty"`
	CtaOrdenante       *string               `xml:"CtaOrdenante,attr" bson:"CtaOrdenante,omitempty" json:"CtaOrdenante,omitempty"`
	RfcEmisorCtaBen    *string               `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCtaBen,omitempty" json:"RfcEmisorCtaBen,omitempty"`
	CtaBeneficiario    *string               `xml:"CtaBeneficiario,attr" bson:"CtaBeneficiario,omitempty" json:"CtaBeneficiario,omitempty"`
	TipoCadPago        *string               `xml:"TipoCadPago,attr" bson:"TipoCadPago,omitempty" json:"TipoCadPago,omitempty"`
	CertPago           *string               `xml:"CertPago,attr" bson:"CertPago,omitempty" json:"CertPago,omitempty"`
	CadPago            *string               `xml:"CadPago,attr" bson:"CadPago,omitempty" json:"CadPago,omitempty"`
	SelloPago          *string               `xml:"SelloPago,attr" bson:"SelloPago,omitempty" json:"SelloPago,omitempty"`
	DoctoRelacionado10 *[]DoctoRelacionado10 `xml:"DoctoRelacionado" bson:"DoctoRelacionado,omitempty" json:"DoctoRelacionado,omitempty"`
}

type DoctoRelacionado10 struct {
	IdDocumento      datatypes.Uuid `xml:"IdDocumento,attr" bson:"IdDocumento" json:"IdDocumento"`
	Serie            *string        `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie,omitempty"`
	Folio            *string        `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	MonedaDR         string         `xml:"MonedaDR,attr" bson:"MonedaDR" json:"MonedaDR"`
	TipoCambioDR     *float64       `xml:"TipoCambioDR,attr" bson:"TipoCambioDR,omitempty" json:"TipoCambioDR,omitempty"`
	MetodoDePagoDR   string         `xml:"MetodoDePagoDR,attr" bson:"MetodoDePagoDR" json:"MetodoDePagoDR"`
	NumParcialidad   float64        `xml:"NumParcialidad,attr" bson:"NumParcialidad" json:"NumParcialidad"`
	ImpSaldoAnt      float64        `xml:"ImpSaldoAnt,attr" bson:"ImpSaldoAnt" json:"ImpSaldoAnt"`
	ImpPagado        float64        `xml:"ImpPagado,attr" bson:"ImpPagado" json:"ImpPagado"`
	ImpSaldoInsoluto float64        `xml:"ImpSaldoInsoluto,attr" bson:"ImpSaldoInsoluto" json:"ImpSaldoInsoluto"`
	Impuestos        *ImpuestosP10  `xml:"Impuestos,attr" bson:"Impuestos,omitempty" json:"Impuestos,omitempty"`
}

type ImpuestosP10 struct {
	TotalImpuestosRetenidos   *float64        `xml:"TotalImpuestosRetenidos" bson:"TotalImpuestosRetenidos" json:"TotalImpuestosRetenidos"`
	TotalImpuestosTrasladados *float64        `xml:"TotalImpuestosTrasladados" bson:"TotalImpuestosTrasladados" json:"TotalImpuestosTrasladados"`
	Retenciones               *[]RetencionP10 `xml:"Retenciones>Retencion" bson:"Retenciones" json:"Retenciones"`
	Traslados                 *[]TrasladoP10  `xml:"Traslados>Traslado" bson:"Traslados" json:"Traslados"`
}

type RetencionP10 struct {
	Impuesto string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	Importe  float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type TrasladoP10 struct {
	Impuesto   string  `xml:"Impuesto,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactor,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuota,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}
