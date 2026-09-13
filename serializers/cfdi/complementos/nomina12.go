package complementos

import (
	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/datatypes"
)

type Nomina12 struct {
	Version           string                 `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoNomina        string                 `xml:"TipoNomina,attr" bson:"TipoNomina" json:"TipoNomina"`
	FechaPago         datatypes.ISODateTime  `xml:"FechaPago,attr" bson:"FechaPagoString" json:"FechaPagoString"`
	FechaInicialPago  datatypes.ISODateTime  `xml:"FechaInicialPago,attr" bson:"FechaInicialPagoString" json:"FechaInicialPagoString"`
	FechaFinalPago    datatypes.ISODateTime  `xml:"FechaFinalPago,attr" bson:"FechaFinalPagoString" json:"FechaFinalPagoString"`
	NumDiasPagados    float64                `xml:"NumDiasPagados,attr" bson:"NumDiasPagados" json:"NumDiasPagados"`
	TotalPercepciones *float64               `xml:"TotalPercepciones,attr" bson:"TotalPercepciones,omitempty" json:"TotalPercepciones,omitempty"`
	TotalDeducciones  *float64               `xml:"TotalDeducciones,attr" bson:"TotalDeducciones,omitempty" json:"TotalDeducciones,omitempty"`
	TotalOtrosPagos   *float64               `xml:"TotalOtrosPagos,attr" bson:"TotalOtrosPagos,omitempty" json:"TotalOtrosPagos,omitempty"`
	Emisor            *Emisor12              `xml:"Emisor" bson:"Emisor,omitempty" json:"Emisor,omitempty"`
	Receptor          Receptor12             `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Percepciones      *Percepciones12        `xml:"Percepciones" bson:"Percepciones,omitempty" json:"Percepciones,omitempty"`
	Deducciones       *Deducciones12         `xml:"Deducciones" bson:"Deducciones,omitempty" json:"Deducciones,omitempty"`
	OtrosPagos        *[]OtroPagoNomina12    `xml:"OtrosPagos>OtroPago" bson:"OtrosPagos,omitempty" json:"OtrosPagos,omitempty"`
	Incapacidades     *[]IncapacidadNomina12 `xml:"Incapacidades>Incapacidad" bson:"Incapacidades,omitempty" json:"Incapacidades,omitempty"`
}

type Emisor12 struct {
	Curp             *string        `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`
	RegistroPatronal *string        `xml:"RegistroPatronal,attr" bson:"RegistroPatronal,omitempty" json:"RegistroPatronal,omitempty"`
	RfcPatronOrigen  *string        `xml:"RfcPatronOrigen,attr" bson:"RfcPatronOrigen,omitempty" json:"RfcPatronOrigen,omitempty"`
	EntidadSncf      *EntidadSNCF12 `xml:"EntidadSNCF" bson:"EntidadSncf,omitempty" json:"EntidadSncf,omitempty"`
}

type EntidadSNCF12 struct {
	OrigenRecurso      string   `xml:"OrigenRecurso,attr" bson:"OrigenRecurso" json:"OrigenRecurso"`
	MontoRecursoPropio *float64 `xml:"MontoRecursoPropio,attr" bson:"MontoRecursoPropio,omitempty" json:"MontoRecursoPropio,omitempty"`
}

type Receptor12 struct {
	Curp                   string                 `xml:"Curp,attr" bson:"Curp" json:"Curp"`
	NumSeguridadSocial     *string                `xml:"NumSeguridadSocial,attr" bson:"NumSeguridadSocial,omitempty" json:"NumSeguridadSocial,omitempty"`
	FechaInicioRelLaboral  *datatypes.ISODateTime `xml:"FechaInicioRelLaboral,attr" bson:"FechaInicioRelLaboral,omitempty" json:"FechaInicioRelLaboral,omitempty"`
	Antigüedad             *string                `xml:"Antigüedad,attr" bson:"Antigüedad,omitempty" json:"Antigüedad,omitempty"`
	TipoContrato           string                 `xml:"TipoContrato,attr" bson:"TipoContrato" json:"TipoContrato"`
	Sindicalizado          *string                `xml:"Sindicalizado,attr" bson:"Sindicalizado,omitempty" json:"Sindicalizado,omitempty"`
	TipoJornada            *string                `xml:"TipoJornada,attr" bson:"TipoJornada,omitempty" json:"TipoJornada,omitempty"`
	TipoRegimen            string                 `xml:"TipoRegimen,attr" bson:"TipoRegimen" json:"TipoRegimen"`
	NumEmpleado            string                 `xml:"NumEmpleado,attr" bson:"NumEmpleado" json:"NumEmpleado"`
	Departamento           *string                `xml:"Departamento,attr" bson:"Departamento,omitempty" json:"Departamento,omitempty"`
	Puesto                 *string                `xml:"Puesto,attr" bson:"Puesto,omitempty" json:"Puesto,omitempty"`
	RiesgoPuesto           *string                `xml:"RiesgoPuesto,attr" bson:"RiesgoPuesto,omitempty" json:"RiesgoPuesto,omitempty"`
	PeriodicidadPago       string                 `xml:"PeriodicidadPago,attr" bson:"PeriodicidadPago" json:"PeriodicidadPago"`
	Banco                  *string                `xml:"Banco,attr" bson:"Banco,omitempty" json:"Banco,omitempty"`
	CuentaBancaria         *string                `xml:"CuentaBancaria,attr" bson:"CuentaBancaria,omitempty" json:"CuentaBancaria,omitempty"`
	SalarioBaseCotApor     *float64               `xml:"SalarioBaseCotApor,attr" bson:"SalarioBaseCotApor,omitempty" json:"SalarioBaseCotApor,omitempty"`
	SalarioDiarioIntegrado *float64               `xml:"SalarioDiarioIntegrado,attr" bson:"SalarioDiarioIntegrado,omitempty" json:"SalarioDiarioIntegrado,omitempty"`
	ClaveEntFed            string                 `xml:"ClaveEntFed,attr" bson:"ClaveEntFed" json:"ClaveEntFed"`
	SubContratacion        *[]SubContratacion12   `xml:"SubContratacion" bson:"SubContratacion,omitempty" json:"SubContratacion,omitempty"`
}

