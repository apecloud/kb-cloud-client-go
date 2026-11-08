// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBTokenUpdate struct {
	Status InfluxDBTokenStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBTokenUpdate instantiates a new InfluxDBTokenUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBTokenUpdate(status InfluxDBTokenStatus) *InfluxDBTokenUpdate {
	this := InfluxDBTokenUpdate{}
	this.Status = status
	return &this
}

// NewInfluxDBTokenUpdateWithDefaults instantiates a new InfluxDBTokenUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBTokenUpdateWithDefaults() *InfluxDBTokenUpdate {
	this := InfluxDBTokenUpdate{}
	return &this
}

// GetStatus returns the Status field value.
func (o *InfluxDBTokenUpdate) GetStatus() InfluxDBTokenStatus {
	if o == nil {
		var ret InfluxDBTokenStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenUpdate) GetStatusOk() (*InfluxDBTokenStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *InfluxDBTokenUpdate) SetStatus(v InfluxDBTokenStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBTokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBTokenUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *InfluxDBTokenStatus `json:"status"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
