// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterJarReplica struct {
	Component  string     `json:"component"`
	PodName    string     `json:"podName"`
	PodUid     string     `json:"podUID"`
	Status     string     `json:"status"`
	Error      *string    `json:"error,omitempty"`
	Generation int64      `json:"generation"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterJarReplica instantiates a new ClusterJarReplica object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterJarReplica(component string, podName string, podUid string, status string, generation int64) *ClusterJarReplica {
	this := ClusterJarReplica{}
	this.Component = component
	this.PodName = podName
	this.PodUid = podUid
	this.Status = status
	this.Generation = generation
	return &this
}

// NewClusterJarReplicaWithDefaults instantiates a new ClusterJarReplica object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterJarReplicaWithDefaults() *ClusterJarReplica {
	this := ClusterJarReplica{}
	return &this
}

// GetComponent returns the Component field value.
func (o *ClusterJarReplica) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ClusterJarReplica) SetComponent(v string) {
	o.Component = v
}

// GetPodName returns the PodName field value.
func (o *ClusterJarReplica) GetPodName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetPodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodName, true
}

// SetPodName sets field value.
func (o *ClusterJarReplica) SetPodName(v string) {
	o.PodName = v
}

// GetPodUid returns the PodUid field value.
func (o *ClusterJarReplica) GetPodUid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PodUid
}

// GetPodUidOk returns a tuple with the PodUid field value
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetPodUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodUid, true
}

// SetPodUid sets field value.
func (o *ClusterJarReplica) SetPodUid(v string) {
	o.PodUid = v
}

// GetStatus returns the Status field value.
func (o *ClusterJarReplica) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClusterJarReplica) SetStatus(v string) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ClusterJarReplica) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ClusterJarReplica) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ClusterJarReplica) SetError(v string) {
	o.Error = &v
}

// GetGeneration returns the Generation field value.
func (o *ClusterJarReplica) GetGeneration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetGenerationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Generation, true
}

// SetGeneration sets field value.
func (o *ClusterJarReplica) SetGeneration(v int64) {
	o.Generation = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ClusterJarReplica) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterJarReplica) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ClusterJarReplica) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ClusterJarReplica) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterJarReplica) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["podName"] = o.PodName
	toSerialize["podUID"] = o.PodUid
	toSerialize["status"] = o.Status
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	toSerialize["generation"] = o.Generation
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
func (o *ClusterJarReplica) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component  *string    `json:"component"`
		PodName    *string    `json:"podName"`
		PodUid     *string    `json:"podUID"`
		Status     *string    `json:"status"`
		Error      *string    `json:"error,omitempty"`
		Generation *int64     `json:"generation"`
		UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.PodName == nil {
		return fmt.Errorf("required field podName missing")
	}
	if all.PodUid == nil {
		return fmt.Errorf("required field podUID missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Generation == nil {
		return fmt.Errorf("required field generation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "podName", "podUID", "status", "error", "generation", "updatedAt"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.PodName = *all.PodName
	o.PodUid = *all.PodUid
	o.Status = *all.Status
	o.Error = all.Error
	o.Generation = *all.Generation
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
