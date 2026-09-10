package complementos

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/helpers"
)

type Pagos20 struct {
	XMLName xml.Name  `xml:"http://www.sat.gob.mx/Pagos20 Pagos"`
	Version string    `xml:"Version,attr" bson:"Version" json:"Version"`
	Totales Totales20 `xml:"Totales" bson:"Totales" json:"Totales"`
	Pago    []Pago20  `xml:"Pago" bson:"Pago" json:"Pago"`
}

type Totales20 struct {
	TotalRetencionesIVA         *float64 `xml:"TotalRetencionesIVA,attr" bson:"TotalRetencionesIVA,omitempty" json:"TotalRetencionesIVA,omitempty"`
	TotalRetencionesISR         *float64 `xml:"TotalRetencionesISR,attr" bson:"TotalRetencionesISR,omitempty" json:"TotalRetencionesISR,omitempty"`
	TotalRetencionesIEPS        *float64 `xml:"TotalRetencionesIEPS,attr" bson:"TotalRetencionesIEPS,omitempty" json:"TotalRetencionesIEPS,omitempty"`
	TotalTrasladosBaseIVA16     *float64 `xml:"TotalTrasladosBaseIVA16,attr" bson:"TotalTrasladosBaseIVA16,omitempty" json:"TotalTrasladosBaseIVA16,omitempty"`
	TotalTrasladosImpuestoIVA16 *float64 `xml:"TotalTrasladosImpuestoIVA16,attr" bson:"TotalTrasladosImpuestoIVA16,omitempty" json:"TotalTrasladosImpuestoIVA16,omitempty"`
	TotalTrasladosBaseIVA8      *float64 `xml:"TotalTrasladosBaseIVA8,attr" bson:"TotalTrasladosBaseIVA8,omitempty" json:"TotalTrasladosBaseIVA8,omitempty"`
	TotalTrasladosImpuestoIVA8  *float64 `xml:"TotalTrasladosImpuestoIVA8,attr" bson:"TotalTrasladosImpuestoIVA8,omitempty" json:"TotalTrasladosImpuestoIVA8,omitempty"`
	TotalTrasladosBaseIVA0      *float64 `xml:"TotalTrasladosBaseIVA0,attr" bson:"TotalTrasladosBaseIVA0,omitempty" json:"TotalTrasladosBaseIVA0,omitempty"`
	TotalTrasladosImpuestoIVA0  *float64 `xml:"TotalTrasladosImpuestoIVA0,attr" bson:"TotalTrasladosImpuestoIVA0,omitempty" json:"TotalTrasladosImpuestoIVA0,omitempty"`
	TotalTrasladosBaseIVAExento *float64 `xml:"TotalTrasladosBaseIVAExento,attr" bson:"TotalTrasladosBaseIVAExento,omitempty" json:"TotalTrasladosBaseIVAExento,omitempty"`
	MontoTotalPagos             float64  `xml:"MontoTotalPagos,attr" bson:"MontoTotalPagos" json:"MontoTotalPagos"`
}

type Pago20 struct {
	FechaPago                      time.Time            `xml:"FechaPago,attr" bson:"FechaPago" json:"FechaPago"`
	FormaDePagoP                   string               `xml:"FormaDePagoP,attr" bson:"FormaDePagoP" json:"FormaDePagoP"`
	MonedaP                        string               `xml:"MonedaP,attr" bson:"MonedaP" json:"MonedaP"`
	TipoCambioP                    *float64             `xml:"TipoCambioP,attr" bson:"TipoCambioP,omitempty" json:"TipoCambioP,omitempty"`
	Monto                          float64              `xml:"Monto,attr" bson:"Monto" json:"Monto"`
	NumOperacion                   *string              `xml:"NumOperacion,attr" bson:"NumOperacion,omitempty" json:"NumOperacion,omitempty"`
	NomBancoOrdExt                 *string              `xml:"RfcEmisorCtaOrd,attr" bson:"RfcEmisorCtaOrd,omitempty" json:"RfcEmisorCtaOrd,omitempty"`
	NombreBancoOrdenanteExtranjero *string              `xml:"NomBancoOrdExt,attr" bson:"NomBancoOrdExt,omitempty" json:"NomBancoOrdExt,omitempty"`
	CtaOrdenante                   *string              `xml:"CtaOrdenante,attr" bson:"CtaOrdenante,omitempty" json:"CtaOrdenante,omitempty"`
	RfcEmisorCtaBen                *string              `xml:"RfcEmisorCtaBen,attr" bson:"RfcEmisorCtaBen,omitempty" json:"RfcEmisorCtaBen,omitempty"`
	CtaBeneficiario                *string              `xml:"CtaBeneficiario,attr" bson:"CtaBeneficiario,omitempty" json:"CtaBeneficiario,omitempty"`
	TipoCadPago                    *string              `xml:"TipoCadPago,attr" bson:"TipoCadPago,omitempty" json:"TipoCadPago,omitempty"`
	CertPago                       *string              `xml:"CertPago,attr" bson:"CertPago,omitempty" json:"CertPago,omitempty"`
	CadPago                        *string              `xml:"CadPago,attr" bson:"CadPago,omitempty" json:"CadPago,omitempty"`
	SelloPago                      *string              `xml:"SelloPago,attr" bson:"SelloPago,omitempty" json:"SelloPago,omitempty"`
	DoctoRelacionado               []DoctoRelacionado20 `xml:"DoctoRelacionado" bson:"DoctoRelacionado" json:"DoctoRelacionado"`
	ImpuestosP                     *ImpuestosDR20       `xml:"ImpuestosP" bson:"ImpuestosP,omitempty" json:"ImpuestosP,omitempty"`
}

