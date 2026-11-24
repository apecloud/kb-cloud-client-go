// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PlatformComponentPodStats struct {
	Running int64 `json:"running"`
	Pending int64 `json:"pending"`
	Failed  int64 `json:"failed"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformComponentPodStats instantiates a new PlatformComponentPodStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformComponentPodStats(running int64, pending int64, failed int64) *PlatformComponentPodStats {
	this := PlatformComponentPodStats{}
	this.Running = running
	this.Pending = pending
	this.Failed = failed
	return &this
}

// NewPlatformComponentPodStatsWithDefaults instantiates a new PlatformComponentPodStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformComponentPodStatsWithDefaults() *PlatformComponentPodStats {
	this := PlatformComponentPodStats{}
	return &this
}

// GetRunning returns the Running field value.
func (o *PlatformComponentPodStats) GetRunning() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Running
}

// GetRunningOk returns a tuple with the Running field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentPodStats) GetRunningOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Running, true
}

// SetRunning sets field value.
func (o *PlatformComponentPodStats) SetRunning(v int64) {
	o.Running = v
}

// GetPending returns the Pending field value.
func (o *PlatformComponentPodStats) GetPending() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentPodStats) GetPendingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value.
func (o *PlatformComponentPodStats) SetPending(v int64) {
	o.Pending = v
}

// GetFailed returns the Failed field value.
func (o *PlatformComponentPodStats) GetFailed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *PlatformComponentPodStats) GetFailedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value.
func (o *PlatformComponentPodStats) SetFailed(v int64) {
	o.Failed = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformComponentPodStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["running"] = o.Running
	toSerialize["pending"] = o.Pending
	toSerialize["failed"] = o.Failed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformComponentPodStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Running *int64 `json:"running"`
		Pending *int64 `json:"pending"`
		Failed  *int64 `json:"failed"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Running == nil {
		return fmt.Errorf("required field running missing")
	}
	if all.Pending == nil {
		return fmt.Errorf("required field pending missing")
	}
	if all.Failed == nil {
		return fmt.Errorf("required field failed missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"running", "pending", "failed"})
	} else {
		return err
	}
	o.Running = *all.Running
	o.Pending = *all.Pending
	o.Failed = *all.Failed

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
