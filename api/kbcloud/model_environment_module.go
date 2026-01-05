// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModule Single environment module information
type EnvironmentModule struct {
	// Environment module name
	Name string `json:"name"`
	// Environment module version
	Version *string `json:"version,omitempty"`
	// Status of environment module
	Status EnvironmentModuleStatus `json:"status"`
	// Number of replicas
	Replicas *int32 `json:"replicas,omitempty"`
	// Deployment location
	Location *string `json:"location,omitempty"`
	// Cluster information
	ClusterInfo *ClusterInfo          `json:"clusterInfo,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	DisplayName *LocalizedDescription `json:"displayName,omitempty"`
	// indicate module is optional. If false, the module is required and must be installed
	Optional *bool `json:"optional,omitempty"`
	// indicate whether module is installed by default when creating environment. Only effective when optional is true
	DefaultInstalled *bool `json:"defaultInstalled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModule instantiates a new EnvironmentModule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModule(name string, status EnvironmentModuleStatus) *EnvironmentModule {
	this := EnvironmentModule{}
	this.Name = name
	this.Status = status
	return &this
}

// NewEnvironmentModuleWithDefaults instantiates a new EnvironmentModule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleWithDefaults() *EnvironmentModule {
	this := EnvironmentModule{}
	return &this
}

// GetName returns the Name field value.
func (o *EnvironmentModule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentModule) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EnvironmentModule) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EnvironmentModule) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EnvironmentModule) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value.
func (o *EnvironmentModule) GetStatus() EnvironmentModuleStatus {
	if o == nil {
		var ret EnvironmentModuleStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetStatusOk() (*EnvironmentModuleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *EnvironmentModule) SetStatus(v EnvironmentModuleStatus) {
	o.Status = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *EnvironmentModule) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *EnvironmentModule) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *EnvironmentModule) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *EnvironmentModule) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *EnvironmentModule) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *EnvironmentModule) SetLocation(v string) {
	o.Location = &v
}

// GetClusterInfo returns the ClusterInfo field value if set, zero value otherwise.
func (o *EnvironmentModule) GetClusterInfo() ClusterInfo {
	if o == nil || o.ClusterInfo == nil {
		var ret ClusterInfo
		return ret
	}
	return *o.ClusterInfo
}

// GetClusterInfoOk returns a tuple with the ClusterInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetClusterInfoOk() (*ClusterInfo, bool) {
	if o == nil || o.ClusterInfo == nil {
		return nil, false
	}
	return o.ClusterInfo, true
}

// HasClusterInfo returns a boolean if a field has been set.
func (o *EnvironmentModule) HasClusterInfo() bool {
	return o != nil && o.ClusterInfo != nil
}

// SetClusterInfo gets a reference to the given ClusterInfo and assigns it to the ClusterInfo field.
func (o *EnvironmentModule) SetClusterInfo(v ClusterInfo) {
	o.ClusterInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentModule) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentModule) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *EnvironmentModule) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EnvironmentModule) GetDisplayName() LocalizedDescription {
	if o == nil || o.DisplayName == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EnvironmentModule) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given LocalizedDescription and assigns it to the DisplayName field.
func (o *EnvironmentModule) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *EnvironmentModule) GetOptional() bool {
	if o == nil || o.Optional == nil {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetOptionalOk() (*bool, bool) {
	if o == nil || o.Optional == nil {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *EnvironmentModule) HasOptional() bool {
	return o != nil && o.Optional != nil
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *EnvironmentModule) SetOptional(v bool) {
	o.Optional = &v
}

// GetDefaultInstalled returns the DefaultInstalled field value if set, zero value otherwise.
func (o *EnvironmentModule) GetDefaultInstalled() bool {
	if o == nil || o.DefaultInstalled == nil {
		var ret bool
		return ret
	}
	return *o.DefaultInstalled
}

// GetDefaultInstalledOk returns a tuple with the DefaultInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetDefaultInstalledOk() (*bool, bool) {
	if o == nil || o.DefaultInstalled == nil {
		return nil, false
	}
	return o.DefaultInstalled, true
}

// HasDefaultInstalled returns a boolean if a field has been set.
func (o *EnvironmentModule) HasDefaultInstalled() bool {
	return o != nil && o.DefaultInstalled != nil
}

// SetDefaultInstalled gets a reference to the given bool and assigns it to the DefaultInstalled field.
func (o *EnvironmentModule) SetDefaultInstalled(v bool) {
	o.DefaultInstalled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	toSerialize["status"] = o.Status
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.ClusterInfo != nil {
		toSerialize["clusterInfo"] = o.ClusterInfo
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Optional != nil {
		toSerialize["optional"] = o.Optional
	}
	if o.DefaultInstalled != nil {
		toSerialize["defaultInstalled"] = o.DefaultInstalled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string                  `json:"name"`
		Version          *string                  `json:"version,omitempty"`
		Status           *EnvironmentModuleStatus `json:"status"`
		Replicas         *int32                   `json:"replicas,omitempty"`
		Location         *string                  `json:"location,omitempty"`
		ClusterInfo      *ClusterInfo             `json:"clusterInfo,omitempty"`
		Description      *LocalizedDescription    `json:"description,omitempty"`
		DisplayName      *LocalizedDescription    `json:"displayName,omitempty"`
		Optional         *bool                    `json:"optional,omitempty"`
		DefaultInstalled *bool                    `json:"defaultInstalled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "version", "status", "replicas", "location", "clusterInfo", "description", "displayName", "optional", "defaultInstalled"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Version = all.Version
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Replicas = all.Replicas
	o.Location = all.Location
	if all.ClusterInfo != nil && all.ClusterInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClusterInfo = all.ClusterInfo
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	if all.DisplayName != nil && all.DisplayName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayName = all.DisplayName
	o.Optional = all.Optional
	o.DefaultInstalled = all.DefaultInstalled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
