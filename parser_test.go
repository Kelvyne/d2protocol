package d2protocol

import (
	"reflect"
	"testing"
)

func TestParseMessage(t *testing.T) {
	type args struct {
		id     uint16
		packet []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Message
		wantErr bool
	}{
		{
			"protocolRequired",
			args{1, []byte{}},
			&ProtocolRequired{0, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMessage(tt.args.id, tt.args.packet)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setWrappedFlag(t *testing.T) {
	type args struct {
		b        uint8
		position uint
		v        bool
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			"simple true",
			args{0, 2, true},
			0x04,
		},
		{
			"simple fals",
			args{0, 2, false},
			0x00,
		},
		{
			"non zero flag false",
			args{0x5f, 0, false},
			0x5e,
		},
		{
			"non zero flag true",
			args{0x5e, 0, true},
			0x5f,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setWrappedFlag(tt.args.b, tt.args.position, tt.args.v); got != tt.want {
				t.Errorf("setWrappedFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWrappedFlag(t *testing.T) {
	type args struct {
		b        uint8
		position uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWrappedFlag(tt.args.b, tt.args.position); got != tt.want {
				t.Errorf("getWrappedFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
