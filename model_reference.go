/*
Fipe API

API de Consulta Tabela FIPE fornece preços médios de veículos no mercado nacional. Atualizada mensalmente com dados extraidos da tabela FIPE.    Essa API Fipe utiliza banco de dados próprio, onde todas as requisições acontecem internamente, sem sobrecarregar o Web Service da Fipe, evitando assim bloqueios por múltiplos acessos.    A API está online desde 2015 e totalmente gratuíta. Gostaria que ele continuasse gratuíta? O que acha de me pagar uma cerveja? 🍺    [![Make a donation](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=QUPMYWH6XAC5G)   ## Available SDKs  * [Fipe Go SDK](https://pkg.go.dev/github.com/parallelum/fipe-go)  * [Fipe .NetCore Nuget SDK](https://www.nuget.org/packages/Br.Com.Parallelum.Fipe/)  * [Fipe Javascript SDK](https://github.com/deividfortuna/fipe-promise)  

API version: 2.0.0
Contact: deividfortuna@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fipe

import (
	"encoding/json"
)

// Reference struct for Reference
type Reference struct {
	Code *string `json:"code,omitempty"`
	Month *string `json:"month,omitempty"`
}

// NewReference instantiates a new Reference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReference() *Reference {
	this := Reference{}
	return &this
}

// NewReferenceWithDefaults instantiates a new Reference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceWithDefaults() *Reference {
	this := Reference{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Reference) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reference) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Reference) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Reference) SetCode(v string) {
	o.Code = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *Reference) GetMonth() string {
	if o == nil || o.Month == nil {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reference) GetMonthOk() (*string, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *Reference) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *Reference) SetMonth(v string) {
	o.Month = &v
}

func (o Reference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	return json.Marshal(toSerialize)
}

type NullableReference struct {
	value *Reference
	isSet bool
}

func (v NullableReference) Get() *Reference {
	return v.value
}

func (v *NullableReference) Set(val *Reference) {
	v.value = val
	v.isSet = true
}

func (v NullableReference) IsSet() bool {
	return v.isSet
}

func (v *NullableReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReference(val *Reference) *NullableReference {
	return &NullableReference{value: val, isSet: true}
}

func (v NullableReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


