// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterTags struct {
	// The cluster id corresponding to the tag
	ClusterId *string `json:"clusterId,omitempty"`
	Tags      []Tag   `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterTags instantiates a new ClusterTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterTags() *ClusterTags {
	this := ClusterTags{}
	return &this
}

// NewClusterTagsWithDefaults instantiates a new ClusterTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterTagsWithDefaults() *ClusterTags {
	this := ClusterTags{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ClusterTags) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTags) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ClusterTags) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ClusterTags) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ClusterTags) GetTags() []Tag {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTags) GetTagsOk() (*[]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClusterTags) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *ClusterTags) SetTags(v []Tag) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId *string `json:"clusterId,omitempty"`
		Tags      []Tag   `json:"tags,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterId", "tags"})
	} else {
		return err
	}
	o.ClusterId = all.ClusterId
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
