package request

type ShortLinkRequest struct {
	Url        string `json:"url" validate:"required,url"`
	UrlType    string `json:"url_type" validate:"required,oneof=short long"`
	CustomName string `json:"custom_name,omitempty" validate:"omitempty,alphanum"`
}
