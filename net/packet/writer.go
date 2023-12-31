package packet

import (
	"errors"
	"io"
)

var (
	ErrStringTooBig    = errors.New("string too big")
	ErrByteArrayTooBig = errors.New("byte array too big")
)

func WritePacket(w io.Writer, pk Packet) {
	w.Write([]byte{pk.ID()})
	pk.Encode(Writer{w})
}

type Writer struct {
	buf io.Writer
}

func (w Writer) Byte(i uint8) {
	w.buf.Write([]byte{i})
}

func (w Writer) SByte(i int8) {
	w.buf.Write([]byte{byte(i)})
}

func (w Writer) Short(i int16) {
	w.buf.Write([]byte{byte(i >> 8), byte(i)})
}

func (w Writer) Int(i int32) {
	w.buf.Write([]byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)})
}

func (w Writer) String(s string) error {
	b := []byte(s)
	if len(b) > 64 {
		return ErrStringTooBig
	}
	for len(b) < 64 {
		b = append(b, 0x20)
	}
	w.buf.Write(b)
	return nil
}

func (w Writer) ByteArray(b []byte) error {
	if len(b) > 1024 {
		return ErrByteArrayTooBig
	}
	if len(b) < 1024 {
		b = append(b, make([]byte, 1024-len(b))...)
	}
	w.buf.Write(b)
	return nil
}

func (w Writer) Auto(b any) {
	switch t := b.(type) {
	case byte:
		w.Byte(t)
	case int8:
		w.SByte(t)
	case int16:
		w.Short(t)
	case int32:
		w.Int(t)
	case string:
		w.String(t)
	case []byte:
		w.ByteArray(t)
	}
}
