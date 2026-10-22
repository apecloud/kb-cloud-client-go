// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgDeletionBlocker struct {
	Code         string `json:"code"`
	ResourceType string `json:"resourceType"`
	Count        int64  `json:"count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgDeletionBlocker instantiates a new OrgDeletionBlocker object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgDeletionBlocker(code string, resourceType string, count int64) *OrgDeletionBlocker {
	this := OrgDeletionBlocker{}
	this.Code = code
	this.ResourceType = resourceType
	this.Count = count
	return &this
}

// NewOrgDeletionBlockerWithDefaults instantiates a new OrgDeletionBlocker object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgDeletionBlockerWithDefaults() *OrgDeletionBlocker {
	this := OrgDeletionBlocker{}
	return &this
}

// GetCode returns the Code field value.
func (o *OrgDeletionBlocker) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionBlocker) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *OrgDeletionBlocker) SetCode(v string) {
	o.Code = v
}

// GetResourceType returns the ResourceType field value.
func (o *OrgDeletionBlocker) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionBlocker) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *OrgDeletionBlocker) SetResourceType(v string) {
	o.ResourceType = v
}

// GetCount returns the Count field value.
func (o *OrgDeletionBlocker) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionBlocker) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *OrgDeletionBlocker) SetCount(v int64) {
	o.Count = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgDeletionBlocker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["count"] = o.Count

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgDeletionBlocker) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code         *string `json:"code"`
		ResourceType *string `json:"resourceType"`
		Count        *int64  `json:"count"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resourceType missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "resourceType", "count"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.ResourceType = *all.ResourceType
	o.Count = *all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
