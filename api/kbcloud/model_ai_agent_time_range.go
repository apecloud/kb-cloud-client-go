// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentTimeRange struct {
	From   *time.Time `json:"from,omitempty"`
	To     *time.Time `json:"to,omitempty"`
	Preset *string    `json:"preset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentTimeRange instantiates a new AiAgentTimeRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentTimeRange() *AiAgentTimeRange {
	this := AiAgentTimeRange{}
	return &this
}

// NewAiAgentTimeRangeWithDefaults instantiates a new AiAgentTimeRange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentTimeRangeWithDefaults() *AiAgentTimeRange {
	this := AiAgentTimeRange{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *AiAgentTimeRange) GetFrom() time.Time {
	if o == nil || o.From == nil {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTimeRange) GetFromOk() (*time.Time, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *AiAgentTimeRange) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *AiAgentTimeRange) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *AiAgentTimeRange) GetTo() time.Time {
	if o == nil || o.To == nil {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTimeRange) GetToOk() (*time.Time, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *AiAgentTimeRange) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *AiAgentTimeRange) SetTo(v time.Time) {
	o.To = &v
}

// GetPreset returns the Preset field value if set, zero value otherwise.
func (o *AiAgentTimeRange) GetPreset() string {
	if o == nil || o.Preset == nil {
		var ret string
		return ret
	}
	return *o.Preset
}

// GetPresetOk returns a tuple with the Preset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTimeRange) GetPresetOk() (*string, bool) {
	if o == nil || o.Preset == nil {
		return nil, false
	}
	return o.Preset, true
}

// HasPreset returns a boolean if a field has been set.
func (o *AiAgentTimeRange) HasPreset() bool {
	return o != nil && o.Preset != nil
}

// SetPreset gets a reference to the given string and assigns it to the Preset field.
func (o *AiAgentTimeRange) SetPreset(v string) {
	o.Preset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		if o.From.Nanosecond() == 0 {
			toSerialize["from"] = o.From.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["from"] = o.From.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.To != nil {
		if o.To.Nanosecond() == 0 {
			toSerialize["to"] = o.To.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["to"] = o.To.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Preset != nil {
		toSerialize["preset"] = o.Preset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentTimeRange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From   *time.Time `json:"from,omitempty"`
		To     *time.Time `json:"to,omitempty"`
		Preset *string    `json:"preset,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"from", "to", "preset"})
	} else {
		return err
	}
	o.From = all.From
	o.To = all.To
	o.Preset = all.Preset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
