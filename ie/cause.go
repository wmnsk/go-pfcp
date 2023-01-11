// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// Cause definitions.
const (
	_                                                                    uint8 = 0
	CauseRequestAccepted                                                 uint8 = 1
	CauseMoreUsageReportToSend                                           uint8 = 2
	CauseRequestPartiallyAccepted                                              = 3
	CauseRequestRejected                                                 uint8 = 64
	CauseSessionContextNotFound                                          uint8 = 65
	CauseMandatoryIEMissing                                              uint8 = 66
	CauseConditionalIEMissing                                            uint8 = 67
	CauseInvalidLength                                                   uint8 = 68
	CauseMandatoryIEIncorrect                                            uint8 = 69
	CauseInvalidForwardingPolicy                                         uint8 = 70
	CauseInvalidFTEIDAllocationOption                                    uint8 = 71
	CauseNoEstablishedPFCPAssociation                                    uint8 = 72
	CauseRuleCreationModificationFailure                                 uint8 = 73
	CausePFCPEntityInCongestion                                          uint8 = 74
	CauseNoResourcesAvailable                                            uint8 = 75
	CauseServiceNotSupported                                             uint8 = 76
	CauseSystemFailure                                                   uint8 = 77
	CauseRedirectionRequested                                            uint8 = 78
	CauseAllDynamicAddressesAreOccupied                                  uint8 = 79
	CauseUnknownPredefinedRule                                           uint8 = 80
	CauseUnknownApplicationID                                            uint8 = 81
	CauseL2TPTunnelEstablishmentFailure                                  uint8 = 82
	CauseL2TPSessionEstablishmentFailure                                 uint8 = 83
	CauseL2TPTunnelRelease                                               uint8 = 84
	CauseL2TPSessionRelease                                              uint8 = 85
	CausePFCPSessionRestorationFailureDueToRequestedResourceNotAvailable uint8 = 86
	CauseL2TPTunnelEstablishmentFailureTunnelAuthFailure                 uint8 = 87
	CauseL2TPSessionEstablihmentFailureSessionAuthFailure                uint8 = 88
	CauseL2TPTunnelEstablishmentFailureLNSNotReachable                   uint8 = 89
)

// NewCause creates a new Cause IE.
func NewCause(cause uint8) *IE {
	return newUint8ValIE(Cause, cause)
}

// Cause returns Cause in uint8 if the type of IE matches.
func (i *IE) Cause() (uint8, error) {
	if i.Type != Cause {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
