// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// Environment Environment info 
type Environment struct {
	// Provider
	Provider string `json:"provider"`
	// Region
	Region string `json:"region"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones"`
	// Configuration of networking for this environment
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	CreatedAt time.Time `json:"createdAt"`
	// The description of the organization
	Description *string `json:"description,omitempty"`
	// The display name of the context
	DisplayName *string `json:"displayName,omitempty"`
	// environment id
	Id uuid.UUID `json:"id"`
	// The full, unique name of this Object in the format contexts/{name}, set during creation. name must be a valid RFC 1123 compliant DNS label
	Name string `json:"name"`
	// Organization Name
	OrgName string `json:"orgName"`
	// Output only. State of the Environment resource
	State EnvironmentState `json:"state"`
	// Type of this environment
	Type EnvironmentType `json:"type"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	UpdatedAt time.Time `json:"updatedAt"`
	// Image registry used by the environment
	ImageRegistry *string `json:"imageRegistry,omitempty"`
	// extra info for environment
	ExtraInfo *string `json:"extraInfo,omitempty"`
	// namespace info for environment
	Namespaces []string `json:"namespaces,omitempty"`
	// Enable pod antiaffinity for cluster
	PodAntiAffinityEnabled *bool `json:"podAntiAffinityEnabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewEnvironment instantiates a new Environment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironment(provider string, region string, availabilityZones []string, createdAt time.Time, id uuid.UUID, name string, orgName string, state EnvironmentState, typeVar EnvironmentType, updatedAt time.Time) *Environment {
	this := Environment{}
	this.Provider = provider
	this.Region = region
	this.AvailabilityZones = availabilityZones
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.OrgName = orgName
	this.State = state
	this.Type = typeVar
	this.UpdatedAt = updatedAt
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	return &this
}
// GetProvider returns the Provider field value.
func (o *Environment) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Environment) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *Environment) SetProvider(v string) {
	o.Provider = v
}


// GetRegion returns the Region field value.
func (o *Environment) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Environment) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *Environment) SetRegion(v string) {
	o.Region = v
}


// GetAvailabilityZones returns the AvailabilityZones field value.
func (o *Environment) GetAvailabilityZones() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value
// and a boolean to check if the value has been set.
func (o *Environment) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// SetAvailabilityZones sets field value.
func (o *Environment) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}


// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise.
func (o *Environment) GetNetworkConfig() NetworkConfig {
	if o == nil || o.NetworkConfig == nil {
		var ret NetworkConfig
		return ret
	}
	return *o.NetworkConfig
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetNetworkConfigOk() (*NetworkConfig, bool) {
	if o == nil || o.NetworkConfig == nil {
		return nil, false
	}
	return o.NetworkConfig, true
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *Environment) HasNetworkConfig() bool {
	return o != nil && o.NetworkConfig != nil
}

// SetNetworkConfig gets a reference to the given NetworkConfig and assigns it to the NetworkConfig field.
func (o *Environment) SetNetworkConfig(v NetworkConfig) {
	o.NetworkConfig = &v
}


// GetCreatedAt returns the CreatedAt field value.
func (o *Environment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Environment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Environment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Environment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Environment) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Environment) SetDescription(v string) {
	o.Description = &v
}


// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Environment) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Environment) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Environment) SetDisplayName(v string) {
	o.DisplayName = &v
}


// GetId returns the Id field value.
func (o *Environment) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Environment) SetId(v uuid.UUID) {
	o.Id = v
}


// GetName returns the Name field value.
func (o *Environment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Environment) SetName(v string) {
	o.Name = v
}


// GetOrgName returns the OrgName field value.
func (o *Environment) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *Environment) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *Environment) SetOrgName(v string) {
	o.OrgName = v
}


// GetState returns the State field value.
func (o *Environment) GetState() EnvironmentState {
	if o == nil {
		var ret EnvironmentState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Environment) GetStateOk() (*EnvironmentState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *Environment) SetState(v EnvironmentState) {
	o.State = v
}


// GetType returns the Type field value.
func (o *Environment) GetType() EnvironmentType {
	if o == nil {
		var ret EnvironmentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Environment) GetTypeOk() (*EnvironmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Environment) SetType(v EnvironmentType) {
	o.Type = v
}


// GetUpdatedAt returns the UpdatedAt field value.
func (o *Environment) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Environment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Environment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}


// GetImageRegistry returns the ImageRegistry field value if set, zero value otherwise.
func (o *Environment) GetImageRegistry() string {
	if o == nil || o.ImageRegistry == nil {
		var ret string
		return ret
	}
	return *o.ImageRegistry
}

// GetImageRegistryOk returns a tuple with the ImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetImageRegistryOk() (*string, bool) {
	if o == nil || o.ImageRegistry == nil {
		return nil, false
	}
	return o.ImageRegistry, true
}

