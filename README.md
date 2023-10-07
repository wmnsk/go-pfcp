# go-pfcp

A comprehensive PFCP implementation in the Go programming language.

![CI status](https://github.com/wmnsk/go-pfcp/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wmnsk/go-pfcp?status.svg)](https://godoc.org/github.com/wmnsk/go-pfcp)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/go-pfcp/blob/master/LICENSE)

## What is PFCP?

PFCP (Packet Forwarding Control Protocol) is a signaling protocol used in mobile networking infrastruct, such as EPC and 5GC, to implement the CUPS architecture (Control and User Plane Separation, not a printing system). This architecture and the protocol detail is defined in 3GPP TS 29.244.

Are you also searching for a GTP implementation in Go? If so, [go-gtp](https://github.com/wmnsk/go-gtp) is the perfect choice for you! :p

## Project Status

This project is still WIP.  
Implementation of all the messages and IEs defined in TS 29.244 V16.7.0 (2021-04) has been done, but the exported APIs may still be updated in the future (we add a new tag in that case).

We are now working on implementing networking functionalities (like setting up associations, establish sessions with easy & quick APIs), as well as updating the messages and IE definitions according to the latest specifications.

## Getting Started

### Installation

go-pfcp supports Go Modules. Run `go mod tidy` in your project's directory to collect the required packages automatically. See [Features](#features) for how to handle messages and IEs, and what is supported.

### Running examples

#### Exchanging Heartbeat

Small heartbeat client and server examples are available under [examples/heartheat](./examples/heartheat/).
The client sends a Heartbeat Request, and the server responds to it with a Heartbeat Response.

1. Run the server

By default, the server listens on loopback. The address/port can be specified explicitly by the `-s` flag.

```shell-session
go-pfcp/examples/heartbeat/hb-server$ go run main.go
2019/12/22 20:03:31 waiting for messages to come on: 127.0.0.2:8805
```

2. Run the client

By default, the client sends a Heartbeat to loopback. The address/port can be specified explicitly by the `-s` flag.
It exits after printing the timestamp in the received Heartbeat Response.

The output should look like below on the client side:

```shell-session
go-pfcp/examples/heartbeat/hb-client$ go run main.go
2019/12/22 20:03:36 sent Heartbeat Request to: 127.0.0.2:8805
2019/12/22 20:03:36 got Heartbeat Response with TS: 2019-12-22 20:03:36 +0900 JST, from: 127.0.0.2:8805
go-pfcp/examples/heartbeat/hb-client$
go-pfcp/examples/heartbeat/hb-client$ go run main.go
2019/12/22 20:03:40 sent Heartbeat Request to: 127.0.0.2:8805
2019/12/22 20:03:40 got Heartbeat Response with TS: 2019-12-22 20:03:40 +0900 JST, from: 127.0.0.2:8805
```

On the server side, the output should look like below:

```shell-session
go-pfcp/examples/heartbeat/hb-server$ go run main.go
2019/12/22 20:03:31 waiting for messages to come on: 127.0.0.2:8805
2019/12/22 20:03:36 got Heartbeat Request with TS: 2019-12-22 20:03:36 +0900 JST, from: 127.0.0.1:47305
2019/12/22 20:03:36 sent Heartbeat Response to: 127.0.0.1:47305
2019/12/22 20:03:36 waiting for messages to come on: 127.0.0.2:8805
2019/12/22 20:03:40 got Heartbeat Request with TS: 2019-12-22 20:03:40 +0900 JST, from: 127.0.0.1:55395
2019/12/22 20:03:40 sent Heartbeat Response to: 127.0.0.1:55395
2019/12/22 20:03:40 waiting for messages to come on: 127.0.0.2:8805
^Csignal: interrupt
```

_The server continues to listen on the same port until it is terminated (by Ctrl-C or something)._

## Features

### Messages

All the messages implements the same interface: `message.Message`, and have their own structs named `<MessageName>`.

#### Creating a message and encoding it

To create a message and encode it into binary, use `New<MessageName>()` function, which returns a `*<MessageName>` struct. The parameters for each constructor varies depending on the message type, but in general the first parameter is the sequence number, the second parameter is a SEID the message is session-related one, and the rest are the IEs contained in the message. See the godoc for the details.

_Handling of IEs is described in the [Information Elements](#information-elements) section._

```go
assocSetupReq := message.NewAssociationSetupRequest(
	sequenceNumber,
	ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
	ie.NewRecoveryTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
	ie.NewUPFunctionFeatures(0x01, 0x02),
	ie.NewCPFunctionFeatures(0x3f),
	ie.NewAlternativeSMFIPAddress(net.ParseIP("127.0.0.1"), net.ParseIP("2001::1")),
	ie.NewSMFSetID("go-pfcp.epc.3gppnetwork.org"),
	ie.NewPFCPSessionRetentionInformation(
		ie.NewCPPFCPEntityIPAddress(net.ParseIP("127.0.0.1"), nil),
	),
	ie.NewUEIPAddressPoolInformation(
		ie.NewUEIPAddressPoolIdentity("go-pfcp"),
		ie.NewNetworkInstance("some.instance.example"),
	),
	ie.NewGTPUPathQoSControlInformation(
		ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
		ie.NewGTPUPathInterfaceType(1, 1),
		ie.NewQoSReportTrigger(1, 1, 1),
		ie.NewTransportLevelMarking(0x1111),
		ie.NewMeasurementMethod(1, 1, 1),
		ie.NewAveragePacketDelay(10*time.Second),
		ie.NewMinimumPacketDelay(10*time.Second),
		ie.NewMaximumPacketDelay(10*time.Second),
		ie.NewTimer(20*time.Hour),
	),
	ie.NewClockDriftControlInformation(
		ie.NewRequestedClockDriftInformation(1, 1),
		ie.NewTSNTimeDomainNumber(255),
		ie.NewTimeOffsetThreshold(10*time.Second),
		ie.NewCumulativeRateRatioThreshold(0xffffffff),
	),
	ie.NewNFInstanceID([]byte{0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x33, 0x33, 0x33, 0x33, 0x44, 0x44, 0x44, 0x44}),
)
```

This will create a `*AssociationSetupRequest` struct with the given sequence number and IEs. The fields of the struct vary depending on the message type, but in general it has a common `*Header` at the top, and the list of IEs that have no explicit field at the bottom. This means that you can add any type of IEs to any type of messages, even if it is not supported in this library.

```go
type AssociationSetupRequest struct {
	*Header
	NodeID                          *ie.IE
	RecoveryTimeStamp               *ie.IE
	UPFunctionFeatures              *ie.IE
	CPFunctionFeatures              *ie.IE
	UserPlaneIPResourceInformation  []*ie.IE
	AlternativeSMFIPAddress         []*ie.IE
	SMFSetID                        *ie.IE
	PFCPSessionRetentionInformation *ie.IE
	UEIPAddressPoolInformation      []*ie.IE
	GTPUPathQoSControlInformation   []*ie.IE
	ClockDriftControlInformation    []*ie.IE
	UPFInstanceID                   *ie.IE
	PFCPASReqFlags                  *ie.IE
	IEs                             []*ie.IE
}
```

All the `<MessageName>` struct, or the `message.Message` interface, has the `Marshal()` method which returns the binary in `[]byte`. This can be used to send the message to the peer on top of any transport protocol (typically UDP) using the Go's standard `net` library.

```go
// serialize the message
b, err := assocSetupReq.Marshal()
if err != nil {
	// handle error
}

// send `b` to "127.0.0.1:8805"
raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8805")
if err != nil {
	// handle error
}

conn, err := net.DialUDP("udp", nil, raddr)
if err != nil {
	// handle error
}

if _, err := conn.Write(b); err != nil {
	// handle error
}

log.Printf("sent %s to %s", assocSetupReq.MessageTypeName(), raddr)
```

Is the message type you want to create not supported in this library? No problem. You can still create a message of any type by using `NewGeneric()` function. This function takes the message type as the first parameter, and the rest are the same as the `New<MessageName>()` function.

```go
// create a message of type 0x64
yourMessage := message.NewGeneric(
	0x64,
	sequenceNumber,
	seid,
	ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
	ie.NewRecoveryTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// ...
)
```

The created `yourMessage` is of type `*message.Generic`, which is a struct that implements the `message.Message` interface and has the common `*Header` and the list of IEs. You can still call the `Marshal()` method on this struct to get the binary.

#### Decoding a received message

To decode a received message, use `message.Parse()` function. This returns a message in the `message.Message` interface and an error.

```go
// receive a message on a UDP connection
b := make([]byte, 1500)
n, raddr, err := conn.ReadFromUDP(b)
if err != nil {
	// handle error
}

// decode the message
// NOTE: when you do this in another goroutine, copy the `b` to another slice
// as the `b` will be overwritten by the next `ReadFromUDP()` call.
msg, err := message.Parse(b[:n])
if err != nil {
	// handle error
}

log.Printf("got %s from %s", msg.MessageTypeName(), raddr)
```

To access the fields of the message, you need to assert the type of the message to the corresponding struct. For example, to access IEs in the `AssociationSetupResponse` message, you need to assert the type of the message to `*AssociationSetupResponse` first.


```go
// assert the type of the message
assocSetupRes, ok := msg.(*message.AssociationSetupResponse)
if !ok {
	// handle error
}

// get the value of NodeID IE
// see the Information Elements section for what `NodeID()` method does
nodeID, err := assocSetupRes.NodeID.NodeID()
if err != nil {
	// handle error
}
```

If the message type is not supported in this library, it can still be asserted to `*Generic` whose header and IEs are accessible through the `Header` and `IEs` fields.

```go
// assert the type of the message
unknownMsg, ok := msg.(*message.Generic)
if !ok {
	// handle error
}

// iterate over the IEs to get the value of NodeID IE
for _, i := range unknownMsg.IEs {
	switch i.Type {
	case ie.NodeID:
		nodeID, err := i.NodeID()
		if err != nil {
			// handle error
		}
	case ie.SomeOtherIE:
		// ...
	}
}
```

So, what you will be likely to do is to check the type of the message and handle it accordingly using the `switch` statement.

```go
switch m := msg.(type) {
case *message.AssociationSetupRequest:
	// handle AssociationSetupRequest
case *message.AssociationSetupResponse:
	// handle AssociationSetupResponse
case *message.SomeOtherMessage:
	// ...
}

// alternatively, you can use the `MessageType()` method
switch msg.MessageType() {
case message.MsgTypeAssociationSetupRequest:
	// handle AssociationSetupRequest
case message.MsgTypeAssociationSetupResponse:
	// handle AssociationSetupResponse
case message.MsgTypeSomeOtherMessage:
	// ...
}

// using a map containing the handlers for each message type may also be a good idea
// NOTE: use with care, as there may be a race condition when accessing the map
handlers := map[uint8]func(message.Message) error{
	message.MsgTypeAssociationSetupRequest: func(m message.Message) error {
		assocSetupReq := m.(*message.AssociationSetupRequest)
		// handle AssociationSetupRequest
	},
	message.MsgTypeAssociationSetupResponse: func(m message.Message) error {
		assocSetupRes := m.(*message.AssociationSetupResponse)
		// handle AssociationSetupResponse
	},
	// ...
}

handle, ok := handlers[msg.MessageType()]
if !ok {
	// handle unsupported message type
} else {
	if err := handle(msg); err != nil {
		// handle error
	}
}
```

#### List of supported messages

Messages are implemented in conformance with TS 29.244 V16.7.0 (2021-04). The word "supported" in the table below means that the struct and the constructor for the message are implemented in this library. As described in the previous section, you can still create a message of any type eve if it is not supported or missing in the table.

##### PFCP Node related messages

| Message Type | Message                        | Sxa | Sxb | Sxc | N4  | Supported? |
| ------------ | ------------------------------ | --- | --- | --- | --- | ---------- |
| 1            | Heartbeat Request              | X   | X   | X   | X   | Yes        |
| 2            | Heartbeat Response             | X   | X   | X   | X   | Yes        |
| 3            | PFD Management Request         | -   | X   | X   | X   | Yes        |
| 4            | PFD Management Response        | -   | X   | X   | X   | Yes        |
| 5            | Association Setup Request      | X   | X   | X   | X   | Yes        |
| 6            | Association Setup Response     | X   | X   | X   | X   | Yes        |
| 7            | Association Update Request     | X   | X   | X   | X   | Yes        |
| 8            | Association Update Response    | X   | X   | X   | X   | Yes        |
| 9            | Association Release Request    | X   | X   | X   | X   | Yes        |
| 10           | Association Release Response   | X   | X   | X   | X   | Yes        |
| 11           | Version Not Supported Response | X   | X   | X   | X   | Yes        |
| 12           | Node Report Request            | X   | X   | X   | X   | Yes        |
| 13           | Node Report Response           | X   | X   | X   | X   | Yes        |
| 14           | Session Set Deletion Request   | X   | X   | -   |     | Yes        |
| 15           | Session Set Deletion Response  | X   | X   | -   |     | Yes        |
| 16 to 49     | _(For future use)_             |     |     |     |     | -          |

##### PFCP Session related messages

| Message Type | Message                        | Sxa | Sxb | Sxc | N4  | Supported? |
| ------------ | ------------------------------ | --- | --- | --- | --- | ---------- |
| 50           | Session Establishment Request  | X   | X   | X   | X   | Yes        |
| 51           | Session Establishment Response | X   | X   | X   | X   | Yes        |
| 52           | Session Modification Request   | X   | X   | X   | X   | Yes        |
| 53           | Session Modification Response  | X   | X   | X   | X   | Yes        |
| 54           | Session Deletion Request       | X   | X   | X   | X   | Yes        |
| 55           | Session Deletion Response      | X   | X   | X   | X   | Yes        |
| 56           | Session Report Request         | X   | X   | X   | X   | Yes        |
| 57           | Session Report Response        | X   | X   | X   | X   | Yes        |
| 58 to 99     | _(For future use)_             |     |     |     |     | -          |

### Information Elements

All the IEs are of the same type: `ie.IE`.

#### Creating IEs

Constructors are available for every type of _supported_ IEs (see [Supported IEs](#list-of-supported-ies) for which are supported).  
To create a `CreatePDR` IE, use `NewCreatePDR()`. Parameters for those "named" constructors are implemented as friendly as possible for the built-in Go types.

```go
// returned type and parameters in NewCreatePDR() are all *ie.IE.
createPDR := ie.NewCreatePDR(
	ie.NewPDRID(0xffff),
	ie.NewPrecedence(0x11111111),
	ie.NewPDI(
		ie.NewSourceInterface(ie.SrcInterfaceAccess),
		ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
		ie.NewNetworkInstance("some.instance.example"),
		ie.NewRedundantTransmissionParametersInPDI(
			ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
			ie.NewNetworkInstance("some.instance.example"),
		),
		ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0),
		ie.NewTrafficEndpointID(0x01),
		ie.NewSDFFilter("aaaaaaaa", "bb", "cccc", "ddd", 0xffffffff),
		ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
		ie.NewEthernetPDUSessionInformation(0x01),
		ie.NewEthernetPacketFilter(
			ie.NewEthernetFilterID(0xffffffff),
			ie.NewEthernetFilterProperties(0x01),
			ie.NewMACAddress(mac1, mac2, mac3, mac4),
			ie.NewEthertype(0xffff),
			ie.NewCTAG(0x07, 1, 1, 4095),
			ie.NewSTAG(0x07, 1, 1, 4095),
			ie.NewSDFFilter("aaaaaaaa", "bb", "cccc", "ddd", 0xffffffff),
		),
	),
	ie.NewOuterHeaderRemoval(0x01, 0x02),
	ie.NewFARID(0xffffffff),
	ie.NewURRID(0xffffffff),
	ie.NewQERID(0xffffffff),
	ie.NewActivatePredefinedRules("go-pfcp"),
	ie.NewActivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
	ie.NewDeactivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
	ie.NewMARID(0x1111),
	ie.NewPacketReplicationAndDetectionCarryOnInformation(0x0f),
	ie.NewIPMulticastAddressingInfo(
		ie.NewIPMulticastAddress(net.ParseIP("127.0.0.1"), nil, net.ParseIP("127.0.0.1"), nil),
		ie.NewSourceIPAddress(net.ParseIP("127.0.0.1"), nil, 24),
	),
	ie.NewUEIPAddressPoolIdentity("go-pfcp"),
)
```

Instead of using the named constructors, you can also create an arbitrary IE of any type with the `SomeType` value where `SomeType` is any of the `Uint8`, `Uint16`, `Uint32`, `Uint64`, `String`, or `FQDN`. This is useful when you don't want to use the named constructors which require the specific type of parameters, or when you want to create an IE that is not supported in this library.

For example, the following two codes will create the same `*ie.IE` instance.

```go
qvTime := ie.NewQuotaValidityTime(10 * time.Second)
```

```go
qvTime := ie.NewUint32(ie.QuotaValidityTime, 0x0000000a)
```

Sometimes you may have a value that is already in `uint32` format. In that case, the latter way of creating an IE is preferable, as converting the `uint32` value to `time.Duration` is not very efficient (compiler optimization possibly helps in the former case? I don't know, but it doesn't _look_ good for humans anyway).

Alternatively, you may want to create an IE with the undecoded or already encoded binary (in `[]byte`), which can be done using `ie.New()`.

```go
qvTime := ie.New(ie.QuotaValidityTime, []byte{0x00, 0x00, 0x00, 0x0a})
```

When you want to create an unsupported IE whose IE type is 999 and the value is of `uint16` type, you can do like below.
If it is vendor-specific, you can also specify the enterprise ID.

```go
ie999 := ie.NewUint16(999, 0x0102)
ie999.EnterpriseID = 0x1234
```

This can instead be done using `ie.NewUint16VendorSpecific()`. There are no type-specific constructors for vendor-specific IEs.

```go
ie999 := ie.NewVendorSpecificIE(999, 0x1234, []byte{0x01, 0x02})
```

#### Retrieving values from IEs

To retrieve values from an IE, you can call helper methods that have the same name as the IE itself on an `*ie.IE`. For example, you can get the value of a `NetworkInstance` IE by calling the `NetworkInstance()` method.

```go
ni := ie.NewNetworkInstance("some.instance.example")
v, err := ni.NetworkInstance()
```

In the example above, calling the `NetworkInstance()` method returns `"some.instance.example", nil`. However, if `ni` is not of type `ie.NetworkInstance` or the payload is not in right format, it returns `"", SomeError`. It's important to always check the returned error, as the value may be empty _as expected_.

The type of returned values are determined as friendly as possible for the built-in Go types. For example, `QuotaValidityTime()` returns the value as `time.Duration`. In the cases that this is undesirable, such as when you need to pass the raw value to another node, you can use `ValueAsUint32()` method instead, which returns the value as `uint32`. This could also be useful for handling vendor-specific IEs whose value retrieval methods are not available in this library. Available `ValueAs`-methods are `ValueAsUin8()`, `ValueAsUint16()`, `ValueAsUint32()`, `ValueAsUint64()`, `ValueAsString()`, and `ValueAsFQDN()`.

```go
qvTime := ie.NewQuotaValidityTime(10 * time.Second)

vDuration, err := qvTime.QuotaValidityTime()
vUint32, err := qvTime.ValueAsUint32()
```

For IEs with more complex payloads, such as F-TEID, calling the `<IE-name>` method returns a `<IE-name>Fields` struct containing the values in its fields.

```go
fteid := ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, nil)
fteidFields, err := fteid.FTEID() // `FTEIDFields` struct

teid := fteidFields.TEID      // TEID as uint32
v4 := fteidFields.IPv4Address // IPv4 address as net.IP
```

For grouped IEs, accesing the `ChildIEs` field is the best way to retrieve the list of IEs contained. If you are not sure that the IE is already parsed, you can use `ValueAsGrouped()` method instead, which parses the payload into `[]*IE` and returns it if the `ChildIEs` field is empty.

```go
cpdrIE := ie.NewCreatePDR(
	ie.NewPDRID(0xffff),
	// ...
)

// most efficient way
cpdrChildren := cpdrIE.ChildIEs
// or use this if you are not sure that the IE is already parsed
cpdrChildren, err := cpdrIE.ValueAsGrouped()
```

_NOTE: if you have called `ie.Parse`, the child IEs are already parsed._

_To determine if an IE is grouped or not, this library uses the `defaultGroupedIEMap` in `ie_grouped.go`, which contains the list of grouped IEs. You can add your own IE type to this map using `ie.AddGroupedIEType()` function, or you can change the entire logic to determine if an IE is grouped or not by setting your own function to `ie.SetIsGroupedFun` function._

`<IE-name>` method is also available for consistency with non-grouped IEs, but it is not recommended to use it as it always parses the payload into `[]*IE` and returns it though the `ChildIEs` field is already populated. In the rare case that the payload can be modified after the IE is created or parsed, this method could be useful.

```go
cpdrChildren, err := cpdrIE.CreatePDR()
```

For convenience, helper methods of the child IEs can be called directly on the grouped IE, or you can call `FindByType()` method to get the child IE of the specified type.

```go
pdrID, err := cpdrIE.PDRID()

// or
pdrID, err := cpdrIE.FindByType(ie.PDRID)
```

This is useful when you only need a small number of IEs within a grouped IE. Since it iterates over all child IEs internally, retrieving everything using these method is very inefficient. When you need to retrieve the value of multiple IEs, it is recommended to iterate over the list of child IEs in your own code.

```go
var (
	pdrID uint16
	farID uint32
	// ...
)

for _, i := range cpdrIE.ChildIEs {
	switch i.Type {
	case ie.PDRID:
		v, err = i.PDRID()
		if err != nil {
			// handle error
		}
		pdrID = v
	case ie.FARID:
		v, err = i.FARID()
		if err != nil {
			// handle error
		}
		farID = v
	case ie.SomeOtherIE:
		// ...
	}
}
```

#### List of supported IEs

IEs are implemented in conformance with TS 29.244 V16.7.0 (2021-04). The word "supported" in the table below means that the constructor and helper method for the IE are implemented in this library. As described in the previous section, you can still create an IE of any type even if it is not supported or missing in the table.

| IE Type        | Information elements                                                       | Supported? |
| -------------- | -------------------------------------------------------------------------- | ---------- |
| 0              | _(Reserved)_                                                               | -          |
| 1              | Create PDR                                                                 | Yes        |
| 2              | PDI                                                                        | Yes        |
| 3              | Create FAR                                                                 | Yes        |
| 4              | Forwarding Parameters                                                      | Yes        |
| 5              | Duplicating Parameters                                                     | Yes        |
| 6              | Create URR                                                                 | Yes        |
| 7              | Create QER                                                                 | Yes        |
| 8              | Created PDR                                                                | Yes        |
| 9              | Update PDR                                                                 | Yes        |
| 10             | Update FAR                                                                 | Yes        |
| 11             | Update Forwarding Parameters                                               | Yes        |
| 12             | Update BAR (PFCP Session Report Response)                                  | Yes        |
| 13             | Update URR                                                                 | Yes        |
| 14             | Update QER                                                                 | Yes        |
| 15             | Remove PDR                                                                 | Yes        |
| 16             | Remove FAR                                                                 | Yes        |
| 17             | Remove URR                                                                 | Yes        |
| 18             | Remove QER                                                                 | Yes        |
| 19             | Cause                                                                      | Yes        |
| 20             | Source Interface                                                           | Yes        |
| 21             | F-TEID                                                                     | Yes        |
| 22             | Network Instance                                                           | Yes        |
| 23             | SDF Filter                                                                 | Yes        |
| 24             | Application ID                                                             | Yes        |
| 25             | Gate Status                                                                | Yes        |
| 26             | MBR                                                                        | Yes        |
| 27             | GBR                                                                        | Yes        |
| 28             | QER Correlation ID                                                         | Yes        |
| 29             | Precedence                                                                 | Yes        |
| 30             | Transport Level Marking                                                    | Yes        |
| 31             | Volume Threshold                                                           | Yes        |
| 32             | Time Threshold                                                             | Yes        |
| 33             | Monitoring Time                                                            | Yes        |
| 34             | Subsequent Volume Threshold                                                | Yes        |
| 35             | Subsequent Time Threshold                                                  | Yes        |
| 36             | Inactivity Detection Time                                                  | Yes        |
| 37             | Reporting Triggers                                                         | Yes        |
| 38             | Redirect Information                                                       | Yes        |
| 39             | Report Type                                                                | Yes        |
| 40             | Offending IE                                                               | Yes        |
| 41             | Forwarding Policy                                                          | Yes        |
| 42             | Destination Interface                                                      | Yes        |
| 43             | UP Function Features                                                       | Yes        |
| 44             | Apply Action                                                               | Yes        |
| 45             | Downlink Data Service Information                                          | Yes        |
| 46             | Downlink Data Notification Delay                                           | Yes        |
| 47             | DL Buffering Duration                                                      | Yes        |
| 48             | DL Buffering Suggested Packet Count                                        | Yes        |
| 49             | PFCPSMReq-Flags                                                            | Yes        |
| 50             | PFCPSRRsp-Flags                                                            | Yes        |
| 51             | Load Control Information                                                   | Yes        |
| 52             | Sequence Number                                                            | Yes        |
| 53             | Metric                                                                     | Yes        |
| 54             | Overload Control Information                                               | Yes        |
| 55             | Timer                                                                      | Yes        |
| 56             | Packet Detection Rule ID                                                   | Yes        |
| 57             | F-SEID                                                                     | Yes        |
| 58             | Application ID's PFDs                                                      | Yes        |
| 59             | PFD context                                                                | Yes        |
| 60             | Node ID                                                                    | Yes        |
| 61             | PFD contents                                                               | Yes        |
| 62             | Measurement Method                                                         | Yes        |
| 63             | Usage Report Trigger                                                       | Yes        |
| 64             | Measurement Period                                                         | Yes        |
| 65             | FQ-CSID                                                                    | Yes        |
| 66             | Volume Measurement                                                         | Yes        |
| 67             | Duration Measurement                                                       | Yes        |
| 68             | Application Detection Information                                          | Yes        |
| 69             | Time of First Packet                                                       | Yes        |
| 70             | Time of Last Packet                                                        | Yes        |
| 71             | Quota Holding Time                                                         | Yes        |
| 72             | Dropped DL Traffic Threshold                                               | Yes        |
| 73             | Volume Quota                                                               | Yes        |
| 74             | Time Quota                                                                 | Yes        |
| 75             | Start Time                                                                 | Yes        |
| 76             | End Time                                                                   | Yes        |
| 77             | Query URR                                                                  | Yes        |
| 78             | Usage Report (Session Modification Response)                               | Yes        |
| 79             | Usage Report (Session Deletion Response)                                   | Yes        |
| 80             | Usage Report (Session Report Request)                                      | Yes        |
| 81             | URR ID                                                                     | Yes        |
| 82             | Linked URR ID                                                              | Yes        |
| 83             | Downlink Data Report                                                       | Yes        |
| 84             | Outer Header Creation                                                      | Yes        |
| 85             | Create BAR                                                                 | Yes        |
| 86             | Update BAR (Session Modification Request)                                  | Yes        |
| 87             | Remove BAR                                                                 | Yes        |
| 88             | BAR ID                                                                     | Yes        |
| 89             | CP Function Features                                                       | Yes        |
| 90             | Usage Information                                                          | Yes        |
| 91             | Application Instance ID                                                    | Yes        |
| 92             | Flow Information                                                           | Yes        |
| 93             | UE IP Address                                                              | Yes        |
| 94             | Packet Rate                                                                | Yes        |
| 95             | Outer Header Removal                                                       | Yes        |
| 96             | Recovery Time Stamp                                                        | Yes        |
| 97             | DL Flow Level Marking                                                      | Yes        |
| 98             | Header Enrichment                                                          | Yes        |
| 99             | Error Indication Report                                                    | Yes        |
| 100            | Measurement Information                                                    | Yes        |
| 101            | Node Report Type                                                           | Yes        |
| 102            | User Plane Path Failure Report                                             | Yes        |
| 103            | Remote GTP-U Peer                                                          | Yes        |
| 104            | UR-SEQN                                                                    | Yes        |
| 105            | Update Duplicating Parameters                                              | Yes        |
| 106            | Activate Predefined Rules                                                  | Yes        |
| 107            | Deactivate Predefined Rules                                                | Yes        |
| 108            | FAR ID                                                                     | Yes        |
| 109            | QER ID                                                                     | Yes        |
| 110            | OCI Flags                                                                  | Yes        |
| 111            | PFCP Association Release Request                                           | Yes        |
| 112            | Graceful Release Period                                                    | Yes        |
| 113            | PDN Type                                                                   | Yes        |
| 114            | Failed Rule ID                                                             | Yes        |
| 115            | Time Quota Mechanism                                                       | Yes        |
| 116            | User Plane IP Resource Information                                         | Yes        |
| 117            | User Plane Inactivity Timer                                                | Yes        |
| 118            | Aggregated URRs                                                            | Yes        |
| 119            | Multiplier                                                                 | Yes        |
| 120            | Aggregated URR ID                                                          | Yes        |
| 121            | Subsequent Volume Quota                                                    | Yes        |
| 122            | Subsequent Time Quota                                                      | Yes        |
| 123            | RQI                                                                        | Yes        |
| 124            | QFI                                                                        | Yes        |
| 125            | Query URR Reference                                                        | Yes        |
| 126            | Additional Usage Reports Information                                       | Yes        |
| 127            | Create Traffic Endpoint                                                    | Yes        |
| 128            | Created Traffic Endpoint                                                   | Yes        |
| 129            | Update Traffic Endpoint                                                    | Yes        |
| 130            | Remove Traffic Endpoint                                                    | Yes        |
| 131            | Traffic Endpoint ID                                                        | Yes        |
| 132            | Ethernet Packet Filter                                                     | Yes        |
| 133            | MAC address                                                                | Yes        |
| 134            | C-TAG                                                                      | Yes        |
| 135            | S-TAG                                                                      | Yes        |
| 136            | Ethertype                                                                  | Yes        |
| 137            | Proxying                                                                   | Yes        |
| 138            | Ethernet Filter ID                                                         | Yes        |
| 139            | Ethernet Filter Properties                                                 | Yes        |
| 140            | Suggested Buffering Packets Count                                          | Yes        |
| 141            | User ID                                                                    | Yes        |
| 142            | Ethernet PDU Session Information                                           | Yes        |
| 143            | Ethernet Traffic Information                                               | Yes        |
| 144            | MAC Addresses Detected                                                     | Yes        |
| 145            | MAC Addresses Removed                                                      | Yes        |
| 146            | Ethernet Inactivity Timer                                                  | Yes        |
| 147            | Additional Monitoring Time                                                 | Yes        |
| 148            | Event Quota                                                                | Yes        |
| 149            | Event Threshold                                                            | Yes        |
| 150            | Subsequent Event Quota                                                     | Yes        |
| 151            | Subsequent Event Threshold                                                 | Yes        |
| 152            | Trace Information                                                          | Yes        |
| 153            | Framed-Route                                                               | Yes        |
| 154            | Framed-Routing                                                             | Yes        |
| 155            | Framed-IPv6-Route                                                          | Yes        |
| 156            | Event Time Stamp                                                           | Yes        |
| 157            | Averaging Window                                                           | Yes        |
| 158            | Paging Policy Indicator                                                    | Yes        |
| 159            | APN/DNN                                                                    | Yes        |
| 160            | 3GPP Interface Type                                                        | Yes        |
| 161            | PFCPSRReq-Flags                                                            | Yes        |
| 162            | PFCPAUReq-Flags                                                            | Yes        |
| 163            | Activation Time                                                            | Yes        |
| 164            | Deactivation Time                                                          | Yes        |
| 165            | Create MAR                                                                 | Yes        |
| 166            | 3GPP Access Forwarding Action Information                                  | Yes        |
| 167            | Non-3GPP Access Forwarding Action Information                              | Yes        |
| 168            | Remove MAR                                                                 | Yes        |
| 169            | Update MAR                                                                 | Yes        |
| 170            | MAR ID                                                                     | Yes        |
| 171            | Steering Functionality                                                     | Yes        |
| 172            | Steering Mode                                                              | Yes        |
| 173            | Weight                                                                     | Yes        |
| 174            | Priority                                                                   | Yes        |
| 175            | Update 3GPP Access Forwarding Action Information                           | Yes        |
| 176            | Update Non 3GPP Access Forwarding Action Information                       | Yes        |
| 177            | UE IP address Pool Identity                                                | Yes        |
| 178            | Alternative SMF IP Address                                                 | Yes        |
| 179            | Packet Replication and Detection Carry-On Information                      | Yes        |
| 180            | SMF Set ID                                                                 | Yes        |
| 181            | Quota Validity Time                                                        | Yes        |
| 182            | Number of Reports                                                          | Yes        |
| 183            | PFCP Session Retention Information (within PFCP Association Setup Request) | Yes        |
| 184            | PFCPASRsp-Flags                                                            | Yes        |
| 185            | CP PFCP Entity IP Address                                                  | Yes        |
| 186            | PFCPSEReq-Flags                                                            | Yes        |
| 187            | User Plane Path Recovery Report                                            | Yes        |
| 188            | IP Multicast Addressing Info within PFCP Session Establishment Request     | Yes        |
| 189            | Join IP Multicast Information IE within Usage Report                       | Yes        |
| 190            | Leave IP Multicast Information IE within Usage Report                      | Yes        |
| 191            | IP Multicast Address                                                       | Yes        |
| 192            | Source IP Address                                                          | Yes        |
| 193            | Packet Rate Status                                                         | Yes        |
| 194            | Create Bridge Info for TSC                                                 | Yes        |
| 195            | Created Bridge Info for TSC                                                | Yes        |
| 196            | DS-TT Port Number                                                          | Yes        |
| 197            | NW-TT Port Number                                                          | Yes        |
| 198            | TSN Bridge ID                                                              | Yes        |
| 199            | TSC Management Information IE within PFCP Session Modification Request     | Yes        |
| 200            | TSC Management Information IE within PFCP Session Modification Response    | Yes        |
| 201            | TSC Management Information IE within PFCP Session Report Request           | Yes        |
| 202            | Port Management Information Container                                      | Yes        |
| 203            | Clock Drift Control Information                                            | Yes        |
| 204            | Requested Clock Drift Information                                          | Yes        |
| 205            | Clock Drift Report                                                         | Yes        |
| 206            | TSN Time Domain Number                                                     | Yes        |
| 207            | Time Offset Threshold                                                      | Yes        |
| 208            | Cumulative rateRatio Threshold                                             | Yes        |
| 209            | Time Offset Measurement                                                    | Yes        |
| 210            | Cumulative rateRatio Measurement                                           | Yes        |
| 211            | Remove SRR                                                                 | Yes        |
| 212            | Create SRR                                                                 | Yes        |
| 213            | Update SRR                                                                 | Yes        |
| 214            | Session Report                                                             | Yes        |
| 215            | SRR ID                                                                     | Yes        |
| 216            | Access Availability Control Information                                    | Yes        |
| 217            | Requested Access Availability Information                                  | Yes        |
| 218            | Access Availability Report                                                 | Yes        |
| 219            | Access Availability Information                                            | Yes        |
| 220            | Provide ATSSS Control Information                                          | Yes        |
| 221            | ATSSS Control Parameters                                                   | Yes        |
| 222            | MPTCP Control Information                                                  | Yes        |
| 223            | ATSSS-LL Control Information                                               | Yes        |
| 224            | PMF Control Information                                                    | Yes        |
| 225            | MPTCP Parameters                                                           | Yes        |
| 226            | ATSSS-LL Parameters                                                        | Yes        |
| 227            | PMF Parameters                                                             | Yes        |
| 228            | MPTCP Address Information                                                  | Yes        |
| 229            | UE Link-Specific IP Address                                                | Yes        |
| 230            | PMF Address Information                                                    | Yes        |
| 231            | ATSSS-LL Information                                                       | Yes        |
| 232            | Data Network Access Identifier                                             | Yes        |
| 233            | UE IP address Pool Information                                             | Yes        |
| 234            | Average Packet Delay                                                       | Yes        |
| 235            | Minimum Packet Delay                                                       | Yes        |
| 236            | Maximum Packet Delay                                                       | Yes        |
| 237            | QoS Report Trigger                                                         | Yes        |
| 238            | GTP-U Path QoS Control Information                                         | Yes        |
| 239            | GTP-U Path QoS Report (PFCP Node Report Request)                           | Yes        |
| 240            | QoS Information in GTP-U Path QoS Report                                   | Yes        |
| 241            | GTP-U Path Interface Type                                                  | Yes        |
| 242            | QoS Monitoring per QoS flow Control Information                            | Yes        |
| 243            | Requested QoS Monitoring                                                   | Yes        |
| 244            | Reporting Frequency                                                        | Yes        |
| 245            | Packet Delay Thresholds                                                    | Yes        |
| 246            | Minimum Wait Time                                                          | Yes        |
| 247            | QoS Monitoring Report                                                      | Yes        |
| 248            | QoS Monitoring Measurement                                                 | Yes        |
| 249            | MT-EDT Control Information                                                 | Yes        |
| 250            | DL Data Packets Size                                                       | Yes        |
| 251            | QER Control Indications                                                    | Yes        |
| 252            | Packet Rate Status Report                                                  | Yes        |
| 253            | NF Instance ID                                                             | Yes        |
| 254            | Ethernet Context Information                                               | Yes        |
| 255            | Redundant Transmission Parameters                                          | Yes        |
| 256            | Updated PDR                                                                | Yes        |
| 257            | S-NSSAI                                                                    | Yes        |
| 258            | IP version                                                                 | Yes        |
| 259            | PFCPASReq-Flags                                                            | Yes        |
| 260            | Data Status                                                                | Yes        |
| 261            | Provide RDS configuration information                                      | Yes        |
| 262            | RDS configuration information                                              | Yes        |
| 263            | Query Packet Rate Status IE within PFCP Session Modification Request       | Yes        |
| 264            | Packet Rate Status Report IE within PFCP Session Modification Response     | Yes        |
| 265            | MPTCP Applicable Indication                                                | Yes        |
| 266            | Bridge Management Information Container                                    | Yes        |
| 267            | UE IP Address Usage Information                                            | Yes        |
| 268            | Number of UE IP Addresses                                                  | Yes        |
| 269            | Validity Timer                                                             | Yes        |
| 270            | Redundant Transmission Forwarding Parameters                               | Yes        |
| 271            | Transport Delay Reporting                                                  | Yes        |
| 272 to 32767   | _(For future use)_                                                         | -          |
| 32768 to 65535 | Reserved for vendor specific IEs                                           | -          |

## Author(s)

Yoshiyuki Kurauchi ([Website](https://wmnsk.com/))

_I'm always open to welcome co-authors! Please feel free to talk to me._

## LICENSE

[MIT](https://github.com/wmnsk/go-pfcp/blob/master/LICENSE)
