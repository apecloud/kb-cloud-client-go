// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseDiagnosticSource struct {
	Name        string                         `json:"name"`
	Status      KingbaseDiagnosticSourceStatus `json:"status"`
	CollectedAt string                         `json:"collectedAt"`
	Message     *string                        `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKingbaseDiagnosticSource instantiates a new KingbaseDiagnosticSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKingbaseDiagnosticSource(name string, status KingbaseDiagnosticSourceStatus, collectedAt string) *KingbaseDiagnosticSource {
	this := KingbaseDiagnosticSource{}
	this.Name = name
	this.Status = status
	this.CollectedAt = collectedAt
	return &this
}

// NewKingbaseDiagnosticSourceWithDefaults instantiates a new KingbaseDiagnosticSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKingbaseDiagnosticSourceWithDefaults() *KingbaseDiagnosticSource {
	this := KingbaseDiagnosticSource{}
	return &this
}

// GetName returns the Name field value.
func (o *KingbaseDiagnosticSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KingbaseDiagnosticSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *KingbaseDiagnosticSource) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *KingbaseDiagnosticSource) GetStatus() KingbaseDiagnosticSourceStatus {
	if o == nil {
		var ret KingbaseDiagnosticSourceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *KingbaseDiagnosticSource) GetStatusOk() (*KingbaseDiagnosticSourceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *KingbaseDiagnosticSource) SetStatus(v KingbaseDiagnosticSourceStatus) {
	o.Status = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *KingbaseDiagnosticSource) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *KingbaseDiagnosticSource) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *KingbaseDiagnosticSource) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KingbaseDiagnosticSource) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseDiagnosticSource) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KingbaseDiagnosticSource) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KingbaseDiagnosticSource) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KingbaseDiagnosticSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
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
func (o *KingbaseDiagnosticSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                         `json:"name"`
		Status      *KingbaseDiagnosticSourceStatus `json:"status"`
		CollectedAt *string                         `json:"collectedAt"`
		Message     *string                         `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "status", "collectedAt", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.CollectedAt = *all.CollectedAt
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
