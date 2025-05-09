// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterHistory The history of a parameter
type ParameterHistory struct {
	// The name of the parameter
	ParameterName string `json:"parameterName"`
	// The old value of the parameter
	OldValue string `json:"oldValue"`
	// The new value of the parameter
	NewValue string `json:"newValue"`
	// The date and time the parameter was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// event source
	Source *EventSource `json:"source,omitempty"`
	// operator of the event, if source is user, operator is user name; if source is system, operator is system name
	Operator *string `json:"operator,omitempty"`
	// The user ID of the operator
	OperatorId *string `json:"operatorId,omitempty"`
	// The name of the configuration file
	FileName *string `json:"fileName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterHistory instantiates a new ParameterHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterHistory(parameterName string, oldValue string, newValue string, updatedAt time.Time) *ParameterHistory {
	this := ParameterHistory{}
	this.ParameterName = parameterName
	this.OldValue = oldValue
	this.NewValue = newValue
	this.UpdatedAt = updatedAt
	return &this
}

// NewParameterHistoryWithDefaults instantiates a new ParameterHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterHistoryWithDefaults() *ParameterHistory {
	this := ParameterHistory{}
	return &this
}

// GetParameterName returns the ParameterName field value.
func (o *ParameterHistory) GetParameterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetParameterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterName, true
}

// SetParameterName sets field value.
func (o *ParameterHistory) SetParameterName(v string) {
	o.ParameterName = v
}

// GetOldValue returns the OldValue field value.
func (o *ParameterHistory) GetOldValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value.
func (o *ParameterHistory) SetOldValue(v string) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value.
func (o *ParameterHistory) GetNewValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value.
func (o *ParameterHistory) SetNewValue(v string) {
	o.NewValue = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ParameterHistory) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ParameterHistory) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ParameterHistory) GetSource() EventSource {
	if o == nil || o.Source == nil {
		var ret EventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetSourceOk() (*EventSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ParameterHistory) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given EventSource and assigns it to the Source field.
func (o *ParameterHistory) SetSource(v EventSource) {
	o.Source = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ParameterHistory) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ParameterHistory) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ParameterHistory) SetOperator(v string) {
	o.Operator = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *ParameterHistory) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *ParameterHistory) HasOperatorId() bool {
	return o != nil && o.OperatorId != nil
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *ParameterHistory) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ParameterHistory) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ParameterHistory) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ParameterHistory) SetFileName(v string) {
	o.FileName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["parameterName"] = o.ParameterName
	toSerialize["oldValue"] = o.OldValue
	toSerialize["newValue"] = o.NewValue
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ParameterName *string      `json:"parameterName"`
		OldValue      *string      `json:"oldValue"`
		NewValue      *string      `json:"newValue"`
		UpdatedAt     *time.Time   `json:"updatedAt"`
		Source        *EventSource `json:"source,omitempty"`
		Operator      *string      `json:"operator,omitempty"`
		OperatorId    *string      `json:"operatorId,omitempty"`
		FileName      *string      `json:"fileName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ParameterName == nil {
		return fmt.Errorf("required field parameterName missing")
	}
	if all.OldValue == nil {
		return fmt.Errorf("required field oldValue missing")
	}
	if all.NewValue == nil {
		return fmt.Errorf("required field newValue missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"parameterName", "oldValue", "newValue", "updatedAt", "source", "operator", "operatorId", "fileName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ParameterName = *all.ParameterName
	o.OldValue = *all.OldValue
	o.NewValue = *all.NewValue
	o.UpdatedAt = *all.UpdatedAt
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.Operator = all.Operator
	o.OperatorId = all.OperatorId
	o.FileName = all.FileName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
