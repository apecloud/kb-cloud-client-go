// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/google/uuid"
)

// EnvironmentCreate Environment creation info
type EnvironmentCreate struct {
	// The full, unique name of this Object in the format environments/{name}, set during creation. name must be a valid RFC 1123 compliant DNS label
	Name string `json:"name"`
	// Type of this environment
	Type EnvironmentType `json:"type"`
	// Configuration of resource scheduling for this environment
	SchedulingConfig *SchedulingConfig `json:"schedulingConfig,omitempty"`
	// Configuration to provision infrastructure for this environment
	ProvisionConfig ProvisionConfig `json:"provisionConfig"`
	// Organizations that have access for this environment
	Organizations []string `json:"organizations"`
	// Cloud Provider
	Provider string `json:"provider"`
	// Cloud Region
	Region string `json:"region"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// The description of the organization
	Description *string `json:"description,omitempty"`
	// The display name of the environment
	DisplayName string     `json:"displayName"`
	Id          *uuid.UUID `json:"id,omitempty"`
	// Extra info for environment
	ExtraInfo *string `json:"extraInfo,omitempty"`
	// Environment delete policy to protect environment from false delete
	DeletePolicy *EnvironmentDeletePolicy `json:"deletePolicy,omitempty"`
	// overwrite a environment if it has been added before
	Overwrite *bool `json:"overwrite,omitempty"`
	Dns       *Dns  `json:"dns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCreate instantiates a new EnvironmentCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCreate(name string, typeVar EnvironmentType, provisionConfig ProvisionConfig, organizations []string, provider string, region string, displayName string) *EnvironmentCreate {
	this := EnvironmentCreate{}
	this.Name = name
	this.Type = typeVar
	this.ProvisionConfig = provisionConfig
	this.Organizations = organizations
	this.Provider = provider
	this.Region = region
	this.DisplayName = displayName
	var deletePolicy EnvironmentDeletePolicy = EnvironmentDeletePolicyDoNotDelete
	this.DeletePolicy = &deletePolicy
	var overwrite bool = false
	this.Overwrite = &overwrite
	return &this
}

// NewEnvironmentCreateWithDefaults instantiates a new EnvironmentCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCreateWithDefaults() *EnvironmentCreate {
	this := EnvironmentCreate{}
	var deletePolicy EnvironmentDeletePolicy = EnvironmentDeletePolicyDoNotDelete
	this.DeletePolicy = &deletePolicy
	var overwrite bool = false
	this.Overwrite = &overwrite
	return &this
}

// GetName returns the Name field value.
func (o *EnvironmentCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentCreate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *EnvironmentCreate) GetType() EnvironmentType {
	if o == nil {
		var ret EnvironmentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetTypeOk() (*EnvironmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EnvironmentCreate) SetType(v EnvironmentType) {
	o.Type = v
}

// GetSchedulingConfig returns the SchedulingConfig field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetSchedulingConfig() SchedulingConfig {
	if o == nil || o.SchedulingConfig == nil {
		var ret SchedulingConfig
		return ret
	}
	return *o.SchedulingConfig
}

// GetSchedulingConfigOk returns a tuple with the SchedulingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetSchedulingConfigOk() (*SchedulingConfig, bool) {
	if o == nil || o.SchedulingConfig == nil {
		return nil, false
	}
	return o.SchedulingConfig, true
}

// HasSchedulingConfig returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasSchedulingConfig() bool {
	return o != nil && o.SchedulingConfig != nil
}

// SetSchedulingConfig gets a reference to the given SchedulingConfig and assigns it to the SchedulingConfig field.
func (o *EnvironmentCreate) SetSchedulingConfig(v SchedulingConfig) {
	o.SchedulingConfig = &v
}

// GetProvisionConfig returns the ProvisionConfig field value.
func (o *EnvironmentCreate) GetProvisionConfig() ProvisionConfig {
	if o == nil {
		var ret ProvisionConfig
		return ret
	}
	return o.ProvisionConfig
}

// GetProvisionConfigOk returns a tuple with the ProvisionConfig field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetProvisionConfigOk() (*ProvisionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisionConfig, true
}

// SetProvisionConfig sets field value.
func (o *EnvironmentCreate) SetProvisionConfig(v ProvisionConfig) {
	o.ProvisionConfig = v
}

// GetOrganizations returns the Organizations field value.
func (o *EnvironmentCreate) GetOrganizations() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetOrganizationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organizations, true
}

// SetOrganizations sets field value.
func (o *EnvironmentCreate) SetOrganizations(v []string) {
	o.Organizations = v
}

// GetProvider returns the Provider field value.
func (o *EnvironmentCreate) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *EnvironmentCreate) SetProvider(v string) {
	o.Provider = v
}

