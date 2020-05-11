// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationID creates a new ApplicationID IE.
func NewApplicationID(instance string) *IE {
	return newStringIE(ApplicationID, instance)
}

// ApplicationID returns ApplicationID in string if the type of IE matches.
func (i *IE) ApplicationID() (string, error) {
	switch i.Type {
	case ApplicationID:
		return string(i.Payload), nil
	case ApplicationIDsPFDs:
		ies, err := i.ApplicationIDsPFDs()
		if err != nil {
			return "", err
		}
		for _, x := range ies {
			if x.Type == ApplicationID {
				return x.ApplicationID()
			}
		}
		return "", ErrIENotFound
	case ApplicationDetectionInformation:
		ies, err := i.ApplicationDetectionInformation()
		if err != nil {
			return "", err
		}
		for _, x := range ies {
			if x.Type == ApplicationID {
				return x.ApplicationID()
			}
		}
		return "", ErrIENotFound
	default:
		return "", ErrInvalidType
	}
}
