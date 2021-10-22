# Go API client for fipe

API de Consulta Tabela FIPE fornece preços médios de veículos no mercado nacional. Atualizada mensalmente com dados extraidos da tabela FIPE

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./fipe"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://parallelum.com.br/fipe/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FipeApi* | [**GetBrandsByType**](docs/FipeApi.md#getbrandsbytype) | **Get** /{vehicleType}/brands | Brands by type
*FipeApi* | [**GetFipeInfo**](docs/FipeApi.md#getfipeinfo) | **Get** /{vehicleType}/brands/{brandId}/models/{modelId}/years/{yearId} | Fipe info
*FipeApi* | [**GetModelsByBrand**](docs/FipeApi.md#getmodelsbybrand) | **Get** /{vehicleType}/brands/{brandId}/models | Models by brand
*FipeApi* | [**GetReferences**](docs/FipeApi.md#getreferences) | **Get** /references | Fipe month references
*FipeApi* | [**GetYearByModel**](docs/FipeApi.md#getyearbymodel) | **Get** /{vehicleType}/brands/{brandId}/models/{modelId}/years | Years for model


## Documentation For Models

 - [FipeResult](docs/FipeResult.md)
 - [NamedCode](docs/NamedCode.md)
 - [Reference](docs/Reference.md)
 - [VehiclesType](docs/VehiclesType.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

deividfortuna@gmail.com

