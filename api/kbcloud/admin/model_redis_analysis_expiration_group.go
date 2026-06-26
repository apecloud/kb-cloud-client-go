// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisExpirationGroup struct {
	Threshold *int64  `json:"threshold,omitempty"`
	Label     *string `json:"label,omitempty"`
	// Memory bytes likely to be freed in this TTL group.
	Total       *int64   `json:"total,omitempty"`
	Keys        *int64   `json:"keys,omitempty"`
	KeyRatio    *float64 `json:"keyRatio,omitempty"`
	MemoryRatio *float64 `json:"memoryRatio,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisExpirationGroup instantiates a new RedisAnalysisExpirationGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisExpirationGroup() *RedisAnalysisExpirationGroup {
	this := RedisAnalysisExpirationGroup{}
	return &this
}

// NewRedisAnalysisExpirationGroupWithDefaults instantiates a new RedisAnalysisExpirationGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisExpirationGroupWithDefaults() *RedisAnalysisExpirationGroup {
	this := RedisAnalysisExpirationGroup{}
	return &this
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RedisAnalysisExpirationGroup) GetThreshold() int64 {
	if o == nil || o.Threshold == nil {
		var ret int64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisExpirationGroup) GetThresholdOk() (*int64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RedisAnalysisExpirationGroup) HasThreshold() bool {
	return o != nil && o.Threshold != nil
}

// SetThreshold gets a reference to the given int64 and assigns it to the Threshold field.
func (o *RedisAnalysisExpirationGroup) SetThreshold(v int64) {
	o.Threshold = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RedisAnalysisExpirationGroup) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisExpirationGroup) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RedisAnalysisExpirationGroup) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RedisAnalysisExpirationGroup) SetLabel(v string) {
	o.Label = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RedisAnalysisExpirationGroup) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisExpirationGroup) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RedisAnalysisExpirationGroup) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RedisAnalysisExpirationGroup) SetTotal(v int64) {
	o.Total = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *RedisAnalysisExpirationGroup) GetKeys() int64 {
	if o == nil || o.Keys == nil {
		var ret int64
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisExpirationGroup) GetKeysOk() (*int64, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *RedisAnalysisExpirationGroup) HasKeys() bool {
	return o != nil && o.Keys != nil
}

// SetKeys gets a reference to the given int64 and assigns it to the Keys field.
func (o *RedisAnalysisExpirationGroup) SetKeys(v int64) {
	o.Keys = &v
}

// GetKeyRatio returns the KeyRatio field value if set, zero value otherwise.
func (o *RedisAnalysisExpirationGroup) GetKeyRatio() float64 {
	if o == nil || o.KeyRatio == nil {
		var ret float64
		return ret
	}
	return *o.KeyRatio
}

// GetKeyRatioOk returns a tuple with the KeyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisExpirationGroup) GetKeyRatioOk() (*float64, bool) {
	if o == nil || o.KeyRatio == nil {
		return nil, false
	}
	return o.KeyRatio, true
}

// HasKeyRatio returns a boolean if a field has been set.
func (o *RedisAnalysisExpirationGroup) HasKeyRatio() bool {
	return o != nil && o.KeyRatio != nil
}

// SetKeyRatio gets a reference to the given float64 and assigns it to the KeyRatio field.
func (o *RedisAnalysisExpirationGroup) SetKeyRatio(v float64) {
	o.KeyRatio = &v
}

// GetMemoryRatio returns the MemoryRatio field value if set, zero value otherwise.
func (o *RedisAnalysisExpirationGroup) GetMemoryRatio() float64 {
	if o == nil || o.MemoryRatio == nil {
		var ret float64
		return ret
	}
	return *o.MemoryRatio
}

// GetMemoryRatioOk returns a tuple with the MemoryRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisExpirationGroup) GetMemoryRatioOk() (*float64, bool) {
	if o == nil || o.MemoryRatio == nil {
		return nil, false
	}
	return o.MemoryRatio, true
}

// HasMemoryRatio returns a boolean if a field has been set.
func (o *RedisAnalysisExpirationGroup) HasMemoryRatio() bool {
	return o != nil && o.MemoryRatio != nil
}

// SetMemoryRatio gets a reference to the given float64 and assigns it to the MemoryRatio field.
func (o *RedisAnalysisExpirationGroup) SetMemoryRatio(v float64) {
	o.MemoryRatio = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisExpirationGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.KeyRatio != nil {
		toSerialize["keyRatio"] = o.KeyRatio
	}
	if o.MemoryRatio != nil {
		toSerialize["memoryRatio"] = o.MemoryRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisExpirationGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Threshold   *int64   `json:"threshold,omitempty"`
		Label       *string  `json:"label,omitempty"`
		Total       *int64   `json:"total,omitempty"`
		Keys        *int64   `json:"keys,omitempty"`
		KeyRatio    *float64 `json:"keyRatio,omitempty"`
		MemoryRatio *float64 `json:"memoryRatio,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"threshold", "label", "total", "keys", "keyRatio", "memoryRatio"})
	} else {
		return err
	}
	o.Threshold = all.Threshold
	o.Label = all.Label
	o.Total = all.Total
	o.Keys = all.Keys
	o.KeyRatio = all.KeyRatio
	o.MemoryRatio = all.MemoryRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
