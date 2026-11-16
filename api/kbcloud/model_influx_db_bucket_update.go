// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// InfluxDBBucketUpdate Partial update. Omitted fields are unchanged; empty description clears it. Absent or empty retentionRules does not change retention. To retain forever send everySeconds: 0. System buckets cannot be renamed.
type InfluxDBBucketUpdate struct {
	Name           *string                       `json:"name,omitempty"`
	Description    *string                       `json:"description,omitempty"`
	RetentionRules []InfluxDBBucketRetentionRule `json:"retentionRules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBBucketUpdate instantiates a new InfluxDBBucketUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBBucketUpdate() *InfluxDBBucketUpdate {
	this := InfluxDBBucketUpdate{}
	return &this
}

// NewInfluxDBBucketUpdateWithDefaults instantiates a new InfluxDBBucketUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBBucketUpdateWithDefaults() *InfluxDBBucketUpdate {
	this := InfluxDBBucketUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InfluxDBBucketUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InfluxDBBucketUpdate) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InfluxDBBucketUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfluxDBBucketUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfluxDBBucketUpdate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfluxDBBucketUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetRetentionRules returns the RetentionRules field value if set, zero value otherwise.
func (o *InfluxDBBucketUpdate) GetRetentionRules() []InfluxDBBucketRetentionRule {
	if o == nil || o.RetentionRules == nil {
		var ret []InfluxDBBucketRetentionRule
		return ret
	}
	return o.RetentionRules
}

// GetRetentionRulesOk returns a tuple with the RetentionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketUpdate) GetRetentionRulesOk() (*[]InfluxDBBucketRetentionRule, bool) {
	if o == nil || o.RetentionRules == nil {
		return nil, false
	}
	return &o.RetentionRules, true
}

// HasRetentionRules returns a boolean if a field has been set.
func (o *InfluxDBBucketUpdate) HasRetentionRules() bool {
	return o != nil && o.RetentionRules != nil
}

// SetRetentionRules gets a reference to the given []InfluxDBBucketRetentionRule and assigns it to the RetentionRules field.
func (o *InfluxDBBucketUpdate) SetRetentionRules(v []InfluxDBBucketRetentionRule) {
	o.RetentionRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBBucketUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
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
func (o *InfluxDBBucketUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string                       `json:"name,omitempty"`
		Description    *string                       `json:"description,omitempty"`
		RetentionRules []InfluxDBBucketRetentionRule `json:"retentionRules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "retentionRules"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Description = all.Description
	o.RetentionRules = all.RetentionRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
