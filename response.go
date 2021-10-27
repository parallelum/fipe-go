/*
Fipe API

API de Consulta Tabela FIPE fornece preços médios de veículos no mercado nacional. Atualizada mensalmente com dados extraidos da tabela FIPE.    Essa API Fipe utiliza banco de dados próprio, onde todas as requisições acontecem internamente, sem sobrecarregar o Web Service da Fipe, evitando assim bloqueios por múltiplos acessos.    A API está online desde 2015 e totalmente gratuíta. Gostaria que ele continuasse gratuíta? O que acha de me pagar uma cerveja? 🍺    [![Make a donation](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=QUPMYWH6XAC5G)   ## Available SDKs  * [Fipe Go SDK](https://pkg.go.dev/github.com/parallelum/fipe-go)  * [Fipe .NetCore Nuget SDK](https://www.nuget.org/packages/Br.Com.Parallelum.Fipe/)  * [Fipe Javascript SDK](https://github.com/deividfortuna/fipe-promise)  

API version: 2.0.0
Contact: deividfortuna@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fipe

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResponse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
