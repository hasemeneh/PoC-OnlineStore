package response

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	BasicResponse
	JsonBody JsonBody
	Error    error
}
type JsonBody struct {
	*ResponseError
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewJsonResponse() *JsonResponse {
	return &JsonResponse{
		BasicResponse: BasicResponse{
			ContentType: JsonContentType,
			StatusCode:  http.StatusOK,
		},
	}
}

func (r *JsonResponse) SetData(data interface{}) *JsonResponse {
	r.JsonBody.Data = data
	return r
}

func (r *JsonResponse) SetMessage(Message string) *JsonResponse {
	r.JsonBody.Message = Message
	return r
}

func (r *JsonResponse) SetError(err error) *JsonResponse {
	if respErr, ok := err.(*ResponseError); ok {
		r.JsonBody.ResponseError = respErr
	} else {
		//when unspecified error is provided it will categorize the response as internal server error
		r.JsonBody.ResponseError = NewResponseError(err.Error(), 500)
	}
	return r
}

func (r *JsonResponse) WriteResponse(w http.ResponseWriter) {
	b, err := json.Marshal(r.JsonBody)
	if err != nil {
		jsonBody := JsonBody{
			ResponseError: NewResponseError(err.Error(), http.StatusInternalServerError),
		}
		b, _ = json.Marshal(jsonBody)
	}
	r.Body = b
	if r.JsonBody.ResponseError != nil {
		r.StatusCode = r.JsonBody.ErrorCode
	}
	r.BasicResponse.WriteResponse(w)
}
