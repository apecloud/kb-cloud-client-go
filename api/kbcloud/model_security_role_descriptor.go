// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SecurityRoleDescriptor struct {
	Description       *string                           `json:"description,omitempty"`
	Cluster           []string                          `json:"cluster,omitempty"`
	Global            map[string]interface{}            `json:"global,omitempty"`
	Indices           []SecurityIndexPrivileges         `json:"indices,omitempty"`
	RemoteIndices     []SecurityRemoteIndexPrivileges   `json:"remote_indices,omitempty"`
	RemoteCluster     []SecurityRemoteClusterPrivileges `json:"remote_cluster,omitempty"`
	Applications      []SecurityApplicationPrivileges   `json:"applications,omitempty"`
	RunAs             []string                          `json:"run_as,omitempty"`
	Metadata          map[string]interface{}            `json:"metadata,omitempty"`
	Restriction       map[string]interface{}            `json:"restriction,omitempty"`
	TransientMetadata map[string]interface{}            `json:"transient_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityRoleDescriptor instantiates a new SecurityRoleDescriptor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityRoleDescriptor() *SecurityRoleDescriptor {
	this := SecurityRoleDescriptor{}
	return &this
}

// NewSecurityRoleDescriptorWithDefaults instantiates a new SecurityRoleDescriptor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityRoleDescriptorWithDefaults() *SecurityRoleDescriptor {
	this := SecurityRoleDescriptor{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityRoleDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetCluster() []string {
	if o == nil || o.Cluster == nil {
		var ret []string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetClusterOk() (*[]string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given []string and assigns it to the Cluster field.
func (o *SecurityRoleDescriptor) SetCluster(v []string) {
	o.Cluster = v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetGlobal() map[string]interface{} {
	if o == nil || o.Global == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetGlobalOk() (*map[string]interface{}, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return &o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasGlobal() bool {
	return o != nil && o.Global != nil
}

// SetGlobal gets a reference to the given map[string]interface{} and assigns it to the Global field.
func (o *SecurityRoleDescriptor) SetGlobal(v map[string]interface{}) {
	o.Global = v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetIndices() []SecurityIndexPrivileges {
	if o == nil || o.Indices == nil {
		var ret []SecurityIndexPrivileges
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetIndicesOk() (*[]SecurityIndexPrivileges, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return &o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasIndices() bool {
	return o != nil && o.Indices != nil
}

// SetIndices gets a reference to the given []SecurityIndexPrivileges and assigns it to the Indices field.
func (o *SecurityRoleDescriptor) SetIndices(v []SecurityIndexPrivileges) {
	o.Indices = v
}

// GetRemoteIndices returns the RemoteIndices field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetRemoteIndices() []SecurityRemoteIndexPrivileges {
	if o == nil || o.RemoteIndices == nil {
		var ret []SecurityRemoteIndexPrivileges
		return ret
	}
	return o.RemoteIndices
}

// GetRemoteIndicesOk returns a tuple with the RemoteIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetRemoteIndicesOk() (*[]SecurityRemoteIndexPrivileges, bool) {
	if o == nil || o.RemoteIndices == nil {
		return nil, false
	}
	return &o.RemoteIndices, true
}

// HasRemoteIndices returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasRemoteIndices() bool {
	return o != nil && o.RemoteIndices != nil
}

// SetRemoteIndices gets a reference to the given []SecurityRemoteIndexPrivileges and assigns it to the RemoteIndices field.
func (o *SecurityRoleDescriptor) SetRemoteIndices(v []SecurityRemoteIndexPrivileges) {
	o.RemoteIndices = v
}

// GetRemoteCluster returns the RemoteCluster field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetRemoteCluster() []SecurityRemoteClusterPrivileges {
	if o == nil || o.RemoteCluster == nil {
		var ret []SecurityRemoteClusterPrivileges
		return ret
	}
	return o.RemoteCluster
}

// GetRemoteClusterOk returns a tuple with the RemoteCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetRemoteClusterOk() (*[]SecurityRemoteClusterPrivileges, bool) {
	if o == nil || o.RemoteCluster == nil {
		return nil, false
	}
	return &o.RemoteCluster, true
}

// HasRemoteCluster returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasRemoteCluster() bool {
	return o != nil && o.RemoteCluster != nil
}

// SetRemoteCluster gets a reference to the given []SecurityRemoteClusterPrivileges and assigns it to the RemoteCluster field.
func (o *SecurityRoleDescriptor) SetRemoteCluster(v []SecurityRemoteClusterPrivileges) {
	o.RemoteCluster = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetApplications() []SecurityApplicationPrivileges {
	if o == nil || o.Applications == nil {
		var ret []SecurityApplicationPrivileges
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetApplicationsOk() (*[]SecurityApplicationPrivileges, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return &o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasApplications() bool {
	return o != nil && o.Applications != nil
}

// SetApplications gets a reference to the given []SecurityApplicationPrivileges and assigns it to the Applications field.
func (o *SecurityRoleDescriptor) SetApplications(v []SecurityApplicationPrivileges) {
	o.Applications = v
}

// GetRunAs returns the RunAs field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetRunAs() []string {
	if o == nil || o.RunAs == nil {
		var ret []string
		return ret
	}
	return o.RunAs
}

// GetRunAsOk returns a tuple with the RunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetRunAsOk() (*[]string, bool) {
	if o == nil || o.RunAs == nil {
		return nil, false
	}
	return &o.RunAs, true
}

// HasRunAs returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasRunAs() bool {
	return o != nil && o.RunAs != nil
}

// SetRunAs gets a reference to the given []string and assigns it to the RunAs field.
func (o *SecurityRoleDescriptor) SetRunAs(v []string) {
	o.RunAs = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SecurityRoleDescriptor) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRestriction returns the Restriction field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetRestriction() map[string]interface{} {
	if o == nil || o.Restriction == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Restriction
}

// GetRestrictionOk returns a tuple with the Restriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetRestrictionOk() (*map[string]interface{}, bool) {
	if o == nil || o.Restriction == nil {
		return nil, false
	}
	return &o.Restriction, true
}

// HasRestriction returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasRestriction() bool {
	return o != nil && o.Restriction != nil
}

// SetRestriction gets a reference to the given map[string]interface{} and assigns it to the Restriction field.
func (o *SecurityRoleDescriptor) SetRestriction(v map[string]interface{}) {
	o.Restriction = v
}

// GetTransientMetadata returns the TransientMetadata field value if set, zero value otherwise.
func (o *SecurityRoleDescriptor) GetTransientMetadata() map[string]interface{} {
	if o == nil || o.TransientMetadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientMetadata
}

// GetTransientMetadataOk returns a tuple with the TransientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleDescriptor) GetTransientMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.TransientMetadata == nil {
		return nil, false
	}
	return &o.TransientMetadata, true
}

// HasTransientMetadata returns a boolean if a field has been set.
func (o *SecurityRoleDescriptor) HasTransientMetadata() bool {
	return o != nil && o.TransientMetadata != nil
}

// SetTransientMetadata gets a reference to the given map[string]interface{} and assigns it to the TransientMetadata field.
func (o *SecurityRoleDescriptor) SetTransientMetadata(v map[string]interface{}) {
	o.TransientMetadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityRoleDescriptor) MarshalJSON() ([]byte, error) {
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
func (o *SecurityRoleDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       *string                           `json:"description,omitempty"`
		Cluster           []string                          `json:"cluster,omitempty"`
		Global            map[string]interface{}            `json:"global,omitempty"`
		Indices           []SecurityIndexPrivileges         `json:"indices,omitempty"`
		RemoteIndices     []SecurityRemoteIndexPrivileges   `json:"remote_indices,omitempty"`
		RemoteCluster     []SecurityRemoteClusterPrivileges `json:"remote_cluster,omitempty"`
		Applications      []SecurityApplicationPrivileges   `json:"applications,omitempty"`
		RunAs             []string                          `json:"run_as,omitempty"`
		Metadata          map[string]interface{}            `json:"metadata,omitempty"`
		Restriction       map[string]interface{}            `json:"restriction,omitempty"`
		TransientMetadata map[string]interface{}            `json:"transient_metadata,omitempty"`
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
