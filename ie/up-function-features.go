// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUPFunctionFeatures creates a new UPFunctionFeatures IE.
// Each feature should be given by octets (5th to 8th octet). It expects 4 octets
// as input, excessive ones are ignored.
func NewUPFunctionFeatures(features ...uint8) *IE {
	var l int
	if len(features) >= 3 {
		l = 4
	} else {
		l = 2
	}

	ie := New(UPFunctionFeatures, make([]byte, l))
	for i, feature := range features {
		if i > 3 {
			break
		}
		ie.Payload[i] = feature
	}

	return ie
}

// UPFunctionFeatures returns UPFunctionFeatures in []byte if the type of IE matches.
func (i *IE) UPFunctionFeatures() ([]byte, error) {
	if i.Type != UPFunctionFeatures {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasBUCP reports whether an IE has BUCP bit.
func (i *IE) HasBUCP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasDDND reports whether an IE has DDND bit.
func (i *IE) HasDDND() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasDLBD reports whether an IE has DLBD bit.
func (i *IE) HasDLBD() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has3rdBit(i.Payload[0])
}

// HasTRST reports whether an IE has TRST bit.
func (i *IE) HasTRST() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has4thBit(i.Payload[0])
}

// HasFTUP reports whether an IE has FTUP bit.
func (i *IE) HasFTUP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has5thBit(i.Payload[0])
}

// HasPFDM reports whether an IE has PFDM bit.
func (i *IE) HasPFDM() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has6thBit(i.Payload[0])
}

// HasHEEU reports whether an IE has HEEU bit.
func (i *IE) HasHEEU() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has7thBit(i.Payload[0])
}

// HasTREU reports whether an IE has TREU bit.
func (i *IE) HasTREU() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has8thBit(i.Payload[0])
}

// HasEMPU reports whether an IE has EMPU bit.
func (i *IE) HasEMPU() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has1stBit(i.Payload[1])
}

// HasPDIU reports whether an IE has PDIU bit.
func (i *IE) HasPDIU() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has2ndBit(i.Payload[1])
}

// HasUDBC reports whether an IE has UDBC bit.
func (i *IE) HasUDBC() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has3rdBit(i.Payload[1])
}

// HasQUOAC reports whether an IE has QUOAC bit.
func (i *IE) HasQUOAC() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has4thBit(i.Payload[1])
}

// HasTRACE reports whether an IE has TRACE bit.
func (i *IE) HasTRACE() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has5thBit(i.Payload[1])
}

// HasFRRT reports whether an IE has FRRT bit.
func (i *IE) HasFRRT() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has6thBit(i.Payload[1])
}

// HasPFDE reports whether an IE has PFDE bit.
func (i *IE) HasPFDE() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has7thBit(i.Payload[1])
}

// HasEPFAR reports whether an IE has EPFAR bit.
func (i *IE) HasEPFAR() bool {
	switch i.Type {
	case UPFunctionFeatures:
		if len(i.Payload) < 2 {
			return false
		}

		return has8thBit(i.Payload[1])
	case CPFunctionFeatures:
		if len(i.Payload) < 1 {
			return false
		}

		return has3rdBit(i.Payload[0])
	default:
		return false
	}
}

// HasDPDRA reports whether an IE has DPDRA bit.
func (i *IE) HasDPDRA() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 3 {
		return false
	}

	return has1stBit(i.Payload[2])
}

// HasADPDP reports whether an IE has ADPDP bit.
func (i *IE) HasADPDP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 3 {
		return false
	}

	return has2ndBit(i.Payload[2])
}

// HasUEIP reports whether an IE has UEIP bit.
func (i *IE) HasUEIP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 3 {
		return false
	}

	return has3rdBit(i.Payload[2])
}

// HasSSET reports whether an IE has SSET bit.
func (i *IE) HasSSET() bool {
	switch i.Type {
	case UPFunctionFeatures:
		if len(i.Payload) < 3 {
			return false
		}

		return has4thBit(i.Payload[2])
	case CPFunctionFeatures:
		if len(i.Payload) < 1 {
			return false
		}

		return has4thBit(i.Payload[0])
	default:
		return false
	}
}

// HasMNOP reports whether an IE has MNOP bit.
func (i *IE) HasMNOP() bool {
	switch i.Type {
	case UPFunctionFeatures:
		if len(i.Payload) < 3 {
			return false
		}

		return has5thBit(i.Payload[2])
	case MeasurementInformation:
		if len(i.Payload) < 1 {
			return false
		}

		return has5thBit(i.Payload[0])
	default:
		return false
	}
}

// HasMTE reports whether an IE has MTE bit.
func (i *IE) HasMTE() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 3 {
		return false
	}

	return has6thBit(i.Payload[2])
}

// HasBUNDL reports whether an IE has BUNDL bit.
func (i *IE) HasBUNDL() bool {
	switch i.Type {
	case UPFunctionFeatures:
		if len(i.Payload) < 3 {
			return false
		}

		return has7thBit(i.Payload[2])
	case CPFunctionFeatures:
		if len(i.Payload) < 1 {
			return false
		}

		return has5thBit(i.Payload[0])
	default:
		return false
	}
}

