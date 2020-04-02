package message

import "github.com/wmnsk/go-pfcp/ie"

// SessionEstablishmentRequset is a SessionEstablishmentRequset formed PFCP Header and its IEs above.
type SessionEstablishmentRequset struct {
	*Header
	NodeID *ie.IE
	IEs    []*ie.IE
}

// NewSessionEstablishmentRequset creates a new SessionEstablishmentRequset.
func NewSessionEstablishmentRequset(nid *ie.IE, IEs ...*ie.IE) *SessionEstablishmentRequset {
	m := &SessionEstablishmentRequset{
		Header: NewHeader(
			1, 0, 0, 0,
			MsgTypeSessionEstablishmentRequest, 0, 0, 0,
			nil,
		),
		NodeID: nid,
		IEs:    IEs,
	}
	m.SetLength()

	return m
}

// Marshal returns the byte sequence generated from a SessionEstablishmentRequset.
func (m *SessionEstablishmentRequset) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *SessionEstablishmentRequset) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if i := m.NodeID; i != nil {
		if err := i.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParseSessionEstablishmentRequset decodes a given byte sequence as a SessionEstablishmentRequset.
func ParseSessionEstablishmentRequset(b []byte) (*SessionEstablishmentRequset, error) {
	m := &SessionEstablishmentRequset{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes a given byte sequence as a SessionEstablishmentRequset.
func (m *SessionEstablishmentRequset) UnmarshalBinary(b []byte) error {
	var err error
	m.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(m.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(m.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		switch i.Type {
		case ie.NodeID:
			m.NodeID = i
		default:
			m.IEs = append(m.IEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (m *SessionEstablishmentRequset) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if i := m.NodeID; i != nil {
		l += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (m *SessionEstablishmentRequset) SetLength() {
	l := m.Header.MarshalLen() - len(m.Header.Payload) - 4

	if i := m.NodeID; i != nil {
		l += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		l += ie.MarshalLen()
	}
	m.Header.Length = uint16(l)
}

// MessageTypeName returns the name of protocol.
func (m *SessionEstablishmentRequset) MessageTypeName() string {
	return "PFCP Session Establishment  Request"
}

// SEID returns the SEID in uint64.
func (m *SessionEstablishmentRequset) SEID() uint64 {
	return m.Header.seid()
}
