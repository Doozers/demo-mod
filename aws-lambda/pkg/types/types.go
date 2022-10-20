package types

type Request struct {
	RequestType string `json:"requestType"`
	Sum         *Sum   `json:"sum"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Result  int    `json:"result,omitempty"`
}
