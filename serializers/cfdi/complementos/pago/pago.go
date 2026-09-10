package pago

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Pagos interface {
	GetVersion() string
}

func (p *Pagos10) GetVersion() string { return p.Version }
func (p *Pagos20) GetVersion() string { return p.Version }

var (
	ErrNamespaceDesconocido = errors.New("namespace de Pagos no reconocido")
	ErrVersionInconsistente = errors.New("Version no coincide con el namespace")
)

const (
	NS_Pagos10 = "http://www.sat.gob.mx/Pagos"
	NS_Pagos20 = "http://www.sat.gob.mx/Pagos20"
)

type pagosResolver struct {
	ns        string
	version   string
	unmarshal func([]byte) (Pagos, error)
}

var resolvers = []pagosResolver{
	{NS_Pagos10, "1.0", func(b []byte) (Pagos, error) {
		var p Pagos10
		return &p, xml.Unmarshal(b, &p)
	}},
	{NS_Pagos20, "2.0", func(b []byte) (Pagos, error) {
		var p Pagos20
		return &p, xml.Unmarshal(b, &p)
	}},
}

func UnmarshalPagos(data []byte) (Pagos, error) {
	var probe struct {
		XMLName xml.Name
		Version string `xml:"Version,attr"`
	}
	if err := xml.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	for _, r := range resolvers {
		if r.ns == probe.XMLName.Space {
			if probe.Version != "" && probe.Version != r.version {
				return nil, fmt.Errorf("%w: %s/%s", ErrVersionInconsistente, r.ns, probe.Version)
			}
			return r.unmarshal(data)
		}
	}
	return nil, fmt.Errorf("%w: %q", ErrNamespaceDesconocido, probe.XMLName.Space)
}
