// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterStorageUsageHistoryInstance struct {
	// Kubernetes pod name for this storage usage series.
	InstanceName *string `json:"instanceName,omitempty"`
	// Kubernetes pod name for compatibility with Kubernetes-level troubleshooting.
	PodName *string `json:"podName,omitempty"`
	// PersistentVolumeClaim name when it can be read from metrics labels.
	PvcName *string `json:"pvcName,omitempty"`
	// Database replica role normalized from metrics labels.
	Role *ClusterStorageUsageHistoryInstanceRole `json:"role,omitempty"`
	// Optional KubeBlocks component name when it can be read from metrics labels.
	ComponentName *string                           `json:"componentName,omitempty"`
	Points        []ClusterStorageUsageHistoryPoint `json:"points"`
	Source        *string                           `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterStorageUsageHistoryInstance instantiates a new ClusterStorageUsageHistoryInstance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterStorageUsageHistoryInstance(points []ClusterStorageUsageHistoryPoint) *ClusterStorageUsageHistoryInstance {
	this := ClusterStorageUsageHistoryInstance{}
	this.Points = points
	return &this
}

// NewClusterStorageUsageHistoryInstanceWithDefaults instantiates a new ClusterStorageUsageHistoryInstance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterStorageUsageHistoryInstanceWithDefaults() *ClusterStorageUsageHistoryInstance {
	this := ClusterStorageUsageHistoryInstance{}
	return &this
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *ClusterStorageUsageHistoryInstance) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *ClusterStorageUsageHistoryInstance) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *ClusterStorageUsageHistoryInstance) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *ClusterStorageUsageHistoryInstance) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *ClusterStorageUsageHistoryInstance) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *ClusterStorageUsageHistoryInstance) SetPodName(v string) {
	o.PodName = &v
}

// GetPvcName returns the PvcName field value if set, zero value otherwise.
func (o *ClusterStorageUsageHistoryInstance) GetPvcName() string {
	if o == nil || o.PvcName == nil {
		var ret string
		return ret
	}
	return *o.PvcName
}

// GetPvcNameOk returns a tuple with the PvcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetPvcNameOk() (*string, bool) {
	if o == nil || o.PvcName == nil {
		return nil, false
	}
	return o.PvcName, true
}

// HasPvcName returns a boolean if a field has been set.
func (o *ClusterStorageUsageHistoryInstance) HasPvcName() bool {
	return o != nil && o.PvcName != nil
}

// SetPvcName gets a reference to the given string and assigns it to the PvcName field.
func (o *ClusterStorageUsageHistoryInstance) SetPvcName(v string) {
	o.PvcName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ClusterStorageUsageHistoryInstance) GetRole() ClusterStorageUsageHistoryInstanceRole {
	if o == nil || o.Role == nil {
		var ret ClusterStorageUsageHistoryInstanceRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetRoleOk() (*ClusterStorageUsageHistoryInstanceRole, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ClusterStorageUsageHistoryInstance) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given ClusterStorageUsageHistoryInstanceRole and assigns it to the Role field.
func (o *ClusterStorageUsageHistoryInstance) SetRole(v ClusterStorageUsageHistoryInstanceRole) {
	o.Role = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *ClusterStorageUsageHistoryInstance) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *ClusterStorageUsageHistoryInstance) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *ClusterStorageUsageHistoryInstance) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetPoints returns the Points field value.
func (o *ClusterStorageUsageHistoryInstance) GetPoints() []ClusterStorageUsageHistoryPoint {
	if o == nil {
		var ret []ClusterStorageUsageHistoryPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetPointsOk() (*[]ClusterStorageUsageHistoryPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *ClusterStorageUsageHistoryInstance) SetPoints(v []ClusterStorageUsageHistoryPoint) {
	o.Points = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ClusterStorageUsageHistoryInstance) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryInstance) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ClusterStorageUsageHistoryInstance) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ClusterStorageUsageHistoryInstance) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterStorageUsageHistoryInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.PodName != nil {
		toSerialize["podName"] = o.PodName
	}
	if o.PvcName != nil {
		toSerialize["pvcName"] = o.PvcName
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	toSerialize["points"] = o.Points
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterStorageUsageHistoryInstance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName  *string                                 `json:"instanceName,omitempty"`
		PodName       *string                                 `json:"podName,omitempty"`
		PvcName       *string                                 `json:"pvcName,omitempty"`
		Role          *ClusterStorageUsageHistoryInstanceRole `json:"role,omitempty"`
		ComponentName *string                                 `json:"componentName,omitempty"`
		Points        *[]ClusterStorageUsageHistoryPoint      `json:"points"`
		Source        *string                                 `json:"source,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceName", "podName", "pvcName", "role", "componentName", "points", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.InstanceName = all.InstanceName
	o.PodName = all.PodName
	o.PvcName = all.PvcName
	if all.Role != nil && !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = all.Role
	}
	o.ComponentName = all.ComponentName
	o.Points = *all.Points
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
