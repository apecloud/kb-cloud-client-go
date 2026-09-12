// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowResponse struct {
	Window  *PostgresqlSQLWindow       `json:"window,omitempty"`
	Status  *PostgresqlSQLWindowStatus `json:"status,omitempty"`
	Reasons []string                   `json:"reasons,omitempty"`
	// Collection-level completeness before database, user or SQL filtering. expectedCollections is estimated from the recorded cadence; zero means cadence is unavailable when no coverage exists.
	Coverage *PostgresqlSQLWindowCoverage `json:"coverage,omitempty"`
	Items    []PostgresqlSQLWindowItem    `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowResponse instantiates a new PostgresqlSQLWindowResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowResponse() *PostgresqlSQLWindowResponse {
	this := PostgresqlSQLWindowResponse{}
	return &this
}

// NewPostgresqlSQLWindowResponseWithDefaults instantiates a new PostgresqlSQLWindowResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowResponseWithDefaults() *PostgresqlSQLWindowResponse {
	this := PostgresqlSQLWindowResponse{}
	return &this
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowResponse) GetWindow() PostgresqlSQLWindow {
	if o == nil || o.Window == nil {
		var ret PostgresqlSQLWindow
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowResponse) GetWindowOk() (*PostgresqlSQLWindow, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResponse) HasWindow() bool {
	return o != nil && o.Window != nil
}

// SetWindow gets a reference to the given PostgresqlSQLWindow and assigns it to the Window field.
func (o *PostgresqlSQLWindowResponse) SetWindow(v PostgresqlSQLWindow) {
	o.Window = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowResponse) GetStatus() PostgresqlSQLWindowStatus {
	if o == nil || o.Status == nil {
		var ret PostgresqlSQLWindowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowResponse) GetStatusOk() (*PostgresqlSQLWindowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PostgresqlSQLWindowStatus and assigns it to the Status field.
func (o *PostgresqlSQLWindowResponse) SetStatus(v PostgresqlSQLWindowStatus) {
	o.Status = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowResponse) GetReasons() []string {
	if o == nil || o.Reasons == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowResponse) GetReasonsOk() (*[]string, bool) {
	if o == nil || o.Reasons == nil {
		return nil, false
	}
	return &o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResponse) HasReasons() bool {
	return o != nil && o.Reasons != nil
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *PostgresqlSQLWindowResponse) SetReasons(v []string) {
	o.Reasons = v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowResponse) GetCoverage() PostgresqlSQLWindowCoverage {
	if o == nil || o.Coverage == nil {
		var ret PostgresqlSQLWindowCoverage
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowResponse) GetCoverageOk() (*PostgresqlSQLWindowCoverage, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResponse) HasCoverage() bool {
	return o != nil && o.Coverage != nil
}

// SetCoverage gets a reference to the given PostgresqlSQLWindowCoverage and assigns it to the Coverage field.
func (o *PostgresqlSQLWindowResponse) SetCoverage(v PostgresqlSQLWindowCoverage) {
	o.Coverage = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowResponse) GetItems() []PostgresqlSQLWindowItem {
	if o == nil || o.Items == nil {
		var ret []PostgresqlSQLWindowItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowResponse) GetItemsOk() (*[]PostgresqlSQLWindowItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []PostgresqlSQLWindowItem and assigns it to the Items field.
func (o *PostgresqlSQLWindowResponse) SetItems(v []PostgresqlSQLWindowItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
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
func (o *PostgresqlSQLWindowResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Window   *PostgresqlSQLWindow         `json:"window,omitempty"`
		Status   *PostgresqlSQLWindowStatus   `json:"status,omitempty"`
		Reasons  []string                     `json:"reasons,omitempty"`
		Coverage *PostgresqlSQLWindowCoverage `json:"coverage,omitempty"`
		Items    []PostgresqlSQLWindowItem    `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"window", "status", "reasons", "coverage", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Window != nil && all.Window.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Window = all.Window
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Reasons = all.Reasons
	if all.Coverage != nil && all.Coverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Coverage = all.Coverage
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
