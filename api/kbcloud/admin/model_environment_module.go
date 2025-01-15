// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

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
	// Environment module status (running,updating, error, disabled and status for cluster)
	Status string `json:"status"`
	// Hosting status (Hostable, Non-hostable, Hosted). When hosting_status is Hosted, cluster_info will be returned
	HostingStatus *HostingStatus `json:"hostingStatus,omitempty"`
	// Number of replicas
	Replicas *int32 `json:"replicas,omitempty"`
	// Deployment location
	Location *string `json:"location,omitempty"`
	// Cluster information
	ClusterInfo *ClusterInfo `json:"clusterInfo,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModule instantiates a new EnvironmentModule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModule(name string, status string) *EnvironmentModule {
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
func (o *EnvironmentModule) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *EnvironmentModule) SetStatus(v string) {
	o.Status = v
}

// GetHostingStatus returns the HostingStatus field value if set, zero value otherwise.
func (o *EnvironmentModule) GetHostingStatus() HostingStatus {
	if o == nil || o.HostingStatus == nil {
		var ret HostingStatus
		return ret
	}
	return *o.HostingStatus
}

// GetHostingStatusOk returns a tuple with the HostingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModule) GetHostingStatusOk() (*HostingStatus, bool) {
	if o == nil || o.HostingStatus == nil {
		return nil, false
	}
	return o.HostingStatus, true
}

// HasHostingStatus returns a boolean if a field has been set.
func (o *EnvironmentModule) HasHostingStatus() bool {
	return o != nil && o.HostingStatus != nil
}

// SetHostingStatus gets a reference to the given HostingStatus and assigns it to the HostingStatus field.
func (o *EnvironmentModule) SetHostingStatus(v HostingStatus) {
	o.HostingStatus = &v
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
	if o.HostingStatus != nil {
		toSerialize["hostingStatus"] = o.HostingStatus
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.ClusterInfo != nil {
		toSerialize["clusterInfo"] = o.ClusterInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string        `json:"name"`
		Version       *string        `json:"version,omitempty"`
		Status        *string        `json:"status"`
		HostingStatus *HostingStatus `json:"hostingStatus,omitempty"`
		Replicas      *int32         `json:"replicas,omitempty"`
		Location      *string        `json:"location,omitempty"`
		ClusterInfo   *ClusterInfo   `json:"clusterInfo,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "version", "status", "hostingStatus", "replicas", "location", "clusterInfo"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Version = all.Version
	o.Status = *all.Status
	if all.HostingStatus != nil && !all.HostingStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.HostingStatus = all.HostingStatus
	}
	o.Replicas = all.Replicas
	o.Location = all.Location
	if all.ClusterInfo != nil && all.ClusterInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClusterInfo = all.ClusterInfo

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
