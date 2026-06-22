// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlPerformanceTrendSource struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	// Coarse source status. Values are success, partial, failed, skipped, or noData.
	Status      string  `json:"status"`
	CollectedAt string  `json:"collectedAt"`
	Message     *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlPerformanceTrendSource instantiates a new PostgresqlPerformanceTrendSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlPerformanceTrendSource(name string, kind string, status string, collectedAt string) *PostgresqlPerformanceTrendSource {
	this := PostgresqlPerformanceTrendSource{}
	this.Name = name
	this.Kind = kind
	this.Status = status
	this.CollectedAt = collectedAt
	return &this
}

// NewPostgresqlPerformanceTrendSourceWithDefaults instantiates a new PostgresqlPerformanceTrendSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlPerformanceTrendSourceWithDefaults() *PostgresqlPerformanceTrendSource {
	this := PostgresqlPerformanceTrendSource{}
	return &this
}

// GetName returns the Name field value.
func (o *PostgresqlPerformanceTrendSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PostgresqlPerformanceTrendSource) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value.
func (o *PostgresqlPerformanceTrendSource) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSource) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *PostgresqlPerformanceTrendSource) SetKind(v string) {
	o.Kind = v
}

// GetStatus returns the Status field value.
func (o *PostgresqlPerformanceTrendSource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PostgresqlPerformanceTrendSource) SetStatus(v string) {
	o.Status = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PostgresqlPerformanceTrendSource) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSource) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PostgresqlPerformanceTrendSource) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostgresqlPerformanceTrendSource) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendSource) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostgresqlPerformanceTrendSource) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostgresqlPerformanceTrendSource) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlPerformanceTrendSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["kind"] = o.Kind
	toSerialize["status"] = o.Status
	toSerialize["collectedAt"] = o.CollectedAt
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlPerformanceTrendSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string `json:"name"`
		Kind        *string `json:"kind"`
		Status      *string `json:"status"`
		CollectedAt *string `json:"collectedAt"`
		Message     *string `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "kind", "status", "collectedAt", "message"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Kind = *all.Kind
	o.Status = *all.Status
	o.CollectedAt = *all.CollectedAt
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
