package response

type Result struct {
	Result  bool   `json:"result"`
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

type DataResult struct {
	Result

	Data interface{} `json:"data,omitempty"`
}
