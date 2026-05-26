// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeComponent struct {
	Component    string `json:"component"`
	HideEnpoints bool   `json:"hideEnpoints"`
	HideOnCreate bool   `json:"hideOnCreate"`
	// the name of the serviceRef defined in mode's serviceRefs.
	// If set, it means a serviceRef can be used to replace this component.
	//
	ServiceRef *string `json:"serviceRef,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeComponent instantiates a new ModeComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeComponent(component string, hideEnpoints bool, hideOnCreate bool) *ModeComponent {
	this := ModeComponent{}
	this.Component = component
	this.HideEnpoints = hideEnpoints
	this.HideOnCreate = hideOnCreate
	return &this
}

// NewModeComponentWithDefaults instantiates a new ModeComponent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeComponentWithDefaults() *ModeComponent {
	this := ModeComponent{}
	return &this
}

// GetComponent returns the Component field value.
func (o *ModeComponent) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ModeComponent) SetComponent(v string) {
	o.Component = v
}

// GetHideEnpoints returns the HideEnpoints field value.
func (o *ModeComponent) GetHideEnpoints() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HideEnpoints
}

// GetHideEnpointsOk returns a tuple with the HideEnpoints field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetHideEnpointsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideEnpoints, true
}

// SetHideEnpoints sets field value.
func (o *ModeComponent) SetHideEnpoints(v bool) {
	o.HideEnpoints = v
}

// GetHideOnCreate returns the HideOnCreate field value.
func (o *ModeComponent) GetHideOnCreate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HideOnCreate
}

// GetHideOnCreateOk returns a tuple with the HideOnCreate field value
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetHideOnCreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideOnCreate, true
}

// SetHideOnCreate sets field value.
func (o *ModeComponent) SetHideOnCreate(v bool) {
	o.HideOnCreate = v
}

// GetServiceRef returns the ServiceRef field value if set, zero value otherwise.
func (o *ModeComponent) GetServiceRef() string {
	if o == nil || o.ServiceRef == nil {
		var ret string
		return ret
	}
	return *o.ServiceRef
}

// GetServiceRefOk returns a tuple with the ServiceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponent) GetServiceRefOk() (*string, bool) {
	if o == nil || o.ServiceRef == nil {
		return nil, false
	}
	return o.ServiceRef, true
}

// HasServiceRef returns a boolean if a field has been set.
func (o *ModeComponent) HasServiceRef() bool {
	return o != nil && o.ServiceRef != nil
}

// SetServiceRef gets a reference to the given string and assigns it to the ServiceRef field.
func (o *ModeComponent) SetServiceRef(v string) {
	o.ServiceRef = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["hideEnpoints"] = o.HideEnpoints
	toSerialize["hideOnCreate"] = o.HideOnCreate
	if o.ServiceRef != nil {
		toSerialize["serviceRef"] = o.ServiceRef
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component    *string `json:"component"`
		HideEnpoints *bool   `json:"hideEnpoints"`
		HideOnCreate *bool   `json:"hideOnCreate"`
		ServiceRef   *string `json:"serviceRef,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.HideEnpoints == nil {
		return fmt.Errorf("required field hideEnpoints missing")
	}
	if all.HideOnCreate == nil {
		return fmt.Errorf("required field hideOnCreate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "hideEnpoints", "hideOnCreate", "serviceRef"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.HideEnpoints = *all.HideEnpoints
	o.HideOnCreate = *all.HideOnCreate
	o.ServiceRef = all.ServiceRef

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