// HasGCOM reports whether an IE has GCOM bit.
func (i *IE) HasGCOM() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 3 {
		return false
	}

	return has8thBit(i.Payload[2])
}

// HasMPAS reports whether an IE has MPAS bit.
func (i *IE) HasMPAS() bool {
	switch i.Type {
	case UPFunctionFeatures:
		if len(i.Payload) < 4 {
			return false
		}

		return has1stBit(i.Payload[3])
	case CPFunctionFeatures:
		if len(i.Payload) < 1 {
			return false
		}

		return has6thBit(i.Payload[0])
	default:
		return false
	}
}

// HasRTTL reports whether an IE has RTTL bit.
func (i *IE) HasRTTL() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 4 {
		return false
	}

	return has2ndBit(i.Payload[3])
}

// HasVTIME reports whether an IE has VTIME bit.
func (i *IE) HasVTIME() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 4 {
		return false
	}

	return has3rdBit(i.Payload[3])
}

// HasNORP reports wether an IE has NORP bit.
func (i *IE) HasNORP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 4 {
		return false
	}

	return has4thBit(i.Payload[3])
}

// HasIPTV reports wether an IE has IPTV bit.
func (i *IE) HasIPTV() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 4 {
		return false
	}

	return has5thBit(i.Payload[3])
}

// HasTSCU reports wether an IE has TSCU bit.
func (i *IE) HasTSCU() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 4 {
		return false
	}

	return has7thBit(i.Payload[3])
}

// HasMPTCP reports wether an IE has MPTCP bit.
func (i *IE) HasMPTCP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 4 {
		return false
	}

	return has8thBit(i.Payload[3])
}

// HasATSSSLL reports wether an IE has ATSSS-LL bit.
func (i *IE) HasATSSSLL() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has1stBit(i.Payload[4])
}

// HasQFQM reports wether an IE has QFQM bit.
func (i *IE) HasQFQM() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has2ndBit(i.Payload[4])
}

// HasGPQM reports wether an IE has GPQM bit.
func (i *IE) HasGPQM() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has3rdBit(i.Payload[4])
}

// HasMTEDT reports wether an IE has MT-EDT bit.
func (i *IE) HasMTEDT() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has4thBit(i.Payload[4])
}

// HasCIOT reports wether an IE has CIOT bit.
func (i *IE) HasCIOT() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has5thBit(i.Payload[4])
}

// HasETHAR reports wether an IE has ETHAR bit.
func (i *IE) HasETHAR() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has6thBit(i.Payload[4])
}

// HasDDDS reports wether an IE has DDDS bit.
func (i *IE) HasDDDS() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 5 {
		return false
	}

	return has7thBit(i.Payload[4])
}

// HasRTTWP reports wether an IE has RTTWP bit.
func (i *IE) HasRTTWP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has1stBit(i.Payload[5])
}

// HasQUASF reports wether an IE has QUASF bit.
func (i *IE) HasQUASF() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has2ndBit(i.Payload[5])
}

// HasNSPOC reports wether an IE has NSPOC bit.
func (i *IE) HasNSPOC() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has3rdBit(i.Payload[5])
}

// HasL2TP reports wether an IE has L2TP bit.
func (i *IE) HasL2TP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has4thBit(i.Payload[5])
}

// HasUPBER reports wether an IE has UPBER bit.
func (i *IE) HasUPBER() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has5thBit(i.Payload[5])
}

// HasRESPS reports wether an IE has RESPS bit.
func (i *IE) HasRESPS() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has6thBit(i.Payload[5])
}

// HasIPREP reports wether an IE has IPREP bit.
func (i *IE) HasIPREP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has7thBit(i.Payload[5])
}

// HasDNSTS reports wether an IE has DNSTS bit.
func (i *IE) HasDNSTS() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 6 {
		return false
	}

	return has8thBit(i.Payload[5])
}

// HasDRQOS reports wether an IE has DRQOS bit.
func (i *IE) HasDRQOS() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 7 {
		return false
	}

	return has1stBit(i.Payload[6])
}

// HasMBSN4 reports wether an IE has MBSN4 bit.
func (i *IE) HasMBSN4() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 7 {
		return false
	}

	return has2ndBit(i.Payload[6])
}

// HasPSUPRM reports wether an IE has PSUPRM bit.
func (i *IE) HasPSUPRM() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 7 {
		return false
	}

	return has3rdBit(i.Payload[6])
}

// HasEPPPI reports wether an IE has EPPPI bit.
func (i *IE) HasEPPPI() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 7 {
		return false
	}

	return has4thBit(i.Payload[6])
}

// HasRATP reports wether an IE has RATP bit.
func (i *IE) HasRATP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 7 {
		return false
	}

	return has5thBit(i.Payload[6])
}

// HasUPIDP reports wether an IE has UPIDP bit.
func (i *IE) HasUPIDP() bool {
	if i.Type != UPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 7 {
		return false
	}

	return has6thBit(i.Payload[6])
}
