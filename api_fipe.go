/*
Fipe API

API de Consulta Tabela FIPE fornece preços médios de veículos no mercado nacional. Atualizada mensalmente com dados extraidos da tabela FIPE.    Essa API Fipe utiliza banco de dados próprio, onde todas as requisições acontecem internamente, sem sobrecarregar o Web Service da Fipe, evitando assim bloqueios por múltiplos acessos.    A API está online desde 2015 e totalmente gratuíta. Gostaria que ele continuasse gratuíta? O que acha de me pagar uma cerveja? 🍺    [![Make a donation](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=QUPMYWH6XAC5G)   ### Fipe API SDKs  - [Fipe Go SDK](https://pkg.go.dev/github.com/parallelum/fipe-go)  - [Fipe .NetCore Nuget SDK](https://www.nuget.org/packages/Br.Com.Parallelum.Fipe/)  - [Fipe Javascript SDK](https://github.com/deividfortuna/fipe-promise)  

API version: 2.0.0
Contact: deividfortuna@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fipe

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type FipeApi interface {

	/*
	GetBrandsByType Brands by type

	Returns brands for the type of vehicle

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@return ApiGetBrandsByTypeRequest
	*/
	GetBrandsByType(ctx context.Context, vehicleType VehiclesType) ApiGetBrandsByTypeRequest

	// GetBrandsByTypeExecute executes the request
	//  @return []NamedCode
	GetBrandsByTypeExecute(r ApiGetBrandsByTypeRequest) ([]NamedCode, *http.Response, error)

	/*
	GetFipeInfo Fipe info

	Returns the Fipe information for the vehicle (price estimation)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@param brandId Brand of the vehicle
	@param modelId Model of the vehicle
	@param yearId Year for the vehicle
	@return ApiGetFipeInfoRequest
	*/
	GetFipeInfo(ctx context.Context, vehicleType VehiclesType, brandId int32, modelId int32, yearId string) ApiGetFipeInfoRequest

	// GetFipeInfoExecute executes the request
	//  @return FipeResult
	GetFipeInfoExecute(r ApiGetFipeInfoRequest) (*FipeResult, *http.Response, error)

	/*
	GetHistoryByFipeCode Fipe price history by Fipe code

	Returns the price history for the vehicle (price estimation)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@param fipeCode Fipe internal reference code
	@param yearId Year for the vehicle
	@return ApiGetHistoryByFipeCodeRequest
	*/
	GetHistoryByFipeCode(ctx context.Context, vehicleType VehiclesType, fipeCode string, yearId string) ApiGetHistoryByFipeCodeRequest

	// GetHistoryByFipeCodeExecute executes the request
	//  @return FipeHistoryResult
	GetHistoryByFipeCodeExecute(r ApiGetHistoryByFipeCodeRequest) (*FipeHistoryResult, *http.Response, error)

	/*
	GetInfoByFipeCode Fipe info by Fipe code

	Returns the Fipe information for the vehicle (price estimation)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@param fipeCode Fipe internal reference code
	@param yearId Year for the vehicle
	@return ApiGetInfoByFipeCodeRequest
	*/
	GetInfoByFipeCode(ctx context.Context, vehicleType VehiclesType, fipeCode string, yearId string) ApiGetInfoByFipeCodeRequest

	// GetInfoByFipeCodeExecute executes the request
	//  @return FipeResult
	GetInfoByFipeCodeExecute(r ApiGetInfoByFipeCodeRequest) (*FipeResult, *http.Response, error)

	/*
	GetModelsByBrand Models by brand

	Returns models for the brand

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@param brandId Brand of the vehicle
	@return ApiGetModelsByBrandRequest
	*/
	GetModelsByBrand(ctx context.Context, vehicleType VehiclesType, brandId int32) ApiGetModelsByBrandRequest

	// GetModelsByBrandExecute executes the request
	//  @return []NamedCode
	GetModelsByBrandExecute(r ApiGetModelsByBrandRequest) ([]NamedCode, *http.Response, error)

	/*
	GetReferences Fipe month references

	Returns months and codes reference from Fipe

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetReferencesRequest
	*/
	GetReferences(ctx context.Context) ApiGetReferencesRequest

	// GetReferencesExecute executes the request
	//  @return []Reference
	GetReferencesExecute(r ApiGetReferencesRequest) ([]Reference, *http.Response, error)

	/*
	GetYearByModel Years by model

	Returns years for the specific model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@param brandId Brand of the vehicle
	@param modelId Model of the vehicle
	@return ApiGetYearByModelRequest
	*/
	GetYearByModel(ctx context.Context, vehicleType VehiclesType, brandId int32, modelId int32) ApiGetYearByModelRequest

	// GetYearByModelExecute executes the request
	//  @return []NamedCode
	GetYearByModelExecute(r ApiGetYearByModelRequest) ([]NamedCode, *http.Response, error)

	/*
	GetYearsByFipeCode Years by Fipe code

	Returns years available for vehicle by fipe code

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vehicleType Type of vehicle
	@param fipeCode Fipe internal reference code
	@return ApiGetYearsByFipeCodeRequest
	*/
	GetYearsByFipeCode(ctx context.Context, vehicleType VehiclesType, fipeCode string) ApiGetYearsByFipeCodeRequest

	// GetYearsByFipeCodeExecute executes the request
	//  @return []NamedCode
	GetYearsByFipeCodeExecute(r ApiGetYearsByFipeCodeRequest) ([]NamedCode, *http.Response, error)
}

