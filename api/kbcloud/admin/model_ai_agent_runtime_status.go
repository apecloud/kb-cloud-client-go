// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRuntimeStatus struct {
	Status  AiAgentRuntimeStatusCode `json:"status"`
	Message *string                  `json:"message,omitempty"`
	Model   *string                  `json:"model,omitempty"`
	Scope   *AiAgentScope            `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRuntimeStatus instantiates a new AiAgentRuntimeStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRuntimeStatus(status AiAgentRuntimeStatusCode) *AiAgentRuntimeStatus {
	this := AiAgentRuntimeStatus{}
	this.Status = status
	return &this
}

// NewAiAgentRuntimeStatusWithDefaults instantiates a new AiAgentRuntimeStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRuntimeStatusWithDefaults() *AiAgentRuntimeStatus {
	this := AiAgentRuntimeStatus{}
	return &this
}

// GetStatus returns the Status field value.
func (o *AiAgentRuntimeStatus) GetStatus() AiAgentRuntimeStatusCode {
	if o == nil {
		var ret AiAgentRuntimeStatusCode
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentRuntimeStatus) GetStatusOk() (*AiAgentRuntimeStatusCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentRuntimeStatus) SetStatus(v AiAgentRuntimeStatusCode) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AiAgentRuntimeStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRuntimeStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AiAgentRuntimeStatus) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AiAgentRuntimeStatus) SetMessage(v string) {
	o.Message = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiAgentRuntimeStatus) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRuntimeStatus) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiAgentRuntimeStatus) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiAgentRuntimeStatus) SetModel(v string) {
	o.Model = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiAgentRuntimeStatus) GetScope() AiAgentScope {
	if o == nil || o.Scope == nil {
		var ret AiAgentScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRuntimeStatus) GetScopeOk() (*AiAgentScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiAgentRuntimeStatus) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiAgentScope and assigns it to the Scope field.
func (o *AiAgentRuntimeStatus) SetScope(v AiAgentScope) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRuntimeStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRuntimeStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status  *AiAgentRuntimeStatusCode `json:"status"`
		Message *string                   `json:"message,omitempty"`
		Model   *string                   `json:"model,omitempty"`
		Scope   *AiAgentScope             `json:"scope,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "message", "model", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Message = all.Message
	o.Model = all.Model
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
