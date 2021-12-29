package api

type err struct {
	Code      int    `json:"code"`
	ErrorText string `json:"error_text"`
	Status    bool   `json:"status"`
}

var httperrs = struct {
	InternalServerError500 err
}{
	InternalServerError500: err{500, "500 Internal Server Error", false},
}

var apierrs = struct {
	ErrorJsonData err
	ErrorJsonRpc  err
}{
	ErrorJsonData: err{1, "Error Json Data", false},
	ErrorJsonRpc:  err{2, "Json RPC return an unexpected error", false},
}
