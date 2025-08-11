// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type StandardResource struct {
	Limit   common.NullableFloat64 `json:"limit,omitempty"`
	Request common.NullableFloat64 `json:"request,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStandardResource instantiates a new StandardResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStandardResource() *StandardResource {
	this := StandardResource{}
	return &this
}

// NewStandardResourceWithDefaults instantiates a new StandardResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStandardResourceWithDefaults() *StandardResource {
	this := StandardResource{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardResource) GetLimit() float64 {
	if o == nil || o.Limit.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StandardResource) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *StandardResource) HasLimit() bool {
	return o != nil && o.Limit.IsSet()
}

// SetLimit gets a reference to the given common.NullableFloat64 and assigns it to the Limit field.
func (o *StandardResource) SetLimit(v float64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil.
func (o *StandardResource) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil.
func (o *StandardResource) UnsetLimit() {
	o.Limit.Unset()
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardResource) GetRequest() float64 {
	if o == nil || o.Request.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Request.Get()
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StandardResource) GetRequestOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Request.Get(), o.Request.IsSet()
}

// HasRequest returns a boolean if a field has been set.
func (o *StandardResource) HasRequest() bool {
	return o != nil && o.Request.IsSet()
}

// SetRequest gets a reference to the given common.NullableFloat64 and assigns it to the Request field.
func (o *StandardResource) SetRequest(v float64) {
	o.Request.Set(&v)
}

// SetRequestNil sets the value for Request to be an explicit nil.
func (o *StandardResource) SetRequestNil() {
	o.Request.Set(nil)
}

// UnsetRequest ensures that no value is present for Request, not even an explicit nil.
func (o *StandardResource) UnsetRequest() {
	o.Request.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o StandardResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Request.IsSet() {
		toSerialize["request"] = o.Request.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StandardResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit   common.NullableFloat64 `json:"limit,omitempty"`
		Request common.NullableFloat64 `json:"request,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"limit", "request"})
	} else {
		return err
	}
	o.Limit = all.Limit
	o.Request = all.Request

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
