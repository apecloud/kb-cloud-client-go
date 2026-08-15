// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeObjectStorageServiceRef Defines one object storage dependency option for a mode.
type ModeObjectStorageServiceRef struct {
	// Defines a ServiceRef for a cluster, enabling access to both external services and
	// Services provided by other Clusters. The defined serviceRef must be provided when creating cluster.
	//
	ServiceRef ModeServiceRef `json:"serviceRef"`
	// The path in helm values that some object storage config will use. If empty, the values will not be set.
	AdditionalHelmValuePath ModeObjectStorageAdditionalHelmValuePath `json:"additionalHelmValuePath"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeObjectStorageServiceRef instantiates a new ModeObjectStorageServiceRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeObjectStorageServiceRef(serviceRef ModeServiceRef, additionalHelmValuePath ModeObjectStorageAdditionalHelmValuePath) *ModeObjectStorageServiceRef {
	this := ModeObjectStorageServiceRef{}
	this.ServiceRef = serviceRef
	this.AdditionalHelmValuePath = additionalHelmValuePath
	return &this
}

// NewModeObjectStorageServiceRefWithDefaults instantiates a new ModeObjectStorageServiceRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeObjectStorageServiceRefWithDefaults() *ModeObjectStorageServiceRef {
	this := ModeObjectStorageServiceRef{}
	return &this
}

// GetServiceRef returns the ServiceRef field value.
func (o *ModeObjectStorageServiceRef) GetServiceRef() ModeServiceRef {
	if o == nil {
		var ret ModeServiceRef
		return ret
	}
	return o.ServiceRef
}

// GetServiceRefOk returns a tuple with the ServiceRef field value
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageServiceRef) GetServiceRefOk() (*ModeServiceRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceRef, true
}

// SetServiceRef sets field value.
func (o *ModeObjectStorageServiceRef) SetServiceRef(v ModeServiceRef) {
	o.ServiceRef = v
}

// GetAdditionalHelmValuePath returns the AdditionalHelmValuePath field value.
func (o *ModeObjectStorageServiceRef) GetAdditionalHelmValuePath() ModeObjectStorageAdditionalHelmValuePath {
	if o == nil {
		var ret ModeObjectStorageAdditionalHelmValuePath
		return ret
	}
	return o.AdditionalHelmValuePath
}

// GetAdditionalHelmValuePathOk returns a tuple with the AdditionalHelmValuePath field value
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageServiceRef) GetAdditionalHelmValuePathOk() (*ModeObjectStorageAdditionalHelmValuePath, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalHelmValuePath, true
}

// SetAdditionalHelmValuePath sets field value.
func (o *ModeObjectStorageServiceRef) SetAdditionalHelmValuePath(v ModeObjectStorageAdditionalHelmValuePath) {
	o.AdditionalHelmValuePath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeObjectStorageServiceRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["serviceRef"] = o.ServiceRef
	toSerialize["additionalHelmValuePath"] = o.AdditionalHelmValuePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeObjectStorageServiceRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServiceRef              *ModeServiceRef                           `json:"serviceRef"`
		AdditionalHelmValuePath *ModeObjectStorageAdditionalHelmValuePath `json:"additionalHelmValuePath"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ServiceRef == nil {
		return fmt.Errorf("required field serviceRef missing")
	}
	if all.AdditionalHelmValuePath == nil {
		return fmt.Errorf("required field additionalHelmValuePath missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"serviceRef", "additionalHelmValuePath"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ServiceRef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServiceRef = *all.ServiceRef
	if all.AdditionalHelmValuePath.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AdditionalHelmValuePath = *all.AdditionalHelmValuePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
