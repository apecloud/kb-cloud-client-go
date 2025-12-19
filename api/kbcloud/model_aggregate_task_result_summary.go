// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AggregateTaskResultSummary struct {
	Error   *int32 `json:"error,omitempty"`
	Warning *int32 `json:"warning,omitempty"`
	Normal  *int32 `json:"normal,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregateTaskResultSummary instantiates a new AggregateTaskResultSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateTaskResultSummary() *AggregateTaskResultSummary {
	this := AggregateTaskResultSummary{}
	return &this
}

// NewAggregateTaskResultSummaryWithDefaults instantiates a new AggregateTaskResultSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateTaskResultSummaryWithDefaults() *AggregateTaskResultSummary {
	this := AggregateTaskResultSummary{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AggregateTaskResultSummary) GetError() int32 {
	if o == nil || o.Error == nil {
		var ret int32
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTaskResultSummary) GetErrorOk() (*int32, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AggregateTaskResultSummary) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given int32 and assigns it to the Error field.
func (o *AggregateTaskResultSummary) SetError(v int32) {
	o.Error = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *AggregateTaskResultSummary) GetWarning() int32 {
	if o == nil || o.Warning == nil {
		var ret int32
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTaskResultSummary) GetWarningOk() (*int32, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *AggregateTaskResultSummary) HasWarning() bool {
	return o != nil && o.Warning != nil
}

// SetWarning gets a reference to the given int32 and assigns it to the Warning field.
func (o *AggregateTaskResultSummary) SetWarning(v int32) {
	o.Warning = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *AggregateTaskResultSummary) GetNormal() int32 {
	if o == nil || o.Normal == nil {
		var ret int32
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTaskResultSummary) GetNormalOk() (*int32, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *AggregateTaskResultSummary) HasNormal() bool {
	return o != nil && o.Normal != nil
}

// SetNormal gets a reference to the given int32 and assigns it to the Normal field.
func (o *AggregateTaskResultSummary) SetNormal(v int32) {
	o.Normal = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateTaskResultSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateTaskResultSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error   *int32 `json:"error,omitempty"`
		Warning *int32 `json:"warning,omitempty"`
		Normal  *int32 `json:"normal,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"error", "warning", "normal"})
	} else {
		return err
	}
	o.Error = all.Error
	o.Warning = all.Warning
	o.Normal = all.Normal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
