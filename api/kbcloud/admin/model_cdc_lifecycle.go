// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CdcLifecycle struct {
	PreStart  *CdcLifecycleAction `json:"preStart,omitempty"`
	PostStart *CdcLifecycleAction `json:"postStart,omitempty"`
	PreStop   *CdcLifecycleAction `json:"preStop,omitempty"`
	PostStop  *CdcLifecycleAction `json:"postStop,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcLifecycle instantiates a new CdcLifecycle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcLifecycle() *CdcLifecycle {
	this := CdcLifecycle{}
	return &this
}

// NewCdcLifecycleWithDefaults instantiates a new CdcLifecycle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcLifecycleWithDefaults() *CdcLifecycle {
	this := CdcLifecycle{}
	return &this
}

// GetPreStart returns the PreStart field value if set, zero value otherwise.
func (o *CdcLifecycle) GetPreStart() CdcLifecycleAction {
	if o == nil || o.PreStart == nil {
		var ret CdcLifecycleAction
		return ret
	}
	return *o.PreStart
}

// GetPreStartOk returns a tuple with the PreStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcLifecycle) GetPreStartOk() (*CdcLifecycleAction, bool) {
	if o == nil || o.PreStart == nil {
		return nil, false
	}
	return o.PreStart, true
}

// HasPreStart returns a boolean if a field has been set.
func (o *CdcLifecycle) HasPreStart() bool {
	return o != nil && o.PreStart != nil
}

// SetPreStart gets a reference to the given CdcLifecycleAction and assigns it to the PreStart field.
func (o *CdcLifecycle) SetPreStart(v CdcLifecycleAction) {
	o.PreStart = &v
}

// GetPostStart returns the PostStart field value if set, zero value otherwise.
func (o *CdcLifecycle) GetPostStart() CdcLifecycleAction {
	if o == nil || o.PostStart == nil {
		var ret CdcLifecycleAction
		return ret
	}
	return *o.PostStart
}

// GetPostStartOk returns a tuple with the PostStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcLifecycle) GetPostStartOk() (*CdcLifecycleAction, bool) {
	if o == nil || o.PostStart == nil {
		return nil, false
	}
	return o.PostStart, true
}

// HasPostStart returns a boolean if a field has been set.
func (o *CdcLifecycle) HasPostStart() bool {
	return o != nil && o.PostStart != nil
}

// SetPostStart gets a reference to the given CdcLifecycleAction and assigns it to the PostStart field.
func (o *CdcLifecycle) SetPostStart(v CdcLifecycleAction) {
	o.PostStart = &v
}

// GetPreStop returns the PreStop field value if set, zero value otherwise.
func (o *CdcLifecycle) GetPreStop() CdcLifecycleAction {
	if o == nil || o.PreStop == nil {
		var ret CdcLifecycleAction
		return ret
	}
	return *o.PreStop
}

// GetPreStopOk returns a tuple with the PreStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcLifecycle) GetPreStopOk() (*CdcLifecycleAction, bool) {
	if o == nil || o.PreStop == nil {
		return nil, false
	}
	return o.PreStop, true
}

// HasPreStop returns a boolean if a field has been set.
func (o *CdcLifecycle) HasPreStop() bool {
	return o != nil && o.PreStop != nil
}

// SetPreStop gets a reference to the given CdcLifecycleAction and assigns it to the PreStop field.
func (o *CdcLifecycle) SetPreStop(v CdcLifecycleAction) {
	o.PreStop = &v
}

// GetPostStop returns the PostStop field value if set, zero value otherwise.
func (o *CdcLifecycle) GetPostStop() CdcLifecycleAction {
	if o == nil || o.PostStop == nil {
		var ret CdcLifecycleAction
		return ret
	}
	return *o.PostStop
}

// GetPostStopOk returns a tuple with the PostStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcLifecycle) GetPostStopOk() (*CdcLifecycleAction, bool) {
	if o == nil || o.PostStop == nil {
		return nil, false
	}
	return o.PostStop, true
}

// HasPostStop returns a boolean if a field has been set.
func (o *CdcLifecycle) HasPostStop() bool {
	return o != nil && o.PostStop != nil
}

// SetPostStop gets a reference to the given CdcLifecycleAction and assigns it to the PostStop field.
func (o *CdcLifecycle) SetPostStop(v CdcLifecycleAction) {
	o.PostStop = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PreStart != nil {
		toSerialize["preStart"] = o.PreStart
	}
	if o.PostStart != nil {
		toSerialize["postStart"] = o.PostStart
	}
	if o.PreStop != nil {
		toSerialize["preStop"] = o.PreStop
	}
	if o.PostStop != nil {
		toSerialize["postStop"] = o.PostStop
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcLifecycle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PreStart  *CdcLifecycleAction `json:"preStart,omitempty"`
		PostStart *CdcLifecycleAction `json:"postStart,omitempty"`
		PreStop   *CdcLifecycleAction `json:"preStop,omitempty"`
		PostStop  *CdcLifecycleAction `json:"postStop,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"preStart", "postStart", "preStop", "postStop"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.PreStart != nil && all.PreStart.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PreStart = all.PreStart
	if all.PostStart != nil && all.PostStart.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PostStart = all.PostStart
	if all.PreStop != nil && all.PreStop.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PreStop = all.PreStop
	if all.PostStop != nil && all.PostStop.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PostStop = all.PostStop

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
