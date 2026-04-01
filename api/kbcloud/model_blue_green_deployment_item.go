// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentItem blueGreenDeploymentItem is the response to a blue-green deployment creation request.
type BlueGreenDeploymentItem struct {
	// The ID of the blue-green deployment.
	DeploymentId string `json:"deploymentID"`
	// The name of the blue-green deployment.
	DeploymentName *string `json:"deploymentName,omitempty"`
	// KubeBlocks cluster information
	BlueCluster *ClusterListItem `json:"blueCluster,omitempty"`
	// KubeBlocks cluster information
	GreenCluster *ClusterListItem `json:"greenCluster,omitempty"`
	// The status of a blue-green deployment.
	DeploymentStatus *BlueGreenDeploymentStatus `json:"deploymentStatus,omitempty"`
	// The lag of the replication.
	ReplicationLag common.NullableString `json:"replicationLag,omitempty"`
	// The creation time of the blue-green deployment.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The update time of the blue-green deployment.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentItem instantiates a new BlueGreenDeploymentItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentItem(deploymentId string) *BlueGreenDeploymentItem {
	this := BlueGreenDeploymentItem{}
	this.DeploymentId = deploymentId
	return &this
}

// NewBlueGreenDeploymentItemWithDefaults instantiates a new BlueGreenDeploymentItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentItemWithDefaults() *BlueGreenDeploymentItem {
	this := BlueGreenDeploymentItem{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value.
func (o *BlueGreenDeploymentItem) GetDeploymentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentId, true
}

// SetDeploymentId sets field value.
func (o *BlueGreenDeploymentItem) SetDeploymentId(v string) {
	o.DeploymentId = v
}

// GetDeploymentName returns the DeploymentName field value if set, zero value otherwise.
func (o *BlueGreenDeploymentItem) GetDeploymentName() string {
	if o == nil || o.DeploymentName == nil {
		var ret string
		return ret
	}
	return *o.DeploymentName
}

// GetDeploymentNameOk returns a tuple with the DeploymentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetDeploymentNameOk() (*string, bool) {
	if o == nil || o.DeploymentName == nil {
		return nil, false
	}
	return o.DeploymentName, true
}

// HasDeploymentName returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasDeploymentName() bool {
	return o != nil && o.DeploymentName != nil
}

// SetDeploymentName gets a reference to the given string and assigns it to the DeploymentName field.
func (o *BlueGreenDeploymentItem) SetDeploymentName(v string) {
	o.DeploymentName = &v
}

// GetBlueCluster returns the BlueCluster field value if set, zero value otherwise.
func (o *BlueGreenDeploymentItem) GetBlueCluster() ClusterListItem {
	if o == nil || o.BlueCluster == nil {
		var ret ClusterListItem
		return ret
	}
	return *o.BlueCluster
}

// GetBlueClusterOk returns a tuple with the BlueCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetBlueClusterOk() (*ClusterListItem, bool) {
	if o == nil || o.BlueCluster == nil {
		return nil, false
	}
	return o.BlueCluster, true
}

// HasBlueCluster returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasBlueCluster() bool {
	return o != nil && o.BlueCluster != nil
}

// SetBlueCluster gets a reference to the given ClusterListItem and assigns it to the BlueCluster field.
func (o *BlueGreenDeploymentItem) SetBlueCluster(v ClusterListItem) {
	o.BlueCluster = &v
}

// GetGreenCluster returns the GreenCluster field value if set, zero value otherwise.
func (o *BlueGreenDeploymentItem) GetGreenCluster() ClusterListItem {
	if o == nil || o.GreenCluster == nil {
		var ret ClusterListItem
		return ret
	}
	return *o.GreenCluster
}

// GetGreenClusterOk returns a tuple with the GreenCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetGreenClusterOk() (*ClusterListItem, bool) {
	if o == nil || o.GreenCluster == nil {
		return nil, false
	}
	return o.GreenCluster, true
}

// HasGreenCluster returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasGreenCluster() bool {
	return o != nil && o.GreenCluster != nil
}

// SetGreenCluster gets a reference to the given ClusterListItem and assigns it to the GreenCluster field.
func (o *BlueGreenDeploymentItem) SetGreenCluster(v ClusterListItem) {
	o.GreenCluster = &v
}

