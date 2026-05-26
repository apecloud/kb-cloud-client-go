// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoValidationInfo struct {
	Validator        interface{} `json:"validator,omitempty"`
	ValidationLevel  *string     `json:"validationLevel,omitempty"`
	ValidationAction *string     `json:"validationAction,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoValidationInfo instantiates a new MongoValidationInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoValidationInfo() *MongoValidationInfo {
	this := MongoValidationInfo{}
	return &this
}

// NewMongoValidationInfoWithDefaults instantiates a new MongoValidationInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoValidationInfoWithDefaults() *MongoValidationInfo {
	this := MongoValidationInfo{}
	return &this
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *MongoValidationInfo) GetValidator() interface{} {
	if o == nil || o.Validator == nil {
		var ret interface{}
		return ret
	}
	return o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoValidationInfo) GetValidatorOk() (*interface{}, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return &o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *MongoValidationInfo) HasValidator() bool {
	return o != nil && o.Validator != nil
}

// SetValidator gets a reference to the given interface{} and assigns it to the Validator field.
func (o *MongoValidationInfo) SetValidator(v interface{}) {
	o.Validator = v
}

// GetValidationLevel returns the ValidationLevel field value if set, zero value otherwise.
func (o *MongoValidationInfo) GetValidationLevel() string {
	if o == nil || o.ValidationLevel == nil {
		var ret string
		return ret
	}
	return *o.ValidationLevel
}

// GetValidationLevelOk returns a tuple with the ValidationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoValidationInfo) GetValidationLevelOk() (*string, bool) {
	if o == nil || o.ValidationLevel == nil {
		return nil, false
	}
	return o.ValidationLevel, true
}

// HasValidationLevel returns a boolean if a field has been set.
func (o *MongoValidationInfo) HasValidationLevel() bool {
	return o != nil && o.ValidationLevel != nil
}

// SetValidationLevel gets a reference to the given string and assigns it to the ValidationLevel field.
func (o *MongoValidationInfo) SetValidationLevel(v string) {
	o.ValidationLevel = &v
}

// GetValidationAction returns the ValidationAction field value if set, zero value otherwise.
func (o *MongoValidationInfo) GetValidationAction() string {
	if o == nil || o.ValidationAction == nil {
		var ret string
		return ret
	}
	return *o.ValidationAction
}

// GetValidationActionOk returns a tuple with the ValidationAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoValidationInfo) GetValidationActionOk() (*string, bool) {
	if o == nil || o.ValidationAction == nil {
		return nil, false
	}
	return o.ValidationAction, true
}

// HasValidationAction returns a boolean if a field has been set.
func (o *MongoValidationInfo) HasValidationAction() bool {
	return o != nil && o.ValidationAction != nil
}

// SetValidationAction gets a reference to the given string and assigns it to the ValidationAction field.
func (o *MongoValidationInfo) SetValidationAction(v string) {
	o.ValidationAction = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoValidationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	if o.ValidationLevel != nil {
		toSerialize["validationLevel"] = o.ValidationLevel
	}
	if o.ValidationAction != nil {
		toSerialize["validationAction"] = o.ValidationAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoValidationInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Validator        interface{} `json:"validator,omitempty"`
		ValidationLevel  *string     `json:"validationLevel,omitempty"`
		ValidationAction *string     `json:"validationAction,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"validator", "validationLevel", "validationAction"})
	} else {
		return err
	}
	o.Validator = all.Validator
	o.ValidationLevel = all.ValidationLevel
	o.ValidationAction = all.ValidationAction

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
