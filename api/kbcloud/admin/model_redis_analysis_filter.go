// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisFilter struct {
	Match *string `json:"match,omitempty"`
	Count *int64  `json:"count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisFilter instantiates a new RedisAnalysisFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisFilter() *RedisAnalysisFilter {
	this := RedisAnalysisFilter{}
	return &this
}

// NewRedisAnalysisFilterWithDefaults instantiates a new RedisAnalysisFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisFilterWithDefaults() *RedisAnalysisFilter {
	this := RedisAnalysisFilter{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *RedisAnalysisFilter) GetMatch() string {
	if o == nil || o.Match == nil {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisFilter) GetMatchOk() (*string, bool) {
	if o == nil || o.Match == nil {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *RedisAnalysisFilter) HasMatch() bool {
	return o != nil && o.Match != nil
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *RedisAnalysisFilter) SetMatch(v string) {
	o.Match = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RedisAnalysisFilter) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisFilter) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RedisAnalysisFilter) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RedisAnalysisFilter) SetCount(v int64) {
	o.Count = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Match *string `json:"match,omitempty"`
		Count *int64  `json:"count,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"match", "count"})
	} else {
		return err
	}
	o.Match = all.Match
	o.Count = all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
