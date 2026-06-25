// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisKeyMutateResponse struct {
	OperationSummary *RedisOperationSummary `json:"operationSummary,omitempty"`
	Metadata         *RedisKeySummary       `json:"metadata,omitempty"`
	// Discriminated Redis key value payload for string, hash, list, set, zset, stream, and json editors.
	ValuePreview *RedisKeyValue `json:"valuePreview,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyMutateResponse instantiates a new RedisKeyMutateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyMutateResponse() *RedisKeyMutateResponse {
	this := RedisKeyMutateResponse{}
	return &this
}

// NewRedisKeyMutateResponseWithDefaults instantiates a new RedisKeyMutateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyMutateResponseWithDefaults() *RedisKeyMutateResponse {
	this := RedisKeyMutateResponse{}
	return &this
}

// GetOperationSummary returns the OperationSummary field value if set, zero value otherwise.
func (o *RedisKeyMutateResponse) GetOperationSummary() RedisOperationSummary {
	if o == nil || o.OperationSummary == nil {
		var ret RedisOperationSummary
		return ret
	}
	return *o.OperationSummary
}

// GetOperationSummaryOk returns a tuple with the OperationSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateResponse) GetOperationSummaryOk() (*RedisOperationSummary, bool) {
	if o == nil || o.OperationSummary == nil {
		return nil, false
	}
	return o.OperationSummary, true
}

// HasOperationSummary returns a boolean if a field has been set.
func (o *RedisKeyMutateResponse) HasOperationSummary() bool {
	return o != nil && o.OperationSummary != nil
}

// SetOperationSummary gets a reference to the given RedisOperationSummary and assigns it to the OperationSummary field.
func (o *RedisKeyMutateResponse) SetOperationSummary(v RedisOperationSummary) {
	o.OperationSummary = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RedisKeyMutateResponse) GetMetadata() RedisKeySummary {
	if o == nil || o.Metadata == nil {
		var ret RedisKeySummary
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateResponse) GetMetadataOk() (*RedisKeySummary, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RedisKeyMutateResponse) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given RedisKeySummary and assigns it to the Metadata field.
func (o *RedisKeyMutateResponse) SetMetadata(v RedisKeySummary) {
	o.Metadata = &v
}

// GetValuePreview returns the ValuePreview field value if set, zero value otherwise.
func (o *RedisKeyMutateResponse) GetValuePreview() RedisKeyValue {
	if o == nil || o.ValuePreview == nil {
		var ret RedisKeyValue
		return ret
	}
	return *o.ValuePreview
}

// GetValuePreviewOk returns a tuple with the ValuePreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateResponse) GetValuePreviewOk() (*RedisKeyValue, bool) {
	if o == nil || o.ValuePreview == nil {
		return nil, false
	}
	return o.ValuePreview, true
}

// HasValuePreview returns a boolean if a field has been set.
func (o *RedisKeyMutateResponse) HasValuePreview() bool {
	return o != nil && o.ValuePreview != nil
}

// SetValuePreview gets a reference to the given RedisKeyValue and assigns it to the ValuePreview field.
func (o *RedisKeyMutateResponse) SetValuePreview(v RedisKeyValue) {
	o.ValuePreview = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyMutateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OperationSummary != nil {
		toSerialize["operationSummary"] = o.OperationSummary
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ValuePreview != nil {
		toSerialize["valuePreview"] = o.ValuePreview
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyMutateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OperationSummary *RedisOperationSummary `json:"operationSummary,omitempty"`
		Metadata         *RedisKeySummary       `json:"metadata,omitempty"`
		ValuePreview     *RedisKeyValue         `json:"valuePreview,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"operationSummary", "metadata", "valuePreview"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OperationSummary != nil && all.OperationSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OperationSummary = all.OperationSummary
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.ValuePreview != nil && all.ValuePreview.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ValuePreview = all.ValuePreview

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
