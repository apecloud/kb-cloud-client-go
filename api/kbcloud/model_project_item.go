// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ProjectItem Project item.
type ProjectItem struct {
	// The namee of the project.
	Name string `json:"name"`
	// The total number of clusters in the project.
	ClusterTotal int32 `json:"clusterTotal"`
	// The time the project was created.
	CreatedAt time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectItem instantiates a new ProjectItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectItem(name string, clusterTotal int32, createdAt time.Time) *ProjectItem {
	this := ProjectItem{}
	this.Name = name
	this.ClusterTotal = clusterTotal
	this.CreatedAt = createdAt
	return &this
}

// NewProjectItemWithDefaults instantiates a new ProjectItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectItemWithDefaults() *ProjectItem {
	this := ProjectItem{}
	return &this
}

// GetName returns the Name field value.
func (o *ProjectItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProjectItem) SetName(v string) {
	o.Name = v
}

// GetClusterTotal returns the ClusterTotal field value.
func (o *ProjectItem) GetClusterTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ClusterTotal
}

// GetClusterTotalOk returns a tuple with the ClusterTotal field value
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetClusterTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterTotal, true
}

// SetClusterTotal sets field value.
func (o *ProjectItem) SetClusterTotal(v int32) {
	o.ClusterTotal = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ProjectItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ProjectItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["clusterTotal"] = o.ClusterTotal
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string    `json:"name"`
		ClusterTotal *int32     `json:"clusterTotal"`
		CreatedAt    *time.Time `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ClusterTotal == nil {
		return fmt.Errorf("required field clusterTotal missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "clusterTotal", "createdAt"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.ClusterTotal = *all.ClusterTotal
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
