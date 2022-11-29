package spec

type SayHelloRequest struct {
	Name string `json:"name"`
}

type SayHello200Response struct {
	Data SayHello200ResponseData `json:"data"`
}
type SayHello200ResponseData struct {
	Sentence string `json:"sentence"`
}

type SayHello400Response struct {
	Error SayHello400ResponseError `json:"error"`
}
type SayHello400ResponseError struct {
	Message string `json:"message"`
}

type SayHello500Response struct {
	Error SayHello500ResponseError `json:"error"`
}
type SayHello500ResponseError struct {
	Message string `json:"message"`
}