// FipeApiService FipeApi service
type FipeApiService service

type ApiGetBrandsByTypeRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
}

func (r ApiGetBrandsByTypeRequest) Execute() ([]NamedCode, *http.Response, error) {
	return r.ApiService.GetBrandsByTypeExecute(r)
}

/*
GetBrandsByType Brands by type

Returns brands for the type of vehicle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @return ApiGetBrandsByTypeRequest
*/
func (a *FipeApiService) GetBrandsByType(ctx context.Context, vehicleType VehiclesType) ApiGetBrandsByTypeRequest {
	return ApiGetBrandsByTypeRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
	}
}

// Execute executes the request
//  @return []NamedCode
func (a *FipeApiService) GetBrandsByTypeExecute(r ApiGetBrandsByTypeRequest) ([]NamedCode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []NamedCode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetBrandsByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/brands"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFipeInfoRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
	brandId int32
	modelId int32
	yearId string
	reference *int32
}

// Month reference code
func (r ApiGetFipeInfoRequest) Reference(reference int32) ApiGetFipeInfoRequest {
	r.reference = &reference
	return r
}

func (r ApiGetFipeInfoRequest) Execute() (*FipeResult, *http.Response, error) {
	return r.ApiService.GetFipeInfoExecute(r)
}

/*
GetFipeInfo Fipe info

Returns the Fipe information for the vehicle (price estimation)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @param brandId Brand of the vehicle
 @param modelId Model of the vehicle
 @param yearId Year for the vehicle
 @return ApiGetFipeInfoRequest
*/
func (a *FipeApiService) GetFipeInfo(ctx context.Context, vehicleType VehiclesType, brandId int32, modelId int32, yearId string) ApiGetFipeInfoRequest {
	return ApiGetFipeInfoRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
		brandId: brandId,
		modelId: modelId,
		yearId: yearId,
	}
}

// Execute executes the request
//  @return FipeResult
func (a *FipeApiService) GetFipeInfoExecute(r ApiGetFipeInfoRequest) (*FipeResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FipeResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetFipeInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/brands/{brandId}/models/{modelId}/years/{yearId}"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brandId"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelId"+"}", url.PathEscape(parameterValueToString(r.modelId, "modelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"yearId"+"}", url.PathEscape(parameterValueToString(r.yearId, "yearId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.reference != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reference", r.reference, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetHistoryByFipeCodeRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
	fipeCode string
	yearId string
}

func (r ApiGetHistoryByFipeCodeRequest) Execute() (*FipeHistoryResult, *http.Response, error) {
	return r.ApiService.GetHistoryByFipeCodeExecute(r)
}

/*
GetHistoryByFipeCode Fipe price history by Fipe code

Returns the price history for the vehicle (price estimation)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @param fipeCode Fipe internal reference code
 @param yearId Year for the vehicle
 @return ApiGetHistoryByFipeCodeRequest
*/
func (a *FipeApiService) GetHistoryByFipeCode(ctx context.Context, vehicleType VehiclesType, fipeCode string, yearId string) ApiGetHistoryByFipeCodeRequest {
	return ApiGetHistoryByFipeCodeRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
		fipeCode: fipeCode,
		yearId: yearId,
	}
}

// Execute executes the request
//  @return FipeHistoryResult
func (a *FipeApiService) GetHistoryByFipeCodeExecute(r ApiGetHistoryByFipeCodeRequest) (*FipeHistoryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FipeHistoryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetHistoryByFipeCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/{fipeCode}/years/{yearId}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fipeCode"+"}", url.PathEscape(parameterValueToString(r.fipeCode, "fipeCode")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"yearId"+"}", url.PathEscape(parameterValueToString(r.yearId, "yearId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInfoByFipeCodeRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
	fipeCode string
	yearId string
}

func (r ApiGetInfoByFipeCodeRequest) Execute() (*FipeResult, *http.Response, error) {
	return r.ApiService.GetInfoByFipeCodeExecute(r)
}

