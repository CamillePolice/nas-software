package types

type BinaryResponse struct {
	Binary      []byte `json:"binary"`
	ArchiveName string `json:"archiveName"`
}
