// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EnvironmentModule Single environment module information
type EnvironmentModule struct {
	// Environment module name
	Name *string `json:"name,omitempty"`
	// Environment module version
	Version *string `json:"version,omitempty"`
	// Environment module status (running, stopped, error, not_installed)
	Status *string `json:"status,omitempty"`
	// Management type (unmanageable, manageable, kb-cluster). When management_type is kb-cluster, cluster_name, cluster_definition and org_name will be returned
	ManagementType *string `json:"management_type,omitempty"`
	// Number of replicas
	Replicas *int32 `json:"replicas,omitempty"`
	// Deployment location
	Location *string `json:"location,omitempty"`
	// The name of the cluster, only returned when management_type is kb-cluster
	ClusterName *string `json:"cluster_name,omitempty"`
	// The definition of the cluster, only returned when management_type is kb-cluster
	ClusterDefinition *string `json:"cluster_definition,omitempty"`
	// The organization name, only returned when management_type is kb-cluster
	OrgName *string `json:"org_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModule instantiates a new EnvironmentModule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModule() *EnvironmentModule {
	this := EnvironmentModule{}
	return &this
}

// NewEnvironmentModuleWithDefaults instantiates a new EnvironmentModule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleWithDefaults() *EnvironmentModule {
	this := EnvironmentModule{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentModule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentModule) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentModule) SetName(v string) {
	o.Name = &v
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

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvironmentModule) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvironmentModule) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnvironmentModule) SetStatus(v string) {
	o.Status = &v
}

// GetManagementType returns the ManagementType field value if set, zero value otherwise.
func (o *EnvironmentModule) GetManagementType() string {
	if o == nil || o.ManagementType == nil {
		var ret string
		return ret
	}
	return *o.ManagementType
}

// GetManagementTypeOk returns a tuple with the ManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetManagementTypeOk() (*string, bool) {
	if o == nil || o.ManagementType == nil {
		return nil, false
	}
	return o.ManagementType, true
}

// HasManagementType returns a boolean if a field has been set.
func (o *EnvironmentModule) HasManagementType() bool {
	return o != nil && o.ManagementType != nil
}

// SetManagementType gets a reference to the given string and assigns it to the ManagementType field.
func (o *EnvironmentModule) SetManagementType(v string) {
	o.ManagementType = &v
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

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *EnvironmentModule) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *EnvironmentModule) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *EnvironmentModule) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterDefinition returns the ClusterDefinition field value if set, zero value otherwise.
func (o *EnvironmentModule) GetClusterDefinition() string {
	if o == nil || o.ClusterDefinition == nil {
		var ret string
		return ret
	}
	return *o.ClusterDefinition
}

// GetClusterDefinitionOk returns a tuple with the ClusterDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetClusterDefinitionOk() (*string, bool) {
	if o == nil || o.ClusterDefinition == nil {
		return nil, false
	}
	return o.ClusterDefinition, true
}

// HasClusterDefinition returns a boolean if a field has been set.
func (o *EnvironmentModule) HasClusterDefinition() bool {
	return o != nil && o.ClusterDefinition != nil
}

// SetClusterDefinition gets a reference to the given string and assigns it to the ClusterDefinition field.
func (o *EnvironmentModule) SetClusterDefinition(v string) {
	o.ClusterDefinition = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *EnvironmentModule) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *EnvironmentModule) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *EnvironmentModule) SetOrgName(v string) {
	o.OrgName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ManagementType != nil {
		toSerialize["management_type"] = o.ManagementType
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.ClusterDefinition != nil {
		toSerialize["cluster_definition"] = o.ClusterDefinition
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string `json:"name,omitempty"`
		Version           *string `json:"version,omitempty"`
		Status            *string `json:"status,omitempty"`
		ManagementType    *string `json:"management_type,omitempty"`
		Replicas          *int32  `json:"replicas,omitempty"`
		Location          *string `json:"location,omitempty"`
		ClusterName       *string `json:"cluster_name,omitempty"`
		ClusterDefinition *string `json:"cluster_definition,omitempty"`
		OrgName           *string `json:"org_name,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "version", "status", "management_type", "replicas", "location", "cluster_name", "cluster_definition", "org_name"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Version = all.Version
	o.Status = all.Status
	o.ManagementType = all.ManagementType
	o.Replicas = all.Replicas
	o.Location = all.Location
	o.ClusterName = all.ClusterName
	o.ClusterDefinition = all.ClusterDefinition
	o.OrgName = all.OrgName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
