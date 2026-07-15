// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type SecurityApplicationPrivileges struct {
	Application *string  `json:"application,omitempty"`
	Privileges  []string `json:"privileges,omitempty"`
	Resources   []string `json:"resources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityApplicationPrivileges instantiates a new SecurityApplicationPrivileges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityApplicationPrivileges() *SecurityApplicationPrivileges {
	this := SecurityApplicationPrivileges{}
	return &this
}

// NewSecurityApplicationPrivilegesWithDefaults instantiates a new SecurityApplicationPrivileges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityApplicationPrivilegesWithDefaults() *SecurityApplicationPrivileges {
	this := SecurityApplicationPrivileges{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SecurityApplicationPrivileges) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityApplicationPrivileges) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SecurityApplicationPrivileges) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SecurityApplicationPrivileges) SetApplication(v string) {
	o.Application = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *SecurityApplicationPrivileges) GetPrivileges() []string {
	if o == nil || o.Privileges == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityApplicationPrivileges) GetPrivilegesOk() (*[]string, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *SecurityApplicationPrivileges) HasPrivileges() bool {
	return o != nil && o.Privileges != nil
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *SecurityApplicationPrivileges) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *SecurityApplicationPrivileges) GetResources() []string {
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityApplicationPrivileges) GetResourcesOk() (*[]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *SecurityApplicationPrivileges) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *SecurityApplicationPrivileges) SetResources(v []string) {
	o.Resources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityApplicationPrivileges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityApplicationPrivileges) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Application *string  `json:"application,omitempty"`
		Privileges  []string `json:"privileges,omitempty"`
		Resources   []string `json:"resources,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"application", "privileges", "resources"})
	} else {
		return err
	}
	o.Application = all.Application
	o.Privileges = all.Privileges
	o.Resources = all.Resources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
