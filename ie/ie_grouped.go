package ie

import "sync"

// We're using map to avoid iterating over a list.
// The value `true` is not actually used.
// TODO: consider using a slice with utils in slices package introduced in Go 1.21.
var (
	mu                  sync.RWMutex
	defaultGroupedIEMap = map[uint16]bool{
		CreatePDR:                            true,
		PDI:                                  true,
		CreateFAR:                            true,
		ForwardingParameters:                 true,
		DuplicatingParameters:                true,
		CreateURR:                            true,
		CreateQER:                            true,
		CreatedPDR:                           true,
		UpdatePDR:                            true,
		UpdateFAR:                            true,
		UpdateForwardingParameters:           true,
		UpdateBARWithinSessionReportResponse: true,
		UpdateURR:                            true,
		UpdateQER:                            true,
		RemovePDR:                            true,
		RemoveFAR:                            true,
		RemoveURR:                            true,
		RemoveQER:                            true,
		LoadControlInformation:               true,
		OverloadControlInformation:           true,
		ApplicationIDsPFDs:                   true,
		PFDContext:                           true,
		ApplicationDetectionInformation:      true,
		QueryURR:                             true,
		UsageReportWithinSessionModificationResponse: true,
		UsageReportWithinSessionDeletionResponse:     true,
		UsageReportWithinSessionReportRequest:        true,
		DownlinkDataReport:                           true,
		CreateBAR:                                    true,
		UpdateBARWithinSessionModificationRequest:    true,
		RemoveBAR:                                                 true,
		ErrorIndicationReport:                                     true,
		UserPlanePathFailureReport:                                true,
		UpdateDuplicatingParameters:                               true,
		AggregatedURRs:                                            true,
		CreateTrafficEndpoint:                                     true,
		CreatedTrafficEndpoint:                                    true,
		UpdateTrafficEndpoint:                                     true,
		RemoveTrafficEndpoint:                                     true,
		EthernetPacketFilter:                                      true,
		EthernetTrafficInformation:                                true,
		AdditionalMonitoringTime:                                  true,
		CreateMAR:                                                 true,
		TGPPAccessForwardingActionInformation:                     true,
		NonTGPPAccessForwardingActionInformation:                  true,
		RemoveMAR:                                                 true,
		UpdateMAR:                                                 true,
		UpdateTGPPAccessForwardingActionInformation:               true,
		UpdateNonTGPPAccessForwardingActionInformation:            true,
		PFCPSessionRetentionInformation:                           true,
		UserPlanePathRecoveryReport:                               true,
		IPMulticastAddressingInfo:                                 true,
		JoinIPMulticastInformationWithinUsageReport:               true,
		LeaveIPMulticastInformationWithinUsageReport:              true,
		CreatedBridgeInfoForTSC:                                   true,
		TSCManagementInformationWithinSessionModificationRequest:  true,
		TSCManagementInformationWithinSessionModificationResponse: true,
		TSCManagementInformationWithinSessionReportRequest:        true,
		ClockDriftControlInformation:                              true,
		ClockDriftReport:                                          true,
		RemoveSRR:                                                 true,
		CreateSRR:                                                 true,
		UpdateSRR:                                                 true,
		SessionReport:                                             true,
		AccessAvailabilityControlInformation:                      true,
		AccessAvailabilityReport:                                  true,
		ProvideATSSSControlInformation:                            true,
		ATSSSControlParameters:                                    true,
		MPTCPParameters:                                           true,
		ATSSSLLParameters:                                         true,
		PMFParameters:                                             true,
		UEIPAddressPoolInformation:                                true,
		GTPUPathQoSControlInformation:                             true,
		GTPUPathQoSReport:                                         true,
		QoSInformationInGTPUPathQoSReport:                         true,
		QoSMonitoringPerQoSFlowControlInformation:                 true,
		QoSMonitoringReport:                                       true,
		PacketRateStatusReport:                                    true,
		EthernetContextInformation:                                true,
		RedundantTransmissionParameters:                           true,
		UpdatedPDR:                                                true,
		ProvideRDSConfigurationInformation:                        true,
		QueryPacketRateStatusWithinSessionModificationRequest:     true,
		PacketRateStatusReportWithinSessionModificationResponse:   true,
		UEIPAddressUsageInformation:                               true,
		RedundantTransmissionForwardingParameters:                 true,
		TransportDelayReporting:                                   true,
	}
	isGroupedFun = func(t uint16) bool {
		mu.RLock()
		defer mu.RUnlock()
		_, ok := defaultGroupedIEMap[t]
		return ok
	}
)

// SetIsGroupedFun sets a function to check if an IE is grouped type or not.
func SetIsGroupedFun(fun func(t uint16) bool) {
	mu.Lock()
	defer mu.Unlock()
	isGroupedFun = fun
}

// AddGroupedIEType adds IE type(s) to the defaultGroupedIEMap.
// This is useful when you want to add new IE types to the defaultGroupedIEMap,
// e.g., to handle vendor-specific IEs as grouped type.
func AddGroupedIEType(ts ...uint16) {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range ts {
		defaultGroupedIEMap[t] = true
	}
}

// IsGrouped reports whether an IE is grouped type or not.
//
// By default, this package determines if an IE is grouped type or not by checking
// if the IE type is in the defaultGroupedIEMap.
// You can change this entire behavior by calling SetIsGroupedFun(), or you can add
// new IE types to the defaultGroupedIEMap by calling AddGroupedIEType().
func (i *IE) IsGrouped() bool {
	return isGroupedFun(i.Type)
}

// Add adds variable number of IEs to a IE if the IE is grouped type and update length.
// Otherwise, this does nothing (no errors).
func (i *IE) Add(ies ...*IE) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	i.ChildIEs = append(i.ChildIEs, ies...)
	for _, ie := range i.ChildIEs {
		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.SetLength()
}

// Remove removes an IE looked up by type.
func (i *IE) Remove(typ uint16) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	newChildren := make([]*IE, len(i.ChildIEs))
	idx := 0
	for _, ie := range i.ChildIEs {
		if ie.Type == typ {
			newChildren = newChildren[:len(newChildren)-1]
			continue
		}
		newChildren[idx] = ie
		idx++

		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.ChildIEs = newChildren
	i.SetLength()
}

// FindByType returns IE looked up by type.
//
// The program may be slower when calling this method multiple times
// because this ranges over a ChildIEs each time it is called.
func (i *IE) FindByType(typ uint16) (*IE, error) {
	if !i.IsGrouped() {
		return nil, ErrInvalidType
	}

	for _, ie := range i.ChildIEs {
		if ie.Type == typ {
			return ie, nil
		}
	}
	return nil, ErrIENotFound
}
