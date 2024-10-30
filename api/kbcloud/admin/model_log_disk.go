// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type LogDisk struct {
	LogDiskCapacity *string `json:"log_disk_capacity,omitempty"`
	LogDiskAssigned *string `json:"log_disk_assigned,omitempty"`
	LogDiskInUse *string `json:"log_disk_in_use,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewLogDisk instantiates a new LogDisk object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogDisk() *LogDisk {
	this := LogDisk{}
	return &this
}

// NewLogDiskWithDefaults instantiates a new LogDisk object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogDiskWithDefaults() *LogDisk {
	this := LogDisk{}
	return &this
}
// GetLogDiskCapacity returns the LogDiskCapacity field value if set, zero value otherwise.
func (o *LogDisk) GetLogDiskCapacity() string {
	if o == nil || o.LogDiskCapacity == nil {
		var ret string
		return ret
	}
	return *o.LogDiskCapacity
}

// GetLogDiskCapacityOk returns a tuple with the LogDiskCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDisk) GetLogDiskCapacityOk() (*string, bool) {
	if o == nil || o.LogDiskCapacity == nil {
		return nil, false
	}
	return o.LogDiskCapacity, true
}

// HasLogDiskCapacity returns a boolean if a field has been set.
func (o *LogDisk) HasLogDiskCapacity() bool {
	return o != nil && o.LogDiskCapacity != nil
}

// SetLogDiskCapacity gets a reference to the given string and assigns it to the LogDiskCapacity field.
func (o *LogDisk) SetLogDiskCapacity(v string) {
	o.LogDiskCapacity = &v
}


// GetLogDiskAssigned returns the LogDiskAssigned field value if set, zero value otherwise.
func (o *LogDisk) GetLogDiskAssigned() string {
	if o == nil || o.LogDiskAssigned == nil {
		var ret string
		return ret
	}
	return *o.LogDiskAssigned
}

// GetLogDiskAssignedOk returns a tuple with the LogDiskAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDisk) GetLogDiskAssignedOk() (*string, bool) {
	if o == nil || o.LogDiskAssigned == nil {
		return nil, false
	}
	return o.LogDiskAssigned, true
}

// HasLogDiskAssigned returns a boolean if a field has been set.
func (o *LogDisk) HasLogDiskAssigned() bool {
	return o != nil && o.LogDiskAssigned != nil
}

// SetLogDiskAssigned gets a reference to the given string and assigns it to the LogDiskAssigned field.
func (o *LogDisk) SetLogDiskAssigned(v string) {
	o.LogDiskAssigned = &v
}


// GetLogDiskInUse returns the LogDiskInUse field value if set, zero value otherwise.
func (o *LogDisk) GetLogDiskInUse() string {
	if o == nil || o.LogDiskInUse == nil {
		var ret string
		return ret
	}
	return *o.LogDiskInUse
}

// GetLogDiskInUseOk returns a tuple with the LogDiskInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDisk) GetLogDiskInUseOk() (*string, bool) {
	if o == nil || o.LogDiskInUse == nil {
		return nil, false
	}
	return o.LogDiskInUse, true
}

// HasLogDiskInUse returns a boolean if a field has been set.
func (o *LogDisk) HasLogDiskInUse() bool {
	return o != nil && o.LogDiskInUse != nil
}

// SetLogDiskInUse gets a reference to the given string and assigns it to the LogDiskInUse field.
func (o *LogDisk) SetLogDiskInUse(v string) {
	o.LogDiskInUse = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o LogDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.LogDiskCapacity != nil {
		toSerialize["log_disk_capacity"] = o.LogDiskCapacity
	}
	if o.LogDiskAssigned != nil {
		toSerialize["log_disk_assigned"] = o.LogDiskAssigned
	}
	if o.LogDiskInUse != nil {
		toSerialize["log_disk_in_use"] = o.LogDiskInUse
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogDisk) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogDiskCapacity *string `json:"log_disk_capacity,omitempty"`
		LogDiskAssigned *string `json:"log_disk_assigned,omitempty"`
		LogDiskInUse *string `json:"log_disk_in_use,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "log_disk_capacity", "log_disk_assigned", "log_disk_in_use",  })
	} else {
		return err
	}
	o.LogDiskCapacity = all.LogDiskCapacity
	o.LogDiskAssigned = all.LogDiskAssigned
	o.LogDiskInUse = all.LogDiskInUse

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
