// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewBridgeManagementInformationContainer creates a new BridgeManagementInformationContainer IE.
func NewBridgeManagementInformationContainer(info string) *IE {
	return newStringIE(BridgeManagementInformationContainer, info)
}

// BridgeManagementInformationContainer returns BridgeManagementInformationContainer in string if the type of IE matches.
func (i *IE) BridgeManagementInformationContainer() (string, error) {
	switch i.Type {
	case BridgeManagementInformationContainer:
		return i.ValueAsString()
	case TSCManagementInformationWithinSessionModificationRequest,
		TSCManagementInformationWithinSessionModificationResponse,
		TSCManagementInformationWithinSessionReportRequest:
		ies, err := i.TSCManagementInformation()
		if err != nil {
			return "", err
		}
		for _, x := range ies {
			if x.Type == BridgeManagementInformationContainer {
				return x.BridgeManagementInformationContainer()
			}
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
