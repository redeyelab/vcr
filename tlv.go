package main

// TLV Type, Length and Value
type TLV struct {
	buffer []byte
}

type TLVCallbacks struct {
	Type     int
	Callback func(tlv *TLV)
}

const (
	TLVZero = iota
	TLVTerm

	TLVPlay
	TLVPause
)

// NewTLV gets a new TLV ready to go
func NewTLV(typ, l byte) (t TLV) {
	if l < 2 {
		l = 2
	}
	t.buffer = make([]byte, l)
	t.buffer[0] = typ
	t.buffer[1] = l

	return t
}

// Type of TLV
func (tlv *TLV) Type() int {
	return int(tlv.buffer[0])
}

// Type of TLV
func (tlv *TLV) Len() int {
	return int(tlv.buffer[1])
}

// TypeLen of TLV
func (tlv *TLV) TypeLen() (t int, l int) {
	return int(tlv.buffer[0]), int(tlv.buffer[1])
}

// Value of the TLV
func (tlv *TLV) Value() []byte {
	return tlv.buffer[2:]
}

func (tlv *TLV) Str() string {
	return string(tlv.buffer)
}