// GetRegion returns the Region field value.
func (o *EnvironmentCreate) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *EnvironmentCreate) SetRegion(v string) {
	o.Region = v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasAvailabilityZones() bool {
	return o != nil && o.AvailabilityZones != nil
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *EnvironmentCreate) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentCreate) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value.
func (o *EnvironmentCreate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *EnvironmentCreate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *EnvironmentCreate) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetExtraInfo() string {
	if o == nil || o.ExtraInfo == nil {
		var ret string
		return ret
	}
	return *o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetExtraInfoOk() (*string, bool) {
	if o == nil || o.ExtraInfo == nil {
		return nil, false
	}
	return o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasExtraInfo() bool {
	return o != nil && o.ExtraInfo != nil
}

// SetExtraInfo gets a reference to the given string and assigns it to the ExtraInfo field.
func (o *EnvironmentCreate) SetExtraInfo(v string) {
	o.ExtraInfo = &v
}

// GetDeletePolicy returns the DeletePolicy field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetDeletePolicy() EnvironmentDeletePolicy {
	if o == nil || o.DeletePolicy == nil {
		var ret EnvironmentDeletePolicy
		return ret
	}
	return *o.DeletePolicy
}

// GetDeletePolicyOk returns a tuple with the DeletePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetDeletePolicyOk() (*EnvironmentDeletePolicy, bool) {
	if o == nil || o.DeletePolicy == nil {
		return nil, false
	}
	return o.DeletePolicy, true
}

// HasDeletePolicy returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasDeletePolicy() bool {
	return o != nil && o.DeletePolicy != nil
}

// SetDeletePolicy gets a reference to the given EnvironmentDeletePolicy and assigns it to the DeletePolicy field.
func (o *EnvironmentCreate) SetDeletePolicy(v EnvironmentDeletePolicy) {
	o.DeletePolicy = &v
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetOverwrite() bool {
	if o == nil || o.Overwrite == nil {
		var ret bool
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetOverwriteOk() (*bool, bool) {
	if o == nil || o.Overwrite == nil {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasOverwrite() bool {
	return o != nil && o.Overwrite != nil
}

// SetOverwrite gets a reference to the given bool and assigns it to the Overwrite field.
func (o *EnvironmentCreate) SetOverwrite(v bool) {
	o.Overwrite = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetDns() Dns {
	if o == nil || o.Dns == nil {
		var ret Dns
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetDnsOk() (*Dns, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasDns() bool {
	return o != nil && o.Dns != nil
}

// SetDns gets a reference to the given Dns and assigns it to the Dns field.
func (o *EnvironmentCreate) SetDns(v Dns) {
	o.Dns = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.SchedulingConfig != nil {
		toSerialize["schedulingConfig"] = o.SchedulingConfig
	}
	toSerialize["provisionConfig"] = o.ProvisionConfig
	toSerialize["organizations"] = o.Organizations
	toSerialize["provider"] = o.Provider
	toSerialize["region"] = o.Region
	if o.AvailabilityZones != nil {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExtraInfo != nil {
		toSerialize["extraInfo"] = o.ExtraInfo
	}
	if o.DeletePolicy != nil {
		toSerialize["deletePolicy"] = o.DeletePolicy
	}
	if o.Overwrite != nil {
		toSerialize["overwrite"] = o.Overwrite
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string                  `json:"name"`
		Type              *EnvironmentType         `json:"type"`
		SchedulingConfig  *SchedulingConfig        `json:"schedulingConfig,omitempty"`
		ProvisionConfig   *ProvisionConfig         `json:"provisionConfig"`
		Organizations     *[]string                `json:"organizations"`
		Provider          *string                  `json:"provider"`
		Region            *string                  `json:"region"`
		AvailabilityZones []string                 `json:"availabilityZones,omitempty"`
		Description       *string                  `json:"description,omitempty"`
		DisplayName       *string                  `json:"displayName"`
		Id                *uuid.UUID               `json:"id,omitempty"`
		ExtraInfo         *string                  `json:"extraInfo,omitempty"`
		DeletePolicy      *EnvironmentDeletePolicy `json:"deletePolicy,omitempty"`
		Overwrite         *bool                    `json:"overwrite,omitempty"`
		Dns               *Dns                     `json:"dns,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ProvisionConfig == nil {
		return fmt.Errorf("required field provisionConfig missing")
	}
	if all.Organizations == nil {
		return fmt.Errorf("required field organizations missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "schedulingConfig", "provisionConfig", "organizations", "provider", "region", "availabilityZones", "description", "displayName", "id", "extraInfo", "deletePolicy", "overwrite", "dns"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.SchedulingConfig != nil && all.SchedulingConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SchedulingConfig = all.SchedulingConfig
	if all.ProvisionConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProvisionConfig = *all.ProvisionConfig
	o.Organizations = *all.Organizations
	o.Provider = *all.Provider
	o.Region = *all.Region
	o.AvailabilityZones = all.AvailabilityZones
	o.Description = all.Description
	o.DisplayName = *all.DisplayName
	o.Id = all.Id
	o.ExtraInfo = all.ExtraInfo
	if all.DeletePolicy != nil && !all.DeletePolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.DeletePolicy = all.DeletePolicy
	}
	o.Overwrite = all.Overwrite
	if all.Dns != nil && all.Dns.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dns = all.Dns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