type SubContratacion12 struct {
	RfcLabora        string  `xml:"RfcLabora,attr" bson:"RfcLabora" json:"RfcLabora"`
	PorcentajeTiempo float64 `xml:"PorcentajeTiempo,attr" bson:"PorcentajeTiempo" json:"PorcentajeTiempo"`
}

type Percepciones12 struct {
	TotalSueldos                 *float64                   `xml:"TotalSueldos,attr" bson:"TotalSueldos,omitempty" json:"TotalSueldos,omitempty"`
	TotalSeparacionIndemnizacion *float64                   `xml:"TotalSeparacionIndemnizacion,attr" bson:"TotalSeparacionIndemnizacion,omitempty" json:"TotalSeparacionIndemnizacion,omitempty"`
	TotalJubilacionPensionRetiro *float64                   `xml:"TotalJubilacionPensionRetiro,attr" bson:"TotalJubilacionPensionRetiro,omitempty" json:"TotalJubilacionPensionRetiro,omitempty"`
	TotalGravado                 float64                    `xml:"TotalGravado,attr" bson:"TotalGravado" json:"TotalGravado"`
	TotalExento                  float64                    `xml:"TotalExento,attr" bson:"TotalExento" json:"TotalExento"`
	Percepcion                   []Percepcion12             `xml:"Percepcion" bson:"Percepcion" json:"Percepcion"`
	JubilacionPensionRetiro      *JubilacionPensionRetiro12 `xml:"JubilacionPensionRetiro" bson:"JubilacionPensionRetiro,omitempty" json:"JubilacionPensionRetiro,omitempty"`
	SeparacionIndemnizacion      *SeparacionIndemnizacion12 `xml:"SeparacionIndemnizacion" bson:"SeparacionIndemnizacion,omitempty" json:"SeparacionIndemnizacion,omitempty"`
}

type Percepcion12 struct {
	TipoPercepcion   string              `xml:"TipoPercepcion,attr" bson:"TipoPercepcion" json:"TipoPercepcion"`
	Clave            string              `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto         string              `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	ImporteGravado   float64             `xml:"ImporteGravado,attr" bson:"ImporteGravado" json:"ImporteGravado"`
	ImporteExento    float64             `xml:"ImporteExento,attr" bson:"ImporteExento" json:"ImporteExento"`
	AccionesOTitulos *AccionesOTitulos12 `xml:"AccionesOTitulos" bson:"AccionesOTitulos,omitempty" json:"AccionesOTitulos,omitempty"`
	HorasExtra       *[]HorasExtra12     `xml:"HorasExtra" bson:"HorasExtra,omitempty" json:"HorasExtra,omitempty"`
}

type AccionesOTitulos12 struct {
	ValorMercado      float64 `xml:"ValorMercado,attr" bson:"ValorMercado" json:"ValorMercado"`
	PrecioAlOtorgarse float64 `xml:"PrecioAlOtorgarse,attr" bson:"PrecioAlOtorgarse" json:"PrecioAlOtorgarse"`
}

type HorasExtra12 struct {
	Dias          int     `xml:"Dias,attr" bson:"Dias" json:"Dias"`
	TipoHoras     string  `xml:"TipoHoras,attr" bson:"TipoHoras" json:"TipoHoras"`
	HorasExtra    int     `xml:"HorasExtra,attr" bson:"HorasExtra" json:"HorasExtra"`
	ImportePagado float64 `xml:"ImportePagado,attr" bson:"ImportePagado" json:"ImportePagado"`
}

