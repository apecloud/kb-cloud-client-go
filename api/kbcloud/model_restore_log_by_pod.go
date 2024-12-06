// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "time"

// RestoreLogByPod info of restore workload log
type RestoreLogByPod struct {
	// restore pod log
	Log *string `json:"log,omitempty"`
	// pod ownerReference objectkey
	OwnerObjectKey *string `json:"ownerObjectKey,omitempty"`
	// pod name
	PodName *string `json:"podName,omitempty"`
	// Date/time when the restore pod was created.
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreLogByPod instantiates a new RestoreLogByPod object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreLogByPod() *RestoreLogByPod {
	this := RestoreLogByPod{}
	return &this
}

// NewRestoreLogByPodWithDefaults instantiates a new RestoreLogByPod object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreLogByPodWithDefaults() *RestoreLogByPod {
	this := RestoreLogByPod{}
	return &this
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *RestoreLogByPod) GetLog() string {
	if o == nil || o.Log == nil {
		var ret string
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreLogByPod) GetLogOk() (*string, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *RestoreLogByPod) HasLog() bool {
	return o != nil && o.Log != nil
}

// SetLog gets a reference to the given string and assigns it to the Log field.
func (o *RestoreLogByPod) SetLog(v string) {
	o.Log = &v
}

// GetOwnerObjectKey returns the OwnerObjectKey field value if set, zero value otherwise.
func (o *RestoreLogByPod) GetOwnerObjectKey() string {
	if o == nil || o.OwnerObjectKey == nil {
		var ret string
		return ret
	}
	return *o.OwnerObjectKey
}

// GetOwnerObjectKeyOk returns a tuple with the OwnerObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreLogByPod) GetOwnerObjectKeyOk() (*string, bool) {
	if o == nil || o.OwnerObjectKey == nil {
		return nil, false
	}
	return o.OwnerObjectKey, true
}

// HasOwnerObjectKey returns a boolean if a field has been set.
func (o *RestoreLogByPod) HasOwnerObjectKey() bool {
	return o != nil && o.OwnerObjectKey != nil
}

// SetOwnerObjectKey gets a reference to the given string and assigns it to the OwnerObjectKey field.
func (o *RestoreLogByPod) SetOwnerObjectKey(v string) {
	o.OwnerObjectKey = &v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *RestoreLogByPod) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreLogByPod) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *RestoreLogByPod) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *RestoreLogByPod) SetPodName(v string) {
	o.PodName = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *RestoreLogByPod) GetCreationTimestamp() time.Time {
	if o == nil || o.CreationTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreLogByPod) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *RestoreLogByPod) HasCreationTimestamp() bool {
	return o != nil && o.CreationTimestamp != nil
}

// SetCreationTimestamp gets a reference to the given time.Time and assigns it to the CreationTimestamp field.
func (o *RestoreLogByPod) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreLogByPod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.OwnerObjectKey != nil {
		toSerialize["ownerObjectKey"] = o.OwnerObjectKey
	}
	if o.PodName != nil {
		toSerialize["podName"] = o.PodName
	}
	if o.CreationTimestamp != nil {
		if o.CreationTimestamp.Nanosecond() == 0 {
			toSerialize["creationTimestamp"] = o.CreationTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["creationTimestamp"] = o.CreationTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreLogByPod) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Log               *string    `json:"log,omitempty"`
		OwnerObjectKey    *string    `json:"ownerObjectKey,omitempty"`
		PodName           *string    `json:"podName,omitempty"`
		CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"log", "ownerObjectKey", "podName", "creationTimestamp"})
	} else {
		return err
	}
	o.Log = all.Log
	o.OwnerObjectKey = all.OwnerObjectKey
	o.PodName = all.PodName
	o.CreationTimestamp = all.CreationTimestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
