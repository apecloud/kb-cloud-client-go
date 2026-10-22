// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgRetainedResource struct {
	Type            string `json:"type"`
	Count           int64  `json:"count"`
	RetentionPolicy string `json:"retentionPolicy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgRetainedResource instantiates a new OrgRetainedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgRetainedResource(typeVar string, count int64, retentionPolicy string) *OrgRetainedResource {
	this := OrgRetainedResource{}
	this.Type = typeVar
	this.Count = count
	this.RetentionPolicy = retentionPolicy
	return &this
}

// NewOrgRetainedResourceWithDefaults instantiates a new OrgRetainedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgRetainedResourceWithDefaults() *OrgRetainedResource {
	this := OrgRetainedResource{}
	return &this
}

// GetType returns the Type field value.
func (o *OrgRetainedResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgRetainedResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgRetainedResource) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value.
func (o *OrgRetainedResource) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrgRetainedResource) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *OrgRetainedResource) SetCount(v int64) {
	o.Count = v
}

// GetRetentionPolicy returns the RetentionPolicy field value.
func (o *OrgRetainedResource) GetRetentionPolicy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *OrgRetainedResource) GetRetentionPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value.
func (o *OrgRetainedResource) SetRetentionPolicy(v string) {
	o.RetentionPolicy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgRetainedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["count"] = o.Count
	toSerialize["retentionPolicy"] = o.RetentionPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgRetainedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type            *string `json:"type"`
		Count           *int64  `json:"count"`
		RetentionPolicy *string `json:"retentionPolicy"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.RetentionPolicy == nil {
		return fmt.Errorf("required field retentionPolicy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "count", "retentionPolicy"})
	} else {
		return err
	}
	o.Type = *all.Type
	o.Count = *all.Count
	o.RetentionPolicy = *all.RetentionPolicy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
