// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseSQLAnalysisRange struct {
	// Requested range label (1h, 6h, 24h, 7d).
	Label *string `json:"label,omitempty"`
	// Range start in RFC3339 UTC.
	Start *string `json:"start,omitempty"`
	// Range end in RFC3339 UTC.
	End *string `json:"end,omitempty"`
	// Range duration in seconds.
	DurationSeconds *int64 `json:"durationSeconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSQLAnalysisRange instantiates a new ClickhouseSQLAnalysisRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSQLAnalysisRange() *ClickhouseSQLAnalysisRange {
	this := ClickhouseSQLAnalysisRange{}
	return &this
}

// NewClickhouseSQLAnalysisRangeWithDefaults instantiates a new ClickhouseSQLAnalysisRange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSQLAnalysisRangeWithDefaults() *ClickhouseSQLAnalysisRange {
	this := ClickhouseSQLAnalysisRange{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysisRange) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysisRange) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysisRange) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ClickhouseSQLAnalysisRange) SetLabel(v string) {
	o.Label = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysisRange) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysisRange) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysisRange) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *ClickhouseSQLAnalysisRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysisRange) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysisRange) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysisRange) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *ClickhouseSQLAnalysisRange) SetEnd(v string) {
	o.End = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *ClickhouseSQLAnalysisRange) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSQLAnalysisRange) GetDurationSecondsOk() (*int64, bool) {
	if o == nil || o.DurationSeconds == nil {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *ClickhouseSQLAnalysisRange) HasDurationSeconds() bool {
	return o != nil && o.DurationSeconds != nil
}

// SetDurationSeconds gets a reference to the given int64 and assigns it to the DurationSeconds field.
func (o *ClickhouseSQLAnalysisRange) SetDurationSeconds(v int64) {
	o.DurationSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSQLAnalysisRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.DurationSeconds != nil {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSQLAnalysisRange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label           *string `json:"label,omitempty"`
		Start           *string `json:"start,omitempty"`
		End             *string `json:"end,omitempty"`
		DurationSeconds *int64  `json:"durationSeconds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"label", "start", "end", "durationSeconds"})
	} else {
		return err
	}
	o.Label = all.Label
	o.Start = all.Start
	o.End = all.End
	o.DurationSeconds = all.DurationSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
