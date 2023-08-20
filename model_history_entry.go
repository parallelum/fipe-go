/*
Fipe API

API de Consulta Tabela FIPE fornece preços médios de veículos no mercado nacional. Atualizada mensalmente com dados extraidos da tabela FIPE.    Essa API Fipe utiliza banco de dados próprio, onde todas as requisições acontecem internamente, sem sobrecarregar o Web Service da Fipe, evitando assim bloqueios por múltiplos acessos.    A API está online desde 2015 e totalmente gratuíta. Gostaria que ele continuasse gratuíta? O que acha de me pagar uma cerveja? 🍺    [![Make a donation](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=QUPMYWH6XAC5G)   ### Fipe API SDKs  - [Fipe Go SDK](https://pkg.go.dev/github.com/parallelum/fipe-go)  - [Fipe .NetCore Nuget SDK](https://www.nuget.org/packages/Br.Com.Parallelum.Fipe/)  - [Fipe Javascript SDK](https://github.com/deividfortuna/fipe-promise)  

API version: 2.0.0
Contact: deividfortuna@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fipe

import (
	"encoding/json"
)

// checks if the HistoryEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryEntry{}

// HistoryEntry struct for HistoryEntry
type HistoryEntry struct {
	Price *string `json:"price,omitempty"`
	Month *string `json:"month,omitempty"`
	Reference *string `json:"reference,omitempty"`
}

// NewHistoryEntry instantiates a new HistoryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryEntry() *HistoryEntry {
	this := HistoryEntry{}
	return &this
}

// NewHistoryEntryWithDefaults instantiates a new HistoryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryEntryWithDefaults() *HistoryEntry {
	this := HistoryEntry{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *HistoryEntry) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryEntry) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *HistoryEntry) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *HistoryEntry) SetPrice(v string) {
	o.Price = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *HistoryEntry) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryEntry) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *HistoryEntry) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *HistoryEntry) SetMonth(v string) {
	o.Month = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *HistoryEntry) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryEntry) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *HistoryEntry) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *HistoryEntry) SetReference(v string) {
	o.Reference = &v
}

func (o HistoryEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableHistoryEntry struct {
	value *HistoryEntry
	isSet bool
}

func (v NullableHistoryEntry) Get() *HistoryEntry {
	return v.value
}

func (v *NullableHistoryEntry) Set(val *HistoryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryEntry(val *HistoryEntry) *NullableHistoryEntry {
	return &NullableHistoryEntry{value: val, isSet: true}
}

func (v NullableHistoryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

