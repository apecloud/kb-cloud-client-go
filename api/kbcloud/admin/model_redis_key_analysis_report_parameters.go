// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RedisKeyAnalysisReportParameters Parameters for a Redis keys analysis task. topBigKey and topSlot are capped at 500.
// numOfElemMin filters keys whose element count is below the threshold. bigKeySizeMin
// filters keys whose memory usage is below the threshold (for example "10kb").
type RedisKeyAnalysisReportParameters struct {
	TopBigKey     *int64  `json:"topBigKey,omitempty"`
	NumOfElemMin  *int64  `json:"numOfElemMin,omitempty"`
	BigKeySizeMin *string `json:"bigKeySizeMin,omitempty"`
	TopSlot       *int64  `json:"topSlot,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyAnalysisReportParameters instantiates a new RedisKeyAnalysisReportParameters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyAnalysisReportParameters() *RedisKeyAnalysisReportParameters {
	this := RedisKeyAnalysisReportParameters{}
	var topBigKey int64 = 100
	this.TopBigKey = &topBigKey
	var numOfElemMin int64 = 3000
	this.NumOfElemMin = &numOfElemMin
	var bigKeySizeMin string = "10kb"
	this.BigKeySizeMin = &bigKeySizeMin
	var topSlot int64 = 100
	this.TopSlot = &topSlot
	return &this
}

// NewRedisKeyAnalysisReportParametersWithDefaults instantiates a new RedisKeyAnalysisReportParameters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyAnalysisReportParametersWithDefaults() *RedisKeyAnalysisReportParameters {
	this := RedisKeyAnalysisReportParameters{}
	var topBigKey int64 = 100
	this.TopBigKey = &topBigKey
	var numOfElemMin int64 = 3000
	this.NumOfElemMin = &numOfElemMin
	var bigKeySizeMin string = "10kb"
	this.BigKeySizeMin = &bigKeySizeMin
	var topSlot int64 = 100
	this.TopSlot = &topSlot
	return &this
}

// GetTopBigKey returns the TopBigKey field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReportParameters) GetTopBigKey() int64 {
	if o == nil || o.TopBigKey == nil {
		var ret int64
		return ret
	}
	return *o.TopBigKey
}

// GetTopBigKeyOk returns a tuple with the TopBigKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReportParameters) GetTopBigKeyOk() (*int64, bool) {
	if o == nil || o.TopBigKey == nil {
		return nil, false
	}
	return o.TopBigKey, true
}

// HasTopBigKey returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReportParameters) HasTopBigKey() bool {
	return o != nil && o.TopBigKey != nil
}

// SetTopBigKey gets a reference to the given int64 and assigns it to the TopBigKey field.
func (o *RedisKeyAnalysisReportParameters) SetTopBigKey(v int64) {
	o.TopBigKey = &v
}

// GetNumOfElemMin returns the NumOfElemMin field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReportParameters) GetNumOfElemMin() int64 {
	if o == nil || o.NumOfElemMin == nil {
		var ret int64
		return ret
	}
	return *o.NumOfElemMin
}

// GetNumOfElemMinOk returns a tuple with the NumOfElemMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReportParameters) GetNumOfElemMinOk() (*int64, bool) {
	if o == nil || o.NumOfElemMin == nil {
		return nil, false
	}
	return o.NumOfElemMin, true
}

// HasNumOfElemMin returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReportParameters) HasNumOfElemMin() bool {
	return o != nil && o.NumOfElemMin != nil
}

// SetNumOfElemMin gets a reference to the given int64 and assigns it to the NumOfElemMin field.
func (o *RedisKeyAnalysisReportParameters) SetNumOfElemMin(v int64) {
	o.NumOfElemMin = &v
}

// GetBigKeySizeMin returns the BigKeySizeMin field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReportParameters) GetBigKeySizeMin() string {
	if o == nil || o.BigKeySizeMin == nil {
		var ret string
		return ret
	}
	return *o.BigKeySizeMin
}

// GetBigKeySizeMinOk returns a tuple with the BigKeySizeMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReportParameters) GetBigKeySizeMinOk() (*string, bool) {
	if o == nil || o.BigKeySizeMin == nil {
		return nil, false
	}
	return o.BigKeySizeMin, true
}

// HasBigKeySizeMin returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReportParameters) HasBigKeySizeMin() bool {
	return o != nil && o.BigKeySizeMin != nil
}

// SetBigKeySizeMin gets a reference to the given string and assigns it to the BigKeySizeMin field.
func (o *RedisKeyAnalysisReportParameters) SetBigKeySizeMin(v string) {
	o.BigKeySizeMin = &v
}

// GetTopSlot returns the TopSlot field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReportParameters) GetTopSlot() int64 {
	if o == nil || o.TopSlot == nil {
		var ret int64
		return ret
	}
	return *o.TopSlot
}

// GetTopSlotOk returns a tuple with the TopSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReportParameters) GetTopSlotOk() (*int64, bool) {
	if o == nil || o.TopSlot == nil {
		return nil, false
	}
	return o.TopSlot, true
}

// HasTopSlot returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReportParameters) HasTopSlot() bool {
	return o != nil && o.TopSlot != nil
}

// SetTopSlot gets a reference to the given int64 and assigns it to the TopSlot field.
func (o *RedisKeyAnalysisReportParameters) SetTopSlot(v int64) {
	o.TopSlot = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyAnalysisReportParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TopBigKey != nil {
		toSerialize["topBigKey"] = o.TopBigKey
	}
	if o.NumOfElemMin != nil {
		toSerialize["numOfElemMin"] = o.NumOfElemMin
	}
	if o.BigKeySizeMin != nil {
		toSerialize["bigKeySizeMin"] = o.BigKeySizeMin
	}
	if o.TopSlot != nil {
		toSerialize["topSlot"] = o.TopSlot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyAnalysisReportParameters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TopBigKey     *int64  `json:"topBigKey,omitempty"`
		NumOfElemMin  *int64  `json:"numOfElemMin,omitempty"`
		BigKeySizeMin *string `json:"bigKeySizeMin,omitempty"`
		TopSlot       *int64  `json:"topSlot,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topBigKey", "numOfElemMin", "bigKeySizeMin", "topSlot"})
	} else {
		return err
	}
	o.TopBigKey = all.TopBigKey
	o.NumOfElemMin = all.NumOfElemMin
	o.BigKeySizeMin = all.BigKeySizeMin
	o.TopSlot = all.TopSlot

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
