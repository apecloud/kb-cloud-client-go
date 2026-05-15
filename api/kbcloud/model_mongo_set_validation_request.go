// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoSetValidationRequest struct {
	Validator interface{} `json:"validator,omitempty"`
	// validationLevel value
	Level *string `json:"level,omitempty"`
	// validationAction value
	Action *string `json:"action,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoSetValidationRequest instantiates a new MongoSetValidationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoSetValidationRequest() *MongoSetValidationRequest {
	this := MongoSetValidationRequest{}
	return &this
}

// NewMongoSetValidationRequestWithDefaults instantiates a new MongoSetValidationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoSetValidationRequestWithDefaults() *MongoSetValidationRequest {
	this := MongoSetValidationRequest{}
	return &this
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *MongoSetValidationRequest) GetValidator() interface{} {
	if o == nil || o.Validator == nil {
		var ret interface{}
		return ret
	}
	return o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSetValidationRequest) GetValidatorOk() (*interface{}, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return &o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *MongoSetValidationRequest) HasValidator() bool {
	return o != nil && o.Validator != nil
}

// SetValidator gets a reference to the given interface{} and assigns it to the Validator field.
func (o *MongoSetValidationRequest) SetValidator(v interface{}) {
	o.Validator = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *MongoSetValidationRequest) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSetValidationRequest) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *MongoSetValidationRequest) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *MongoSetValidationRequest) SetLevel(v string) {
	o.Level = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *MongoSetValidationRequest) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSetValidationRequest) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *MongoSetValidationRequest) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *MongoSetValidationRequest) SetAction(v string) {
	o.Action = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoSetValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoSetValidationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Validator interface{} `json:"validator,omitempty"`
		Level     *string     `json:"level,omitempty"`
		Action    *string     `json:"action,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"validator", "level", "action"})
	} else {
		return err
	}
	o.Validator = all.Validator
	o.Level = all.Level
	o.Action = all.Action

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
