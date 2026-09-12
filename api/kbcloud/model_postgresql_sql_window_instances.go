// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PostgresqlSQLWindowInstances struct {
	Status  *PostgresqlSQLWindowInstanceStatus `json:"status,omitempty"`
	Reasons []string                           `json:"reasons,omitempty"`
	Items   []PostgresqlSQLWindowInstance      `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowInstances instantiates a new PostgresqlSQLWindowInstances object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowInstances() *PostgresqlSQLWindowInstances {
	this := PostgresqlSQLWindowInstances{}
	return &this
}

// NewPostgresqlSQLWindowInstancesWithDefaults instantiates a new PostgresqlSQLWindowInstances object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowInstancesWithDefaults() *PostgresqlSQLWindowInstances {
	this := PostgresqlSQLWindowInstances{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowInstances) GetStatus() PostgresqlSQLWindowInstanceStatus {
	if o == nil || o.Status == nil {
		var ret PostgresqlSQLWindowInstanceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowInstances) GetStatusOk() (*PostgresqlSQLWindowInstanceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowInstances) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PostgresqlSQLWindowInstanceStatus and assigns it to the Status field.
func (o *PostgresqlSQLWindowInstances) SetStatus(v PostgresqlSQLWindowInstanceStatus) {
	o.Status = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowInstances) GetReasons() []string {
	if o == nil || o.Reasons == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowInstances) GetReasonsOk() (*[]string, bool) {
	if o == nil || o.Reasons == nil {
		return nil, false
	}
	return &o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowInstances) HasReasons() bool {
	return o != nil && o.Reasons != nil
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *PostgresqlSQLWindowInstances) SetReasons(v []string) {
	o.Reasons = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowInstances) GetItems() []PostgresqlSQLWindowInstance {
	if o == nil || o.Items == nil {
		var ret []PostgresqlSQLWindowInstance
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowInstances) GetItemsOk() (*[]PostgresqlSQLWindowInstance, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowInstances) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []PostgresqlSQLWindowInstance and assigns it to the Items field.
func (o *PostgresqlSQLWindowInstances) SetItems(v []PostgresqlSQLWindowInstance) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowInstances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
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
func (o *PostgresqlSQLWindowInstances) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status  *PostgresqlSQLWindowInstanceStatus `json:"status,omitempty"`
		Reasons []string                           `json:"reasons,omitempty"`
		Items   []PostgresqlSQLWindowInstance      `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "reasons", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Reasons = all.Reasons
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