func (p *Pago20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		FechaPago string `xml:"FechaPago,attr"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	fechaPago, err := helpers.ParseDatetime(aux.FechaPago)
	if err != nil {
		return err
	}
	p.FechaPago = fechaPago
	return nil
}

type DoctoRelacionado20 struct {
	IdDocumento      string         `xml:"IdDocumento,attr" bson:"IdDocumento" json:"IdDocumento"`
	Serie            *string        `xml:"Serie,attr" bson:"Serie,omitempty" json:"Serie,omitempty"`
	Folio            *string        `xml:"Folio,attr" bson:"Folio,omitempty" json:"Folio,omitempty"`
	MonedaDR         string         `xml:"MonedaDR,attr" bson:"MonedaDR" json:"MonedaDR"`
	EquivalenciaDR   *float64       `xml:"EquivalenciaDR,attr" bson:"EquivalenciaDR,omitempty" json:"EquivalenciaDR,omitempty"`
	NumParcialidad   float64        `xml:"NumParcialidad,attr" bson:"NumParcialidad" json:"NumParcialidad"`
	ImpSaldoAnt      float64        `xml:"ImpSaldoAnt,attr" bson:"ImpSaldoAnt" json:"ImpSaldoAnt"`
	ImpPagado        float64        `xml:"ImpPagado,attr" bson:"ImpPagado" json:"ImpPagado"`
	ImpSaldoInsoluto float64        `xml:"ImpSaldoInsoluto,attr" bson:"ImpSaldoInsoluto" json:"ImpSaldoInsoluto"`
	ObjetoImpDR      string         `xml:"ObjetoImpDR,attr" bson:"ObjetoImpDR" json:"ObjetoImpDR"`
	ImpuestosDR      *ImpuestosDR20 `xml:"ImpuestosDR" bson:"ImpuestosDR,omitempty" json:"ImpuestosDR,omitempty"`
}

func (dr *DoctoRelacionado20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmp struct {
		IdDocumento string `xml:"IdDocumento,attr"`
	}
	var aux tmp
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	dr.IdDocumento = strings.ToUpper(aux.IdDocumento)
	return nil
}

type ImpuestosDR20 struct {
	RetencionesDR *[]RetencionesDR20 `xml:"RetencionesDR>RetencionDR" bson:"RetencionesDR,omitempty" json:"RetencionesDR,omitempty"`
	TrasladosDR   *[]TrasladosDR20   `xml:"TrasladosDR>TrasladoDR" bson:"TrasladosDR,omitempty" json:"TrasladosDR,omitempty"`
}

type RetencionesDR20 struct {
	Base       float64 `xml:"BaseDR,attr" bson:"Base" json:"Base"`
	Impuesto   string  `xml:"ImpuestoDR,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string  `xml:"TipoFactorDR,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota float64 `xml:"TasaOCuotaDR,attr" bson:"TasaOCuota" json:"TasaOCuota"`
	Importe    float64 `xml:"ImporteDR,attr" bson:"Importe" json:"Importe"`
}

type TrasladosDR20 struct {
	Base       float64  `xml:"BaseDR,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"ImpuestoDR,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactorDR,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuotaDR,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"ImporteDR,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
}

type ImpuestosPagos20 struct {
	RetencionesP *[]RetencionesP20 `xml:"RetencionesP>RetencionP" bson:"RetencionesP,omitempty" json:"RetencionesP,omitempty"`
	TrasladosP   *[]TrasladosP20   `xml:"TrasladosP>TrasladoP" bson:"TrasladosP,omitempty" json:"TrasladosP,omitempty"`
}

type RetencionesP20 struct {
	Impuesto string  `xml:"ImpuestoP,attr" bson:"Impuesto" json:"Impuesto"`
	Importe  float64 `xml:"ImporteP,attr" bson:"Importe" json:"Importe"`
}

type TrasladosP20 struct {
	Base       float64  `xml:"BaseP,attr" bson:"Base" json:"Base"`
	Impuesto   string   `xml:"ImpuestoP,attr" bson:"Impuesto" json:"Impuesto"`
	TipoFactor string   `xml:"TipoFactorP,attr" bson:"TipoFactor" json:"TipoFactor"`
	TasaOCuota *float64 `xml:"TasaOCuotaP,attr" bson:"TasaOCuota,omitempty" json:"TasaOCuota,omitempty"`
	Importe    *float64 `xml:"ImporteP,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
}
