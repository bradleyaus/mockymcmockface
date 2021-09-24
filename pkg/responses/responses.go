package responses

import (
	"encoding/json"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Responder interface {
	Get()(protoiface.MessageV1, error)
}

type response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type responseGetterKey struct {
	service string
	method string
}

func NewService() *ResponseService {
	return &ResponseService{}
}

type ResponseService struct {

}


func (r * ResponseService) GetResponse(service string, method string, input interface{}) (resp *response, err error) {

	jsonString := `
		{
			"data": {
				"resp": "lol"
			}
		}
	`

	resp = &response{}
	err = json.Unmarshal([]byte(jsonString), resp)

	return resp, nil
}