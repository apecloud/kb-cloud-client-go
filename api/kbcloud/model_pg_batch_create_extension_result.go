// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PgBatchCreateExtensionResult struct {
	Database string                       `json:"database"`
	Name     string                       `json:"name"`
	Status   PgBatchCreateExtensionStatus `json:"status"`
	// The error for an extension that failed to install or schedule.
	Error *string `json:"error,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPgBatchCreateExtensionResult instantiates a new PgBatchCreateExtensionResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPgBatchCreateExtensionResult(database string, name string, status PgBatchCreateExtensionStatus) *PgBatchCreateExtensionResult {
	this := PgBatchCreateExtensionResult{}
	this.Database = database
	this.Name = name
	this.Status = status
	return &this
}

// NewPgBatchCreateExtensionResultWithDefaults instantiates a new PgBatchCreateExtensionResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPgBatchCreateExtensionResultWithDefaults() *PgBatchCreateExtensionResult {
	this := PgBatchCreateExtensionResult{}
	return &this
}

// GetDatabase returns the Database field value.
func (o *PgBatchCreateExtensionResult) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionResult) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PgBatchCreateExtensionResult) SetDatabase(v string) {
	o.Database = v
}

// GetName returns the Name field value.
func (o *PgBatchCreateExtensionResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PgBatchCreateExtensionResult) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *PgBatchCreateExtensionResult) GetStatus() PgBatchCreateExtensionStatus {
	if o == nil {
		var ret PgBatchCreateExtensionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionResult) GetStatusOk() (*PgBatchCreateExtensionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PgBatchCreateExtensionResult) SetStatus(v PgBatchCreateExtensionStatus) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PgBatchCreateExtensionResult) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionResult) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PgBatchCreateExtensionResult) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *PgBatchCreateExtensionResult) SetError(v string) {
	o.Error = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PgBatchCreateExtensionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["database"] = o.Database
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PgBatchCreateExtensionResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database *string                       `json:"database"`
		Name     *string                       `json:"name"`
		Status   *PgBatchCreateExtensionStatus `json:"status"`
		Error    *string                       `json:"error,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "name", "status", "error"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Database = *all.Database
	o.Name = *all.Name
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Error = all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
