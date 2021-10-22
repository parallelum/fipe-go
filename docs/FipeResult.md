# FipeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **string** | Price in Brazilian Real | [optional] 
**Brand** | Pointer to **string** | Brand of the vehicle | [optional] 
**Model** | Pointer to **string** | Model of the vehicle | [optional] 
**ModelYear** | Pointer to **int32** | Year of vehicle production | [optional] 
**Fuel** | Pointer to **string** | Fuel used by the vehicle | [optional] 
**CodeFipe** | Pointer to **string** | Unique Fipe identifier | [optional] 
**ReferenceMonth** | Pointer to **string** | Month of the price | [optional] 
**VehicleType** | Pointer to **int32** | Type of the vehicle | [optional] 
**FuelAcronym** | Pointer to **string** | Fuel acronym (eg. G) | [optional] 

## Methods

### NewFipeResult

`func NewFipeResult() *FipeResult`

NewFipeResult instantiates a new FipeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFipeResultWithDefaults

`func NewFipeResultWithDefaults() *FipeResult`

NewFipeResultWithDefaults instantiates a new FipeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *FipeResult) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FipeResult) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FipeResult) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FipeResult) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBrand

`func (o *FipeResult) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *FipeResult) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *FipeResult) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *FipeResult) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetModel

`func (o *FipeResult) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FipeResult) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FipeResult) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FipeResult) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelYear

`func (o *FipeResult) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *FipeResult) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *FipeResult) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *FipeResult) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### GetFuel

`func (o *FipeResult) GetFuel() string`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *FipeResult) GetFuelOk() (*string, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *FipeResult) SetFuel(v string)`

SetFuel sets Fuel field to given value.

### HasFuel

`func (o *FipeResult) HasFuel() bool`

HasFuel returns a boolean if a field has been set.

### GetCodeFipe

`func (o *FipeResult) GetCodeFipe() string`

GetCodeFipe returns the CodeFipe field if non-nil, zero value otherwise.

### GetCodeFipeOk

`func (o *FipeResult) GetCodeFipeOk() (*string, bool)`

GetCodeFipeOk returns a tuple with the CodeFipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeFipe

`func (o *FipeResult) SetCodeFipe(v string)`

SetCodeFipe sets CodeFipe field to given value.

### HasCodeFipe

`func (o *FipeResult) HasCodeFipe() bool`

HasCodeFipe returns a boolean if a field has been set.

### GetReferenceMonth

`func (o *FipeResult) GetReferenceMonth() string`

GetReferenceMonth returns the ReferenceMonth field if non-nil, zero value otherwise.

### GetReferenceMonthOk

`func (o *FipeResult) GetReferenceMonthOk() (*string, bool)`

GetReferenceMonthOk returns a tuple with the ReferenceMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMonth

`func (o *FipeResult) SetReferenceMonth(v string)`

SetReferenceMonth sets ReferenceMonth field to given value.

### HasReferenceMonth

`func (o *FipeResult) HasReferenceMonth() bool`

HasReferenceMonth returns a boolean if a field has been set.

### GetVehicleType

`func (o *FipeResult) GetVehicleType() int32`

GetVehicleType returns the VehicleType field if non-nil, zero value otherwise.

### GetVehicleTypeOk

`func (o *FipeResult) GetVehicleTypeOk() (*int32, bool)`

GetVehicleTypeOk returns a tuple with the VehicleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleType

`func (o *FipeResult) SetVehicleType(v int32)`

SetVehicleType sets VehicleType field to given value.

### HasVehicleType

`func (o *FipeResult) HasVehicleType() bool`

HasVehicleType returns a boolean if a field has been set.

### GetFuelAcronym

`func (o *FipeResult) GetFuelAcronym() string`

GetFuelAcronym returns the FuelAcronym field if non-nil, zero value otherwise.

### GetFuelAcronymOk

`func (o *FipeResult) GetFuelAcronymOk() (*string, bool)`

GetFuelAcronymOk returns a tuple with the FuelAcronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelAcronym

`func (o *FipeResult) SetFuelAcronym(v string)`

SetFuelAcronym sets FuelAcronym field to given value.

### HasFuelAcronym

`func (o *FipeResult) HasFuelAcronym() bool`

HasFuelAcronym returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


