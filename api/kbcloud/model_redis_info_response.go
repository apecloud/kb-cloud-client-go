// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisInfoResponse struct {
	Section *string                `json:"section,omitempty"`
	NodeId  *string                `json:"nodeId,omitempty"`
	Raw     *string                `json:"raw,omitempty"`
	Values  map[string]interface{} `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisInfoResponse instantiates a new RedisInfoResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisInfoResponse() *RedisInfoResponse {
	this := RedisInfoResponse{}
	return &this
}

// NewRedisInfoResponseWithDefaults instantiates a new RedisInfoResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisInfoResponseWithDefaults() *RedisInfoResponse {
	this := RedisInfoResponse{}
	return &this
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *RedisInfoResponse) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisInfoResponse) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *RedisInfoResponse) HasSection() bool {
	return o != nil && o.Section != nil
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *RedisInfoResponse) SetSection(v string) {
	o.Section = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *RedisInfoResponse) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisInfoResponse) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *RedisInfoResponse) HasNodeId() bool {
	return o != nil && o.NodeId != nil
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *RedisInfoResponse) SetNodeId(v string) {
	o.NodeId = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *RedisInfoResponse) GetRaw() string {
	if o == nil || o.Raw == nil {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisInfoResponse) GetRawOk() (*string, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *RedisInfoResponse) HasRaw() bool {
	return o != nil && o.Raw != nil
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *RedisInfoResponse) SetRaw(v string) {
	o.Raw = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *RedisInfoResponse) GetValues() map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisInfoResponse) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RedisInfoResponse) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *RedisInfoResponse) SetValues(v map[string]interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Section != nil {
		toSerialize["section"] = o.Section
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisInfoResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Section *string                `json:"section,omitempty"`
		NodeId  *string                `json:"nodeId,omitempty"`
		Raw     *string                `json:"raw,omitempty"`
		Values  map[string]interface{} `json:"values,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"section", "nodeId", "raw", "values"})
	} else {
		return err
	}
	o.Section = all.Section
	o.NodeId = all.NodeId
	o.Raw = all.Raw
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
