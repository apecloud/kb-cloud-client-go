// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type ComponentOptionVersionMinorVersion struct {
	// determine whether minor version can be rolled back.
	Rollback *bool `json:"rollback,omitempty"`
	// disable roll back the preRelease minor version.
	DisableRollbackPreRelease *bool `json:"disableRollbackPreRelease,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewComponentOptionVersionMinorVersion instantiates a new ComponentOptionVersionMinorVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOptionVersionMinorVersion() *ComponentOptionVersionMinorVersion {
	this := ComponentOptionVersionMinorVersion{}
	return &this
}

// NewComponentOptionVersionMinorVersionWithDefaults instantiates a new ComponentOptionVersionMinorVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOptionVersionMinorVersionWithDefaults() *ComponentOptionVersionMinorVersion {
	this := ComponentOptionVersionMinorVersion{}
	return &this
}
// GetRollback returns the Rollback field value if set, zero value otherwise.
func (o *ComponentOptionVersionMinorVersion) GetRollback() bool {
	if o == nil || o.Rollback == nil {
		var ret bool
		return ret
	}
	return *o.Rollback
}

// GetRollbackOk returns a tuple with the Rollback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMinorVersion) GetRollbackOk() (*bool, bool) {
	if o == nil || o.Rollback == nil {
		return nil, false
	}
	return o.Rollback, true
}

// HasRollback returns a boolean if a field has been set.
func (o *ComponentOptionVersionMinorVersion) HasRollback() bool {
	return o != nil && o.Rollback != nil
}

// SetRollback gets a reference to the given bool and assigns it to the Rollback field.
func (o *ComponentOptionVersionMinorVersion) SetRollback(v bool) {
	o.Rollback = &v
}


// GetDisableRollbackPreRelease returns the DisableRollbackPreRelease field value if set, zero value otherwise.
func (o *ComponentOptionVersionMinorVersion) GetDisableRollbackPreRelease() bool {
	if o == nil || o.DisableRollbackPreRelease == nil {
		var ret bool
		return ret
	}
	return *o.DisableRollbackPreRelease
}

// GetDisableRollbackPreReleaseOk returns a tuple with the DisableRollbackPreRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMinorVersion) GetDisableRollbackPreReleaseOk() (*bool, bool) {
	if o == nil || o.DisableRollbackPreRelease == nil {
		return nil, false
	}
	return o.DisableRollbackPreRelease, true
}

// HasDisableRollbackPreRelease returns a boolean if a field has been set.
func (o *ComponentOptionVersionMinorVersion) HasDisableRollbackPreRelease() bool {
	return o != nil && o.DisableRollbackPreRelease != nil
}

// SetDisableRollbackPreRelease gets a reference to the given bool and assigns it to the DisableRollbackPreRelease field.
func (o *ComponentOptionVersionMinorVersion) SetDisableRollbackPreRelease(v bool) {
	o.DisableRollbackPreRelease = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ComponentOptionVersionMinorVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Rollback != nil {
		toSerialize["rollback"] = o.Rollback
	}
	if o.DisableRollbackPreRelease != nil {
		toSerialize["disableRollbackPreRelease"] = o.DisableRollbackPreRelease
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOptionVersionMinorVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rollback *bool `json:"rollback,omitempty"`
		DisableRollbackPreRelease *bool `json:"disableRollbackPreRelease,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "rollback", "disableRollbackPreRelease",  })
	} else {
		return err
	}
	o.Rollback = all.Rollback
	o.DisableRollbackPreRelease = all.DisableRollbackPreRelease

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