type JubilacionPensionRetiro12 struct {
	TotalUnaExhibicion  *float64 `xml:"TotalUnaExhibicion,attr" bson:"TotalUnaExhibicion,omitempty" json:"TotalUnaExhibicion,omitempty"`
	TotalParcialidad    *float64 `xml:"TotalParcialidad,attr" bson:"TotalParcialidad,omitempty" json:"TotalParcialidad,omitempty"`
	MontoDiario         *float64 `xml:"MontoDiario,attr" bson:"MontoDiario,omitempty" json:"MontoDiario,omitempty"`
	IngresoAcumulable   float64  `xml:"IngresoAcumulable,attr" bson:"IngresoAcumulable" json:"IngresoAcumulable"`
	IngresoNoAcumulable float64  `xml:"IngresoNoAcumulable,attr" bson:"IngresoNoAcumulable" json:"IngresoNoAcumulable"`
}

type SeparacionIndemnizacion12 struct {
	TotalPagado         float64 `xml:"TotalPagado,attr" bson:"TotalPagado" json:"TotalPagado"`
	NumAñosServicio     int     `xml:"NumAñosServicio,attr" bson:"NumAñosServicio" json:"NumAñosServicio"`
	UltimoSueldoMensOrd float64 `xml:"UltimoSueldoMensOrd,attr" bson:"UltimoSueldoMensOrd" json:"UltimoSueldoMensOrd"`
	IngresoAcumulable   float64 `xml:"IngresoAcumulable,attr" bson:"IngresoAcumulable" json:"IngresoAcumulable"`
	IngresoNoAcumulable float64 `xml:"IngresoNoAcumulable,attr" bson:"IngresoNoAcumulable" json:"IngresoNoAcumulable"`
}

type Deducciones12 struct {
	TotalOtrasDeducciones   *float64      `xml:"TotalOtrasDeducciones,attr" bson:"TotalOtrasDeducciones,omitempty" json:"TotalOtrasDeducciones,omitempty"`
	TotalImpuestosRetenidos *float64      `xml:"TotalImpuestosRetenidos,attr" bson:"TotalImpuestosRetenidos,omitempty" json:"TotalImpuestosRetenidos,omitempty"`
	Deduccion               []Deduccion12 `xml:"Deduccion" bson:"Deduccion" json:"Deduccion"`
}

type Deduccion12 struct {
	TipoDeduccion string  `xml:"TipoDeduccion,attr" bson:"TipoDeduccion" json:"TipoDeduccion"`
	Clave         string  `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto      string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	Importe       float64 `xml:"Importe,attr" bson:"Importe" json:"Importe"`
}

type OtroPagoNomina12 struct {
	TipoOtroPago             string                            `xml:"TipoOtroPago,attr" bson:"TipoOtroPago" json:"TipoOtroPago"`
	Clave                    string                            `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto                 string                            `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	Importe                  float64                           `xml:"Importe,attr" bson:"Importe" json:"Importe"`
	SubsidioAlEmpleo         *SubsidioAlEmpleoNomina12         `xml:"SubsidioAlEmpleo" bson:"SubsidioAlEmpleo,omitempty" json:"SubsidioAlEmpleo,omitempty"`
	CompensacionSaldosAFavor *CompensacionSaldosAFavorNomina12 `xml:"CompensacionSaldosAFavor" bson:"CompensacionSaldosAFavor,omitempty" json:"CompensacionSaldosAFavor,omitempty"`
}

type SubsidioAlEmpleoNomina12 struct {
	SubsidioCausado float64 `xml:"SubsidioCausado,attr" bson:"SubsidioCausado" json:"SubsidioCausado"`
}

type CompensacionSaldosAFavorNomina12 struct {
	SaldoAFavor     float64 `xml:"SaldoAFavor,attr" bson:"SaldoAFavor" json:"SaldoAFavor"`
	Año             string  `xml:"Año,attr" bson:"Año" json:"Año"`
	RemanenteSalFav float64 `xml:"RemanenteSalFav,attr" bson:"RemanenteSalFav" json:"RemanenteSalFav"`
}

type IncapacidadNomina12 struct {
	DiasIncapacidad  int      `xml:"DiasIncapacidad,attr" bson:"DiasIncapacidad" json:"DiasIncapacidad"`
	TipoIncapacidad  string   `xml:"TipoIncapacidad,attr" bson:"TipoIncapacidad" json:"TipoIncapacidad"`
	ImporteMonetario *float64 `xml:"ImporteMonetario,attr" bson:"ImporteMonetario,omitempty" json:"ImporteMonetario,omitempty"`
}
