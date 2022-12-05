package webapijson

type Register struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// web login return json
type WebUserResponse struct {
	Lists []*Register `json:"lists"`
}
