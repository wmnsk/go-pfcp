# go-pfcp

PFCP implementation in Golang.

[![CircleCI](https://circleci.com/gh/wmnsk/go-pfcp.svg?style=shield)](https://circleci.com/gh/wmnsk/go-pfcp)
[![GolangCI](https://golangci.com/badges/github.com/wmnsk/go-pfcp.svg)](https://golangci.com/r/github.com/wmnsk/go-pfcp)
[![GoDoc](https://godoc.org/github.com/wmnsk/go-pfcp?status.svg)](https://godoc.org/github.com/wmnsk/go-pfcp)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/go-pfcp/blob/master/LICENSE)

## What is PFCP?

PFCP(Packet Forwarding Control Protocol) is a signaling protocol used in mobile networking infrastructure(LTE EPC, 5GC) to realize CUPS architecture(Control and User Plane Separation, not a printing system).

Looking for GTP implementation in Golang? [go-gtp](https://github.com/wmnsk/go-gtp) would be a good choice for you :P

## Project Status

This project is still WIP; I'm in the middle of implementing messages and IEs, networking functionalities(like setting up associations) are not available currently.

## Getting Started

### Installation

go-pfcp supports Go Modules. Just run `go mod tidy` in your project's directory to collect required packages automatically.

```shell-session
go mod tidy
```

If your Go version is older than 1.11, run `go get` manually the following packages instead.

```shell-session
go get -u github.com/google/go-cmp
go get -u github.com/pascaldekloe/goe
```

### Running examples

_(to be updated)_

## Supported Features

### Messages 

#### PFCP Node related messages

| Message Type | Message                        | Sxa | Sxb | Sxc | N4 | Supported? |
|--------------|--------------------------------|-----|-----|-----|----|------------|
| 1            | Heartbeat Request              | X   | X   | X   | X  | Yes        |
| 2            | Heartbeat Response             | X   | X   | X   | X  | Yes        |
| 3            | PFD Management Request         | -   | X   | X   | X  |            |
| 4            | PFD Management Response        | -   | X   | X   | X  |            |
| 5            | Association Setup Request      | X   | X   | X   | X  |            |
| 6            | Association Setup Response     | X   | X   | X   | X  |            |
| 7            | Association Update Request     | X   | X   | X   | X  |            |
| 8            | Association Update Response    | X   | X   | X   | X  |            |
| 9            | Association Release Request    | X   | X   | X   | X  |            |
| 10           | Association Release Response   | X   | X   | X   | X  |            |
| 11           | Version Not Supported Response | X   | X   | X   | X  |            |
| 12           | Node Report Request            | X   | X   | X   | X  |            |
| 13           | Node Report Response           | X   | X   | X   | X  |            |
| 14           | Session Set Deletion Request   | X   | X   | -   |    |            |
| 15           | Session Set Deletion Response  | X   | X   | -   |    |            |
| 16 to 49     | _(For future use)_             |     |     |     |    | -          |

#### PFCP Session related messages

| Message Type | Message                        | Sxa | Sxb | Sxc | N4 | Supported? |
|--------------|--------------------------------|-----|-----|-----|----|------------|
| 50           | Session Establishment Request  | X   | X   | X   | X  |            |
| 51           | Session Establishment Response | X   | X   | X   | X  |            |
| 52           | Session Modification Request   | X   | X   | X   | X  |            |
| 53           | Session Modification Response  | X   | X   | X   | X  |            |
| 54           | Session Deletion Request       | X   | X   | X   | X  |            |
| 55           | Session Deletion Response      | X   | X   | X   | X  |            |
| 56           | Session Report Request         | X   | X   | X   | X  |            |
| 57           | Session Report Response        | X   | X   | X   | X  |            |
| 58 to 99     | _(For future use)_             |     |     |     |    | -          |

### Information Elements


| IE Type      | Information elements                                  | Supported? |
|--------------|-------------------------------------------------------|------------|
| 0            | _(Reserved)_                                          | -          |
| 1            | Create PDR                                            |            |
| 2            | PDI                                                   |            |
| 3            | Create FAR                                            |            |
| 4            | Forwarding Parameters                                 |            |
| 5            | Duplicating Parameters                                |            |
| 6            | Create URR                                            |            |
| 7            | Create QER                                            |            |
| 8            | Created PDR                                           |            |
| 9            | Update PDR                                            |            |
| 10           | Update FAR                                            |            |
| 11           | Update Forwarding Parameters                          |            |
| 12           | Update BAR (Sx Session Report Response)               |            |
| 13           | Update URR                                            |            |
| 14           | Update QER                                            |            |
| 15           | Remove PDR                                            |            |
| 16           | Remove FAR                                            |            |
| 17           | Remove URR                                            |            |
| 18           | Remove QER                                            |            |
| 19           | Cause                                                 | Yes        |
| 20           | Source Interface                                      | Yes        |
| 21           | F-TEID                                                | Yes        |
| 22           | Network Instance                                      | Yes        |
| 23           | SDF Filter                                            |            |
| 24           | Application ID                                        | Yes        |
| 25           | Gate Status                                           | Yes        |
| 26           | MBR                                                   | Yes        |
| 27           | GBR                                                   | Yes        |
| 28           | QER Correlation ID                                    | Yes        |
| 29           | Precedence                                            | Yes        |
| 30           | Transport Level Marking                               | Yes        |
| 31           | Volume Threshold                                      | Yes        |
| 32           | Time Threshold                                        | Yes        |
| 33           | Monitoring Time                                       | Yes        |
| 34           | Subsequent Volume Threshold                           | Yes        |
| 35           | Subsequent Time Threshold                             | Yes        |
| 36           | Inactivity Detection Time                             | Yes        |
| 37           | Reporting Triggers                                    | Yes        |
| 38           | Redirect Information                                  | Yes        |
| 39           | Report Type                                           | Yes        |
| 40           | Offending IE                                          | Yes        |
| 41           | Forwarding Policy                                     | Yes        |
| 42           | Destination Interface                                 | Yes        |
| 43           | UP Function Features                                  |            |
| 44           | Apply Action                                          |            |
| 45           | Downlink Data Service Information                     |            |
| 46           | Downlink Data Notification Delay                      |            |
| 47           | DL Buffering Duration                                 |            |
| 48           | DL Buffering Suggested Packet Count                   |            |
| 49           | SxSMReq-Flags                                         |            |
| 50           | SxSRRsp-Flags                                         |            |
| 51           | Load Control Information                              |            |
| 52           | Sequence Number                                       |            |
| 53           | Metric                                                |            |
| 54           | Overload Control Information                          |            |
| 55           | Timer                                                 |            |
| 56           | Packet Detection Rule ID                              |            |
| 57           | F-SEID                                                |            |
| 58           | Application ID's PFDs                                 |            |
| 59           | PFD context                                           |            |
| 60           | Node ID                                               |            |
| 61           | PFD contents                                          |            |
| 62           | Measurement Method                                    |            |
| 63           | Usage Report Trigger                                  |            |
| 64           | Measurement Period                                    |            |
| 65           | FQ-CSID                                               |            |
| 66           | Volume Measurement                                    |            |
| 67           | Duration Measurement                                  |            |
| 68           | Application Detection Information                     |            |
| 69           | Time of First Packet                                  |            |
| 70           | Time of Last Packet                                   |            |
| 71           | Quota Holding Time                                    |            |
| 72           | Dropped DL Traffic Threshold                          |            |
| 73           | Volume Quota                                          |            |
| 74           | Time Quota                                            |            |
| 75           | Start Time                                            |            |
| 76           | End Time                                              |            |
| 77           | Query URR                                             |            |
| 78           | Usage Report (in Session Modification Response)       |            |
| 79           | Usage Report (Session Deletion Response)              |            |
| 80           | Usage Report (Session Report Request)                 |            |
| 81           | URR ID                                                |            |
| 82           | Linked URR ID                                         |            |
| 83           | Downlink Data Report                                  |            |
| 84           | Outer Header Creation                                 |            |
| 85           | Create BAR                                            |            |
| 86           | Update BAR (Session Modification Request)             |            |
| 87           | Remove BAR                                            |            |
| 88           | BAR ID                                                |            |
| 89           | CP Function Features                                  |            |
| 90           | Usage Information                                     |            |
| 91           | Application Instance ID                               |            |
| 92           | Flow Information                                      |            |
| 93           | UE IP Address                                         |            |
| 94           | Packet Rate                                           |            |
| 95           | Outer Header Removal                                  |            |
| 96           | Recovery Time Stamp                                   |            |
| 97           | DL Flow Level Marking                                 |            |
| 98           | Header Enrichment                                     |            |
| 99           | Error Indication Report                               |            |
| 100          | Measurement Information                               |            |
| 101          | Node Report Type                                      |            |
| 102          | User Plane Path Failure Report                        |            |
| 103          | Remote GTP-U Peer                                     |            |
| 104          | UR-SEQN                                               |            |
| 105          | Update Duplicating Parameters                         |            |
| 106          | Activate Predefined Rules                             |            |
| 107          | Deactivate Predefined Rules                           |            |
| 108          | FAR ID                                                |            |
| 109          | QER ID                                                |            |
| 110          | OCI Flags                                             |            |
| 111          | Sx Association Release Request                        |            |
| 112          | Graceful Release Period                               |            |
| 113          | PDN Type                                              |            |
| 114          | Failed Rule ID                                        |            |
| 115          | Time Quota Mechanism                                  |            |
| 116          | User Plane IP Resource Information                    |            |
| 117          | User Plane Inactivity Timer                           |            |
| 118          | Aggregated URRs                                       |            |
| 119          | Multiplier                                            |            |
| 120          | Aggregated URR ID                                     |            |
| 121          | Subsequent Volume Quota                               |            |
| 122          | Subsequent Time Quota                                 |            |
| 123          | RQI                                                   |            |
| 124          | QFI                                                   |            |
| 125          | Query URR Reference                                   |            |
| 126          | Additional Usage Reports Information                  |            |
| 127          | Create Traffic Endpoint                               |            |
| 128          | Created Traffic Endpoint                              |            |
| 129          | Update Traffic Endpoint                               |            |
| 130          | Remove Traffic Endpoint                               |            |
| 131          | Traffic Endpoint ID                                   |            |
| 132          | Ethernet Packet Filter                                |            |
| 133          | MAC address                                           |            |
| 134          | C-TAG                                                 |            |
| 135          | S-TAG                                                 |            |
| 136          | Ethertype                                             |            |
| 137          | Proxying                                              |            |
| 138          | Ethernet Filter ID                                    |            |
| 139          | Ethernet Filter Properties                            |            |
| 140          | Suggested Buffering Packets Count                     |            |
| 141          | User ID                                               |            |
| 142          | Ethernet PDU Session Information                      |            |
| 143          | Ethernet Traffic Information                          |            |
| 144          | MAC Addresses Detected                                |            |
| 145          | MAC Addresses Removed                                 |            |
| 146          | Ethernet Inactivity Timer                             |            |
| 147          | Additional Monitoring Time                            |            |
| 148          | Event Quota                                           |            |
| 149          | Event Threshold                                       |            |
| 150          | Subsequent Event Quota                                |            |
| 151          | Subsequent Event Threshold                            |            |
| 152          | Trace Information                                     |            |
| 153          | Framed-Route                                          |            |
| 154          | Framed-Routing                                        |            |
| 155          | Framed-IPv6-Route                                     |            |
| 156          | Event Time Stamp                                      |            |
| 157          | Averaging Window                                      |            |
| 158          | Paging Policy Indicator                               |            |
| 159          | APN/DNN                                               |            |
| 160          | 3GPP Interface Type                                   |            |
| 161          | PFCPSRReq-Flags                                       |            |
| 162          | PFCPAUReq-Flags                                       |            |
| 163          | Activation Time                                       |            |
| 164          | Deactivation Time                                     |            |
| 165          | Create MAR                                            |            |
| 166          | Access Forwarding Action Information 1                |            |
| 167          | Access Forwarding Action Information 2                |            |
| 168          | Remove MAR                                            |            |
| 169          | Update MAR                                            |            |
| 170          | MAR ID                                                |            |
| 171          | Steering Functionality                                |            |
| 172          | Steering Mode                                         |            |
| 173          | Weight                                                |            |
| 174          | Priority                                              |            |
| 175          | Update Access Forwarding Action Information 1         |            |
| 176          | Update Access Forwarding Action Information 2         |            |
| 177          | UE IP address Pool Identity                           |            |
| 178          | Alternative SMF IP Address                            |            |
| 179          | Packet Replication and Detection Carry-On Information |            |
| 180          | SMF Set ID                                            |            |
| 181          | Quota Validity Time                                   |            |
| 182 to 65535 | _(Spare. For future use.)_                            | -          |

## Disclaimer

This is still an experimental project. Any part of implementations(including exported APIs) may be changed before released as v1.0.0.

## Author(s)

Yoshiyuki Kurauchi ([Website](https://wmnsk.com/) / [LinkedIn](https://www.linkedin.com/in/yoshiyuki-kurauchi/))

_I'm always open to welcome co-authors! Please feel free to talk to me._

## LICENSE

[MIT](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)
