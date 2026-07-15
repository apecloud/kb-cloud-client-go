// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ESSecurityRoleDescriptor struct {
	Description       *string                             `json:"description,omitempty"`
	Cluster           []string                            `json:"cluster,omitempty"`
	Global            map[string]interface{}              `json:"global,omitempty"`
	Indices           []ESSecurityIndexPrivileges         `json:"indices,omitempty"`
	RemoteIndices     []ESSecurityRemoteIndexPrivileges   `json:"remote_indices,omitempty"`
	RemoteCluster     []ESSecurityRemoteClusterPrivileges `json:"remote_cluster,omitempty"`
	Applications      []ESSecurityApplicationPrivileges   `json:"applications,omitempty"`
	RunAs             []string                            `json:"run_as,omitempty"`
	Metadata          map[string]interface{}              `json:"metadata,omitempty"`
	Restriction       map[string]interface{}              `json:"restriction,omitempty"`
	TransientMetadata map[string]interface{}              `json:"transient_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityRoleDescriptor instantiates a new ESSecurityRoleDescriptor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityRoleDescriptor() *ESSecurityRoleDescriptor {
	this := ESSecurityRoleDescriptor{}
	return &this
}

// NewESSecurityRoleDescriptorWithDefaults instantiates a new ESSecurityRoleDescriptor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityRoleDescriptorWithDefaults() *ESSecurityRoleDescriptor {
	this := ESSecurityRoleDescriptor{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ESSecurityRoleDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetCluster() []string {
	if o == nil || o.Cluster == nil {
		var ret []string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetClusterOk() (*[]string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given []string and assigns it to the Cluster field.
func (o *ESSecurityRoleDescriptor) SetCluster(v []string) {
	o.Cluster = v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetGlobal() map[string]interface{} {
	if o == nil || o.Global == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetGlobalOk() (*map[string]interface{}, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return &o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasGlobal() bool {
	return o != nil && o.Global != nil
}

// SetGlobal gets a reference to the given map[string]interface{} and assigns it to the Global field.
func (o *ESSecurityRoleDescriptor) SetGlobal(v map[string]interface{}) {
	o.Global = v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetIndices() []ESSecurityIndexPrivileges {
	if o == nil || o.Indices == nil {
		var ret []ESSecurityIndexPrivileges
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetIndicesOk() (*[]ESSecurityIndexPrivileges, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return &o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasIndices() bool {
	return o != nil && o.Indices != nil
}

// SetIndices gets a reference to the given []ESSecurityIndexPrivileges and assigns it to the Indices field.
func (o *ESSecurityRoleDescriptor) SetIndices(v []ESSecurityIndexPrivileges) {
	o.Indices = v
}

// GetRemoteIndices returns the RemoteIndices field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetRemoteIndices() []ESSecurityRemoteIndexPrivileges {
	if o == nil || o.RemoteIndices == nil {
		var ret []ESSecurityRemoteIndexPrivileges
		return ret
	}
	return o.RemoteIndices
}

// GetRemoteIndicesOk returns a tuple with the RemoteIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetRemoteIndicesOk() (*[]ESSecurityRemoteIndexPrivileges, bool) {
	if o == nil || o.RemoteIndices == nil {
		return nil, false
	}
	return &o.RemoteIndices, true
}

// HasRemoteIndices returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasRemoteIndices() bool {
	return o != nil && o.RemoteIndices != nil
}

// SetRemoteIndices gets a reference to the given []ESSecurityRemoteIndexPrivileges and assigns it to the RemoteIndices field.
func (o *ESSecurityRoleDescriptor) SetRemoteIndices(v []ESSecurityRemoteIndexPrivileges) {
	o.RemoteIndices = v
}

// GetRemoteCluster returns the RemoteCluster field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetRemoteCluster() []ESSecurityRemoteClusterPrivileges {
	if o == nil || o.RemoteCluster == nil {
		var ret []ESSecurityRemoteClusterPrivileges
		return ret
	}
	return o.RemoteCluster
}

// GetRemoteClusterOk returns a tuple with the RemoteCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetRemoteClusterOk() (*[]ESSecurityRemoteClusterPrivileges, bool) {
	if o == nil || o.RemoteCluster == nil {
		return nil, false
	}
	return &o.RemoteCluster, true
}

// HasRemoteCluster returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasRemoteCluster() bool {
	return o != nil && o.RemoteCluster != nil
}

// SetRemoteCluster gets a reference to the given []ESSecurityRemoteClusterPrivileges and assigns it to the RemoteCluster field.
func (o *ESSecurityRoleDescriptor) SetRemoteCluster(v []ESSecurityRemoteClusterPrivileges) {
	o.RemoteCluster = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetApplications() []ESSecurityApplicationPrivileges {
	if o == nil || o.Applications == nil {
		var ret []ESSecurityApplicationPrivileges
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetApplicationsOk() (*[]ESSecurityApplicationPrivileges, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return &o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasApplications() bool {
	return o != nil && o.Applications != nil
}

// SetApplications gets a reference to the given []ESSecurityApplicationPrivileges and assigns it to the Applications field.
func (o *ESSecurityRoleDescriptor) SetApplications(v []ESSecurityApplicationPrivileges) {
	o.Applications = v
}

// GetRunAs returns the RunAs field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetRunAs() []string {
	if o == nil || o.RunAs == nil {
		var ret []string
		return ret
	}
	return o.RunAs
}

// GetRunAsOk returns a tuple with the RunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetRunAsOk() (*[]string, bool) {
	if o == nil || o.RunAs == nil {
		return nil, false
	}
	return &o.RunAs, true
}

// HasRunAs returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasRunAs() bool {
	return o != nil && o.RunAs != nil
}

// SetRunAs gets a reference to the given []string and assigns it to the RunAs field.
func (o *ESSecurityRoleDescriptor) SetRunAs(v []string) {
	o.RunAs = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ESSecurityRoleDescriptor) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRestriction returns the Restriction field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetRestriction() map[string]interface{} {
	if o == nil || o.Restriction == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Restriction
}

// GetRestrictionOk returns a tuple with the Restriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetRestrictionOk() (*map[string]interface{}, bool) {
	if o == nil || o.Restriction == nil {
		return nil, false
	}
	return &o.Restriction, true
}

// HasRestriction returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasRestriction() bool {
	return o != nil && o.Restriction != nil
}

// SetRestriction gets a reference to the given map[string]interface{} and assigns it to the Restriction field.
func (o *ESSecurityRoleDescriptor) SetRestriction(v map[string]interface{}) {
	o.Restriction = v
}

// GetTransientMetadata returns the TransientMetadata field value if set, zero value otherwise.
func (o *ESSecurityRoleDescriptor) GetTransientMetadata() map[string]interface{} {
	if o == nil || o.TransientMetadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientMetadata
}

// GetTransientMetadataOk returns a tuple with the TransientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleDescriptor) GetTransientMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.TransientMetadata == nil {
		return nil, false
	}
	return &o.TransientMetadata, true
}

// HasTransientMetadata returns a boolean if a field has been set.
func (o *ESSecurityRoleDescriptor) HasTransientMetadata() bool {
	return o != nil && o.TransientMetadata != nil
}

// SetTransientMetadata gets a reference to the given map[string]interface{} and assigns it to the TransientMetadata field.
func (o *ESSecurityRoleDescriptor) SetTransientMetadata(v map[string]interface{}) {
	o.TransientMetadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityRoleDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Global != nil {
		toSerialize["global"] = o.Global
	}
	if o.Indices != nil {
		toSerialize["indices"] = o.Indices
	}
	if o.RemoteIndices != nil {
		toSerialize["remote_indices"] = o.RemoteIndices
	}
	if o.RemoteCluster != nil {
		toSerialize["remote_cluster"] = o.RemoteCluster
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.RunAs != nil {
		toSerialize["run_as"] = o.RunAs
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Restriction != nil {
		toSerialize["restriction"] = o.Restriction
	}
	if o.TransientMetadata != nil {
		toSerialize["transient_metadata"] = o.TransientMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ESSecurityRoleDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       *string                             `json:"description,omitempty"`
		Cluster           []string                            `json:"cluster,omitempty"`
		Global            map[string]interface{}              `json:"global,omitempty"`
		Indices           []ESSecurityIndexPrivileges         `json:"indices,omitempty"`
		RemoteIndices     []ESSecurityRemoteIndexPrivileges   `json:"remote_indices,omitempty"`
		RemoteCluster     []ESSecurityRemoteClusterPrivileges `json:"remote_cluster,omitempty"`
		Applications      []ESSecurityApplicationPrivileges   `json:"applications,omitempty"`
		RunAs             []string                            `json:"run_as,omitempty"`
		Metadata          map[string]interface{}              `json:"metadata,omitempty"`
		Restriction       map[string]interface{}              `json:"restriction,omitempty"`
		TransientMetadata map[string]interface{}              `json:"transient_metadata,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "cluster", "global", "indices", "remote_indices", "remote_cluster", "applications", "run_as", "metadata", "restriction", "transient_metadata"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Cluster = all.Cluster
	o.Global = all.Global
	o.Indices = all.Indices
	o.RemoteIndices = all.RemoteIndices
	o.RemoteCluster = all.RemoteCluster
	o.Applications = all.Applications
	o.RunAs = all.RunAs
	o.Metadata = all.Metadata
	o.Restriction = all.Restriction
	o.TransientMetadata = all.TransientMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
