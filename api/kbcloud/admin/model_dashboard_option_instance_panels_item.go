// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DashboardOptionInstancePanelsItem struct {
	// NODESCRIPTION Role
	Role *string `json:"role,omitempty"`
	// NODESCRIPTION Panels
	Panels []DashboardOptionInstancePanelsItemPanelsItem `json:"panels,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardOptionInstancePanelsItem instantiates a new DashboardOptionInstancePanelsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardOptionInstancePanelsItem() *DashboardOptionInstancePanelsItem {
	this := DashboardOptionInstancePanelsItem{}
	return &this
}

// NewDashboardOptionInstancePanelsItemWithDefaults instantiates a new DashboardOptionInstancePanelsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardOptionInstancePanelsItemWithDefaults() *DashboardOptionInstancePanelsItem {
	this := DashboardOptionInstancePanelsItem{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *DashboardOptionInstancePanelsItem) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardOptionInstancePanelsItem) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *DashboardOptionInstancePanelsItem) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *DashboardOptionInstancePanelsItem) SetRole(v string) {
	o.Role = &v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *DashboardOptionInstancePanelsItem) GetPanels() []DashboardOptionInstancePanelsItemPanelsItem {
	if o == nil || o.Panels == nil {
		var ret []DashboardOptionInstancePanelsItemPanelsItem
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardOptionInstancePanelsItem) GetPanelsOk() (*[]DashboardOptionInstancePanelsItemPanelsItem, bool) {
	if o == nil || o.Panels == nil {
		return nil, false
	}
	return &o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *DashboardOptionInstancePanelsItem) HasPanels() bool {
	return o != nil && o.Panels != nil
}

// SetPanels gets a reference to the given []DashboardOptionInstancePanelsItemPanelsItem and assigns it to the Panels field.
func (o *DashboardOptionInstancePanelsItem) SetPanels(v []DashboardOptionInstancePanelsItemPanelsItem) {
	o.Panels = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardOptionInstancePanelsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Panels != nil {
		toSerialize["panels"] = o.Panels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardOptionInstancePanelsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role   *string                                       `json:"role,omitempty"`
		Panels []DashboardOptionInstancePanelsItemPanelsItem `json:"panels,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "panels"})
	} else {
		return err
	}
	o.Role = all.Role
	o.Panels = all.Panels

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
