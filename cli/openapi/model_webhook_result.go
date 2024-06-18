/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WebhookResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResult{}

// WebhookResult struct for WebhookResult
type WebhookResult struct {
	Request  *WebhookResultRequest  `json:"request,omitempty"`
	Response *WebhookResultResponse `json:"response,omitempty"`
}

// NewWebhookResult instantiates a new WebhookResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResult() *WebhookResult {
	this := WebhookResult{}
	return &this
}

// NewWebhookResultWithDefaults instantiates a new WebhookResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResultWithDefaults() *WebhookResult {
	this := WebhookResult{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *WebhookResult) GetRequest() WebhookResultRequest {
	if o == nil || isNil(o.Request) {
		var ret WebhookResultRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResult) GetRequestOk() (*WebhookResultRequest, bool) {
	if o == nil || isNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *WebhookResult) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given WebhookResultRequest and assigns it to the Request field.
func (o *WebhookResult) SetRequest(v WebhookResultRequest) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *WebhookResult) GetResponse() WebhookResultResponse {
	if o == nil || isNil(o.Response) {
		var ret WebhookResultResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResult) GetResponseOk() (*WebhookResultResponse, bool) {
	if o == nil || isNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *WebhookResult) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given WebhookResultResponse and assigns it to the Response field.
func (o *WebhookResult) SetResponse(v WebhookResultResponse) {
	o.Response = &v
}

func (o WebhookResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableWebhookResult struct {
	value *WebhookResult
	isSet bool
}

func (v NullableWebhookResult) Get() *WebhookResult {
	return v.value
}

func (v *NullableWebhookResult) Set(val *WebhookResult) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResult) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResult(val *WebhookResult) *NullableWebhookResult {
	return &NullableWebhookResult{value: val, isSet: true}
}

func (v NullableWebhookResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}