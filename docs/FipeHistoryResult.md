# FipeHistoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand of the vehicle | [optional] 
**Model** | Pointer to **string** | Model of the vehicle | [optional] 
**ModelYear** | Pointer to **int32** | Year of vehicle production | [optional] 
**Fuel** | Pointer to **string** | Fuel used by the vehicle | [optional] 
**CodeFipe** | Pointer to **string** | Unique Fipe identifier | [optional] 
**VehicleType** | Pointer to **int32** | Type of the vehicle | [optional] 
**FuelAcronym** | Pointer to **string** | Fuel acronym (eg. G) | [optional] 
**PriceHistory** | Pointer to [**[]HistoryEntry**](HistoryEntry.md) |  | [optional] 

## Methods

### NewFipeHistoryResult

`func NewFipeHistoryResult() *FipeHistoryResult`

NewFipeHistoryResult instantiates a new FipeHistoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFipeHistoryResultWithDefaults

`func NewFipeHistoryResultWithDefaults() *FipeHistoryResult`

NewFipeHistoryResultWithDefaults instantiates a new FipeHistoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *FipeHistoryResult) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *FipeHistoryResult) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *FipeHistoryResult) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *FipeHistoryResult) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetModel

`func (o *FipeHistoryResult) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FipeHistoryResult) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FipeHistoryResult) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FipeHistoryResult) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelYear

`func (o *FipeHistoryResult) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *FipeHistoryResult) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *FipeHistoryResult) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *FipeHistoryResult) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### GetFuel

`func (o *FipeHistoryResult) GetFuel() string`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *FipeHistoryResult) GetFuelOk() (*string, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *FipeHistoryResult) SetFuel(v string)`

SetFuel sets Fuel field to given value.

### HasFuel

`func (o *FipeHistoryResult) HasFuel() bool`

HasFuel returns a boolean if a field has been set.

### GetCodeFipe

`func (o *FipeHistoryResult) GetCodeFipe() string`

GetCodeFipe returns the CodeFipe field if non-nil, zero value otherwise.

### GetCodeFipeOk

`func (o *FipeHistoryResult) GetCodeFipeOk() (*string, bool)`

GetCodeFipeOk returns a tuple with the CodeFipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeFipe

`func (o *FipeHistoryResult) SetCodeFipe(v string)`

SetCodeFipe sets CodeFipe field to given value.

### HasCodeFipe

`func (o *FipeHistoryResult) HasCodeFipe() bool`

HasCodeFipe returns a boolean if a field has been set.

### GetVehicleType

`func (o *FipeHistoryResult) GetVehicleType() int32`

GetVehicleType returns the VehicleType field if non-nil, zero value otherwise.

### GetVehicleTypeOk

`func (o *FipeHistoryResult) GetVehicleTypeOk() (*int32, bool)`

GetVehicleTypeOk returns a tuple with the VehicleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleType

`func (o *FipeHistoryResult) SetVehicleType(v int32)`

SetVehicleType sets VehicleType field to given value.

### HasVehicleType

`func (o *FipeHistoryResult) HasVehicleType() bool`

HasVehicleType returns a boolean if a field has been set.

### GetFuelAcronym

`func (o *FipeHistoryResult) GetFuelAcronym() string`

GetFuelAcronym returns the FuelAcronym field if non-nil, zero value otherwise.

### GetFuelAcronymOk

`func (o *FipeHistoryResult) GetFuelAcronymOk() (*string, bool)`

GetFuelAcronymOk returns a tuple with the FuelAcronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelAcronym

`func (o *FipeHistoryResult) SetFuelAcronym(v string)`

SetFuelAcronym sets FuelAcronym field to given value.

### HasFuelAcronym

`func (o *FipeHistoryResult) HasFuelAcronym() bool`

HasFuelAcronym returns a boolean if a field has been set.

### GetPriceHistory

`func (o *FipeHistoryResult) GetPriceHistory() []HistoryEntry`

GetPriceHistory returns the PriceHistory field if non-nil, zero value otherwise.

### GetPriceHistoryOk

`func (o *FipeHistoryResult) GetPriceHistoryOk() (*[]HistoryEntry, bool)`

GetPriceHistoryOk returns a tuple with the PriceHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHistory

`func (o *FipeHistoryResult) SetPriceHistory(v []HistoryEntry)`

SetPriceHistory sets PriceHistory field to given value.

### HasPriceHistory

`func (o *FipeHistoryResult) HasPriceHistory() bool`

HasPriceHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


