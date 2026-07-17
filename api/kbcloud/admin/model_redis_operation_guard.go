// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisOperationGuard struct {
	AclAllowed *bool   `json:"aclAllowed,omitempty"`
	CrossSlot  *bool   `json:"crossSlot,omitempty"`
	Dangerous  *bool   `json:"dangerous,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisOperationGuard instantiates a new RedisOperationGuard object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisOperationGuard() *RedisOperationGuard {
	this := RedisOperationGuard{}
	return &this
}

// NewRedisOperationGuardWithDefaults instantiates a new RedisOperationGuard object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisOperationGuardWithDefaults() *RedisOperationGuard {
	this := RedisOperationGuard{}
	return &this
}

// GetAclAllowed returns the AclAllowed field value if set, zero value otherwise.
func (o *RedisOperationGuard) GetAclAllowed() bool {
	if o == nil || o.AclAllowed == nil {
		var ret bool
		return ret
	}
	return *o.AclAllowed
}

// GetAclAllowedOk returns a tuple with the AclAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationGuard) GetAclAllowedOk() (*bool, bool) {
	if o == nil || o.AclAllowed == nil {
		return nil, false
	}
	return o.AclAllowed, true
}

// HasAclAllowed returns a boolean if a field has been set.
func (o *RedisOperationGuard) HasAclAllowed() bool {
	return o != nil && o.AclAllowed != nil
}

// SetAclAllowed gets a reference to the given bool and assigns it to the AclAllowed field.
func (o *RedisOperationGuard) SetAclAllowed(v bool) {
	o.AclAllowed = &v
}

// GetCrossSlot returns the CrossSlot field value if set, zero value otherwise.
func (o *RedisOperationGuard) GetCrossSlot() bool {
	if o == nil || o.CrossSlot == nil {
		var ret bool
		return ret
	}
	return *o.CrossSlot
}

// GetCrossSlotOk returns a tuple with the CrossSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationGuard) GetCrossSlotOk() (*bool, bool) {
	if o == nil || o.CrossSlot == nil {
		return nil, false
	}
	return o.CrossSlot, true
}

// HasCrossSlot returns a boolean if a field has been set.
func (o *RedisOperationGuard) HasCrossSlot() bool {
	return o != nil && o.CrossSlot != nil
}

// SetCrossSlot gets a reference to the given bool and assigns it to the CrossSlot field.
func (o *RedisOperationGuard) SetCrossSlot(v bool) {
	o.CrossSlot = &v
}

// GetDangerous returns the Dangerous field value if set, zero value otherwise.
func (o *RedisOperationGuard) GetDangerous() bool {
	if o == nil || o.Dangerous == nil {
		var ret bool
		return ret
	}
	return *o.Dangerous
}

// GetDangerousOk returns a tuple with the Dangerous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationGuard) GetDangerousOk() (*bool, bool) {
	if o == nil || o.Dangerous == nil {
		return nil, false
	}
	return o.Dangerous, true
}

// HasDangerous returns a boolean if a field has been set.
func (o *RedisOperationGuard) HasDangerous() bool {
	return o != nil && o.Dangerous != nil
}

// SetDangerous gets a reference to the given bool and assigns it to the Dangerous field.
func (o *RedisOperationGuard) SetDangerous(v bool) {
	o.Dangerous = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RedisOperationGuard) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisOperationGuard) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RedisOperationGuard) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RedisOperationGuard) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisOperationGuard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AclAllowed != nil {
		toSerialize["aclAllowed"] = o.AclAllowed
	}
	if o.CrossSlot != nil {
		toSerialize["crossSlot"] = o.CrossSlot
	}
	if o.Dangerous != nil {
		toSerialize["dangerous"] = o.Dangerous
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisOperationGuard) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AclAllowed *bool   `json:"aclAllowed,omitempty"`
		CrossSlot  *bool   `json:"crossSlot,omitempty"`
		Dangerous  *bool   `json:"dangerous,omitempty"`
		Reason     *string `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"aclAllowed", "crossSlot", "dangerous", "reason"})
	} else {
		return err
	}
	o.AclAllowed = all.AclAllowed
	o.CrossSlot = all.CrossSlot
	o.Dangerous = all.Dangerous
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
