// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/google/uuid"
)

// Environment Environment info
type Environment struct {
	// Cloud Provider
	Provider string `json:"provider"`
	// Cloud Region
	Region string `json:"region"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones"`
	// Configuration of resource scheduling for this environment
	SchedulingConfig *SchedulingConfig `json:"schedulingConfig,omitempty"`
	// Configuration of networking for this environment
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`
	// The description of the organization
	Description *string `json:"description,omitempty"`
	// The display name of the context
	DisplayName string    `json:"displayName"`
	Id          uuid.UUID `json:"id"`
	// The full, unique name of this Object in the format contexts/{name}, set during creation. name must be a valid RFC 1123 compliant DNS label
	Name string `json:"name"`
	// Organizations that have access for this environment
	Organizations []string `json:"organizations"`
	// The native VM cluster is deployed in environment or not.
	MetricsMonitorEnabled *bool `json:"metricsMonitorEnabled,omitempty"`
	// Output only. State of the Environment resource
	State EnvironmentState `json:"state"`
	// Type of this environment
	Type EnvironmentType `json:"type"`
	// Configuration to provision infrastructure for this environment
	ProvisionConfig ProvisionConfig `json:"provisionConfig"`
	// cluster instance autohealing process config
	AutohealingConfig *AutohealingConfig `json:"autohealingConfig,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	UpdatedAt time.Time `json:"updatedAt"`
	// Extra info for environment
	ExtraInfo *string `json:"extraInfo,omitempty"`
	// Environment delete policy to protect environment from false delete
	DeletePolicy *EnvironmentDeletePolicy `json:"deletePolicy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironment instantiates a new Environment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironment(provider string, region string, availabilityZones []string, displayName string, id uuid.UUID, name string, organizations []string, state EnvironmentState, typeVar EnvironmentType, provisionConfig ProvisionConfig, createdAt time.Time, updatedAt time.Time) *Environment {
	this := Environment{}
	this.Provider = provider
	this.Region = region
	this.AvailabilityZones = availabilityZones
	this.DisplayName = displayName
	this.Id = id
	this.Name = name
	this.Organizations = organizations
	this.State = state
	this.Type = typeVar
	this.ProvisionConfig = provisionConfig
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	var deletePolicy EnvironmentDeletePolicy = ENVIRONMENTDELETEPOLICY_DONOTDELETE
	this.DeletePolicy = &deletePolicy
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	var deletePolicy EnvironmentDeletePolicy = ENVIRONMENTDELETEPOLICY_DONOTDELETE
	this.DeletePolicy = &deletePolicy
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

// GetSchedulingConfig returns the SchedulingConfig field value if set, zero value otherwise.
func (o *Environment) GetSchedulingConfig() SchedulingConfig {
	if o == nil || o.SchedulingConfig == nil {
		var ret SchedulingConfig
		return ret
	}
	return *o.SchedulingConfig
}

// GetSchedulingConfigOk returns a tuple with the SchedulingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetSchedulingConfigOk() (*SchedulingConfig, bool) {
	if o == nil || o.SchedulingConfig == nil {
		return nil, false
	}
	return o.SchedulingConfig, true
}

// HasSchedulingConfig returns a boolean if a field has been set.
func (o *Environment) HasSchedulingConfig() bool {
	return o != nil && o.SchedulingConfig != nil
}

// SetSchedulingConfig gets a reference to the given SchedulingConfig and assigns it to the SchedulingConfig field.
func (o *Environment) SetSchedulingConfig(v SchedulingConfig) {
	o.SchedulingConfig = &v
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

// GetDisplayName returns the DisplayName field value.
func (o *Environment) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Environment) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *Environment) SetDisplayName(v string) {
	o.DisplayName = v
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

// GetOrganizations returns the Organizations field value.
func (o *Environment) GetOrganizations() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value
// and a boolean to check if the value has been set.
func (o *Environment) GetOrganizationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organizations, true
}

// SetOrganizations sets field value.
func (o *Environment) SetOrganizations(v []string) {
	o.Organizations = v
}

// GetMetricsMonitorEnabled returns the MetricsMonitorEnabled field value if set, zero value otherwise.
func (o *Environment) GetMetricsMonitorEnabled() bool {
	if o == nil || o.MetricsMonitorEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MetricsMonitorEnabled
}

// GetMetricsMonitorEnabledOk returns a tuple with the MetricsMonitorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetMetricsMonitorEnabledOk() (*bool, bool) {
	if o == nil || o.MetricsMonitorEnabled == nil {
		return nil, false
	}
	return o.MetricsMonitorEnabled, true
}

// HasMetricsMonitorEnabled returns a boolean if a field has been set.
func (o *Environment) HasMetricsMonitorEnabled() bool {
	return o != nil && o.MetricsMonitorEnabled != nil
}

// SetMetricsMonitorEnabled gets a reference to the given bool and assigns it to the MetricsMonitorEnabled field.
func (o *Environment) SetMetricsMonitorEnabled(v bool) {
	o.MetricsMonitorEnabled = &v
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

// GetProvisionConfig returns the ProvisionConfig field value.
func (o *Environment) GetProvisionConfig() ProvisionConfig {
	if o == nil {
		var ret ProvisionConfig
		return ret
	}
	return o.ProvisionConfig
}

// GetProvisionConfigOk returns a tuple with the ProvisionConfig field value
// and a boolean to check if the value has been set.
func (o *Environment) GetProvisionConfigOk() (*ProvisionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisionConfig, true
}

// SetProvisionConfig sets field value.
func (o *Environment) SetProvisionConfig(v ProvisionConfig) {
	o.ProvisionConfig = v
}

// GetAutohealingConfig returns the AutohealingConfig field value if set, zero value otherwise.
func (o *Environment) GetAutohealingConfig() AutohealingConfig {
	if o == nil || o.AutohealingConfig == nil {
		var ret AutohealingConfig
		return ret
	}
	return *o.AutohealingConfig
}

// GetAutohealingConfigOk returns a tuple with the AutohealingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetAutohealingConfigOk() (*AutohealingConfig, bool) {
	if o == nil || o.AutohealingConfig == nil {
		return nil, false
	}
	return o.AutohealingConfig, true
}

// HasAutohealingConfig returns a boolean if a field has been set.
func (o *Environment) HasAutohealingConfig() bool {
	return o != nil && o.AutohealingConfig != nil
}

// SetAutohealingConfig gets a reference to the given AutohealingConfig and assigns it to the AutohealingConfig field.
func (o *Environment) SetAutohealingConfig(v AutohealingConfig) {
	o.AutohealingConfig = &v
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

// GetDeletePolicy returns the DeletePolicy field value if set, zero value otherwise.
func (o *Environment) GetDeletePolicy() EnvironmentDeletePolicy {
	if o == nil || o.DeletePolicy == nil {
		var ret EnvironmentDeletePolicy
		return ret
	}
	return *o.DeletePolicy
}

// GetDeletePolicyOk returns a tuple with the DeletePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetDeletePolicyOk() (*EnvironmentDeletePolicy, bool) {
	if o == nil || o.DeletePolicy == nil {
		return nil, false
	}
	return o.DeletePolicy, true
}

// HasDeletePolicy returns a boolean if a field has been set.
func (o *Environment) HasDeletePolicy() bool {
	return o != nil && o.DeletePolicy != nil
}

// SetDeletePolicy gets a reference to the given EnvironmentDeletePolicy and assigns it to the DeletePolicy field.
func (o *Environment) SetDeletePolicy(v EnvironmentDeletePolicy) {
	o.DeletePolicy = &v
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
	if o.SchedulingConfig != nil {
		toSerialize["schedulingConfig"] = o.SchedulingConfig
	}
	if o.NetworkConfig != nil {
		toSerialize["networkConfig"] = o.NetworkConfig
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["organizations"] = o.Organizations
	if o.MetricsMonitorEnabled != nil {
		toSerialize["metricsMonitorEnabled"] = o.MetricsMonitorEnabled
	}
	toSerialize["state"] = o.State
	toSerialize["type"] = o.Type
	toSerialize["provisionConfig"] = o.ProvisionConfig
	if o.AutohealingConfig != nil {
		toSerialize["autohealingConfig"] = o.AutohealingConfig
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ExtraInfo != nil {
		toSerialize["extraInfo"] = o.ExtraInfo
	}
	if o.DeletePolicy != nil {
		toSerialize["deletePolicy"] = o.DeletePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Environment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider              *string                  `json:"provider"`
		Region                *string                  `json:"region"`
		AvailabilityZones     *[]string                `json:"availabilityZones"`
		SchedulingConfig      *SchedulingConfig        `json:"schedulingConfig,omitempty"`
		NetworkConfig         *NetworkConfig           `json:"networkConfig,omitempty"`
		Description           *string                  `json:"description,omitempty"`
		DisplayName           *string                  `json:"displayName"`
		Id                    *uuid.UUID               `json:"id"`
		Name                  *string                  `json:"name"`
		Organizations         *[]string                `json:"organizations"`
		MetricsMonitorEnabled *bool                    `json:"metricsMonitorEnabled,omitempty"`
		State                 *EnvironmentState        `json:"state"`
		Type                  *EnvironmentType         `json:"type"`
		ProvisionConfig       *ProvisionConfig         `json:"provisionConfig"`
		AutohealingConfig     *AutohealingConfig       `json:"autohealingConfig,omitempty"`
		CreatedAt             *time.Time               `json:"createdAt"`
		UpdatedAt             *time.Time               `json:"updatedAt"`
		ExtraInfo             *string                  `json:"extraInfo,omitempty"`
		DeletePolicy          *EnvironmentDeletePolicy `json:"deletePolicy,omitempty"`
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
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Organizations == nil {
		return fmt.Errorf("required field organizations missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ProvisionConfig == nil {
		return fmt.Errorf("required field provisionConfig missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provider", "region", "availabilityZones", "schedulingConfig", "networkConfig", "description", "displayName", "id", "name", "organizations", "metricsMonitorEnabled", "state", "type", "provisionConfig", "autohealingConfig", "createdAt", "updatedAt", "extraInfo", "deletePolicy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Provider = *all.Provider
	o.Region = *all.Region
	o.AvailabilityZones = *all.AvailabilityZones
	if all.SchedulingConfig != nil && all.SchedulingConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SchedulingConfig = all.SchedulingConfig
	if all.NetworkConfig != nil && all.NetworkConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NetworkConfig = all.NetworkConfig
	o.Description = all.Description
	o.DisplayName = *all.DisplayName
	o.Id = *all.Id
	o.Name = *all.Name
	o.Organizations = *all.Organizations
	o.MetricsMonitorEnabled = all.MetricsMonitorEnabled
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
	if all.ProvisionConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProvisionConfig = *all.ProvisionConfig
	if all.AutohealingConfig != nil && all.AutohealingConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutohealingConfig = all.AutohealingConfig
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.ExtraInfo = all.ExtraInfo
	if all.DeletePolicy != nil && !all.DeletePolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.DeletePolicy = all.DeletePolicy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
