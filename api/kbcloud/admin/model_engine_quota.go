// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineQuota struct {
	// The engine name
	Name string `json:"name"`
	// The quota of the engine
	Quota *float64 `json:"quota,omitempty"`
	// The used count of the engine
	Used *float64 `json:"used,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineQuota instantiates a new EngineQuota object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineQuota(name string) *EngineQuota {
	this := EngineQuota{}
	this.Name = name
	return &this
}

// NewEngineQuotaWithDefaults instantiates a new EngineQuota object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineQuotaWithDefaults() *EngineQuota {
	this := EngineQuota{}
	return &this
}

// GetName returns the Name field value.
func (o *EngineQuota) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineQuota) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EngineQuota) SetName(v string) {
	o.Name = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *EngineQuota) GetQuota() float64 {
	if o == nil || o.Quota == nil {
		var ret float64
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineQuota) GetQuotaOk() (*float64, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *EngineQuota) HasQuota() bool {
	return o != nil && o.Quota != nil
}

// SetQuota gets a reference to the given float64 and assigns it to the Quota field.
func (o *EngineQuota) SetQuota(v float64) {
	o.Quota = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *EngineQuota) GetUsed() float64 {
	if o == nil || o.Used == nil {
		var ret float64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineQuota) GetUsedOk() (*float64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *EngineQuota) HasUsed() bool {
	return o != nil && o.Used != nil
}

// SetUsed gets a reference to the given float64 and assigns it to the Used field.
func (o *EngineQuota) SetUsed(v float64) {
	o.Used = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineQuota) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string  `json:"name"`
		Quota *float64 `json:"quota,omitempty"`
		Used  *float64 `json:"used,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "quota", "used"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Quota = all.Quota
	o.Used = all.Used

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
