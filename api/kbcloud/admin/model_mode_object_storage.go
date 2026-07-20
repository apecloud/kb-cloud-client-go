// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ModeObjectStorage object storage related configs
type ModeObjectStorage struct {
	// if object storage is enabled for this mode
	Enabled *bool `json:"enabled,omitempty"`
	// Object storage serviceRef configs supported by this mode. Use this when different object
	// storage engines may require different serviceRef or Helm value mappings. When it is set,
	// the cluster create request selects one item by objectStorageConfig.serviceRef.name.
	//
	ServiceRefs []ModeObjectStorageServiceRef `json:"serviceRefs,omitempty"`
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

// GetServiceRefs returns the ServiceRefs field value if set, zero value otherwise.
func (o *ModeObjectStorage) GetServiceRefs() []ModeObjectStorageServiceRef {
	if o == nil || o.ServiceRefs == nil {
		var ret []ModeObjectStorageServiceRef
		return ret
	}
	return o.ServiceRefs
}

// GetServiceRefsOk returns a tuple with the ServiceRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorage) GetServiceRefsOk() (*[]ModeObjectStorageServiceRef, bool) {
	if o == nil || o.ServiceRefs == nil {
		return nil, false
	}
	return &o.ServiceRefs, true
}

// HasServiceRefs returns a boolean if a field has been set.
func (o *ModeObjectStorage) HasServiceRefs() bool {
	return o != nil && o.ServiceRefs != nil
}

// SetServiceRefs gets a reference to the given []ModeObjectStorageServiceRef and assigns it to the ServiceRefs field.
func (o *ModeObjectStorage) SetServiceRefs(v []ModeObjectStorageServiceRef) {
	o.ServiceRefs = v
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
	if o.ServiceRefs != nil {
		toSerialize["serviceRefs"] = o.ServiceRefs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeObjectStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled     *bool                         `json:"enabled,omitempty"`
		ServiceRefs []ModeObjectStorageServiceRef `json:"serviceRefs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "serviceRefs"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.ServiceRefs = all.ServiceRefs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
