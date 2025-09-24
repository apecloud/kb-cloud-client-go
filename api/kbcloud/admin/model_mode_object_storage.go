// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeObjectStorage object storage related configs
type ModeObjectStorage struct {
	// if object storage is enabled for this mode
	Enabled *bool `json:"enabled,omitempty"`
	// Defines a ServiceRef for a cluster, enabling access to both external services and
	// Services provided by other Clusters. The defined serviceRef must be provided when creating cluster.
	//
	ServiceRef *ModeServiceRef `json:"serviceRef,omitempty"`
	// The path in helm values that some object storage config will use. If empty, the values will not be set.
	AdditionalHelmValuePath *ModeObjectStorageAdditionalHelmValuePath `json:"additionalHelmValuePath,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeObjectStorage instantiates a new ModeObjectStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeObjectStorage() *ModeObjectStorage {
	this := ModeObjectStorage{}
	return &this
}

// NewModeObjectStorageWithDefaults instantiates a new ModeObjectStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeObjectStorageWithDefaults() *ModeObjectStorage {
	this := ModeObjectStorage{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ModeObjectStorage) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorage) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ModeObjectStorage) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ModeObjectStorage) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceRef returns the ServiceRef field value if set, zero value otherwise.
func (o *ModeObjectStorage) GetServiceRef() ModeServiceRef {
	if o == nil || o.ServiceRef == nil {
		var ret ModeServiceRef
		return ret
	}
	return *o.ServiceRef
}

// GetServiceRefOk returns a tuple with the ServiceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorage) GetServiceRefOk() (*ModeServiceRef, bool) {
	if o == nil || o.ServiceRef == nil {
		return nil, false
	}
	return o.ServiceRef, true
}

// HasServiceRef returns a boolean if a field has been set.
func (o *ModeObjectStorage) HasServiceRef() bool {
	return o != nil && o.ServiceRef != nil
}

// SetServiceRef gets a reference to the given ModeServiceRef and assigns it to the ServiceRef field.
func (o *ModeObjectStorage) SetServiceRef(v ModeServiceRef) {
	o.ServiceRef = &v
}

// GetAdditionalHelmValuePath returns the AdditionalHelmValuePath field value if set, zero value otherwise.
func (o *ModeObjectStorage) GetAdditionalHelmValuePath() ModeObjectStorageAdditionalHelmValuePath {
	if o == nil || o.AdditionalHelmValuePath == nil {
		var ret ModeObjectStorageAdditionalHelmValuePath
		return ret
	}
	return *o.AdditionalHelmValuePath
}

// GetAdditionalHelmValuePathOk returns a tuple with the AdditionalHelmValuePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorage) GetAdditionalHelmValuePathOk() (*ModeObjectStorageAdditionalHelmValuePath, bool) {
	if o == nil || o.AdditionalHelmValuePath == nil {
		return nil, false
	}
	return o.AdditionalHelmValuePath, true
}

// HasAdditionalHelmValuePath returns a boolean if a field has been set.
func (o *ModeObjectStorage) HasAdditionalHelmValuePath() bool {
	return o != nil && o.AdditionalHelmValuePath != nil
}

// SetAdditionalHelmValuePath gets a reference to the given ModeObjectStorageAdditionalHelmValuePath and assigns it to the AdditionalHelmValuePath field.
func (o *ModeObjectStorage) SetAdditionalHelmValuePath(v ModeObjectStorageAdditionalHelmValuePath) {
	o.AdditionalHelmValuePath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeObjectStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ServiceRef != nil {
		toSerialize["serviceRef"] = o.ServiceRef
	}
	if o.AdditionalHelmValuePath != nil {
		toSerialize["additionalHelmValuePath"] = o.AdditionalHelmValuePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeObjectStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled                 *bool                                     `json:"enabled,omitempty"`
		ServiceRef              *ModeServiceRef                           `json:"serviceRef,omitempty"`
		AdditionalHelmValuePath *ModeObjectStorageAdditionalHelmValuePath `json:"additionalHelmValuePath,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "serviceRef", "additionalHelmValuePath"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.ServiceRef != nil && all.ServiceRef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServiceRef = all.ServiceRef
	if all.AdditionalHelmValuePath != nil && all.AdditionalHelmValuePath.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AdditionalHelmValuePath = all.AdditionalHelmValuePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
