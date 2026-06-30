// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisTypedSummary struct {
	Total *int64                           `json:"total,omitempty"`
	Types []RedisAnalysisSimpleTypeSummary `json:"types,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisTypedSummary instantiates a new RedisAnalysisTypedSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisTypedSummary() *RedisAnalysisTypedSummary {
	this := RedisAnalysisTypedSummary{}
	return &this
}

// NewRedisAnalysisTypedSummaryWithDefaults instantiates a new RedisAnalysisTypedSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisTypedSummaryWithDefaults() *RedisAnalysisTypedSummary {
	this := RedisAnalysisTypedSummary{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RedisAnalysisTypedSummary) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisTypedSummary) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RedisAnalysisTypedSummary) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RedisAnalysisTypedSummary) SetTotal(v int64) {
	o.Total = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *RedisAnalysisTypedSummary) GetTypes() []RedisAnalysisSimpleTypeSummary {
	if o == nil || o.Types == nil {
		var ret []RedisAnalysisSimpleTypeSummary
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisTypedSummary) GetTypesOk() (*[]RedisAnalysisSimpleTypeSummary, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return &o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *RedisAnalysisTypedSummary) HasTypes() bool {
	return o != nil && o.Types != nil
}

// SetTypes gets a reference to the given []RedisAnalysisSimpleTypeSummary and assigns it to the Types field.
func (o *RedisAnalysisTypedSummary) SetTypes(v []RedisAnalysisSimpleTypeSummary) {
	o.Types = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisTypedSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisTypedSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total *int64                           `json:"total,omitempty"`
		Types []RedisAnalysisSimpleTypeSummary `json:"types,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"total", "types"})
	} else {
		return err
	}
	o.Total = all.Total
	o.Types = all.Types

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
