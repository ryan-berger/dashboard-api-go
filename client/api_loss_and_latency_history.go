/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// LossAndLatencyHistoryApiService LossAndLatencyHistoryApi service
type LossAndLatencyHistoryApiService service

type LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest struct {
	ctx context.Context
	ApiService *LossAndLatencyHistoryApiService
	serial string
	ip *string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	uplink *string
}

// The destination IP used to obtain the requested stats. This is required.
func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) Ip(ip string) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	r.ip = &ip
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 60 days from today.
func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) T0(t0 string) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) T1(t1 string) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) Timespan(timespan float32) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60.
func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) Resolution(resolution int32) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	r.resolution = &resolution
	return r
}

// The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, cellular. The default is wan1.
func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) Uplink(uplink string) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	r.uplink = &uplink
	return r
}

func (r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceLossAndLatencyHistoryExecute(r)
}

/*
GetDeviceLossAndLatencyHistory Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for a wired network device.

Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for a wired network device.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest
*/
func (a *LossAndLatencyHistoryApiService) GetDeviceLossAndLatencyHistory(ctx context.Context, serial string) LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest {
	return LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LossAndLatencyHistoryApiService) GetDeviceLossAndLatencyHistoryExecute(r LossAndLatencyHistoryApiGetDeviceLossAndLatencyHistoryRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LossAndLatencyHistoryApiService.GetDeviceLossAndLatencyHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/lossAndLatencyHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ip == nil {
		return localVarReturnValue, nil, reportError("ip is required and must be specified")
	}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
	}
	if r.uplink != nil {
		localVarQueryParams.Add("uplink", parameterToString(*r.uplink, ""))
	}
	localVarQueryParams.Add("ip", parameterToString(*r.ip, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