// GetDeploymentStatus returns the DeploymentStatus field value if set, zero value otherwise.
func (o *BlueGreenDeploymentItem) GetDeploymentStatus() BlueGreenDeploymentStatus {
	if o == nil || o.DeploymentStatus == nil {
		var ret BlueGreenDeploymentStatus
		return ret
	}
	return *o.DeploymentStatus
}

// GetDeploymentStatusOk returns a tuple with the DeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetDeploymentStatusOk() (*BlueGreenDeploymentStatus, bool) {
	if o == nil || o.DeploymentStatus == nil {
		return nil, false
	}
	return o.DeploymentStatus, true
}

// HasDeploymentStatus returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasDeploymentStatus() bool {
	return o != nil && o.DeploymentStatus != nil
}

// SetDeploymentStatus gets a reference to the given BlueGreenDeploymentStatus and assigns it to the DeploymentStatus field.
func (o *BlueGreenDeploymentItem) SetDeploymentStatus(v BlueGreenDeploymentStatus) {
	o.DeploymentStatus = &v
}

// GetReplicationLag returns the ReplicationLag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueGreenDeploymentItem) GetReplicationLag() string {
	if o == nil || o.ReplicationLag.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReplicationLag.Get()
}

// GetReplicationLagOk returns a tuple with the ReplicationLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BlueGreenDeploymentItem) GetReplicationLagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationLag.Get(), o.ReplicationLag.IsSet()
}

// HasReplicationLag returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasReplicationLag() bool {
	return o != nil && o.ReplicationLag.IsSet()
}

// SetReplicationLag gets a reference to the given common.NullableString and assigns it to the ReplicationLag field.
func (o *BlueGreenDeploymentItem) SetReplicationLag(v string) {
	o.ReplicationLag.Set(&v)
}

// SetReplicationLagNil sets the value for ReplicationLag to be an explicit nil.
func (o *BlueGreenDeploymentItem) SetReplicationLagNil() {
	o.ReplicationLag.Set(nil)
}

// UnsetReplicationLag ensures that no value is present for ReplicationLag, not even an explicit nil.
func (o *BlueGreenDeploymentItem) UnsetReplicationLag() {
	o.ReplicationLag.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BlueGreenDeploymentItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BlueGreenDeploymentItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BlueGreenDeploymentItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BlueGreenDeploymentItem) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BlueGreenDeploymentItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["deploymentID"] = o.DeploymentId
	if o.DeploymentName != nil {
		toSerialize["deploymentName"] = o.DeploymentName
	}
	if o.BlueCluster != nil {
		toSerialize["blueCluster"] = o.BlueCluster
	}
	if o.GreenCluster != nil {
		toSerialize["greenCluster"] = o.GreenCluster
	}
	if o.DeploymentStatus != nil {
		toSerialize["deploymentStatus"] = o.DeploymentStatus
	}
	if o.ReplicationLag.IsSet() {
		toSerialize["replicationLag"] = o.ReplicationLag.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeploymentId     *string                    `json:"deploymentID"`
		DeploymentName   *string                    `json:"deploymentName,omitempty"`
		BlueCluster      *ClusterListItem           `json:"blueCluster,omitempty"`
		GreenCluster     *ClusterListItem           `json:"greenCluster,omitempty"`
		DeploymentStatus *BlueGreenDeploymentStatus `json:"deploymentStatus,omitempty"`
		ReplicationLag   common.NullableString      `json:"replicationLag,omitempty"`
		CreatedAt        *time.Time                 `json:"createdAt,omitempty"`
		UpdatedAt        *time.Time                 `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DeploymentId == nil {
		return fmt.Errorf("required field deploymentID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"deploymentID", "deploymentName", "blueCluster", "greenCluster", "deploymentStatus", "replicationLag", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DeploymentId = *all.DeploymentId
	o.DeploymentName = all.DeploymentName
	if all.BlueCluster != nil && all.BlueCluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BlueCluster = all.BlueCluster
	if all.GreenCluster != nil && all.GreenCluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GreenCluster = all.GreenCluster
	if all.DeploymentStatus != nil && !all.DeploymentStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DeploymentStatus = all.DeploymentStatus
	}
	o.ReplicationLag = all.ReplicationLag
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
