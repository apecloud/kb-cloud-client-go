// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisParameter One user-supplied value for a PostgreSQL positional placeholder.
type DiagnosticSQLDiagnosisParameter struct {
	// 1-based placeholder position, matching $1, $2, and so on.
	Position int64 `json:"position"`
	// User-provided parameter value. The backend does not return this raw value in copy_context.
	Value string `json:"value"`
	// Minimal value type used to render a PostgreSQL literal safely for EXPLAIN.
	ValueType *DiagnosticSQLDiagnosisParameterType `json:"valueType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisParameter instantiates a new DiagnosticSQLDiagnosisParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisParameter(position int64, value string) *DiagnosticSQLDiagnosisParameter {
	this := DiagnosticSQLDiagnosisParameter{}
	this.Position = position
	this.Value = value
	return &this
}

// NewDiagnosticSQLDiagnosisParameterWithDefaults instantiates a new DiagnosticSQLDiagnosisParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisParameterWithDefaults() *DiagnosticSQLDiagnosisParameter {
	this := DiagnosticSQLDiagnosisParameter{}
	return &this
}

// GetPosition returns the Position field value.
func (o *DiagnosticSQLDiagnosisParameter) GetPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisParameter) GetPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *DiagnosticSQLDiagnosisParameter) SetPosition(v int64) {
	o.Position = v
}

// GetValue returns the Value field value.
func (o *DiagnosticSQLDiagnosisParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *DiagnosticSQLDiagnosisParameter) SetValue(v string) {
	o.Value = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisParameter) GetValueType() DiagnosticSQLDiagnosisParameterType {
	if o == nil || o.ValueType == nil {
		var ret DiagnosticSQLDiagnosisParameterType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisParameter) GetValueTypeOk() (*DiagnosticSQLDiagnosisParameterType, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisParameter) HasValueType() bool {
	return o != nil && o.ValueType != nil
}

// SetValueType gets a reference to the given DiagnosticSQLDiagnosisParameterType and assigns it to the ValueType field.
func (o *DiagnosticSQLDiagnosisParameter) SetValueType(v DiagnosticSQLDiagnosisParameterType) {
	o.ValueType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["position"] = o.Position
	toSerialize["value"] = o.Value
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Position  *int64                               `json:"position"`
		Value     *string                              `json:"value"`
		ValueType *DiagnosticSQLDiagnosisParameterType `json:"valueType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Position == nil {
		return fmt.Errorf("required field position missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"position", "value", "valueType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Position = *all.Position
	o.Value = *all.Value
	if all.ValueType != nil && !all.ValueType.IsValid() {
		hasInvalidField = true
	} else {
		o.ValueType = all.ValueType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
