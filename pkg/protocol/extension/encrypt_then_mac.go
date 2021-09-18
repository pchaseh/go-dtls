package extension

import (
	"encoding/binary"
)

// EncryptThenMac allows a Client/Server to communicate
// that they would like to use encrypt-then-MAC instead of
// MAC-then-encrypt
type EncryptThenMac struct {
	Length uint16
}

// TypeValue returns the extension TypeValue
func (e EncryptThenMac) TypeValue() TypeValue {
	return EncryptThenMacTypeValue
}

// Marshal encodes the extension
func (e *EncryptThenMac) Marshal() ([]byte, error) {
	out := make([]byte, 2)

	binary.BigEndian.PutUint16(out, uint16(e.TypeValue()))
	binary.BigEndian.PutUint16(out, uint16(e.Length))

	return out, nil
}

// Unmarshal populates the extension from encoded data
func (e *EncryptThenMac) Unmarshal(data []byte) error {
	if len(data) < 4 {
		return errBufferTooSmall
	} else if TypeValue(binary.BigEndian.Uint16(data)) != e.TypeValue() {
		return errInvalidExtensionType
	}

	e.Length = binary.BigEndian.Uint16(data[2:])

	return nil
}
