// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MssqlCachedPlan Cached compiled plan for the selected live request. No captured SQL is executed. This is not an actual runtime plan.
type MssqlCachedPlan struct {
	CapturedAt string `json:"capturedAt"`
	Xml        string `json:"xml"`
	Available  bool   `json:"available"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlCachedPlan instantiates a new MssqlCachedPlan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlCachedPlan(capturedAt string, xml string, available bool) *MssqlCachedPlan {
	this := MssqlCachedPlan{}
	this.CapturedAt = capturedAt
	this.Xml = xml
	this.Available = available
	return &this
}

// NewMssqlCachedPlanWithDefaults instantiates a new MssqlCachedPlan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlCachedPlanWithDefaults() *MssqlCachedPlan {
	this := MssqlCachedPlan{}
	return &this
}

// GetCapturedAt returns the CapturedAt field value.
func (o *MssqlCachedPlan) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlCachedPlan) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *MssqlCachedPlan) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// GetXml returns the Xml field value.
func (o *MssqlCachedPlan) GetXml() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Xml
}

// GetXmlOk returns a tuple with the Xml field value
// and a boolean to check if the value has been set.
func (o *MssqlCachedPlan) GetXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Xml, true
}

// SetXml sets field value.
func (o *MssqlCachedPlan) SetXml(v string) {
	o.Xml = v
}

// GetAvailable returns the Available field value.
func (o *MssqlCachedPlan) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *MssqlCachedPlan) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value.
func (o *MssqlCachedPlan) SetAvailable(v bool) {
	o.Available = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlCachedPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["capturedAt"] = o.CapturedAt
	toSerialize["xml"] = o.Xml
	toSerialize["available"] = o.Available

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlCachedPlan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapturedAt *string `json:"capturedAt"`
		Xml        *string `json:"xml"`
		Available  *bool   `json:"available"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	if all.Xml == nil {
		return fmt.Errorf("required field xml missing")
	}
	if all.Available == nil {
		return fmt.Errorf("required field available missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capturedAt", "xml", "available"})
	} else {
		return err
	}
	o.CapturedAt = *all.CapturedAt
	o.Xml = *all.Xml
	o.Available = *all.Available

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
