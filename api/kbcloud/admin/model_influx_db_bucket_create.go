// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBBucketCreate struct {
	OrgId          string                        `json:"orgID"`
	Name           string                        `json:"name"`
	Description    *string                       `json:"description,omitempty"`
	RetentionRules []InfluxDBBucketRetentionRule `json:"retentionRules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBBucketCreate instantiates a new InfluxDBBucketCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBBucketCreate(orgId string, name string) *InfluxDBBucketCreate {
	this := InfluxDBBucketCreate{}
	this.OrgId = orgId
	this.Name = name
	return &this
}

// NewInfluxDBBucketCreateWithDefaults instantiates a new InfluxDBBucketCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBBucketCreateWithDefaults() *InfluxDBBucketCreate {
	this := InfluxDBBucketCreate{}
	return &this
}

// GetOrgId returns the OrgId field value.
func (o *InfluxDBBucketCreate) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketCreate) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *InfluxDBBucketCreate) SetOrgId(v string) {
	o.OrgId = v
}

// GetName returns the Name field value.
func (o *InfluxDBBucketCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InfluxDBBucketCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfluxDBBucketCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfluxDBBucketCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfluxDBBucketCreate) SetDescription(v string) {
	o.Description = &v
}

// GetRetentionRules returns the RetentionRules field value if set, zero value otherwise.
func (o *InfluxDBBucketCreate) GetRetentionRules() []InfluxDBBucketRetentionRule {
	if o == nil || o.RetentionRules == nil {
		var ret []InfluxDBBucketRetentionRule
		return ret
	}
	return o.RetentionRules
}

// GetRetentionRulesOk returns a tuple with the RetentionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketCreate) GetRetentionRulesOk() (*[]InfluxDBBucketRetentionRule, bool) {
	if o == nil || o.RetentionRules == nil {
		return nil, false
	}
	return &o.RetentionRules, true
}

// HasRetentionRules returns a boolean if a field has been set.
func (o *InfluxDBBucketCreate) HasRetentionRules() bool {
	return o != nil && o.RetentionRules != nil
}

// SetRetentionRules gets a reference to the given []InfluxDBBucketRetentionRule and assigns it to the RetentionRules field.
func (o *InfluxDBBucketCreate) SetRetentionRules(v []InfluxDBBucketRetentionRule) {
	o.RetentionRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBBucketCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgID"] = o.OrgId
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.RetentionRules != nil {
		toSerialize["retentionRules"] = o.RetentionRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBBucketCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgId          *string                       `json:"orgID"`
		Name           *string                       `json:"name"`
		Description    *string                       `json:"description,omitempty"`
		RetentionRules []InfluxDBBucketRetentionRule `json:"retentionRules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgID missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgID", "name", "description", "retentionRules"})
	} else {
		return err
	}
	o.OrgId = *all.OrgId
	o.Name = *all.Name
	o.Description = all.Description
	o.RetentionRules = all.RetentionRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
