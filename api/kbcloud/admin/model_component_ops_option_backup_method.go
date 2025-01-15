// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentOpsOptionBackupMethod indicate the backup method when inplace is true
type ComponentOpsOptionBackupMethod struct {
	// full backup methods
	Full []string `json:"full,omitempty"`
	// incremental backup methods
	Incremental []string `json:"incremental,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOptionBackupMethod instantiates a new ComponentOpsOptionBackupMethod object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOptionBackupMethod() *ComponentOpsOptionBackupMethod {
	this := ComponentOpsOptionBackupMethod{}
	return &this
}

// NewComponentOpsOptionBackupMethodWithDefaults instantiates a new ComponentOpsOptionBackupMethod object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionBackupMethodWithDefaults() *ComponentOpsOptionBackupMethod {
	this := ComponentOpsOptionBackupMethod{}
	return &this
}

// GetFull returns the Full field value if set, zero value otherwise.
func (o *ComponentOpsOptionBackupMethod) GetFull() []string {
	if o == nil || o.Full == nil {
		var ret []string
		return ret
	}
	return o.Full
}

// GetFullOk returns a tuple with the Full field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionBackupMethod) GetFullOk() (*[]string, bool) {
	if o == nil || o.Full == nil {
		return nil, false
	}
	return &o.Full, true
}

// HasFull returns a boolean if a field has been set.
func (o *ComponentOpsOptionBackupMethod) HasFull() bool {
	return o != nil && o.Full != nil
}

// SetFull gets a reference to the given []string and assigns it to the Full field.
func (o *ComponentOpsOptionBackupMethod) SetFull(v []string) {
	o.Full = v
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *ComponentOpsOptionBackupMethod) GetIncremental() []string {
	if o == nil || o.Incremental == nil {
		var ret []string
		return ret
	}
	return o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionBackupMethod) GetIncrementalOk() (*[]string, bool) {
	if o == nil || o.Incremental == nil {
		return nil, false
	}
	return &o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *ComponentOpsOptionBackupMethod) HasIncremental() bool {
	return o != nil && o.Incremental != nil
}

// SetIncremental gets a reference to the given []string and assigns it to the Incremental field.
func (o *ComponentOpsOptionBackupMethod) SetIncremental(v []string) {
	o.Incremental = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOptionBackupMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Full != nil {
		toSerialize["full"] = o.Full
	}
	if o.Incremental != nil {
		toSerialize["incremental"] = o.Incremental
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOptionBackupMethod) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Full        []string `json:"full,omitempty"`
		Incremental []string `json:"incremental,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"full", "incremental"})
	} else {
		return err
	}
	o.Full = all.Full
	o.Incremental = all.Incremental

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
