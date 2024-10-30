// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type DashboardOption struct {
	Component string `json:"component"`
	DashboardUid string `json:"dashboardUid"`
	Variables map[string]string `json:"variables"`
	InstancePanels []DashboardOptionInstancePanelsItem `json:"instancePanels"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewDashboardOption instantiates a new DashboardOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardOption(component string, dashboardUid string, variables map[string]string, instancePanels []DashboardOptionInstancePanelsItem) *DashboardOption {
	this := DashboardOption{}
	this.Component = component
	this.DashboardUid = dashboardUid
	this.Variables = variables
	this.InstancePanels = instancePanels
	return &this
}

// NewDashboardOptionWithDefaults instantiates a new DashboardOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardOptionWithDefaults() *DashboardOption {
	this := DashboardOption{}
	return &this
}
// GetComponent returns the Component field value.
func (o *DashboardOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DashboardOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *DashboardOption) SetComponent(v string) {
	o.Component = v
}


// GetDashboardUid returns the DashboardUid field value.
func (o *DashboardOption) GetDashboardUid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DashboardUid
}

// GetDashboardUidOk returns a tuple with the DashboardUid field value
// and a boolean to check if the value has been set.
func (o *DashboardOption) GetDashboardUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardUid, true
}

// SetDashboardUid sets field value.
func (o *DashboardOption) SetDashboardUid(v string) {
	o.DashboardUid = v
}


// GetVariables returns the Variables field value.
func (o *DashboardOption) GetVariables() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *DashboardOption) GetVariablesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value.
func (o *DashboardOption) SetVariables(v map[string]string) {
	o.Variables = v
}


// GetInstancePanels returns the InstancePanels field value.
func (o *DashboardOption) GetInstancePanels() []DashboardOptionInstancePanelsItem {
	if o == nil {
		var ret []DashboardOptionInstancePanelsItem
		return ret
	}
	return o.InstancePanels
}

// GetInstancePanelsOk returns a tuple with the InstancePanels field value
// and a boolean to check if the value has been set.
func (o *DashboardOption) GetInstancePanelsOk() (*[]DashboardOptionInstancePanelsItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstancePanels, true
}

// SetInstancePanels sets field value.
func (o *DashboardOption) SetInstancePanels(v []DashboardOptionInstancePanelsItem) {
	o.InstancePanels = v
}



// MarshalJSON serializes the struct using spec logic.
func (o DashboardOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["dashboardUid"] = o.DashboardUid
	toSerialize["variables"] = o.Variables
	toSerialize["instancePanels"] = o.InstancePanels

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component"`
		DashboardUid *string `json:"dashboardUid"`
		Variables *map[string]string `json:"variables"`
		InstancePanels *[]DashboardOptionInstancePanelsItem `json:"instancePanels"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.DashboardUid == nil {
		return fmt.Errorf("required field dashboardUid missing")
	}
	if all.Variables == nil {
		return fmt.Errorf("required field variables missing")
	}
	if all.InstancePanels == nil {
		return fmt.Errorf("required field instancePanels missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "component", "dashboardUid", "variables", "instancePanels",  })
	} else {
		return err
	}
	o.Component = *all.Component
	o.DashboardUid = *all.DashboardUid
	o.Variables = *all.Variables
	o.InstancePanels = *all.InstancePanels

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
