// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisOperationSummary struct {
	Status     *RedisOperationStatus `json:"status,omitempty"`
	Message    *string               `json:"message,omitempty"`
	Affected   *int64                `json:"affected,omitempty"`
	DurationMs *int64                `json:"durationMs,omitempty"`
	Guard      *RedisOperationGuard  `json:"guard,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisOperationSummary instantiates a new RedisOperationSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisOperationSummary() *RedisOperationSummary {
	this := RedisOperationSummary{}
	return &this
}

// NewRedisOperationSummaryWithDefaults instantiates a new RedisOperationSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisOperationSummaryWithDefaults() *RedisOperationSummary {
	this := RedisOperationSummary{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RedisOperationSummary) GetStatus() RedisOperationStatus {
	if o == nil || o.Status == nil {
		var ret RedisOperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationSummary) GetStatusOk() (*RedisOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RedisOperationSummary) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given RedisOperationStatus and assigns it to the Status field.
func (o *RedisOperationSummary) SetStatus(v RedisOperationStatus) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RedisOperationSummary) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationSummary) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RedisOperationSummary) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RedisOperationSummary) SetMessage(v string) {
	o.Message = &v
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *RedisOperationSummary) GetAffected() int64 {
	if o == nil || o.Affected == nil {
		var ret int64
		return ret
	}
	return *o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationSummary) GetAffectedOk() (*int64, bool) {
	if o == nil || o.Affected == nil {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *RedisOperationSummary) HasAffected() bool {
	return o != nil && o.Affected != nil
}

// SetAffected gets a reference to the given int64 and assigns it to the Affected field.
func (o *RedisOperationSummary) SetAffected(v int64) {
	o.Affected = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *RedisOperationSummary) GetDurationMs() int64 {
	if o == nil || o.DurationMs == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationSummary) GetDurationMsOk() (*int64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *RedisOperationSummary) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given int64 and assigns it to the DurationMs field.
func (o *RedisOperationSummary) SetDurationMs(v int64) {
	o.DurationMs = &v
}

// GetGuard returns the Guard field value if set, zero value otherwise.
func (o *RedisOperationSummary) GetGuard() RedisOperationGuard {
	if o == nil || o.Guard == nil {
		var ret RedisOperationGuard
		return ret
	}
	return *o.Guard
}

// GetGuardOk returns a tuple with the Guard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationSummary) GetGuardOk() (*RedisOperationGuard, bool) {
	if o == nil || o.Guard == nil {
		return nil, false
	}
	return o.Guard, true
}

// HasGuard returns a boolean if a field has been set.
func (o *RedisOperationSummary) HasGuard() bool {
	return o != nil && o.Guard != nil
}

// SetGuard gets a reference to the given RedisOperationGuard and assigns it to the Guard field.
func (o *RedisOperationSummary) SetGuard(v RedisOperationGuard) {
	o.Guard = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisOperationSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Affected != nil {
		toSerialize["affected"] = o.Affected
	}
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}
	if o.Guard != nil {
		toSerialize["guard"] = o.Guard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisOperationSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status     *RedisOperationStatus `json:"status,omitempty"`
		Message    *string               `json:"message,omitempty"`
		Affected   *int64                `json:"affected,omitempty"`
		DurationMs *int64                `json:"durationMs,omitempty"`
		Guard      *RedisOperationGuard  `json:"guard,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "message", "affected", "durationMs", "guard"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Message = all.Message
	o.Affected = all.Affected
	o.DurationMs = all.DurationMs
	if all.Guard != nil && all.Guard.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Guard = all.Guard

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
