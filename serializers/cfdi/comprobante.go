package cfdi

import (
	"encoding/xml"
	"errors"
	"fmt"
)

const (
	NsCfd2 = "http://www.sat.gob.mx/cfd/2"
	NsCfd3 = "http://www.sat.gob.mx/cfd/3"
	NsCfd4 = "http://www.sat.gob.mx/cfd/4"
	// 1.0 no tiene targetNamespace
)

type Comprobante interface {
	GetVersion() string
}

func (c *Comprobante10) GetVersion() string { return c.Version }
func (c *Comprobante20) GetVersion() string { return c.Version }
func (c *Comprobante22) GetVersion() string { return c.Version }
func (c *Comprobante30) GetVersion() string { return c.Version }
func (c *Comprobante32) GetVersion() string { return c.Version }
func (c *Comprobante33) GetVersion() string { return c.Version }
func (c *Comprobante40) GetVersion() string { return c.Version }

var (
	ErrVersionDesconocida     = errors.New("versión de Comprobante no soportada")
	ErrVersionAusente         = errors.New("atributo Version ausente en Comprobante")
	ErrNamespaceInconsistente = errors.New("namespace no coincide con la versión declarada")
)

// entrada del registro
type comprobanteResolver struct {
	version   string
	namespace string
	decode    func([]byte) (Comprobante, error)
}

var comprobanteRegistry = []comprobanteResolver{
	{"1.0", "", decodeAs[*Comprobante10]},
	{"2.0", NsCfd2, decodeAs[*Comprobante20]},
	{"2.2", NsCfd2, decodeAs[*Comprobante22]},
	{"3.0", NsCfd3, decodeAs[*Comprobante30]},
	{"3.2", NsCfd3, decodeAs[*Comprobante32]},
	{"3.3", NsCfd3, decodeAs[*Comprobante33]},
	{"4.0", NsCfd4, decodeAs[*Comprobante40]},
}

func UnmarshalComprobante(data []byte) (Comprobante, error) {
	// --- Paso 1: sonda ligera para leer namespace y versión ---
	var probe struct {
		XMLName      xml.Name
		VersionLower string `xml:"version,attr"`
		VersionUpper string `xml:"Version,attr"`
	}
	if err := xml.Unmarshal(data, &probe); err != nil {
		return nil, fmt.Errorf("error leyendo cabecera XML: %w", err)
	}

	// --- Paso 2: normalizar el nombre del atributo ---
	// 1.0–3.2 usan "version"; 3.3+ usan "Version".
	version := probe.VersionLower
	if version == "" {
		version = probe.VersionUpper
	}
	if version == "" {
		return nil, ErrVersionAusente
	}

	// --- Paso 3: despachar por versión con validación cruzada de namespace ---
	for _, r := range comprobanteRegistry {
		if r.version != version {
			continue
		}
		if r.namespace != probe.XMLName.Space {
			return nil, fmt.Errorf("%w: versión %s espera ns %q, llegó %q",
				ErrNamespaceInconsistente, version, r.namespace, probe.XMLName.Space)
		}
		return r.decode(data)
	}

	return nil, fmt.Errorf("%w: %q", ErrVersionDesconocida, version)
}

func decodeAs[T Comprobante](data []byte) (Comprobante, error) {
	var t T
	if err := xml.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	return t, nil
}
