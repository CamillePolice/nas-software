package types

type OrganizeRequest struct {
	Path string `json:"path" validate:"required"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
