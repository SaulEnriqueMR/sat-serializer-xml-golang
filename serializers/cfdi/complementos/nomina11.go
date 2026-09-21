package complementos

import (
	"encoding/xml"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/datatypes"
)

type Nomina11 struct {
	XMLName                xml.Name               `xml:"http://www.sat.gob.mx/nomina Nomina"`
	Version                string                 `xml:"Version,attr" bson:"Version" json:"Version"`
	RegistroPatronal       *string                `xml:"RegistroPatronal,attr" bson:"RegistroPatronal,omitempty" json:"RegistroPatronal,omitempty"`
	NumEmpleado            string                 `xml:"NumEmpleado,attr" bson:"NumEmpleado" json:"NumEmpleado"`
	CURP                   string                 `xml:"CURP,attr" bson:"CURP" json:"CURP"`
	TipoRegimen            string                 `xml:"TipoRegimen,attr" bson:"TipoRegimen" json:"TipoRegimen"`
	NumSeguridadSocial     *string                `xml:"NumSeguridadSocial,attr" bson:"NumSeguridadSocial,omitempty" json:"NumSeguridadSocial,omitempty"`
	FechaPago              datatypes.ISODateTime  `xml:"FechaPago,attr" bson:"FechaPagoString" json:"FechaPagoString"`
	FechaInicialPago       datatypes.ISODateTime  `xml:"FechaInicialPago,attr" bson:"FechaInicialPagoString" json:"FechaInicialPagoString"`
	FechaFinalPago         datatypes.ISODateTime  `xml:"FechaFinalPago,attr" bson:"FechaFinalPagoString" json:"FechaFinalPagoString"`
	NumDiasPagados         float64                `xml:"NumDiasPagados,attr" bson:"NumDiasPagados" json:"NumDiasPagados"`
	Departamento           *string                `xml:"Departamento,attr" bson:"Departamento,omitempty" json:"Departamento,omitempty"`
	CLABE                  *string                `xml:"CLABE,attr" bson:"CLABE" json:"CLABE"`
	Banco                  *int64                 `xml:"Banco,attr" bson:"Banco,omitempty" json:"Banco,omitempty"`
	FechaInicioRelLaboral  *datatypes.ISODateTime `xml:"FechaInicioRelLaboral,attr" bson:"FechaInicioRelLaboral,omitempty" json:"FechaInicioRelLaboral,omitempty"`
	Antiguedad             *string                `xml:"Antiguedad,attr" bson:"Antiguedad,omitempty" json:"Antiguedad,omitempty"`
	Puesto                 *string                `xml:"Puesto,attr" bson:"Puesto,omitempty" json:"Puesto,omitempty"`
	TipoContrato           *string                `xml:"TipoContrato,attr" bson:"TipoContrato" json:"TipoContrato"`
	TipoJornada            *string                `xml:"TipoJornada,attr" bson:"TipoJornada,omitempty" json:"TipoJornada,omitempty"`
	PeriodicidadPago       string                 `xml:"PeriodicidadPago,attr" bson:"PeriodicidadPago" json:"PeriodicidadPago"`
	SalarioBaseCotApor     *float64               `xml:"SalarioBaseCotApor,attr" bson:"SalarioBaseCotApor,omitempty" json:"SalarioBaseCotApor,omitempty"`
	RiesgoPuesto           *string                `xml:"RiesgoPuesto,attr" bson:"RiesgoPuesto,omitempty" json:"RiesgoPuesto,omitempty"`
	SalarioDiarioIntegrado *float64               `xml:"SalarioDiarioIntegrado,attr" bson:"SalarioDiarioIntegrado,omitempty" json:"SalarioDiarioIntegrado,omitempty"`
	Percepciones           *Percepciones11        `xml:"Percepciones" bson:"Percepciones,omitempty" json:"Percepciones,omitempty"`
	Deducciones            *Deducciones11         `xml:"Deducciones" bson:"Deducciones,omitempty" json:"Deducciones,omitempty"`
	Incapacidades          *[]Incapacidad11       `xml:"Incapacidades>Incapacidad" bson:"Incapacidades" json:"Incapacidades"`
	HorasExtras            *[]HorasExtra11        `xml:"HorasExtras>HorasExtra" bson:"HorasExtras" json:"HorasExtras"`
}

type Percepciones11 struct {
	TotalGravado float64        `xml:"TotalGravado,attr" bson:"TotalGravado" json:"TotalGravado"`
	TotalExento  float64        `xml:"TotalExento,attr" bson:"TotalExento" json:"TotalExento"`
	Percepcion   []Percepcion11 `xml:"Percepcion" bson:"Percepcion" json:"Percepcion"`
}

type Percepcion11 struct {
	TipoPercepcion int64   `xml:"TipoPercepcion,attr" bson:"TipoPercepcion" json:"TipoPercepcion"`
	Clave          string  `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado" json:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento" json:"ImporteExento"`
}

type Deducciones11 struct {
	TotalGravado float64       `xml:"TotalGravado,attr" bson:"TotalGravado" json:"TotalGravado"`
	TotalExento  float64       `xml:"TotalExento,attr" bson:"TotalExento" json:"TotalExento"`
	Deduccion    []Deduccion11 `xml:"Deduccion" bson:"Deduccion" json:"Deduccion"`
}

type Deduccion11 struct {
	TipoDeduccion  string  `xml:"TipoDeduccion,attr" bson:"TipoDeduccion" json:"TipoDeduccion"`
	Clave          string  `xml:"Clave,attr" bson:"Clave" json:"Clave"`
	Concepto       string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	ImporteGravado float64 `xml:"ImporteGravado,attr" bson:"ImporteGravado" json:"ImporteGravado"`
	ImporteExento  float64 `xml:"ImporteExento,attr" bson:"ImporteExento" json:"ImporteExento"`
}

type Incapacidad11 struct {
	DiasIncapacidad  int     `xml:"DiasIncapacidad,attr" bson:"DiasIncapacidad" json:"DiasIncapacidad"`
	TipoIncapacidad  string  `xml:"TipoIncapacidad,attr" bson:"TipoIncapacidad" json:"TipoIncapacidad"`
	ImporteMonetario float64 `xml:"ImporteMonetario,attr" bson:"ImporteMonetario" json:"ImporteMonetario"`
}

type HorasExtra11 struct {
	Dias          int     `xml:"Dias,attr" bson:"Dias" json:"Dias"`
	TipoHoras     string  `xml:"TipoHoras,attr" bson:"TipoHoras" json:"TipoHoras"`
	HorasExtra    int     `xml:"HorasExtra,attr" bson:"HorasExtra" json:"HorasExtra"`
	ImportePagado float64 `xml:"ImportePagado,attr" bson:"ImportePagado" json:"ImportePagado"`
}
