// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEventThreshold creates a new EventThreshold IE.
func NewEventThreshold(quota uint32) *IE {
	return newUint32ValIE(EventThreshold, quota)
}

// EventThreshold returns EventThreshold in uint32 if the type of IE matches.
func (i *IE) EventThreshold() (uint32, error) {
	switch i.Type {
	case EventThreshold:
		return i.ValueAsUint32()
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == EventThreshold {
				return x.EventThreshold()
			}
		}
		return 0, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == EventThreshold {
				return x.EventThreshold()
			}
		}
		return 0, ErrIENotFound
	case AdditionalMonitoringTime:
		ies, err := i.AdditionalMonitoringTime()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == EventThreshold {
				return x.EventThreshold()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
