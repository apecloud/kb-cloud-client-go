// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterMaintainceWindow the maintenance window for a cluster
type ClusterMaintainceWindow struct {
	// start hour, must be 0-23, use UTC timezone
	StartHour *int32 `json:"startHour,omitempty"`
	// end hour, must be 0-23 and greater than start hour, use UTC timezone
	EndHour *int32 `json:"endHour,omitempty"`
	// weekdays, comma separated values of 1-7
	Weekdays *string `json:"weekdays,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterMaintainceWindow instantiates a new ClusterMaintainceWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterMaintainceWindow() *ClusterMaintainceWindow {
	this := ClusterMaintainceWindow{}
	return &this
}

// NewClusterMaintainceWindowWithDefaults instantiates a new ClusterMaintainceWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterMaintainceWindowWithDefaults() *ClusterMaintainceWindow {
	this := ClusterMaintainceWindow{}
	return &this
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *ClusterMaintainceWindow) GetStartHour() int32 {
	if o == nil || o.StartHour == nil {
		var ret int32
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMaintainceWindow) GetStartHourOk() (*int32, bool) {
	if o == nil || o.StartHour == nil {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *ClusterMaintainceWindow) HasStartHour() bool {
	return o != nil && o.StartHour != nil
}

// SetStartHour gets a reference to the given int32 and assigns it to the StartHour field.
func (o *ClusterMaintainceWindow) SetStartHour(v int32) {
	o.StartHour = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *ClusterMaintainceWindow) GetEndHour() int32 {
	if o == nil || o.EndHour == nil {
		var ret int32
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMaintainceWindow) GetEndHourOk() (*int32, bool) {
	if o == nil || o.EndHour == nil {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *ClusterMaintainceWindow) HasEndHour() bool {
	return o != nil && o.EndHour != nil
}

// SetEndHour gets a reference to the given int32 and assigns it to the EndHour field.
func (o *ClusterMaintainceWindow) SetEndHour(v int32) {
	o.EndHour = &v
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise.
func (o *ClusterMaintainceWindow) GetWeekdays() string {
	if o == nil || o.Weekdays == nil {
		var ret string
		return ret
	}
	return *o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMaintainceWindow) GetWeekdaysOk() (*string, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *ClusterMaintainceWindow) HasWeekdays() bool {
	return o != nil && o.Weekdays != nil
}

// SetWeekdays gets a reference to the given string and assigns it to the Weekdays field.
func (o *ClusterMaintainceWindow) SetWeekdays(v string) {
	o.Weekdays = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterMaintainceWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StartHour != nil {
		toSerialize["startHour"] = o.StartHour
	}
	if o.EndHour != nil {
		toSerialize["endHour"] = o.EndHour
	}
	if o.Weekdays != nil {
		toSerialize["weekdays"] = o.Weekdays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterMaintainceWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartHour *int32  `json:"startHour,omitempty"`
		EndHour   *int32  `json:"endHour,omitempty"`
		Weekdays  *string `json:"weekdays,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startHour", "endHour", "weekdays"})
	} else {
		return err
	}
	o.StartHour = all.StartHour
	o.EndHour = all.EndHour
	o.Weekdays = all.Weekdays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
