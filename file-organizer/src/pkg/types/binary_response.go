package types

type BinaryResponse struct {
	Binary   []byte `json:"binary"`
	Filename string `json:"filename"`
}
