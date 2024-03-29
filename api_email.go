/*
SendPost API

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

API version: 1.0.0
Contact: hello@sendpost.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// EmailAPIService EmailAPI service
type EmailAPIService service

type ApiSendEmailRequest struct {
	ctx context.Context
	ApiService *EmailAPIService
	xSubAccountApiKey *string
	emailMessage *EmailMessage
}

// Sub-Account API Key
func (r ApiSendEmailRequest) XSubAccountApiKey(xSubAccountApiKey string) ApiSendEmailRequest {
	r.xSubAccountApiKey = &xSubAccountApiKey
	return r
}

// Email message
func (r ApiSendEmailRequest) EmailMessage(emailMessage EmailMessage) ApiSendEmailRequest {
	r.emailMessage = &emailMessage
	return r
}

func (r ApiSendEmailRequest) Execute() ([]EmailResponse, *http.Response, error) {
	return r.ApiService.SendEmailExecute(r)
}

/*
SendEmail Method for SendEmail

Send Email To Contacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendEmailRequest
*/
func (a *EmailAPIService) SendEmail(ctx context.Context) ApiSendEmailRequest {
	return ApiSendEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EmailResponse
func (a *EmailAPIService) SendEmailExecute(r ApiSendEmailRequest) ([]EmailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EmailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailAPIService.SendEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subaccount/email/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xSubAccountApiKey == nil {
		return localVarReturnValue, nil, reportError("xSubAccountApiKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-SubAccount-ApiKey", r.xSubAccountApiKey, "")
	// body params
	localVarPostBody = r.emailMessage
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

type ApiSendEmailWithTemplateRequest struct {
	ctx context.Context
	ApiService *EmailAPIService
	xSubAccountApiKey *string
	emailMessage *EmailMessage
}

// Sub-Account API Key
func (r ApiSendEmailWithTemplateRequest) XSubAccountApiKey(xSubAccountApiKey string) ApiSendEmailWithTemplateRequest {
	r.xSubAccountApiKey = &xSubAccountApiKey
	return r
}

// Email message
func (r ApiSendEmailWithTemplateRequest) EmailMessage(emailMessage EmailMessage) ApiSendEmailWithTemplateRequest {
	r.emailMessage = &emailMessage
	return r
}

func (r ApiSendEmailWithTemplateRequest) Execute() ([]EmailResponse, *http.Response, error) {
	return r.ApiService.SendEmailWithTemplateExecute(r)
}

/*
SendEmailWithTemplate Method for SendEmailWithTemplate

Send Email To Contacts With Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendEmailWithTemplateRequest
*/
func (a *EmailAPIService) SendEmailWithTemplate(ctx context.Context) ApiSendEmailWithTemplateRequest {
	return ApiSendEmailWithTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EmailResponse
func (a *EmailAPIService) SendEmailWithTemplateExecute(r ApiSendEmailWithTemplateRequest) ([]EmailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EmailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailAPIService.SendEmailWithTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subaccount/email/template"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xSubAccountApiKey == nil {
		return localVarReturnValue, nil, reportError("xSubAccountApiKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-SubAccount-ApiKey", r.xSubAccountApiKey, "")
	// body params
	localVarPostBody = r.emailMessage
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
