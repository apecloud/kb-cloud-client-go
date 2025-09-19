// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplListItem parameter template information
type ParamTplListItem struct {
	// Name of the organization
	OrgName *string `json:"orgName,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description of parameter template
	Description string `json:"description"`
	// Major version of database engine, eg: 8.0
	MajorVersion string `json:"majorVersion"`
	// Name of database engine
	Engine string `json:"engine"`
	// Name of component
	Component string `json:"component"`
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// the template partition in listParamTpl request
	Partition string `json:"partition"`
	// id of parameter template
	ParamTplId string `json:"paramTplID"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplListItem instantiates a new ParamTplListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplListItem(description string, majorVersion string, engine string, component string, name string, partition string, paramTplId string) *ParamTplListItem {
	this := ParamTplListItem{}
	this.Description = description
	this.MajorVersion = majorVersion
	this.Engine = engine
	this.Component = component
	this.Name = name
	this.Partition = partition
	this.ParamTplId = paramTplId
	return &this
}

// NewParamTplListItemWithDefaults instantiates a new ParamTplListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplListItemWithDefaults() *ParamTplListItem {
	this := ParamTplListItem{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ParamTplListItem) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ParamTplListItem) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ParamTplListItem) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ParamTplListItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ParamTplListItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ParamTplListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value.
func (o *ParamTplListItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ParamTplListItem) SetDescription(v string) {
	o.Description = v
}

// GetMajorVersion returns the MajorVersion field value.
func (o *ParamTplListItem) GetMajorVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetMajorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MajorVersion, true
}

// SetMajorVersion sets field value.
func (o *ParamTplListItem) SetMajorVersion(v string) {
	o.MajorVersion = v
}

// GetEngine returns the Engine field value.
func (o *ParamTplListItem) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *ParamTplListItem) SetEngine(v string) {
	o.Engine = v
}

// GetComponent returns the Component field value.
func (o *ParamTplListItem) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ParamTplListItem) SetComponent(v string) {
	o.Component = v
}

// GetName returns the Name field value.
func (o *ParamTplListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParamTplListItem) SetName(v string) {
	o.Name = v
}

// GetPartition returns the Partition field value.
func (o *ParamTplListItem) GetPartition() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetPartitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partition, true
}

// SetPartition sets field value.
func (o *ParamTplListItem) SetPartition(v string) {
	o.Partition = v
}

// GetParamTplId returns the ParamTplId field value.
func (o *ParamTplListItem) GetParamTplId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParamTplId
}

// GetParamTplIdOk returns a tuple with the ParamTplId field value
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetParamTplIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParamTplId, true
}

// SetParamTplId sets field value.
func (o *ParamTplListItem) SetParamTplId(v string) {
	o.ParamTplId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ParamTplListItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ParamTplListItem) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ParamTplListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["description"] = o.Description
	toSerialize["majorVersion"] = o.MajorVersion
	toSerialize["engine"] = o.Engine
	toSerialize["component"] = o.Component
	toSerialize["name"] = o.Name
	toSerialize["partition"] = o.Partition
	toSerialize["paramTplID"] = o.ParamTplId
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
func (o *ParamTplListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName      *string    `json:"orgName,omitempty"`
		CreatedAt    *time.Time `json:"createdAt,omitempty"`
		Description  *string    `json:"description"`
		MajorVersion *string    `json:"majorVersion"`
		Engine       *string    `json:"engine"`
		Component    *string    `json:"component"`
		Name         *string    `json:"name"`
		Partition    *string    `json:"partition"`
		ParamTplId   *string    `json:"paramTplID"`
		UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.MajorVersion == nil {
		return fmt.Errorf("required field majorVersion missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Partition == nil {
		return fmt.Errorf("required field partition missing")
	}
	if all.ParamTplId == nil {
		return fmt.Errorf("required field paramTplID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "createdAt", "description", "majorVersion", "engine", "component", "name", "partition", "paramTplID", "updatedAt"})
	} else {
		return err
	}
	o.OrgName = all.OrgName
	o.CreatedAt = all.CreatedAt
	o.Description = *all.Description
	o.MajorVersion = *all.MajorVersion
	o.Engine = *all.Engine
	o.Component = *all.Component
	o.Name = *all.Name
	o.Partition = *all.Partition
	o.ParamTplId = *all.ParamTplId
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
