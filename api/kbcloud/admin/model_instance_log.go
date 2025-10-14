// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// InstanceLog Instance log is the container log of the specified cluster instance
type InstanceLog struct {
	// containers is all containers of the instance(pod)
	Containers []string `json:"containers,omitempty"`
	// log is the log of the specified container
	Log *string `json:"log,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceLog instantiates a new InstanceLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceLog() *InstanceLog {
	this := InstanceLog{}
	return &this
}

// NewInstanceLogWithDefaults instantiates a new InstanceLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceLogWithDefaults() *InstanceLog {
	this := InstanceLog{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *InstanceLog) GetContainers() []string {
	if o == nil || o.Containers == nil {
		var ret []string
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceLog) GetContainersOk() (*[]string, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return &o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *InstanceLog) HasContainers() bool {
	return o != nil && o.Containers != nil
}

// SetContainers gets a reference to the given []string and assigns it to the Containers field.
func (o *InstanceLog) SetContainers(v []string) {
	o.Containers = v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *InstanceLog) GetLog() string {
	if o == nil || o.Log == nil {
		var ret string
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceLog) GetLogOk() (*string, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *InstanceLog) HasLog() bool {
	return o != nil && o.Log != nil
}

// SetLog gets a reference to the given string and assigns it to the Log field.
func (o *InstanceLog) SetLog(v string) {
	o.Log = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Containers []string `json:"containers,omitempty"`
		Log        *string  `json:"log,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"containers", "log"})
	} else {
		return err
	}
	o.Containers = all.Containers
	o.Log = all.Log

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
