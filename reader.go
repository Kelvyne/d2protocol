package d2protocol

import "io"

import "encoding/binary"
import "errors"

// Reader is the minimal interface required to deserialiaze any
// protocol message/type
type Reader interface {
	ReadBoolean() (bool, error)
	ReadInt8() (int8, error)
	ReadUInt8() (uint8, error)
	ReadInt16() (int16, error)
	ReadUInt16() (uint16, error)
	ReadInt32() (int32, error)
	ReadUInt32() (uint32, error)
	ReadFloat() (float32, error)
	ReadDouble() (float64, error)
	ReadString() (string, error)
	ReadVarShort() (int16, error)
	ReadVarUhShort() (uint16, error)
	ReadVarInt() (int32, error)
	ReadVarUhInt() (uint32, error)
	ReadVarLong() (int64, error)
	ReadVarUhLong() (uint64, error)
}

type reader struct {
	r io.Reader
}

// ErrReaderMalformedVar means that a variable-length integer (16, 32 or 64 bits)
// is malformed and take too much bits
var ErrReaderMalformedVar = errors.New("malformed variable length integer")

// NewReader creates a Reader
func NewReader(r io.Reader) Reader {
	return &reader{r}
}

func (r *reader) read(x interface{}) error {
	return binary.Read(r.r, binary.BigEndian, x)
}

func (r *reader) ReadBoolean() (bool, error) {
	b, err := r.ReadUInt8()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

func (r *reader) ReadInt8() (int8, error) {
	b, err := r.ReadUInt8()
	if err != nil {
		return 0, err
	}
	return int8(b), nil
}

func (r *reader) ReadUInt8() (uint8, error) {
	var b uint8
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadInt16() (int16, error) {
	var b int16
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadUInt16() (uint16, error) {
	var b uint16
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadInt32() (int32, error) {
	var b int32
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadUInt32() (uint32, error) {
	var b uint32
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadFloat() (float32, error) {
	var b float32
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadDouble() (float64, error) {
	var b float64
	if err := r.read(&b); err != nil {
		return 0, err
	}
	return b, nil
}

func (r *reader) ReadString() (string, error) {
	len, err := r.ReadUInt16()
	if err != nil {
		return "", err
	}

	buf := make([]byte, len)
	_, err = io.ReadFull(r.r, buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (r *reader) ReadVarShort() (int16, error) {
	v, err := r.readVar(16)
	if err != nil {
		return 0, err
	}
	return int16(v), err
}

func (r *reader) ReadVarUhShort() (uint16, error) {
	v, err := r.readVar(16)
	if err != nil {
		return 0, err
	}
	return uint16(v), err
}

func (r *reader) ReadVarInt() (int32, error) {
	v, err := r.readVar(32)
	if err != nil {
		return 0, err
	}
	return int32(v), err
}

func (r *reader) ReadVarUhInt() (uint32, error) {
	v, err := r.readVar(32)
	if err != nil {
		return 0, err
	}
	return uint32(v), err
}

func (r *reader) ReadVarLong() (int64, error) {
	v, err := r.readVar(64)
	if err != nil {
		return 0, err
	}
	return int64(v), err
}

func (r *reader) ReadVarUhLong() (uint64, error) {
	return r.readVar(64)
}

func (r *reader) readVar(bits uint) (uint64, error) {
	var v uint64
	for n := uint(0); n < bits; n += 7 {
		b, err := r.ReadUInt8()
		if err != nil {
			return 0, err
		}

		v |= uint64(b&0x7F) << n

		if b&0x80 == 0 {
			return v, nil
		}
	}
	return 0, ErrReaderMalformedVar
}
