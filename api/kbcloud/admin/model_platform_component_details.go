// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformComponentDetails Platform component detail including pods and observability metadata
type PlatformComponentDetails struct {
	// Component name
	Name        string                `json:"name"`
	Description *LocalizedDescription `json:"description,omitempty"`
	// Platform component status
	Status   PlatformComponentStatus `json:"status"`
	Version  *string                 `json:"version,omitempty"`
	Replicas *int64                  `json:"replicas,omitempty"`
	// Grafana dashboard UID for this component
	DashboardUid *string `json:"dashboardUID,omitempty"`
	// Log query job name for this component
	LogJobName *string `json:"logJobName,omitempty"`
	// List of pods for this component
	Pods []PlatformComponentPod `json:"pods,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformComponentDetails instantiates a new PlatformComponentDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformComponentDetails(name string, status PlatformComponentStatus) *PlatformComponentDetails {
	this := PlatformComponentDetails{}
	this.Name = name
	this.Status = status
	return &this
}

// NewPlatformComponentDetailsWithDefaults instantiates a new PlatformComponentDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformComponentDetailsWithDefaults() *PlatformComponentDetails {
	this := PlatformComponentDetails{}
	return &this
}

// GetName returns the Name field value.
func (o *PlatformComponentDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PlatformComponentDetails) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlatformComponentDetails) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlatformComponentDetails) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *PlatformComponentDetails) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetStatus returns the Status field value.
func (o *PlatformComponentDetails) GetStatus() PlatformComponentStatus {
	if o == nil {
		var ret PlatformComponentStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetStatusOk() (*PlatformComponentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PlatformComponentDetails) SetStatus(v PlatformComponentStatus) {
	o.Status = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PlatformComponentDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PlatformComponentDetails) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PlatformComponentDetails) SetVersion(v string) {
	o.Version = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *PlatformComponentDetails) GetReplicas() int64 {
	if o == nil || o.Replicas == nil {
		var ret int64
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetReplicasOk() (*int64, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *PlatformComponentDetails) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *PlatformComponentDetails) SetReplicas(v int64) {
	o.Replicas = &v
}

// GetDashboardUid returns the DashboardUid field value if set, zero value otherwise.
func (o *PlatformComponentDetails) GetDashboardUid() string {
	if o == nil || o.DashboardUid == nil {
		var ret string
		return ret
	}
	return *o.DashboardUid
}

// GetDashboardUidOk returns a tuple with the DashboardUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetDashboardUidOk() (*string, bool) {
	if o == nil || o.DashboardUid == nil {
		return nil, false
	}
	return o.DashboardUid, true
}

// HasDashboardUid returns a boolean if a field has been set.
func (o *PlatformComponentDetails) HasDashboardUid() bool {
	return o != nil && o.DashboardUid != nil
}

// SetDashboardUid gets a reference to the given string and assigns it to the DashboardUid field.
func (o *PlatformComponentDetails) SetDashboardUid(v string) {
	o.DashboardUid = &v
}

// GetLogJobName returns the LogJobName field value if set, zero value otherwise.
func (o *PlatformComponentDetails) GetLogJobName() string {
	if o == nil || o.LogJobName == nil {
		var ret string
		return ret
	}
	return *o.LogJobName
}

// GetLogJobNameOk returns a tuple with the LogJobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetLogJobNameOk() (*string, bool) {
	if o == nil || o.LogJobName == nil {
		return nil, false
	}
	return o.LogJobName, true
}

// HasLogJobName returns a boolean if a field has been set.
func (o *PlatformComponentDetails) HasLogJobName() bool {
	return o != nil && o.LogJobName != nil
}

// SetLogJobName gets a reference to the given string and assigns it to the LogJobName field.
func (o *PlatformComponentDetails) SetLogJobName(v string) {
	o.LogJobName = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *PlatformComponentDetails) GetPods() []PlatformComponentPod {
	if o == nil || o.Pods == nil {
		var ret []PlatformComponentPod
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentDetails) GetPodsOk() (*[]PlatformComponentPod, bool) {
	if o == nil || o.Pods == nil {
		return nil, false
	}
	return &o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *PlatformComponentDetails) HasPods() bool {
	return o != nil && o.Pods != nil
}

// SetPods gets a reference to the given []PlatformComponentPod and assigns it to the Pods field.
func (o *PlatformComponentDetails) SetPods(v []PlatformComponentPod) {
	o.Pods = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformComponentDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.DashboardUid != nil {
		toSerialize["dashboardUID"] = o.DashboardUid
	}
	if o.LogJobName != nil {
		toSerialize["logJobName"] = o.LogJobName
	}
	if o.Pods != nil {
		toSerialize["pods"] = o.Pods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformComponentDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                  `json:"name"`
		Description  *LocalizedDescription    `json:"description,omitempty"`
		Status       *PlatformComponentStatus `json:"status"`
		Version      *string                  `json:"version,omitempty"`
		Replicas     *int64                   `json:"replicas,omitempty"`
		DashboardUid *string                  `json:"dashboardUID,omitempty"`
		LogJobName   *string                  `json:"logJobName,omitempty"`
		Pods         []PlatformComponentPod   `json:"pods,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "status", "version", "replicas", "dashboardUID", "logJobName", "pods"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Version = all.Version
	o.Replicas = all.Replicas
	o.DashboardUid = all.DashboardUid
	o.LogJobName = all.LogJobName
	o.Pods = all.Pods

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
