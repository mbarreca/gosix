package library

// The header object will hold the header key and value
type Header struct {
	Key   string
	Value string
}

// Constructor
func New(key, value string) *Header {
	return &Header{Key: key, Value: value}
}
