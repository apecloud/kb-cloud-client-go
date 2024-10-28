// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION Result
type Result struct {
	// NODESCRIPTION Pagination
	Pagination *Pagination `json:"pagination,omitempty"`
	// NODESCRIPTION Columns
	Columns []string `json:"columns,omitempty"`
	// NODESCRIPTION Rows
	Rows [][]interface{} `json:"rows,omitempty"`
	// NODESCRIPTION Stats
	Stats *ResultStats `json:"stats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResult instantiates a new Result object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResult() *Result {
	this := Result{}
	return &this
}

// NewResultWithDefaults instantiates a new Result object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResultWithDefaults() *Result {
	this := Result{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *Result) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *Result) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *Result) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *Result) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *Result) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *Result) SetColumns(v []string) {
	o.Columns = v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *Result) GetRows() [][]interface{} {
	if o == nil || o.Rows == nil {
		var ret [][]interface{}
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetRowsOk() (*[][]interface{}, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return &o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *Result) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given [][]interface{} and assigns it to the Rows field.
func (o *Result) SetRows(v [][]interface{}) {
	o.Rows = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Result) GetStats() ResultStats {
	if o == nil || o.Stats == nil {
		var ret ResultStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetStatsOk() (*ResultStats, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Result) HasStats() bool {
	return o != nil && o.Stats != nil
}

// SetStats gets a reference to the given ResultStats and assigns it to the Stats field.
func (o *Result) SetStats(v ResultStats) {
	o.Stats = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Result) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Result) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pagination *Pagination     `json:"pagination,omitempty"`
		Columns    []string        `json:"columns,omitempty"`
		Rows       [][]interface{} `json:"rows,omitempty"`
		Stats      *ResultStats    `json:"stats,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pagination", "columns", "rows", "stats"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination
	o.Columns = all.Columns
	o.Rows = all.Rows
	if all.Stats != nil && all.Stats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Stats = all.Stats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
