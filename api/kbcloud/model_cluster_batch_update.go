// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterBatchUpdate ClusterBatchUpdate is the payload for batch updating multiple clusters with the same update information
type ClusterBatchUpdate struct {
	// List of cluster names to update
	Names []string `json:"names"`
	// ClusterBatchUpdateData contains the update information to be applied to all clusters in the batch
	Update ClusterBatchUpdateData `json:"update"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterBatchUpdate instantiates a new ClusterBatchUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterBatchUpdate(names []string, update ClusterBatchUpdateData) *ClusterBatchUpdate {
	this := ClusterBatchUpdate{}
	this.Names = names
	this.Update = update
	return &this
}

// NewClusterBatchUpdateWithDefaults instantiates a new ClusterBatchUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterBatchUpdateWithDefaults() *ClusterBatchUpdate {
	this := ClusterBatchUpdate{}
	return &this
}

// GetNames returns the Names field value.
func (o *ClusterBatchUpdate) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *ClusterBatchUpdate) GetNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *ClusterBatchUpdate) SetNames(v []string) {
	o.Names = v
}

// GetUpdate returns the Update field value.
func (o *ClusterBatchUpdate) GetUpdate() ClusterBatchUpdateData {
	if o == nil {
		var ret ClusterBatchUpdateData
		return ret
	}
	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value
// and a boolean to check if the value has been set.
func (o *ClusterBatchUpdate) GetUpdateOk() (*ClusterBatchUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Update, true
}

// SetUpdate sets field value.
func (o *ClusterBatchUpdate) SetUpdate(v ClusterBatchUpdateData) {
	o.Update = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterBatchUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["names"] = o.Names
	toSerialize["update"] = o.Update

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterBatchUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Names  *[]string               `json:"names"`
		Update *ClusterBatchUpdateData `json:"update"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Names == nil {
		return fmt.Errorf("required field names missing")
	}
	if all.Update == nil {
		return fmt.Errorf("required field update missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"names", "update"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Names = *all.Names
	if all.Update.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Update = *all.Update

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
