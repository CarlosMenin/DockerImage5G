package ngapType

import "my5G-RANTester/lib/aper"

// Need to import "free5gc/lib/aper" if it uses "aper"

type SecurityKey struct {
	Value aper.BitString `aper:"sizeLB:256,sizeUB:256"`
}
