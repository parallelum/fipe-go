# \FipeApi

All URIs are relative to *https://parallelum.com.br/fipe/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrandsByType**](FipeApi.md#GetBrandsByType) | **Get** /{vehicleType}/brands | Brands by type
[**GetFipeInfo**](FipeApi.md#GetFipeInfo) | **Get** /{vehicleType}/brands/{brandId}/models/{modelId}/years/{yearId} | Fipe info
[**GetHistoryByFipeCode**](FipeApi.md#GetHistoryByFipeCode) | **Get** /{vehicleType}/{fipeCode}/years/{yearId}/history | Fipe price history by Fipe code
[**GetInfoByFipeCode**](FipeApi.md#GetInfoByFipeCode) | **Get** /{vehicleType}/{fipeCode}/years/{yearId} | Fipe info by Fipe code
[**GetModelsByBrand**](FipeApi.md#GetModelsByBrand) | **Get** /{vehicleType}/brands/{brandId}/models | Models by brand
[**GetReferences**](FipeApi.md#GetReferences) | **Get** /references | Fipe month references
[**GetYearByModel**](FipeApi.md#GetYearByModel) | **Get** /{vehicleType}/brands/{brandId}/models/{modelId}/years | Years by model
[**GetYearsByFipeCode**](FipeApi.md#GetYearsByFipeCode) | **Get** /{vehicleType}/{fipeCode}/years | Years by Fipe code



## GetBrandsByType

> []NamedCode GetBrandsByType(ctx, vehicleType).Execute()

Brands by type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetBrandsByType(context.Background(), vehicleType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetBrandsByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandsByType`: []NamedCode
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetBrandsByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandsByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NamedCode**](NamedCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFipeInfo

> FipeResult GetFipeInfo(ctx, vehicleType, brandId, modelId, yearId).Reference(reference).Execute()

Fipe info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle
    brandId := int32(59) // int32 | Brand of the vehicle
    modelId := int32(5940) // int32 | Model of the vehicle
    yearId := "2014-3" // string | Year for the vehicle
    reference := int32(278) // int32 | Month reference code (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetFipeInfo(context.Background(), vehicleType, brandId, modelId, yearId).Reference(reference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetFipeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFipeInfo`: FipeResult
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetFipeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 
**brandId** | **int32** | Brand of the vehicle | 
**modelId** | **int32** | Model of the vehicle | 
**yearId** | **string** | Year for the vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFipeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **reference** | **int32** | Month reference code | 

### Return type

[**FipeResult**](FipeResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryByFipeCode

> FipeHistoryResult GetHistoryByFipeCode(ctx, vehicleType, fipeCode, yearId).Execute()

Fipe price history by Fipe code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle
    fipeCode := "004278-1" // string | Fipe internal reference code
    yearId := "2005-1" // string | Year for the vehicle

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetHistoryByFipeCode(context.Background(), vehicleType, fipeCode, yearId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetHistoryByFipeCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoryByFipeCode`: FipeHistoryResult
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetHistoryByFipeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 
**fipeCode** | **string** | Fipe internal reference code | 
**yearId** | **string** | Year for the vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryByFipeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FipeHistoryResult**](FipeHistoryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfoByFipeCode

> FipeResult GetInfoByFipeCode(ctx, vehicleType, fipeCode, yearId).Execute()

Fipe info by Fipe code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle
    fipeCode := "004278-1" // string | Fipe internal reference code
    yearId := "2005-1" // string | Year for the vehicle

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetInfoByFipeCode(context.Background(), vehicleType, fipeCode, yearId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetInfoByFipeCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfoByFipeCode`: FipeResult
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetInfoByFipeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 
**fipeCode** | **string** | Fipe internal reference code | 
**yearId** | **string** | Year for the vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoByFipeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FipeResult**](FipeResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsByBrand

> []NamedCode GetModelsByBrand(ctx, vehicleType, brandId).Execute()

Models by brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle
    brandId := int32(59) // int32 | Brand of the vehicle

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetModelsByBrand(context.Background(), vehicleType, brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetModelsByBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModelsByBrand`: []NamedCode
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetModelsByBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 
**brandId** | **int32** | Brand of the vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsByBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]NamedCode**](NamedCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferences

> []Reference GetReferences(ctx).Execute()

Fipe month references



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetReferences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferences`: []Reference
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetReferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferencesRequest struct via the builder pattern


### Return type

[**[]Reference**](Reference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYearByModel

> []NamedCode GetYearByModel(ctx, vehicleType, brandId, modelId).Execute()

Years by model



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle
    brandId := int32(59) // int32 | Brand of the vehicle
    modelId := int32(5940) // int32 | Model of the vehicle

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetYearByModel(context.Background(), vehicleType, brandId, modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetYearByModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetYearByModel`: []NamedCode
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetYearByModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 
**brandId** | **int32** | Brand of the vehicle | 
**modelId** | **int32** | Model of the vehicle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetYearByModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]NamedCode**](NamedCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYearsByFipeCode

> []NamedCode GetYearsByFipeCode(ctx, vehicleType, fipeCode).Execute()

Years by Fipe code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/parallelum/fipe-go"
)

func main() {
    vehicleType := openapiclient.VehiclesType("cars") // VehiclesType | Type of vehicle
    fipeCode := "004278-1" // string | Fipe internal reference code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FipeApi.GetYearsByFipeCode(context.Background(), vehicleType, fipeCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FipeApi.GetYearsByFipeCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetYearsByFipeCode`: []NamedCode
    fmt.Fprintf(os.Stdout, "Response from `FipeApi.GetYearsByFipeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleType** | [**VehiclesType**](.md) | Type of vehicle | 
**fipeCode** | **string** | Fipe internal reference code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetYearsByFipeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]NamedCode**](NamedCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

