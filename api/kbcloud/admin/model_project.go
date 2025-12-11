// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Project Project with detail information.
type Project struct {
	// The name of the project.
	Name string `json:"name"`
	// The time the project was created.
	CreatedAt time.Time `json:"createdAt"`
	// The environment ID of the project.
	EnvId string `json:"envID"`
	// The description of the project.
	Description *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProject instantiates a new Project object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProject(name string, createdAt time.Time, envId string) *Project {
	this := Project{}
	this.Name = name
	this.CreatedAt = createdAt
	this.EnvId = envId
	return &this
}

// NewProjectWithDefaults instantiates a new Project object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetName returns the Name field value.
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Project) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvId returns the EnvId field value.
func (o *Project) GetEnvId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value
// and a boolean to check if the value has been set.
func (o *Project) GetEnvIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvId, true
}

// SetEnvId sets field value.
func (o *Project) SetEnvId(v string) {
	o.EnvId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Project) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Project) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Project) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["envID"] = o.EnvId
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Project) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string    `json:"name"`
		CreatedAt   *time.Time `json:"createdAt"`
		EnvId       *string    `json:"envID"`
		Description *string    `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.EnvId == nil {
		return fmt.Errorf("required field envID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "createdAt", "envID", "description"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.CreatedAt = *all.CreatedAt
	o.EnvId = *all.EnvId
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
