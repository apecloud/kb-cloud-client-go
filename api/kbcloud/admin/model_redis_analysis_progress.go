// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisAnalysisProgress struct {
	Total               *int64     `json:"total,omitempty"`
	Scanned             *int64     `json:"scanned,omitempty"`
	Processed           *int64     `json:"processed,omitempty"`
	SampleLimit         *int64     `json:"sampleLimit,omitempty"`
	Cursor              *string    `json:"cursor,omitempty"`
	NextCursor          *string    `json:"nextCursor,omitempty"`
	ScannedPercent      *float64   `json:"scannedPercent,omitempty"`
	Complete            *bool      `json:"complete,omitempty"`
	Sampled             *bool      `json:"sampled,omitempty"`
	ExtrapolationFactor *float64   `json:"extrapolationFactor,omitempty"`
	GeneratedAt         *time.Time `json:"generatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisProgress instantiates a new RedisAnalysisProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisProgress() *RedisAnalysisProgress {
	this := RedisAnalysisProgress{}
	return &this
}

// NewRedisAnalysisProgressWithDefaults instantiates a new RedisAnalysisProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisProgressWithDefaults() *RedisAnalysisProgress {
	this := RedisAnalysisProgress{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RedisAnalysisProgress) SetTotal(v int64) {
	o.Total = &v
}

// GetScanned returns the Scanned field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetScanned() int64 {
	if o == nil || o.Scanned == nil {
		var ret int64
		return ret
	}
	return *o.Scanned
}

// GetScannedOk returns a tuple with the Scanned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetScannedOk() (*int64, bool) {
	if o == nil || o.Scanned == nil {
		return nil, false
	}
	return o.Scanned, true
}

// HasScanned returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasScanned() bool {
	return o != nil && o.Scanned != nil
}

// SetScanned gets a reference to the given int64 and assigns it to the Scanned field.
func (o *RedisAnalysisProgress) SetScanned(v int64) {
	o.Scanned = &v
}

// GetProcessed returns the Processed field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetProcessed() int64 {
	if o == nil || o.Processed == nil {
		var ret int64
		return ret
	}
	return *o.Processed
}

// GetProcessedOk returns a tuple with the Processed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetProcessedOk() (*int64, bool) {
	if o == nil || o.Processed == nil {
		return nil, false
	}
	return o.Processed, true
}

// HasProcessed returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasProcessed() bool {
	return o != nil && o.Processed != nil
}

// SetProcessed gets a reference to the given int64 and assigns it to the Processed field.
func (o *RedisAnalysisProgress) SetProcessed(v int64) {
	o.Processed = &v
}

// GetSampleLimit returns the SampleLimit field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetSampleLimit() int64 {
	if o == nil || o.SampleLimit == nil {
		var ret int64
		return ret
	}
	return *o.SampleLimit
}

// GetSampleLimitOk returns a tuple with the SampleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetSampleLimitOk() (*int64, bool) {
	if o == nil || o.SampleLimit == nil {
		return nil, false
	}
	return o.SampleLimit, true
}

// HasSampleLimit returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasSampleLimit() bool {
	return o != nil && o.SampleLimit != nil
}

// SetSampleLimit gets a reference to the given int64 and assigns it to the SampleLimit field.
func (o *RedisAnalysisProgress) SetSampleLimit(v int64) {
	o.SampleLimit = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *RedisAnalysisProgress) SetCursor(v string) {
	o.Cursor = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *RedisAnalysisProgress) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetScannedPercent returns the ScannedPercent field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetScannedPercent() float64 {
	if o == nil || o.ScannedPercent == nil {
		var ret float64
		return ret
	}
	return *o.ScannedPercent
}

// GetScannedPercentOk returns a tuple with the ScannedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetScannedPercentOk() (*float64, bool) {
	if o == nil || o.ScannedPercent == nil {
		return nil, false
	}
	return o.ScannedPercent, true
}

// HasScannedPercent returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasScannedPercent() bool {
	return o != nil && o.ScannedPercent != nil
}

// SetScannedPercent gets a reference to the given float64 and assigns it to the ScannedPercent field.
func (o *RedisAnalysisProgress) SetScannedPercent(v float64) {
	o.ScannedPercent = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasComplete() bool {
	return o != nil && o.Complete != nil
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *RedisAnalysisProgress) SetComplete(v bool) {
	o.Complete = &v
}

// GetSampled returns the Sampled field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetSampled() bool {
	if o == nil || o.Sampled == nil {
		var ret bool
		return ret
	}
	return *o.Sampled
}

// GetSampledOk returns a tuple with the Sampled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetSampledOk() (*bool, bool) {
	if o == nil || o.Sampled == nil {
		return nil, false
	}
	return o.Sampled, true
}

// HasSampled returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasSampled() bool {
	return o != nil && o.Sampled != nil
}

// SetSampled gets a reference to the given bool and assigns it to the Sampled field.
func (o *RedisAnalysisProgress) SetSampled(v bool) {
	o.Sampled = &v
}

// GetExtrapolationFactor returns the ExtrapolationFactor field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetExtrapolationFactor() float64 {
	if o == nil || o.ExtrapolationFactor == nil {
		var ret float64
		return ret
	}
	return *o.ExtrapolationFactor
}

// GetExtrapolationFactorOk returns a tuple with the ExtrapolationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetExtrapolationFactorOk() (*float64, bool) {
	if o == nil || o.ExtrapolationFactor == nil {
		return nil, false
	}
	return o.ExtrapolationFactor, true
}

// HasExtrapolationFactor returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasExtrapolationFactor() bool {
	return o != nil && o.ExtrapolationFactor != nil
}

// SetExtrapolationFactor gets a reference to the given float64 and assigns it to the ExtrapolationFactor field.
func (o *RedisAnalysisProgress) SetExtrapolationFactor(v float64) {
	o.ExtrapolationFactor = &v
}

// GetGeneratedAt returns the GeneratedAt field value if set, zero value otherwise.
func (o *RedisAnalysisProgress) GetGeneratedAt() time.Time {
	if o == nil || o.GeneratedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.GeneratedAt
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisProgress) GetGeneratedAtOk() (*time.Time, bool) {
	if o == nil || o.GeneratedAt == nil {
		return nil, false
	}
	return o.GeneratedAt, true
}

// HasGeneratedAt returns a boolean if a field has been set.
func (o *RedisAnalysisProgress) HasGeneratedAt() bool {
	return o != nil && o.GeneratedAt != nil
}

// SetGeneratedAt gets a reference to the given time.Time and assigns it to the GeneratedAt field.
func (o *RedisAnalysisProgress) SetGeneratedAt(v time.Time) {
	o.GeneratedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Scanned != nil {
		toSerialize["scanned"] = o.Scanned
	}
	if o.Processed != nil {
		toSerialize["processed"] = o.Processed
	}
	if o.SampleLimit != nil {
		toSerialize["sampleLimit"] = o.SampleLimit
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if o.ScannedPercent != nil {
		toSerialize["scannedPercent"] = o.ScannedPercent
	}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}
	if o.Sampled != nil {
		toSerialize["sampled"] = o.Sampled
	}
	if o.ExtrapolationFactor != nil {
		toSerialize["extrapolationFactor"] = o.ExtrapolationFactor
	}
	if o.GeneratedAt != nil {
		if o.GeneratedAt.Nanosecond() == 0 {
			toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total               *int64     `json:"total,omitempty"`
		Scanned             *int64     `json:"scanned,omitempty"`
		Processed           *int64     `json:"processed,omitempty"`
		SampleLimit         *int64     `json:"sampleLimit,omitempty"`
		Cursor              *string    `json:"cursor,omitempty"`
		NextCursor          *string    `json:"nextCursor,omitempty"`
		ScannedPercent      *float64   `json:"scannedPercent,omitempty"`
		Complete            *bool      `json:"complete,omitempty"`
		Sampled             *bool      `json:"sampled,omitempty"`
		ExtrapolationFactor *float64   `json:"extrapolationFactor,omitempty"`
		GeneratedAt         *time.Time `json:"generatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"total", "scanned", "processed", "sampleLimit", "cursor", "nextCursor", "scannedPercent", "complete", "sampled", "extrapolationFactor", "generatedAt"})
	} else {
		return err
	}
	o.Total = all.Total
	o.Scanned = all.Scanned
	o.Processed = all.Processed
	o.SampleLimit = all.SampleLimit
	o.Cursor = all.Cursor
	o.NextCursor = all.NextCursor
	o.ScannedPercent = all.ScannedPercent
	o.Complete = all.Complete
	o.Sampled = all.Sampled
	o.ExtrapolationFactor = all.ExtrapolationFactor
	o.GeneratedAt = all.GeneratedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
