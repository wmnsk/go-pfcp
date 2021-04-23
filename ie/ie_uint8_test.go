// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"net"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestUint8IEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     uint8
		decoderFunc func(*ie.IE) (uint8, error)
	}{
		{
			description: "AccessAvailabilityInformation",
			structured:  ie.NewAccessAvailabilityInformation(3, 3),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.AccessAvailabilityInformation() },
		}, {
			description: "AccessAvailabilityInformation/AvailabilityStatus",
			structured:  ie.NewAccessAvailabilityInformation(3, 3),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.AvailabilityStatus() },
		}, {
			description: "AccessAvailabilityInformation/AccessType",
			structured:  ie.NewAccessAvailabilityInformation(3, 3),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.AccessType() },
		}, {
			description: "AccessAvailabilityInformation/AccessAvailabilityReport",
			structured: ie.NewAccessAvailabilityReport(
				ie.NewAccessAvailabilityInformation(3, 3),
			),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.AccessAvailabilityInformation() },
		}, {
			description: "ApplyAction",
			structured:  ie.NewApplyAction(4),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(4),
			),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(4),
			),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ApplyAction() },
		}, {
			description: "ATSSSLLControlInformation",
			structured:  ie.NewATSSSLLControlInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLControlInformation() },
		}, {
			description: "ATSSSLLControlInformation/ProvideATSSSControlInformation",
			structured: ie.NewProvideATSSSControlInformation(
				ie.NewMPTCPControlInformation(1),
				ie.NewATSSSLLControlInformation(1),
			),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLControlInformation() },
		}, {
			description: "ATSSSLLInformation",
			structured:  ie.NewATSSSLLInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLInformation() },
		}, {
			description: "ATSSSLLInformation/ATSSSLLParameters",
			structured: ie.NewATSSSLLParameters(
				ie.NewATSSSLLInformation(1),
			),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLInformation() },
		}, {
			description: "ATSSSLLInformation/ATSSSControlParameters",
			structured: ie.NewATSSSControlParameters(
				ie.NewMPTCPParameters(
					ie.NewMPTCPAddressInformation(ie.MPTCPProxyTransportConverter, 8080, net.ParseIP("127.0.0.1"), net.ParseIP("2001::1")),
					ie.NewUELinkSpecificIPAddress(net.ParseIP("127.0.0.1"), net.ParseIP("2001::1"), net.ParseIP("127.0.0.1"), net.ParseIP("2001::1")),
				),
				ie.NewATSSSLLParameters(
					ie.NewATSSSLLInformation(1),
				),
			),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLInformation() },
		}, {
			description: "BARID",
			structured:  ie.NewBARID(0xff),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "BARID/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewBARID(0xff),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "BARID/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewBARID(0xff),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "BARID/CreateBAR",
			structured: ie.NewCreateBAR(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "BARID/UpdateBARWithinSessionReportResponse",
			structured: ie.NewUpdateBARWithinSessionReportResponse(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
				ie.NewDLBufferingDuration(30*time.Second),
				ie.NewDLBufferingSuggestedPacketCount(0xffff),
				ie.NewSuggestedBufferingPacketsCount(0x01),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "BARID/UpdateBARWithinSessionModificationRequest",
			structured: ie.NewUpdateBARWithinSessionModificationRequest(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "BARID/RemoveBAR",
			structured: ie.NewRemoveBAR(
				ie.NewBARID(0xff),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "Cause",
			structured:  ie.NewCause(ie.CauseRequestAccepted),
			decoded:     ie.CauseRequestAccepted,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Cause() },
		}, {
			description: "CPFunctionFeatures",
			structured:  ie.NewCPFunctionFeatures(0x3f),
			decoded:     0x3f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.CPFunctionFeatures() },
		}, {
			description: "CreateBridgeInfoForTSC",
			structured:  ie.NewCreateBridgeInfoForTSC(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.CreateBridgeInfoForTSC() },
		}, {
			description: "DestinationInterface",
			structured:  ie.NewDestinationInterface(ie.DstInterfaceAccess),
			decoded:     ie.DstInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.DestinationInterface() },
		}, {
			description: "DestinationInterface/ForwardingParameters",
			structured: ie.NewForwardingParameters(
				ie.NewDestinationInterface(0xff),
				ie.NewNetworkInstance("some.instance.example"),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.DestinationInterface() },
		}, {
			description: "DestinationInterface/UpdateForwardingParameters",
			structured: ie.NewUpdateForwardingParameters(
				ie.NewDestinationInterface(0xff),
				ie.NewNetworkInstance("some.instance.example"),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.DestinationInterface() },
		}, {
			description: "DestinationInterface/DuplicatingParameters",
			structured: ie.NewDuplicatingParameters(
				ie.NewDestinationInterface(0xff),
				ie.NewNetworkInstance("some.instance.example"),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.DestinationInterface() },
		}, {
			description: "DestinationInterface/UpdateDuplicatingParameters",
			structured: ie.NewUpdateDuplicatingParameters(
				ie.NewDestinationInterface(0xff),
				ie.NewNetworkInstance("some.instance.example"),
			),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.DestinationInterface() },
		}, {
			description: "DLBufferingSuggestedPacketCount",
			structured:  ie.NewDLBufferingSuggestedPacketCount(0xff),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) {
				v, err := i.DLBufferingSuggestedPacketCount()
				return uint8(v), err
			},
		}, {
			description: "DownlinkDataServiceInformation/PPI",
			structured:  ie.NewDownlinkDataServiceInformation(true, true, 1, 1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PPI() },
		}, {
			description: "DownlinkDataServiceInformation/QFI",
			structured:  ie.NewDownlinkDataServiceInformation(true, true, 1, 1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "EthernetFilterProperties",
			structured:  ie.NewEthernetFilterProperties(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetFilterProperties() },
		}, {
			description: "EthernetFilterProperties/PDI",
			structured: ie.NewPDI(
				ie.NewSourceInterface(ie.SrcInterfaceAccess),
				ie.NewEthernetFilterProperties(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetFilterProperties() },
		}, {
			description: "EthernetFilterProperties/EthernetPacketFilter",
			structured: ie.NewEthernetPacketFilter(
				ie.NewEthernetFilterID(0xffffffff),
				ie.NewEthernetFilterProperties(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetFilterProperties() },
		}, {
			description: "EthernetPDUSessionInformation",
			structured:  ie.NewEthernetPDUSessionInformation(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetPDUSessionInformation() },
		}, {
			description: "EthernetPDUSessionInformation/PDI",
			structured: ie.NewPDI(
				ie.NewSourceInterface(ie.SrcInterfaceAccess),
				ie.NewEthernetPDUSessionInformation(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetPDUSessionInformation() },
		}, {
			description: "EthernetPDUSessionInformation/CreateTrafficEndpoint",
			structured: ie.NewCreateTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
				ie.NewEthernetPDUSessionInformation(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetPDUSessionInformation() },
		}, {
			description: "FailedRuleID/RuleIDType",
			structured:  ie.NewFailedRuleID(ie.RuleIDTypePDR, 0xffff),
			decoded:     ie.RuleIDTypePDR,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RuleIDType() },
		}, {
			description: "FlowInformation/FlowDirection",
			structured:  ie.NewFlowInformation(ie.FlowDirectionDownlink, "go-pfcp"),
			decoded:     ie.FlowDirectionDownlink,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.FlowDirection() },
		}, {
			description: "ApplicationDetectionInformation/FlowDirection",
			structured: ie.NewApplicationDetectionInformation(
				ie.NewPDRID(0xffff),
				ie.NewFlowInformation(ie.FlowDirectionDownlink, "go-pfcp"),
			),
			decoded:     ie.FlowDirectionDownlink,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.FlowDirection() },
		}, {
			description: "FQCSID/NodeIDType/IPv4",
			structured:  ie.NewFQCSID("127.0.0.1", 1),
			decoded:     ie.NodeIDIPv4Address,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.NodeIDType() },
		}, {
			description: "FQCSID/NodeIDType/IPv6",
			structured:  ie.NewFQCSID("2001::1", 1),
			decoded:     ie.NodeIDIPv6Address,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.NodeIDType() },
		}, {
			description: "GateStatus/OpenOpen",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenClosed",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusClosed),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/ClosedOpen",
			structured:  ie.NewGateStatus(ie.GateStatusClosed, ie.GateStatusOpen),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/ClosedClosed",
			structured:  ie.NewGateStatus(ie.GateStatusClosed, ie.GateStatusClosed),
			decoded:     5,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenClosed/GateStatusUL",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusClosed),
			decoded:     ie.GateStatusOpen,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatusUL() },
		}, {
			description: "GateStatus/OpenClosed/GateStatusDL",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusClosed),
			decoded:     ie.GateStatusClosed,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatusDL() },
		}, {
			description: "GateStatus/OpenOpen/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenOpen/UpdateQER",
			structured: ie.NewUpdateQER(
				ie.NewQERID(0xffffffff),
				ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GTPUPathInterfaceType",
			structured:  ie.NewGTPUPathInterfaceType(1, 1),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GTPUPathInterfaceType() },
		}, {
			description: "GTPUPathInterfaceType/GTPUPathQoSControlInformation",
			structured: ie.NewGTPUPathQoSControlInformation(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewGTPUPathInterfaceType(1, 1),
			),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GTPUPathInterfaceType() },
		}, {
			description: "GTPUPathInterfaceType/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewGTPUPathInterfaceType(1, 1),
			),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GTPUPathInterfaceType() },
		}, {
			description: "MeasurementInformation",
			structured:  ie.NewMeasurementInformation(0x1f),
			decoded:     0x1f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MeasurementInformation() },
		}, {
			description: "MeasurementInformation/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementInformation(0x1f),
			),
			decoded:     0x1f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MeasurementInformation() },
		}, {
			description: "MeasurementInformation/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementInformation(0x1f),
			),
			decoded:     0x1f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MeasurementInformation() },
		}, {
			description: "Metric",
			structured:  ie.NewMetric(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Metric() },
		}, {
			description: "Metric/LoadControlInformation",
			structured: ie.NewLoadControlInformation(
				ie.NewSequenceNumber(0xffffffff),
				ie.NewMetric(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Metric() },
		}, {
			description: "Metric/OverloadControlInformation",
			structured: ie.NewOverloadControlInformation(
				ie.NewSequenceNumber(0xffffffff),
				ie.NewMetric(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Metric() },
		}, {
			description: "MPTCPControlInformation",
			structured:  ie.NewMPTCPControlInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MPTCPControlInformation() },
		}, {
			description: "MPTCPControlInformation/ProvideATSSSControlInformation",
			structured: ie.NewProvideATSSSControlInformation(
				ie.NewMPTCPControlInformation(0x01),
				ie.NewATSSSLLControlInformation(1),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MPTCPControlInformation() },
		}, {
			description: "MTEDTControlInformation",
			structured:  ie.NewMTEDTControlInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MTEDTControlInformation() },
		}, {
			description: "MTEDTControlInformation/ProvideATSSSControlInformation",
			structured: ie.NewCreateBAR(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
				ie.NewSuggestedBufferingPacketsCount(0x01),
				ie.NewMTEDTControlInformation(1),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MTEDTControlInformation() },
		}, {
			description: "NodeReportType",
			structured:  ie.NewNodeReportType(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.NodeReportType() },
		}, {
			description: "OCIFlags",
			structured:  ie.NewOCIFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.OCIFlags() },
		}, {
			description: "OCIFlags/OverloadControlInformation",
			structured: ie.NewOverloadControlInformation(
				ie.NewSequenceNumber(0xffffffff),
				ie.NewOCIFlags(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.OCIFlags() },
		}, {
			description: "OuterHeaderRemoval/OuterHeaderRemovalDescription",
			structured:  ie.NewOuterHeaderRemoval(0x01, 0x02),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.OuterHeaderRemovalDescription() },
		}, {
			description: "OuterHeaderRemoval/GTPUExternsionHeaderDeletion",
			structured:  ie.NewOuterHeaderRemoval(0x01, 0x02),
			decoded:     0x02,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GTPUExternsionHeaderDeletion() },
		}, {
			description: "PacketReplicationAndDetectionCarryOnInformation",
			structured:  ie.NewPacketReplicationAndDetectionCarryOnInformation(0x0f),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PacketReplicationAndDetectionCarryOnInformation() },
		}, {
			description: "PacketReplicationAndDetectionCarryOnInformation/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff),
				ie.NewPacketReplicationAndDetectionCarryOnInformation(0x0f),
			),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PacketReplicationAndDetectionCarryOnInformation() },
		}, {
			description: "PagingPolicyIndicator",
			structured:  ie.NewPagingPolicyIndicator(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PagingPolicyIndicator() },
		}, {
			description: "PagingPolicyIndicator/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewPagingPolicyIndicator(1),
			),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PagingPolicyIndicator() },
		}, {
			description: "PagingPolicyIndicator/UpdateQER",
			structured: ie.NewUpdateQER(
				ie.NewQERID(0xffffffff),
				ie.NewPagingPolicyIndicator(1),
			),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PagingPolicyIndicator() },
		}, {
			description: "PDNType",
			structured:  ie.NewPDNType(ie.PDNTypeIPv4),
			decoded:     ie.PDNTypeIPv4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PDNType() },
		}, {
			description: "PFCPAssociationReleaseRequest",
			structured:  ie.NewPFCPAssociationReleaseRequest(1, 1),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPAssociationReleaseRequest() },
		}, {
			description: "PFCPASRspFlags",
			structured:  ie.NewPFCPASRspFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPASRspFlags() },
		}, {
			description: "PFCPAUReqFlags",
			structured:  ie.NewPFCPAUReqFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPAUReqFlags() },
		}, {
			description: "PFCPSEReqFlags",
			structured:  ie.NewPFCPSEReqFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPSEReqFlags() },
		}, {
			description: "PFCPSRRspFlags",
			structured:  ie.NewPFCPSRRspFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPSRRspFlags() },
		}, {
			description: "PMFControlInformation",
			structured:  ie.NewPMFControlInformation(0xff),
			decoded:     0x01, // first bit should only be evaluated.
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PMFControlInformation() },
		}, {
			description: "PMFControlInformation/ProvideATSSSControlInformation",
			structured: ie.NewProvideATSSSControlInformation(
				ie.NewMPTCPControlInformation(1),
				ie.NewPMFControlInformation(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PMFControlInformation() },
		}, {
			description: "Priority",
			structured:  ie.NewPriority(ie.PriorityActive),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Priority/CreateMAR",
			structured: ie.NewCreateMAR(
				ie.NewMARID(0xffff),
				ie.NewTGPPAccessForwardingActionInformation(
					ie.NewFARID(0xffffffff),
					ie.NewWeight(0x01),
					ie.NewPriority(ie.PriorityActive),
					ie.NewURRID(0xffffffff),
				),
			),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Priority/UpdateMAR",
			structured: ie.NewUpdateMAR(
				ie.NewMARID(0xffff),
				ie.NewTGPPAccessForwardingActionInformation(
					ie.NewFARID(0xffffffff),
					ie.NewWeight(0x01),
					ie.NewPriority(ie.PriorityActive),
					ie.NewURRID(0xffffffff),
				),
			),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Priority/TGPPAccessForwardingActionInformation",
			structured: ie.NewTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewPriority(ie.PriorityActive),
			),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Priority/NonTGPPAccessForwardingActionInformation",
			structured: ie.NewNonTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewPriority(ie.PriorityActive),
			),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Priority/UpdateTGPPAccessForwardingActionInformation",
			structured: ie.NewUpdateTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewPriority(ie.PriorityActive),
			),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Priority/UpdateNonTGPPAccessForwardingActionInformation",
			structured: ie.NewUpdateNonTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewPriority(ie.PriorityActive),
			),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Proxying",
			structured:  ie.NewProxying(1, 1),
			decoded:     0x03,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Proxying() },
		}, {
			description: "Proxying/ForwardingParameters",
			structured: ie.NewForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewProxying(1, 1),
			),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Proxying() },
		}, {
			description: "Proxying/UpdateForwardingParameters",
			structured: ie.NewUpdateForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewProxying(1, 1),
			),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Proxying() },
		}, {
			description: "QERControlIndications",
			structured:  ie.NewQERControlIndications(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QERControlIndications() },
		}, {
			description: "QERControlIndications/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewQERControlIndications(1, 1, 1),
			),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QERControlIndications() },
		}, {
			description: "QERControlIndications/UpdateQER",
			structured: ie.NewUpdateQER(
				ie.NewQERID(0xffffffff),
				ie.NewQERControlIndications(1, 1, 1),
			),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QERControlIndications() },
		}, {
			description: "QFI",
			structured:  ie.NewQFI(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QFI/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewQFI(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QFI/UpdateQER",
			structured: ie.NewUpdateQER(
				ie.NewQERID(0xffffffff),
				ie.NewQFI(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QFI/CreateTrafficEndpoint",
			structured: ie.NewCreateTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
				ie.NewQFI(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QFI/UpdateTrafficEndpoint",
			structured: ie.NewUpdateTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
				ie.NewQFI(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QFI/QoSMonitoringPerQoSFlowControlInformation",
			structured: ie.NewQoSMonitoringPerQoSFlowControlInformation(
				ie.NewQFI(0x01),
				ie.NewRequestedQoSMonitoring(1, 1, 1),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QFI/QoSMonitoringReport",
			structured: ie.NewQoSMonitoringReport(
				ie.NewQFI(0x01),
				ie.NewQoSMonitoringMeasurement(0x0f, 0x11111111, 0x22222222, 0x33333333),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QoSReportTrigger",
			structured:  ie.NewQoSReportTrigger(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QoSReportTrigger() },
		}, {
			description: "ReportType",
			structured:  ie.NewReportType(1, 1, 1, 1),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ReportType() },
		}, {
			description: "ReportingFrequency",
			structured:  ie.NewReportingFrequency(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ReportingFrequency() },
		}, {
			description: "ReportingFrequency/QoSMonitoringPerQoSFlowControlInformation",
			structured: ie.NewQoSMonitoringPerQoSFlowControlInformation(
				ie.NewRequestedQoSMonitoring(1, 1, 1),
				ie.NewReportingFrequency(1, 1, 1),
			),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ReportingFrequency() },
		}, {
			description: "RequestedAccessAvailabilityInformation",
			structured:  ie.NewRequestedAccessAvailabilityInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedAccessAvailabilityInformation() },
		}, {
			description: "RequestedAccessAvailabilityInformation/AccessAvailabilityControlInformation",
			structured: ie.NewAccessAvailabilityControlInformation(
				ie.NewRequestedAccessAvailabilityInformation(1),
			),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedAccessAvailabilityInformation() },
		}, {
			description: "RequestedClockDriftInformation",
			structured:  ie.NewRequestedClockDriftInformation(1, 1),
			decoded:     0x03,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedClockDriftInformation() },
		}, {
			description: "RequestedClockDriftInformation/ClockDriftControlInformation",
			structured: ie.NewClockDriftControlInformation(
				ie.NewRequestedClockDriftInformation(1, 1),
				ie.NewTSNTimeDomainNumber(255),
			),
			decoded:     0x03,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedClockDriftInformation() },
		}, {
			description: "RequestedQoSMonitoring",
			structured:  ie.NewRequestedQoSMonitoring(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedQoSMonitoring() },
		}, {
			description: "RequestedQoSMonitoring/QoSMonitoringPerQoSFlowControlInformation",
			structured: ie.NewQoSMonitoringPerQoSFlowControlInformation(
				ie.NewQFI(0x01),
				ie.NewRequestedQoSMonitoring(1, 1, 1),
			),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedQoSMonitoring() },
		}, {
			description: "RQI",
			structured:  ie.NewRQI(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RQI() },
		}, {
			description: "RQI/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewRQI(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RQI() },
		}, {
			description: "RQI/UpdateQER",
			structured: ie.NewUpdateQER(
				ie.NewQERID(0xffffffff),
				ie.NewRQI(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RQI() },
		}, {
			description: "SourceInterface",
			structured:  ie.NewSourceInterface(ie.SrcInterfaceAccess),
			decoded:     ie.SrcInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SourceInterface() },
		}, {
			description: "SourceInterface/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff), ie.NewPDI(
					ie.NewSourceInterface(ie.SrcInterfaceAccess),
					ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
				),
			),
			decoded:     ie.SrcInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SourceInterface() },
		}, {
			description: "SourceInterface/UpdatePDR",
			structured: ie.NewUpdatePDR(
				ie.NewPDRID(0xffff), ie.NewPDI(
					ie.NewSourceInterface(ie.SrcInterfaceAccess),
					ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
				),
			),
			decoded:     ie.SrcInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SourceInterface() },
		}, {
			description: "SourceInterface/PDI",
			structured: ie.NewPDI(
				ie.NewSourceInterface(ie.SrcInterfaceAccess),
				ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
			),
			decoded:     ie.SrcInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SourceInterface() },
		}, {
			description: "SRRID",
			structured:  ie.NewSRRID(255),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SRRID() },
		}, {
			description: "SRRID/RemoveSRR",
			structured: ie.NewRemoveSRR(
				ie.NewSRRID(255),
			),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SRRID() },
		}, {
			description: "SRRID/CreateSRR",
			structured: ie.NewCreateSRR(
				ie.NewSRRID(255),
				ie.NewAccessAvailabilityControlInformation(
					ie.NewRequestedAccessAvailabilityInformation(1),
				),
			),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SRRID() },
		}, {
			description: "SRRID/UpdateSRR",
			structured: ie.NewUpdateSRR(
				ie.NewSRRID(255),
				ie.NewAccessAvailabilityControlInformation(
					ie.NewRequestedAccessAvailabilityInformation(1),
				),
			),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SRRID() },
		}, {
			description: "SRRID/SessionReport",
			structured: ie.NewSessionReport(
				ie.NewSRRID(255),
				ie.NewAccessAvailabilityControlInformation(
					ie.NewRequestedAccessAvailabilityInformation(1),
				),
			),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SRRID() },
		}, {
			description: "SteeringFunctionality",
			structured:  ie.NewSteeringFunctionality(ie.SteeringFunctionalityATSSSLL),
			decoded:     ie.SteeringFunctionalityATSSSLL,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringFunctionality() },
		}, {
			description: "SteeringFunctionality/CreateMAR",
			structured: ie.NewCreateMAR(
				ie.NewMARID(0x1111),
				ie.NewSteeringFunctionality(ie.SteeringFunctionalityATSSSLL),
			),
			decoded:     ie.SteeringFunctionalityATSSSLL,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringFunctionality() },
		}, {
			description: "SteeringFunctionality/UpdateMAR",
			structured: ie.NewUpdateMAR(
				ie.NewMARID(0x1111),
				ie.NewSteeringFunctionality(ie.SteeringFunctionalityATSSSLL),
			),
			decoded:     ie.SteeringFunctionalityATSSSLL,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringFunctionality() },
		}, {
			description: "SteeringMode",
			structured:  ie.NewSteeringMode(ie.SteeringModeActiveStandby),
			decoded:     ie.SteeringModeActiveStandby,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringMode() },
		}, {
			description: "SteeringMode/CreateMAR",
			structured: ie.NewCreateMAR(
				ie.NewMARID(0x1111),
				ie.NewSteeringMode(ie.SteeringModeActiveStandby),
			),
			decoded:     ie.SteeringModeActiveStandby,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringMode() },
		}, {
			description: "SteeringMode/UpdateMAR",
			structured: ie.NewUpdateMAR(
				ie.NewMARID(0x1111),
				ie.NewSteeringMode(ie.SteeringModeActiveStandby),
			),
			decoded:     ie.SteeringModeActiveStandby,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringMode() },
		}, {
			description: "SuggestedBufferingPacketsCount",
			structured:  ie.NewSuggestedBufferingPacketsCount(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SuggestedBufferingPacketsCount() },
		}, {
			description: "SuggestedBufferingPacketsCount/CreateBAR",
			structured: ie.NewCreateBAR(
				ie.NewBARID(0xff),
				ie.NewSuggestedBufferingPacketsCount(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SuggestedBufferingPacketsCount() },
		}, {
			description: "SuggestedBufferingPacketsCount/UpdateBARWithinSessionReportResponse",
			structured: ie.NewUpdateBARWithinSessionReportResponse(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
				ie.NewDLBufferingDuration(30*time.Second),
				ie.NewDLBufferingSuggestedPacketCount(0xffff),
				ie.NewSuggestedBufferingPacketsCount(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SuggestedBufferingPacketsCount() },
		}, {
			description: "SuggestedBufferingPacketsCount/UpdateBARWithinSessionModificationRequest",
			structured: ie.NewUpdateBARWithinSessionModificationRequest(
				ie.NewBARID(0xff),
				ie.NewSuggestedBufferingPacketsCount(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SuggestedBufferingPacketsCount() },
		}, {
			description: "TGPPInterfaceType",
			structured:  ie.NewTGPPInterfaceType(ie.TGPPInterfaceTypeS1U),
			decoded:     ie.TGPPInterfaceTypeS1U,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TGPPInterfaceType() },
		}, {
			description: "TGPPInterfaceType/ForwardingParameters",
			structured: ie.NewForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTGPPInterfaceType(ie.TGPPInterfaceTypeS1U),
			),
			decoded:     ie.TGPPInterfaceTypeS1U,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TGPPInterfaceType() },
		}, {
			description: "TGPPInterfaceType/UpdateForwardingParameters",
			structured: ie.NewUpdateForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTGPPInterfaceType(ie.TGPPInterfaceTypeS1U),
			),
			decoded:     ie.TGPPInterfaceTypeS1U,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TGPPInterfaceType() },
		}, {
			description: "TrafficEndpointID",
			structured:  ie.NewTrafficEndpointID(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff),
				ie.NewPDI(
					ie.NewSourceInterface(ie.SrcInterfaceAccess),
					ie.NewTrafficEndpointID(0x01),
				),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/PDI",
			structured: ie.NewPDI(
				ie.NewSourceInterface(ie.SrcInterfaceAccess),
				ie.NewTrafficEndpointID(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/ForwardingParameters",
			structured: ie.NewForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTrafficEndpointID(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/UpdateForwardingParameters",
			structured: ie.NewUpdateForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTrafficEndpointID(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/CreateTrafficEndpoint",
			structured: ie.NewCreateTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
				ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/CreatedTrafficEndpoint",
			structured: ie.NewCreatedTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
				ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/UpdateTrafficEndpoint",
			structured: ie.NewUpdateTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
				ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TrafficEndpointID/RemoveTrafficEndpoint",
			structured: ie.NewRemoveTrafficEndpoint(
				ie.NewTrafficEndpointID(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TSNTimeDomainNumber",
			structured:  ie.NewTSNTimeDomainNumber(255),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TSNTimeDomainNumber() },
		}, {
			description: "TSNTimeDomainNumber/ClockDriftControlInformation",
			structured: ie.NewClockDriftControlInformation(
				ie.NewRequestedClockDriftInformation(1, 1),
				ie.NewTSNTimeDomainNumber(255),
			),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TSNTimeDomainNumber() },
		}, {
			description: "TSNTimeDomainNumber/ClockDriftReport",
			structured: ie.NewClockDriftReport(
				ie.NewTSNTimeDomainNumber(255),
				ie.NewTimeOffsetThreshold(10*time.Second),
			),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TSNTimeDomainNumber() },
		}, {
			description: "UsageInformation",
			structured:  ie.NewUsageInformation(1, 1, 1, 1),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.UsageInformation() },
		}, {
			description: "UsageInformation/UsageReportWithinSessionModificationResponse",
			structured: ie.NewUsageReportWithinSessionModificationResponse(
				ie.NewURRID(0xffffffff),
				ie.NewUsageInformation(1, 1, 1, 1),
			),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.UsageInformation() },
		}, {
			description: "UsageInformation/UsageReportWithinSessionDeletionResponse",
			structured: ie.NewUsageReportWithinSessionDeletionResponse(
				ie.NewURRID(0xffffffff),
				ie.NewUsageInformation(1, 1, 1, 1),
			),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.UsageInformation() },
		}, {
			description: "UsageInformation/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewUsageInformation(1, 1, 1, 1),
			),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.UsageInformation() },
		}, {
			description: "Weight",
			structured:  ie.NewWeight(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		}, {
			description: "Weight/CreateMAR",
			structured: ie.NewCreateMAR(
				ie.NewMARID(0x1111),
				ie.NewTGPPAccessForwardingActionInformation(
					ie.NewFARID(0xffffffff),
					ie.NewWeight(0x01),
				),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		}, {
			description: "Weight/UpdateMAR",
			structured: ie.NewUpdateMAR(
				ie.NewMARID(0x1111),
				ie.NewTGPPAccessForwardingActionInformation(
					ie.NewFARID(0xffffffff),
					ie.NewWeight(0x01),
				),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		}, {
			description: "Weight/TGPPAccessForwardingActionInformation",
			structured: ie.NewTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewWeight(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		}, {
			description: "Weight/NonTGPPAccessForwardingActionInformation",
			structured: ie.NewNonTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewWeight(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		}, {
			description: "Weight/UpdateTGPPAccessForwardingActionInformation",
			structured: ie.NewUpdateTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewWeight(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		}, {
			description: "Weight/UpdateNonTGPPAccessForwardingActionInformation",
			structured: ie.NewUpdateNonTGPPAccessForwardingActionInformation(
				ie.NewFARID(0xffffffff),
				ie.NewWeight(0x01),
			),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.decoderFunc(c.structured)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.decoded); diff != "" {
				t.Error(diff)
			}
		})
	}
}
