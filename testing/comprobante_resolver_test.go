package testing

import (
	"errors"
	"fmt"
	"testing"

	"github.com/SaulEnriqueMR/sat-serializer-xml-golang/serializers/cfdi"
)

func TestUnmarshalComprobante_CasosValidos(t *testing.T) {
	tests := []struct {
		name        string
		xml         string
		wantType    string // nombre del tipo esperado
		wantVersion string
	}{
		{
			name:        "1.0 sin namespace, atributo version",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><Comprobante version="1.0"/>`,
			wantType:    "*cfdi.Comprobante10",
			wantVersion: "1.0",
		},
		{
			name:        "2.0 cfd/2, atributo version",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/2" version="2.0"/>`,
			wantType:    "*cfdi.Comprobante20",
			wantVersion: "2.0",
		},
		{
			name:        "2.2 cfd/2, atributo version",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/2" version="2.2"/>`,
			wantType:    "*cfdi.Comprobante22",
			wantVersion: "2.2",
		},
		{
			name:        "3.0 cfd/3, atributo version",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3" version="3.0"/>`,
			wantType:    "*cfdi.Comprobante30",
			wantVersion: "3.0",
		},
		{
			name:        "3.2 cfd/3, atributo version",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3" version="3.2"/>`,
			wantType:    "*cfdi.Comprobante32",
			wantVersion: "3.2",
		},
		{
			name:        "3.3 cfd/3, atributo Version (mayúscula)",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3" Version="3.3"/>`,
			wantType:    "*cfdi.Comprobante33",
			wantVersion: "3.3",
		},
		{
			name:        "4.0 cfd/4, atributo Version (mayúscula)",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4" Version="4.0"/>`,
			wantType:    "*cfdi.Comprobante40",
			wantVersion: "4.0",
		},
		{
			name:        "prefijo distinto al esperado (p:) pero mismo namespace",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><p:Comprobante xmlns:p="http://www.sat.gob.mx/cfd/4" Version="4.0"/>`,
			wantType:    "*cfdi.Comprobante40",
			wantVersion: "4.0",
		},
		{
			name:        "namespace por defecto (sin prefijo)",
			xml:         `<?xml version="1.0" encoding="UTF-8"?><Comprobante xmlns="http://www.sat.gob.mx/cfd/4" Version="4.0"/>`,
			wantType:    "*cfdi.Comprobante40",
			wantVersion: "4.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfd, err := cfdi.UnmarshalComprobante([]byte(tt.xml))
			if err != nil {
				t.Fatalf("error inesperado: %v", err)
			}
			if cfd == nil {
				t.Fatal("se esperaba un Comprobante, se obtuvo nil")
			}

			gotType := typeName(cfd)
			if gotType != tt.wantType {
				t.Errorf("tipo: got %s, want %s", gotType, tt.wantType)
			}
			if cfd.GetVersion() != tt.wantVersion {
				t.Errorf("versión: got %q, want %q", cfd.GetVersion(), tt.wantVersion)
			}
		})
	}
}

func TestUnmarshalComprobante_Errores(t *testing.T) {
	tests := []struct {
		name    string
		xml     string
		wantErr error
	}{
		{
			name:    "sin atributo version",
			xml:     `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4"/>`,
			wantErr: cfdi.ErrVersionAusente,
		},
		{
			name:    "versión desconocida",
			xml:     `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4" Version="9.9"/>`,
			wantErr: cfdi.ErrVersionDesconocida,
		},
		{
			name:    "versión 4.0 con namespace de cfd/3",
			xml:     `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3" Version="4.0"/>`,
			wantErr: cfdi.ErrNamespaceInconsistente,
		},
		{
			name:    "versión 3.3 con namespace de cfd/2",
			xml:     `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/2" Version="3.3"/>`,
			wantErr: cfdi.ErrNamespaceInconsistente,
		},
		{
			name:    "versión 1.0 con namespace declarado (debería ser sin ns)",
			xml:     `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4" version="1.0"/>`,
			wantErr: cfdi.ErrNamespaceInconsistente,
		},
		{
			name:    "XML malformado",
			xml:     `<Comprobante version="1.0"`,
			wantErr: nil, // cualquier error de parseo; validamos aparte
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cfdi.UnmarshalComprobante([]byte(tt.xml))
			if err == nil {
				t.Fatal("se esperaba error, se obtuvo nil")
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("error: got %v, want que cumpla errors.Is(%v)", err, tt.wantErr)
			}
		})
	}
}

// typeName devuelve el nombre cualificado del tipo dinámico del Comprobante.
func typeName(c cfdi.Comprobante) string {
	// %T imprime "*cfdi.Comprobante40", que es justo lo que queremos.
	return fmt.Sprintf("%T", c)
}
