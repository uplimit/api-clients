/*
Uplimit Organization API

This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

API version: 2025-03-17
Contact: hello@uplimit.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SessionAPIService SessionAPI service
type SessionAPIService service

type ApiV1EnrollUserIntoSessionPostRequest struct {
	ctx context.Context
	ApiService *SessionAPIService
	v1EnrollUserIntoSessionPostRequest *V1EnrollUserIntoSessionPostRequest
}

func (r ApiV1EnrollUserIntoSessionPostRequest) V1EnrollUserIntoSessionPostRequest(v1EnrollUserIntoSessionPostRequest V1EnrollUserIntoSessionPostRequest) ApiV1EnrollUserIntoSessionPostRequest {
	r.v1EnrollUserIntoSessionPostRequest = &v1EnrollUserIntoSessionPostRequest
	return r
}

func (r ApiV1EnrollUserIntoSessionPostRequest) Execute() (*V1EnrollUserIntoSessionPost200Response, *http.Response, error) {
	return r.ApiService.V1EnrollUserIntoSessionPostExecute(r)
}

/*
V1EnrollUserIntoSessionPost Method for V1EnrollUserIntoSessionPost

This API allows developers to add a user into a session. The user must have already been created with the Create User API (see above) before you can add this user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1EnrollUserIntoSessionPostRequest
*/
func (a *SessionAPIService) V1EnrollUserIntoSessionPost(ctx context.Context) ApiV1EnrollUserIntoSessionPostRequest {
	return ApiV1EnrollUserIntoSessionPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1EnrollUserIntoSessionPost200Response
func (a *SessionAPIService) V1EnrollUserIntoSessionPostExecute(r ApiV1EnrollUserIntoSessionPostRequest) (*V1EnrollUserIntoSessionPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1EnrollUserIntoSessionPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionAPIService.V1EnrollUserIntoSessionPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/EnrollUserIntoSession"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.v1EnrollUserIntoSessionPostRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v V1CreateUserPost400Response
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
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiV1ListEnrollmentsInSessionGetRequest struct {
	ctx context.Context
	ApiService *SessionAPIService
	uplimitSessionId *string
	skip *int32
	take *int32
}

func (r ApiV1ListEnrollmentsInSessionGetRequest) UplimitSessionId(uplimitSessionId string) ApiV1ListEnrollmentsInSessionGetRequest {
	r.uplimitSessionId = &uplimitSessionId
	return r
}

func (r ApiV1ListEnrollmentsInSessionGetRequest) Skip(skip int32) ApiV1ListEnrollmentsInSessionGetRequest {
	r.skip = &skip
	return r
}

func (r ApiV1ListEnrollmentsInSessionGetRequest) Take(take int32) ApiV1ListEnrollmentsInSessionGetRequest {
	r.take = &take
	return r
}

func (r ApiV1ListEnrollmentsInSessionGetRequest) Execute() (*V1ListEnrollmentsInSessionGet200Response, *http.Response, error) {
	return r.ApiService.V1ListEnrollmentsInSessionGetExecute(r)
}

/*
V1ListEnrollmentsInSessionGet Method for V1ListEnrollmentsInSessionGet

This API allows developers to list all active enrollments in a session.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ListEnrollmentsInSessionGetRequest
*/
func (a *SessionAPIService) V1ListEnrollmentsInSessionGet(ctx context.Context) ApiV1ListEnrollmentsInSessionGetRequest {
	return ApiV1ListEnrollmentsInSessionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ListEnrollmentsInSessionGet200Response
func (a *SessionAPIService) V1ListEnrollmentsInSessionGetExecute(r ApiV1ListEnrollmentsInSessionGetRequest) (*V1ListEnrollmentsInSessionGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListEnrollmentsInSessionGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionAPIService.V1ListEnrollmentsInSessionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ListEnrollmentsInSession"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uplimitSessionId == nil {
		return localVarReturnValue, nil, reportError("uplimitSessionId is required and must be specified")
	}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "form", "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "take", r.take, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "uplimitSessionId", r.uplimitSessionId, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v V1CreateUserPost400Response
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
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiV1ListSessionsInCourseGetRequest struct {
	ctx context.Context
	ApiService *SessionAPIService
	uplimitCourseId *string
	skip *int32
	take *int32
}

func (r ApiV1ListSessionsInCourseGetRequest) UplimitCourseId(uplimitCourseId string) ApiV1ListSessionsInCourseGetRequest {
	r.uplimitCourseId = &uplimitCourseId
	return r
}

func (r ApiV1ListSessionsInCourseGetRequest) Skip(skip int32) ApiV1ListSessionsInCourseGetRequest {
	r.skip = &skip
	return r
}

func (r ApiV1ListSessionsInCourseGetRequest) Take(take int32) ApiV1ListSessionsInCourseGetRequest {
	r.take = &take
	return r
}

func (r ApiV1ListSessionsInCourseGetRequest) Execute() (*V1ListSessionsInCourseGet200Response, *http.Response, error) {
	return r.ApiService.V1ListSessionsInCourseGetExecute(r)
}

/*
V1ListSessionsInCourseGet Method for V1ListSessionsInCourseGet

This API allows developers to list all sessions of a course.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ListSessionsInCourseGetRequest
*/
func (a *SessionAPIService) V1ListSessionsInCourseGet(ctx context.Context) ApiV1ListSessionsInCourseGetRequest {
	return ApiV1ListSessionsInCourseGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ListSessionsInCourseGet200Response
func (a *SessionAPIService) V1ListSessionsInCourseGetExecute(r ApiV1ListSessionsInCourseGetRequest) (*V1ListSessionsInCourseGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListSessionsInCourseGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionAPIService.V1ListSessionsInCourseGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ListSessionsInCourse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uplimitCourseId == nil {
		return localVarReturnValue, nil, reportError("uplimitCourseId is required and must be specified")
	}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "form", "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "take", r.take, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "uplimitCourseId", r.uplimitCourseId, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v V1CreateUserPost400Response
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
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v V1CreateUserPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
