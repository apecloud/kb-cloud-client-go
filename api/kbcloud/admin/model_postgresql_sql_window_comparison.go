// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowComparison struct {
	Current  *PostgresqlSQLWindowResponse        `json:"current,omitempty"`
	Baseline *PostgresqlSQLWindowResponse        `json:"baseline,omitempty"`
	Items    []PostgresqlSQLWindowComparisonItem `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowComparison instantiates a new PostgresqlSQLWindowComparison object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowComparison() *PostgresqlSQLWindowComparison {
	this := PostgresqlSQLWindowComparison{}
	return &this
}

// NewPostgresqlSQLWindowComparisonWithDefaults instantiates a new PostgresqlSQLWindowComparison object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowComparisonWithDefaults() *PostgresqlSQLWindowComparison {
	this := PostgresqlSQLWindowComparison{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowComparison) GetCurrent() PostgresqlSQLWindowResponse {
	if o == nil || o.Current == nil {
		var ret PostgresqlSQLWindowResponse
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowComparison) GetCurrentOk() (*PostgresqlSQLWindowResponse, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowComparison) HasCurrent() bool {
	return o != nil && o.Current != nil
}

// SetCurrent gets a reference to the given PostgresqlSQLWindowResponse and assigns it to the Current field.
func (o *PostgresqlSQLWindowComparison) SetCurrent(v PostgresqlSQLWindowResponse) {
	o.Current = &v
}

// GetBaseline returns the Baseline field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowComparison) GetBaseline() PostgresqlSQLWindowResponse {
	if o == nil || o.Baseline == nil {
		var ret PostgresqlSQLWindowResponse
		return ret
	}
	return *o.Baseline
}

// GetBaselineOk returns a tuple with the Baseline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowComparison) GetBaselineOk() (*PostgresqlSQLWindowResponse, bool) {
	if o == nil || o.Baseline == nil {
		return nil, false
	}
	return o.Baseline, true
}

// HasBaseline returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowComparison) HasBaseline() bool {
	return o != nil && o.Baseline != nil
}

// SetBaseline gets a reference to the given PostgresqlSQLWindowResponse and assigns it to the Baseline field.
func (o *PostgresqlSQLWindowComparison) SetBaseline(v PostgresqlSQLWindowResponse) {
	o.Baseline = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowComparison) GetItems() []PostgresqlSQLWindowComparisonItem {
	if o == nil || o.Items == nil {
		var ret []PostgresqlSQLWindowComparisonItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowComparison) GetItemsOk() (*[]PostgresqlSQLWindowComparisonItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowComparison) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []PostgresqlSQLWindowComparisonItem and assigns it to the Items field.
func (o *PostgresqlSQLWindowComparison) SetItems(v []PostgresqlSQLWindowComparisonItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowComparison) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Baseline != nil {
		toSerialize["baseline"] = o.Baseline
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowComparison) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Current  *PostgresqlSQLWindowResponse        `json:"current,omitempty"`
		Baseline *PostgresqlSQLWindowResponse        `json:"baseline,omitempty"`
		Items    []PostgresqlSQLWindowComparisonItem `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"current", "baseline", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Current != nil && all.Current.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Current = all.Current
	if all.Baseline != nil && all.Baseline.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Baseline = all.Baseline
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
