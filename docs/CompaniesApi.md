# \CompaniesApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPortfolio**](CompaniesApi.md#ListPortfolio) | **Get** /portfolio | 
[**ListQuotes**](CompaniesApi.md#ListQuotes) | **Get** /quotes/{ticker} | 
[**ListScan**](CompaniesApi.md#ListScan) | **Get** /scan | 



## ListPortfolio

> Portfolio ListPortfolio(ctx, )



List properties of portfolio

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Portfolio**](portfolio.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuotes

> Quotes ListQuotes(ctx, ticker, optional)



List properties of quotes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string**| string ticker (name or id) of the quotes | 
 **optional** | ***ListQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListQuotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routing** | **optional.String**| string routing (name or id) of the quotes | 
 **exchange** | **optional.String**| string exchange (name or id) of the quotes | 
 **currency** | **optional.String**| string currency (name or id) of the quotes | 
 **durationSize** | **optional.String**| string duration_size (name or id) of the quotes | 
 **durationUnit** | **optional.String**| string duration_unit (name or id) of the quotes | 
 **barSize** | **optional.String**| string bar_size (name or id) of the quotes | 
 **barUnit** | **optional.String**| string bar_unit (name or id) of the quotes | 

### Return type

[**Quotes**](quotes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScan

> Scan ListScan(ctx, optional)



List properties of scan

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListScanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListScanOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| string type (name or id) of the scan | 
 **locationCode** | **optional.String**| string locationCode (name or id) of the scan | 
 **scanCode** | **optional.String**| string scanCode (name or id) of the scan | 
 **aboveVolume** | **optional.String**| string aboveVolume (name or id) of the scan | 
 **limit** | **optional.String**| string limit (name or id) of the scan | 

### Return type

[**Scan**](scan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

