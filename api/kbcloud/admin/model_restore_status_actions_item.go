// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RestoreStatusActionsItem struct {
	Message *string `json:"message,omitempty"`
	// action name
	Name *string `json:"name,omitempty"`
	// action status, enum values: [Processing, Completed, Failed]
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreStatusActionsItem instantiates a new RestoreStatusActionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreStatusActionsItem() *RestoreStatusActionsItem {
	this := RestoreStatusActionsItem{}
	return &this
}

// NewRestoreStatusActionsItemWithDefaults instantiates a new RestoreStatusActionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreStatusActionsItemWithDefaults() *RestoreStatusActionsItem {
	this := RestoreStatusActionsItem{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RestoreStatusActionsItem) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestoreStatusActionsItem) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RestoreStatusActionsItem) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreStatusActionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreStatusActionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string `json:"message,omitempty"`
		Name    *string `json:"name,omitempty"`
		Status  *string `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "name", "status"})
	} else {
		return err
	}
	o.Message = all.Message
	o.Name = all.Name
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
