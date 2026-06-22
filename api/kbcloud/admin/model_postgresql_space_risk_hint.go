// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSpaceRiskHint struct {
	Type       string  `json:"type"`
	Severity   string  `json:"severity"`
	Message    string  `json:"message"`
	ObjectKind *string `json:"objectKind,omitempty"`
	Schema     *string `json:"schema,omitempty"`
	ObjectName *string `json:"objectName,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSpaceRiskHint instantiates a new PostgresqlSpaceRiskHint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSpaceRiskHint(typeVar string, severity string, message string) *PostgresqlSpaceRiskHint {
	this := PostgresqlSpaceRiskHint{}
	this.Type = typeVar
	this.Severity = severity
	this.Message = message
	return &this
}

// NewPostgresqlSpaceRiskHintWithDefaults instantiates a new PostgresqlSpaceRiskHint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSpaceRiskHintWithDefaults() *PostgresqlSpaceRiskHint {
	this := PostgresqlSpaceRiskHint{}
	return &this
}

// GetType returns the Type field value.
func (o *PostgresqlSpaceRiskHint) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PostgresqlSpaceRiskHint) SetType(v string) {
	o.Type = v
}

// GetSeverity returns the Severity field value.
func (o *PostgresqlSpaceRiskHint) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *PostgresqlSpaceRiskHint) SetSeverity(v string) {
	o.Severity = v
}

// GetMessage returns the Message field value.
func (o *PostgresqlSpaceRiskHint) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *PostgresqlSpaceRiskHint) SetMessage(v string) {
	o.Message = v
}

// GetObjectKind returns the ObjectKind field value if set, zero value otherwise.
func (o *PostgresqlSpaceRiskHint) GetObjectKind() string {
	if o == nil || o.ObjectKind == nil {
		var ret string
		return ret
	}
	return *o.ObjectKind
}

// GetObjectKindOk returns a tuple with the ObjectKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetObjectKindOk() (*string, bool) {
	if o == nil || o.ObjectKind == nil {
		return nil, false
	}
	return o.ObjectKind, true
}

// HasObjectKind returns a boolean if a field has been set.
func (o *PostgresqlSpaceRiskHint) HasObjectKind() bool {
	return o != nil && o.ObjectKind != nil
}

// SetObjectKind gets a reference to the given string and assigns it to the ObjectKind field.
func (o *PostgresqlSpaceRiskHint) SetObjectKind(v string) {
	o.ObjectKind = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostgresqlSpaceRiskHint) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostgresqlSpaceRiskHint) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostgresqlSpaceRiskHint) SetSchema(v string) {
	o.Schema = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *PostgresqlSpaceRiskHint) GetObjectName() string {
	if o == nil || o.ObjectName == nil {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetObjectNameOk() (*string, bool) {
	if o == nil || o.ObjectName == nil {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *PostgresqlSpaceRiskHint) HasObjectName() bool {
	return o != nil && o.ObjectName != nil
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *PostgresqlSpaceRiskHint) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PostgresqlSpaceRiskHint) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceRiskHint) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PostgresqlSpaceRiskHint) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PostgresqlSpaceRiskHint) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSpaceRiskHint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["severity"] = o.Severity
	toSerialize["message"] = o.Message
	if o.ObjectKind != nil {
		toSerialize["objectKind"] = o.ObjectKind
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.ObjectName != nil {
		toSerialize["objectName"] = o.ObjectName
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSpaceRiskHint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type       *string `json:"type"`
		Severity   *string `json:"severity"`
		Message    *string `json:"message"`
		ObjectKind *string `json:"objectKind,omitempty"`
		Schema     *string `json:"schema,omitempty"`
		ObjectName *string `json:"objectName,omitempty"`
		Reason     *string `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "severity", "message", "objectKind", "schema", "objectName", "reason"})
	} else {
		return err
	}
	o.Type = *all.Type
	o.Severity = *all.Severity
	o.Message = *all.Message
	o.ObjectKind = all.ObjectKind
	o.Schema = all.Schema
	o.ObjectName = all.ObjectName
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
