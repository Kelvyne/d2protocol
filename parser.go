package d2protocol

import (
	"fmt"
	"reflect"
)

// Message represents a Dofus 2 Protocol Message
type Message interface {
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