// HasImageRegistry returns a boolean if a field has been set.
func (o *Environment) HasImageRegistry() bool {
	return o != nil && o.ImageRegistry != nil
}

// SetImageRegistry gets a reference to the given string and assigns it to the ImageRegistry field.
func (o *Environment) SetImageRegistry(v string) {
	o.ImageRegistry = &v
}


// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *Environment) GetExtraInfo() string {
	if o == nil || o.ExtraInfo == nil {
		var ret string
		return ret
	}
	return *o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetExtraInfoOk() (*string, bool) {
	if o == nil || o.ExtraInfo == nil {
		return nil, false
	}
	return o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *Environment) HasExtraInfo() bool {
	return o != nil && o.ExtraInfo != nil
}

// SetExtraInfo gets a reference to the given string and assigns it to the ExtraInfo field.
func (o *Environment) SetExtraInfo(v string) {
	o.ExtraInfo = &v
}


// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *Environment) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *Environment) HasNamespaces() bool {
	return o != nil && o.Namespaces != nil
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *Environment) SetNamespaces(v []string) {
	o.Namespaces = v
}


// GetPodAntiAffinityEnabled returns the PodAntiAffinityEnabled field value if set, zero value otherwise.
func (o *Environment) GetPodAntiAffinityEnabled() bool {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PodAntiAffinityEnabled
}

// GetPodAntiAffinityEnabledOk returns a tuple with the PodAntiAffinityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetPodAntiAffinityEnabledOk() (*bool, bool) {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		return nil, false
	}
	return o.PodAntiAffinityEnabled, true
}

// HasPodAntiAffinityEnabled returns a boolean if a field has been set.
func (o *Environment) HasPodAntiAffinityEnabled() bool {
	return o != nil && o.PodAntiAffinityEnabled != nil
}

// SetPodAntiAffinityEnabled gets a reference to the given bool and assigns it to the PodAntiAffinityEnabled field.
func (o *Environment) SetPodAntiAffinityEnabled(v bool) {
	o.PodAntiAffinityEnabled = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["provider"] = o.Provider
	toSerialize["region"] = o.Region
	toSerialize["availabilityZones"] = o.AvailabilityZones
	if o.NetworkConfig != nil {
		toSerialize["networkConfig"] = o.NetworkConfig
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["orgName"] = o.OrgName
	toSerialize["state"] = o.State
	toSerialize["type"] = o.Type
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ImageRegistry != nil {
		toSerialize["imageRegistry"] = o.ImageRegistry
	}
	if o.ExtraInfo != nil {
		toSerialize["extraInfo"] = o.ExtraInfo
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.PodAntiAffinityEnabled != nil {
		toSerialize["podAntiAffinityEnabled"] = o.PodAntiAffinityEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Environment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider *string `json:"provider"`
		Region *string `json:"region"`
		AvailabilityZones *[]string `json:"availabilityZones"`
		NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`
		CreatedAt *time.Time `json:"createdAt"`
		Description *string `json:"description,omitempty"`
		DisplayName *string `json:"displayName,omitempty"`
		Id *uuid.UUID `json:"id"`
		Name *string `json:"name"`
		OrgName *string `json:"orgName"`
		State *EnvironmentState `json:"state"`
		Type *EnvironmentType `json:"type"`
		UpdatedAt *time.Time `json:"updatedAt"`
		ImageRegistry *string `json:"imageRegistry,omitempty"`
		ExtraInfo *string `json:"extraInfo,omitempty"`
		Namespaces []string `json:"namespaces,omitempty"`
		PodAntiAffinityEnabled *bool `json:"podAntiAffinityEnabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.AvailabilityZones == nil {
		return fmt.Errorf("required field availabilityZones missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "provider", "region", "availabilityZones", "networkConfig", "createdAt", "description", "displayName", "id", "name", "orgName", "state", "type", "updatedAt", "imageRegistry", "extraInfo", "namespaces", "podAntiAffinityEnabled",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Provider = *all.Provider
	o.Region = *all.Region
	o.AvailabilityZones = *all.AvailabilityZones
	if  all.NetworkConfig != nil && all.NetworkConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NetworkConfig = all.NetworkConfig
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.Id = *all.Id
	o.Name = *all.Name
	o.OrgName = *all.OrgName
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UpdatedAt = *all.UpdatedAt
	o.ImageRegistry = all.ImageRegistry
	o.ExtraInfo = all.ExtraInfo
	o.Namespaces = all.Namespaces
	o.PodAntiAffinityEnabled = all.PodAntiAffinityEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
