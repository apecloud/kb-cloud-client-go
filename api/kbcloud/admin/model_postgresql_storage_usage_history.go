// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlStorageUsageHistory struct {
	// Requested history range. Supported values are 24h and 7d.
	TimeRange string `json:"timeRange"`
	// History query start timestamp in UTC.
	Start *string `json:"start,omitempty"`
	// History query end timestamp in UTC.
	End    *string                       `json:"end,omitempty"`
	Points []PostgresqlStorageUsagePoint `json:"points"`
	Source string                        `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlStorageUsageHistory instantiates a new PostgresqlStorageUsageHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlStorageUsageHistory(timeRange string, points []PostgresqlStorageUsagePoint, source string) *PostgresqlStorageUsageHistory {
	this := PostgresqlStorageUsageHistory{}
	this.TimeRange = timeRange
	this.Points = points
	this.Source = source
	return &this
}

// NewPostgresqlStorageUsageHistoryWithDefaults instantiates a new PostgresqlStorageUsageHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlStorageUsageHistoryWithDefaults() *PostgresqlStorageUsageHistory {
	this := PostgresqlStorageUsageHistory{}
	return &this
}

// GetTimeRange returns the TimeRange field value.
func (o *PostgresqlStorageUsageHistory) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsageHistory) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value.
func (o *PostgresqlStorageUsageHistory) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *PostgresqlStorageUsageHistory) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsageHistory) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *PostgresqlStorageUsageHistory) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *PostgresqlStorageUsageHistory) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *PostgresqlStorageUsageHistory) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsageHistory) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *PostgresqlStorageUsageHistory) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *PostgresqlStorageUsageHistory) SetEnd(v string) {
	o.End = &v
}

// GetPoints returns the Points field value.
func (o *PostgresqlStorageUsageHistory) GetPoints() []PostgresqlStorageUsagePoint {
	if o == nil {
		var ret []PostgresqlStorageUsagePoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsageHistory) GetPointsOk() (*[]PostgresqlStorageUsagePoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *PostgresqlStorageUsageHistory) SetPoints(v []PostgresqlStorageUsagePoint) {
	o.Points = v
}

// GetSource returns the Source field value.
func (o *PostgresqlStorageUsageHistory) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsageHistory) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *PostgresqlStorageUsageHistory) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlStorageUsageHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timeRange"] = o.TimeRange
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	toSerialize["points"] = o.Points
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlStorageUsageHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimeRange *string                        `json:"timeRange"`
		Start     *string                        `json:"start,omitempty"`
		End       *string                        `json:"end,omitempty"`
		Points    *[]PostgresqlStorageUsagePoint `json:"points"`
		Source    *string                        `json:"source"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TimeRange == nil {
		return fmt.Errorf("required field timeRange missing")
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timeRange", "start", "end", "points", "source"})
	} else {
		return err
	}
	o.TimeRange = *all.TimeRange
	o.Start = all.Start
	o.End = all.End
	o.Points = *all.Points
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
