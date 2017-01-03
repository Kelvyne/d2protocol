package d2protocol

import (
	"fmt"
	"reflect"
)

var messages = map[uint16]reflect.Type{}

type ProtocolRequired struct {
	RequiredVersion uint32
	CurrentVersion  uint32
}

func (m *ProtocolRequired) ID() uint16 {
	return 1
}

func (m *ProtocolRequired) Serialize(w Writer) error {
	return nil
}

func (m *ProtocolRequired) Deserialize(r Reader) error {
	fmt.Println("deserializing")
	return nil
}
