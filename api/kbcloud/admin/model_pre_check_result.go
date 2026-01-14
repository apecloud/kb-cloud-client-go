// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PreCheckResult struct {
	Name        *string               `json:"name,omitempty"`
	CheckerName *string               `json:"checkerName,omitempty"`
	Message     common.NullableString `json:"message,omitempty"`
	Advice      common.NullableString `json:"advice,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPreCheckResult instantiates a new PreCheckResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPreCheckResult() *PreCheckResult {
	this := PreCheckResult{}
	return &this
}

// NewPreCheckResultWithDefaults instantiates a new PreCheckResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPreCheckResultWithDefaults() *PreCheckResult {
	this := PreCheckResult{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PreCheckResult) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckResult) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PreCheckResult) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PreCheckResult) SetName(v string) {
	o.Name = &v
}

// GetCheckerName returns the CheckerName field value if set, zero value otherwise.
func (o *PreCheckResult) GetCheckerName() string {
	if o == nil || o.CheckerName == nil {
		var ret string
		return ret
	}
	return *o.CheckerName
}

// GetCheckerNameOk returns a tuple with the CheckerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckResult) GetCheckerNameOk() (*string, bool) {
	if o == nil || o.CheckerName == nil {
		return nil, false
	}
	return o.CheckerName, true
}

// HasCheckerName returns a boolean if a field has been set.
func (o *PreCheckResult) HasCheckerName() bool {
	return o != nil && o.CheckerName != nil
}

// SetCheckerName gets a reference to the given string and assigns it to the CheckerName field.
func (o *PreCheckResult) SetCheckerName(v string) {
	o.CheckerName = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckResult) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckResult) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *PreCheckResult) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given common.NullableString and assigns it to the Message field.
func (o *PreCheckResult) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *PreCheckResult) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *PreCheckResult) UnsetMessage() {
	o.Message.Unset()
}

// GetAdvice returns the Advice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckResult) GetAdvice() string {
	if o == nil || o.Advice.Get() == nil {
		var ret string
		return ret
	}
	return *o.Advice.Get()
}

// GetAdviceOk returns a tuple with the Advice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckResult) GetAdviceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Advice.Get(), o.Advice.IsSet()
}

// HasAdvice returns a boolean if a field has been set.
func (o *PreCheckResult) HasAdvice() bool {
	return o != nil && o.Advice.IsSet()
}

// SetAdvice gets a reference to the given common.NullableString and assigns it to the Advice field.
func (o *PreCheckResult) SetAdvice(v string) {
	o.Advice.Set(&v)
}

// SetAdviceNil sets the value for Advice to be an explicit nil.
func (o *PreCheckResult) SetAdviceNil() {
	o.Advice.Set(nil)
}

// UnsetAdvice ensures that no value is present for Advice, not even an explicit nil.
func (o *PreCheckResult) UnsetAdvice() {
	o.Advice.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PreCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CheckerName != nil {
		toSerialize["checkerName"] = o.CheckerName
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Advice.IsSet() {
		toSerialize["advice"] = o.Advice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PreCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string               `json:"name,omitempty"`
		CheckerName *string               `json:"checkerName,omitempty"`
		Message     common.NullableString `json:"message,omitempty"`
		Advice      common.NullableString `json:"advice,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "checkerName", "message", "advice"})
	} else {
		return err
	}
	o.Name = all.Name
	o.CheckerName = all.CheckerName
	o.Message = all.Message
	o.Advice = all.Advice

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
