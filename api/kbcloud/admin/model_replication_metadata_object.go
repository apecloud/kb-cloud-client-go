// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ReplicationMetadataObject struct {
	MetadataType *string                     `json:"metadataType,omitempty"`
	Children     []ReplicationMetadataObject `json:"children,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplicationMetadataObject instantiates a new ReplicationMetadataObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplicationMetadataObject() *ReplicationMetadataObject {
	this := ReplicationMetadataObject{}
	return &this
}

// NewReplicationMetadataObjectWithDefaults instantiates a new ReplicationMetadataObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplicationMetadataObjectWithDefaults() *ReplicationMetadataObject {
	this := ReplicationMetadataObject{}
	return &this
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *ReplicationMetadataObject) GetMetadataType() string {
	if o == nil || o.MetadataType == nil {
		var ret string
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationMetadataObject) GetMetadataTypeOk() (*string, bool) {
	if o == nil || o.MetadataType == nil {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *ReplicationMetadataObject) HasMetadataType() bool {
	return o != nil && o.MetadataType != nil
}

// SetMetadataType gets a reference to the given string and assigns it to the MetadataType field.
func (o *ReplicationMetadataObject) SetMetadataType(v string) {
	o.MetadataType = &v
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationMetadataObject) GetChildren() []ReplicationMetadataObject {
	if o == nil {
		var ret []ReplicationMetadataObject
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplicationMetadataObject) GetChildrenOk() (*[]ReplicationMetadataObject, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ReplicationMetadataObject) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []ReplicationMetadataObject and assigns it to the Children field.
func (o *ReplicationMetadataObject) SetChildren(v []ReplicationMetadataObject) {
	o.Children = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplicationMetadataObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MetadataType != nil {
		toSerialize["metadataType"] = o.MetadataType
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplicationMetadataObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetadataType *string                     `json:"metadataType,omitempty"`
		Children     []ReplicationMetadataObject `json:"children,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metadataType", "children"})
	} else {
		return err
	}
	o.MetadataType = all.MetadataType
	o.Children = all.Children

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
