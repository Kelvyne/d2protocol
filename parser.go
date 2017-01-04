package d2protocol

import (
	"fmt"
	"reflect"
)

// Message represents a Dofus 2 Protocol message
type Message interface {
	ID() uint16
	Serialize(w Writer) error
	Deserialize(r Reader) error
}

// Type represents a Dofus 2 Protocol type
type Type interface {
	ID() uint16
	Serialize(w Writer) error
	Deserialize(r Reader) error
}

// ParseMessage parses a protocol message of the given id
func ParseMessage(id uint16, packet []byte) (Message, error) {
	t, ok := messages[id]
	if !ok {
		return nil, fmt.Errorf("unknown packet %v", id)
	}

	//  This can be done safely because we built the array ourselves ???
	val := reflect.New(t.Elem())
	m := val.Interface().(Message)

	if err := m.Deserialize(nil); err != nil {
		return nil, err
	}
	return m, nil
}

// GetType returns a new Type corresponding to the given id
func GetType(id uint16) (Type, error) {
	t, ok := messages[id]
	if !ok {
		return nil, fmt.Errorf("unknown type %v", id)
	}

	val := reflect.New(t.Elem())
	m := val.Interface().(Type)
	return m, nil
}

func getWrappedFlag(b uint8, position uint) bool {
	return b&(1<<position) != 0
}

func setWrappedFlag(b uint8, position uint, v bool) uint8 {
	mask := uint8(1 << position)
	if !v {
		mask = ^mask
		return b & mask
	}
	return b | mask
}
