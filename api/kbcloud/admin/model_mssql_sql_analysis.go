// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MssqlSQLAnalysis Completed executions of statements still in the plan cache, accumulated since each statement creation time. Not a time-window history. Time values use milliseconds; read/write counts use 8 KiB pages.
type MssqlSQLAnalysis struct {
	CapturedAt string              `json:"capturedAt"`
	Source     string              `json:"source"`
	OrderBy    string              `json:"orderBy"`
	Truncated  bool                `json:"truncated"`
	Items      []MssqlSQLStatement `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlSQLAnalysis instantiates a new MssqlSQLAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlSQLAnalysis(capturedAt string, source string, orderBy string, truncated bool, items []MssqlSQLStatement) *MssqlSQLAnalysis {
	this := MssqlSQLAnalysis{}
	this.CapturedAt = capturedAt
	this.Source = source
	this.OrderBy = orderBy
	this.Truncated = truncated
	this.Items = items
	return &this
}

// NewMssqlSQLAnalysisWithDefaults instantiates a new MssqlSQLAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlSQLAnalysisWithDefaults() *MssqlSQLAnalysis {
	this := MssqlSQLAnalysis{}
	return &this
}

// GetCapturedAt returns the CapturedAt field value.
func (o *MssqlSQLAnalysis) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLAnalysis) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *MssqlSQLAnalysis) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// GetSource returns the Source field value.
func (o *MssqlSQLAnalysis) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLAnalysis) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *MssqlSQLAnalysis) SetSource(v string) {
	o.Source = v
}

// GetOrderBy returns the OrderBy field value.
func (o *MssqlSQLAnalysis) GetOrderBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLAnalysis) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// SetOrderBy sets field value.
func (o *MssqlSQLAnalysis) SetOrderBy(v string) {
	o.OrderBy = v
}

// GetTruncated returns the Truncated field value.
func (o *MssqlSQLAnalysis) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLAnalysis) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *MssqlSQLAnalysis) SetTruncated(v bool) {
	o.Truncated = v
}

// GetItems returns the Items field value.
func (o *MssqlSQLAnalysis) GetItems() []MssqlSQLStatement {
	if o == nil {
		var ret []MssqlSQLStatement
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *MssqlSQLAnalysis) GetItemsOk() (*[]MssqlSQLStatement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *MssqlSQLAnalysis) SetItems(v []MssqlSQLStatement) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlSQLAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["capturedAt"] = o.CapturedAt
	toSerialize["source"] = o.Source
	toSerialize["orderBy"] = o.OrderBy
	toSerialize["truncated"] = o.Truncated
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlSQLAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapturedAt *string              `json:"capturedAt"`
		Source     *string              `json:"source"`
		OrderBy    *string              `json:"orderBy"`
		Truncated  *bool                `json:"truncated"`
		Items      *[]MssqlSQLStatement `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.OrderBy == nil {
		return fmt.Errorf("required field orderBy missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capturedAt", "source", "orderBy", "truncated", "items"})
	} else {
		return err
	}
	o.CapturedAt = *all.CapturedAt
	o.Source = *all.Source
	o.OrderBy = *all.OrderBy
	o.Truncated = *all.Truncated
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
