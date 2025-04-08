/*
Paystack

The OpenAPI specification of the Paystack API that merchants and developers can harness to build financial solutions in Africa.

API version: 1.0.0
Contact: techsupport@paystack.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// PlanAPIService PlanAPI service
type PlanAPIService service

type ApiPlanCreateRequest struct {
	ctx context.Context
	ApiService *PlanAPIService
	name *string
	amount *int32
	interval *string
	description *string
	sendInvoices *bool
	sendSms *bool
	currency *string
	invoiceLimit *int32
}

// Name of plan
func (r ApiPlanCreateRequest) Name(name string) ApiPlanCreateRequest {
	r.name = &name
	return r
}

// Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
func (r ApiPlanCreateRequest) Amount(amount int32) ApiPlanCreateRequest {
	r.amount = &amount
	return r
}

// Interval in words. Valid intervals are daily, weekly, monthly,biannually, annually
func (r ApiPlanCreateRequest) Interval(interval string) ApiPlanCreateRequest {
	r.interval = &interval
	return r
}

// A description for this plan
func (r ApiPlanCreateRequest) Description(description string) ApiPlanCreateRequest {
	r.description = &description
	return r
}

// Set to false if you don&#39;t want invoices to be sent to your customers
func (r ApiPlanCreateRequest) SendInvoices(sendInvoices bool) ApiPlanCreateRequest {
	r.sendInvoices = &sendInvoices
	return r
}

// Set to false if you don&#39;t want text messages to be sent to your customers
func (r ApiPlanCreateRequest) SendSms(sendSms bool) ApiPlanCreateRequest {
	r.sendSms = &sendSms
	return r
}

// Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD
func (r ApiPlanCreateRequest) Currency(currency string) ApiPlanCreateRequest {
	r.currency = &currency
	return r
}

// Number of invoices to raise during subscription to this plan.  Can be overridden by specifying an invoice_limit while subscribing.
func (r ApiPlanCreateRequest) InvoiceLimit(invoiceLimit int32) ApiPlanCreateRequest {
	r.invoiceLimit = &invoiceLimit
	return r
}

func (r ApiPlanCreateRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.PlanCreateExecute(r)
}

/*
PlanCreate Create Plan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPlanCreateRequest
*/
func (a *PlanAPIService) PlanCreate(ctx context.Context) ApiPlanCreateRequest {
	return ApiPlanCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Response
func (a *PlanAPIService) PlanCreateExecute(r ApiPlanCreateRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlanAPIService.PlanCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plan"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "amount", r.amount, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "interval", r.interval, "", "")
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.sendInvoices != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_invoices", r.sendInvoices, "", "")
	}
	if r.sendSms != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_sms", r.sendSms, "", "")
	}
	if r.currency != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "currency", r.currency, "", "")
	}
	if r.invoiceLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invoice_limit", r.invoiceLimit, "", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiPlanFetchRequest struct {
	ctx context.Context
	ApiService *PlanAPIService
	code string
}

func (r ApiPlanFetchRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.PlanFetchExecute(r)
}

/*
PlanFetch Fetch Plan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code
 @return ApiPlanFetchRequest
*/
func (a *PlanAPIService) PlanFetch(ctx context.Context, code string) ApiPlanFetchRequest {
	return ApiPlanFetchRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return Response
func (a *PlanAPIService) PlanFetchExecute(r ApiPlanFetchRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlanAPIService.PlanFetch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plan/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiPlanListRequest struct {
	ctx context.Context
	ApiService *PlanAPIService
	perPage *int32
	page *int32
	interval *string
	amount *int32
	from *time.Time
	to *time.Time
}

// Number of records to fetch per page
func (r ApiPlanListRequest) PerPage(perPage int32) ApiPlanListRequest {
	r.perPage = &perPage
	return r
}

// The section to retrieve
func (r ApiPlanListRequest) Page(page int32) ApiPlanListRequest {
	r.page = &page
	return r
}

// Specify interval of the plan
func (r ApiPlanListRequest) Interval(interval string) ApiPlanListRequest {
	r.interval = &interval
	return r
}

// The amount on the plans to retrieve
func (r ApiPlanListRequest) Amount(amount int32) ApiPlanListRequest {
	r.amount = &amount
	return r
}

// The start date
func (r ApiPlanListRequest) From(from time.Time) ApiPlanListRequest {
	r.from = &from
	return r
}

// The end date
func (r ApiPlanListRequest) To(to time.Time) ApiPlanListRequest {
	r.to = &to
	return r
}

func (r ApiPlanListRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.PlanListExecute(r)
}

/*
PlanList List Plans

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPlanListRequest
*/
func (a *PlanAPIService) PlanList(ctx context.Context) ApiPlanListRequest {
	return ApiPlanListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Response
func (a *PlanAPIService) PlanListExecute(r ApiPlanListRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlanAPIService.PlanList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plan"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	}
	if r.amount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	}
	if r.to != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiPlanUpdateRequest struct {
	ctx context.Context
	ApiService *PlanAPIService
	code string
	name *string
	amount *int32
	interval *string
	description *bool
	sendInvoices *bool
	sendSms *bool
	currency *string
	invoiceLimit *int32
}

// Name of plan
func (r ApiPlanUpdateRequest) Name(name string) ApiPlanUpdateRequest {
	r.name = &name
	return r
}

// Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
func (r ApiPlanUpdateRequest) Amount(amount int32) ApiPlanUpdateRequest {
	r.amount = &amount
	return r
}

// Interval in words. Valid intervals are daily, weekly, monthly,biannually, annually
func (r ApiPlanUpdateRequest) Interval(interval string) ApiPlanUpdateRequest {
	r.interval = &interval
	return r
}

// A description for this plan
func (r ApiPlanUpdateRequest) Description(description bool) ApiPlanUpdateRequest {
	r.description = &description
	return r
}

// Set to false if you don&#39;t want invoices to be sent to your customers
func (r ApiPlanUpdateRequest) SendInvoices(sendInvoices bool) ApiPlanUpdateRequest {
	r.sendInvoices = &sendInvoices
	return r
}

// Set to false if you don&#39;t want text messages to be sent to your customers
func (r ApiPlanUpdateRequest) SendSms(sendSms bool) ApiPlanUpdateRequest {
	r.sendSms = &sendSms
	return r
}

// Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD
func (r ApiPlanUpdateRequest) Currency(currency string) ApiPlanUpdateRequest {
	r.currency = &currency
	return r
}

// Number of invoices to raise during subscription to this plan.  Can be overridden by specifying an invoice_limit while subscribing.
func (r ApiPlanUpdateRequest) InvoiceLimit(invoiceLimit int32) ApiPlanUpdateRequest {
	r.invoiceLimit = &invoiceLimit
	return r
}

func (r ApiPlanUpdateRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.PlanUpdateExecute(r)
}

/*
PlanUpdate Update Plan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code
 @return ApiPlanUpdateRequest
*/
func (a *PlanAPIService) PlanUpdate(ctx context.Context, code string) ApiPlanUpdateRequest {
	return ApiPlanUpdateRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return Response
func (a *PlanAPIService) PlanUpdateExecute(r ApiPlanUpdateRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlanAPIService.PlanUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plan/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.amount != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "amount", r.amount, "", "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "interval", r.interval, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.sendInvoices != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_invoices", r.sendInvoices, "", "")
	}
	if r.sendSms != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_sms", r.sendSms, "", "")
	}
	if r.currency != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "currency", r.currency, "", "")
	}
	if r.invoiceLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invoice_limit", r.invoiceLimit, "", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
