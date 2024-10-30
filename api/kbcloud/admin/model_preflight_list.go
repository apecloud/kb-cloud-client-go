// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// PreflightList The preflight results 
type PreflightList struct {
	Pass []Preflight `json:"pass,omitempty"`
	Warn []Preflight `json:"warn,omitempty"`
	Fail []Preflight `json:"fail,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewPreflightList instantiates a new PreflightList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPreflightList() *PreflightList {
	this := PreflightList{}
	return &this
}

// NewPreflightListWithDefaults instantiates a new PreflightList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPreflightListWithDefaults() *PreflightList {
	this := PreflightList{}
	return &this
}
// GetPass returns the Pass field value if set, zero value otherwise.
func (o *PreflightList) GetPass() []Preflight {
	if o == nil || o.Pass == nil {
		var ret []Preflight
		return ret
	}
	return o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightList) GetPassOk() (*[]Preflight, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return &o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *PreflightList) HasPass() bool {
	return o != nil && o.Pass != nil
}

// SetPass gets a reference to the given []Preflight and assigns it to the Pass field.
func (o *PreflightList) SetPass(v []Preflight) {
	o.Pass = v
}


// GetWarn returns the Warn field value if set, zero value otherwise.
func (o *PreflightList) GetWarn() []Preflight {
	if o == nil || o.Warn == nil {
		var ret []Preflight
		return ret
	}
	return o.Warn
}

// GetWarnOk returns a tuple with the Warn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightList) GetWarnOk() (*[]Preflight, bool) {
	if o == nil || o.Warn == nil {
		return nil, false
	}
	return &o.Warn, true
}

// HasWarn returns a boolean if a field has been set.
func (o *PreflightList) HasWarn() bool {
	return o != nil && o.Warn != nil
}

// SetWarn gets a reference to the given []Preflight and assigns it to the Warn field.
func (o *PreflightList) SetWarn(v []Preflight) {
	o.Warn = v
}


// GetFail returns the Fail field value if set, zero value otherwise.
func (o *PreflightList) GetFail() []Preflight {
	if o == nil || o.Fail == nil {
		var ret []Preflight
		return ret
	}
	return o.Fail
}

// GetFailOk returns a tuple with the Fail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightList) GetFailOk() (*[]Preflight, bool) {
	if o == nil || o.Fail == nil {
		return nil, false
	}
	return &o.Fail, true
}

// HasFail returns a boolean if a field has been set.
func (o *PreflightList) HasFail() bool {
	return o != nil && o.Fail != nil
}

// SetFail gets a reference to the given []Preflight and assigns it to the Fail field.
func (o *PreflightList) SetFail(v []Preflight) {
	o.Fail = v
}



// MarshalJSON serializes the struct using spec logic.
func (o PreflightList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Pass != nil {
		toSerialize["pass"] = o.Pass
	}
	if o.Warn != nil {
		toSerialize["warn"] = o.Warn
	}
	if o.Fail != nil {
		toSerialize["fail"] = o.Fail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PreflightList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pass []Preflight `json:"pass,omitempty"`
		Warn []Preflight `json:"warn,omitempty"`
		Fail []Preflight `json:"fail,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "pass", "warn", "fail",  })
	} else {
		return err
	}
	o.Pass = all.Pass
	o.Warn = all.Warn
	o.Fail = all.Fail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