/*
GetInfoByFipeCode Fipe info by Fipe code

Returns the Fipe information for the vehicle (price estimation)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @param fipeCode Fipe internal reference code
 @param yearId Year for the vehicle
 @return ApiGetInfoByFipeCodeRequest
*/
func (a *FipeApiService) GetInfoByFipeCode(ctx context.Context, vehicleType VehiclesType, fipeCode string, yearId string) ApiGetInfoByFipeCodeRequest {
	return ApiGetInfoByFipeCodeRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
		fipeCode: fipeCode,
		yearId: yearId,
	}
}

// Execute executes the request
//  @return FipeResult
func (a *FipeApiService) GetInfoByFipeCodeExecute(r ApiGetInfoByFipeCodeRequest) (*FipeResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FipeResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetInfoByFipeCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/{fipeCode}/years/{yearId}"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fipeCode"+"}", url.PathEscape(parameterValueToString(r.fipeCode, "fipeCode")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"yearId"+"}", url.PathEscape(parameterValueToString(r.yearId, "yearId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetModelsByBrandRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
	brandId int32
}

func (r ApiGetModelsByBrandRequest) Execute() ([]NamedCode, *http.Response, error) {
	return r.ApiService.GetModelsByBrandExecute(r)
}

/*
GetModelsByBrand Models by brand

Returns models for the brand

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @param brandId Brand of the vehicle
 @return ApiGetModelsByBrandRequest
*/
func (a *FipeApiService) GetModelsByBrand(ctx context.Context, vehicleType VehiclesType, brandId int32) ApiGetModelsByBrandRequest {
	return ApiGetModelsByBrandRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
		brandId: brandId,
	}
}

// Execute executes the request
//  @return []NamedCode
func (a *FipeApiService) GetModelsByBrandExecute(r ApiGetModelsByBrandRequest) ([]NamedCode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []NamedCode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetModelsByBrand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/brands/{brandId}/models"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brandId"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetReferencesRequest struct {
	ctx context.Context
	ApiService FipeApi
}

func (r ApiGetReferencesRequest) Execute() ([]Reference, *http.Response, error) {
	return r.ApiService.GetReferencesExecute(r)
}

/*
GetReferences Fipe month references

Returns months and codes reference from Fipe

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetReferencesRequest
*/
func (a *FipeApiService) GetReferences(ctx context.Context) ApiGetReferencesRequest {
	return ApiGetReferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Reference
func (a *FipeApiService) GetReferencesExecute(r ApiGetReferencesRequest) ([]Reference, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Reference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetReferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/references"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetYearByModelRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
	brandId int32
	modelId int32
}

func (r ApiGetYearByModelRequest) Execute() ([]NamedCode, *http.Response, error) {
	return r.ApiService.GetYearByModelExecute(r)
}

/*
GetYearByModel Years by model

Returns years for the specific model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @param brandId Brand of the vehicle
 @param modelId Model of the vehicle
 @return ApiGetYearByModelRequest
*/
func (a *FipeApiService) GetYearByModel(ctx context.Context, vehicleType VehiclesType, brandId int32, modelId int32) ApiGetYearByModelRequest {
	return ApiGetYearByModelRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
		brandId: brandId,
		modelId: modelId,
	}
}

// Execute executes the request
//  @return []NamedCode
func (a *FipeApiService) GetYearByModelExecute(r ApiGetYearByModelRequest) ([]NamedCode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []NamedCode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetYearByModel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/brands/{brandId}/models/{modelId}/years"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brandId"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelId"+"}", url.PathEscape(parameterValueToString(r.modelId, "modelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetYearsByFipeCodeRequest struct {
	ctx context.Context
	ApiService FipeApi
	vehicleType VehiclesType
	fipeCode string
}

func (r ApiGetYearsByFipeCodeRequest) Execute() ([]NamedCode, *http.Response, error) {
	return r.ApiService.GetYearsByFipeCodeExecute(r)
}

/*
GetYearsByFipeCode Years by Fipe code

Returns years available for vehicle by fipe code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vehicleType Type of vehicle
 @param fipeCode Fipe internal reference code
 @return ApiGetYearsByFipeCodeRequest
*/
func (a *FipeApiService) GetYearsByFipeCode(ctx context.Context, vehicleType VehiclesType, fipeCode string) ApiGetYearsByFipeCodeRequest {
	return ApiGetYearsByFipeCodeRequest{
		ApiService: a,
		ctx: ctx,
		vehicleType: vehicleType,
		fipeCode: fipeCode,
	}
}

// Execute executes the request
//  @return []NamedCode
func (a *FipeApiService) GetYearsByFipeCodeExecute(r ApiGetYearsByFipeCodeRequest) ([]NamedCode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []NamedCode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FipeApiService.GetYearsByFipeCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{vehicleType}/{fipeCode}/years"
	localVarPath = strings.Replace(localVarPath, "{"+"vehicleType"+"}", url.PathEscape(parameterValueToString(r.vehicleType, "vehicleType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fipeCode"+"}", url.PathEscape(parameterValueToString(r.fipeCode, "fipeCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
