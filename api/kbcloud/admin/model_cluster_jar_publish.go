// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterJarPublish struct {
	Generation int64 `json:"generation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterJarPublish instantiates a new ClusterJarPublish object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterJarPublish(generation int64) *ClusterJarPublish {
	this := ClusterJarPublish{}
	this.Generation = generation
	return &this
}

// NewClusterJarPublishWithDefaults instantiates a new ClusterJarPublish object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterJarPublishWithDefaults() *ClusterJarPublish {
	this := ClusterJarPublish{}
	return &this
}

// GetGeneration returns the Generation field value.
func (o *ClusterJarPublish) GetGeneration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPublish) GetGenerationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Generation, true
}

// SetGeneration sets field value.
func (o *ClusterJarPublish) SetGeneration(v int64) {
	o.Generation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterJarPublish) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["generation"] = o.Generation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterJarPublish) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Generation *int64 `json:"generation"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Generation == nil {
		return fmt.Errorf("required field generation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"generation"})
	} else {
		return err
	}
	o.Generation = *all.Generation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
