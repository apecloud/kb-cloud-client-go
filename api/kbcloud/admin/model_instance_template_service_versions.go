// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type InstanceTemplateServiceVersions struct {
	Name                *string  `json:"name,omitempty"`
	CurrentVersion      *string  `json:"currentVersion,omitempty"`
	UpgradeableVersions []string `json:"upgradeableVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateServiceVersions instantiates a new InstanceTemplateServiceVersions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateServiceVersions() *InstanceTemplateServiceVersions {
	this := InstanceTemplateServiceVersions{}
	return &this
}

// NewInstanceTemplateServiceVersionsWithDefaults instantiates a new InstanceTemplateServiceVersions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateServiceVersionsWithDefaults() *InstanceTemplateServiceVersions {
	this := InstanceTemplateServiceVersions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceTemplateServiceVersions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateServiceVersions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceTemplateServiceVersions) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceTemplateServiceVersions) SetName(v string) {
	o.Name = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *InstanceTemplateServiceVersions) GetCurrentVersion() string {
	if o == nil || o.CurrentVersion == nil {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateServiceVersions) GetCurrentVersionOk() (*string, bool) {
	if o == nil || o.CurrentVersion == nil {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *InstanceTemplateServiceVersions) HasCurrentVersion() bool {
	return o != nil && o.CurrentVersion != nil
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *InstanceTemplateServiceVersions) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetUpgradeableVersions returns the UpgradeableVersions field value if set, zero value otherwise.
func (o *InstanceTemplateServiceVersions) GetUpgradeableVersions() []string {
	if o == nil || o.UpgradeableVersions == nil {
		var ret []string
		return ret
	}
	return o.UpgradeableVersions
}

// GetUpgradeableVersionsOk returns a tuple with the UpgradeableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateServiceVersions) GetUpgradeableVersionsOk() (*[]string, bool) {
	if o == nil || o.UpgradeableVersions == nil {
		return nil, false
	}
	return &o.UpgradeableVersions, true
}

// HasUpgradeableVersions returns a boolean if a field has been set.
func (o *InstanceTemplateServiceVersions) HasUpgradeableVersions() bool {
	return o != nil && o.UpgradeableVersions != nil
}

// SetUpgradeableVersions gets a reference to the given []string and assigns it to the UpgradeableVersions field.
func (o *InstanceTemplateServiceVersions) SetUpgradeableVersions(v []string) {
	o.UpgradeableVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateServiceVersions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrentVersion != nil {
		toSerialize["currentVersion"] = o.CurrentVersion
	}
	if o.UpgradeableVersions != nil {
		toSerialize["upgradeableVersions"] = o.UpgradeableVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateServiceVersions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                *string  `json:"name,omitempty"`
		CurrentVersion      *string  `json:"currentVersion,omitempty"`
		UpgradeableVersions []string `json:"upgradeableVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "currentVersion", "upgradeableVersions"})
	} else {
		return err
	}
	o.Name = all.Name
	o.CurrentVersion = all.CurrentVersion
	o.UpgradeableVersions = all.UpgradeableVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
