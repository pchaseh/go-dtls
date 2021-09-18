package extension

import (
	"encoding/binary"
)

// SessionTicket allows for session resummption
// without server-side state
type SessionTicket struct {
	Length uint16
}

// TypeValue returns the extension TypeValue
func (s SessionTicket) TypeValue() TypeValue {
	return SessionTicketTypeValue
}

// Marshal encodes the extension
func (s* SessionTicket) Marshal() ([]byte, error) {
	out := make([]byte, 2)

	binary.BigEndian.PutUint16(out, uint16(s.TypeValue()))
	binary.BigEndian.PutUint16(out, uint16(s.Length))

	return out, nil
}

// Unmarshal populates the extension from encoded data
func (s *SessionTicket) Unmarshal(data []byte) error {
	if len(data) < 4 {
		return errBufferTooSmall
	} else if TypeValue(binary.BigEndian.Uint16(data)) != s.TypeValue() {
		return errInvalidExtensionType
	}

	s.Length = binary.BigEndian.Uint16(data[2:])

	return nil
}
