// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RestoreLog restore workload logs

type RestoreLog struct {
	// items is the list of restoreLogByPod objects
	Items []RestoreLogByPod `json:"items,omitempty"`
	// restore id
	RestoreId *string `json:"restoreId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreLog instantiates a new RestoreLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreLog() *RestoreLog {
	this := RestoreLog{}
	return &this
}

// NewRestoreLogWithDefaults instantiates a new RestoreLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreLogWithDefaults() *RestoreLog {
	this := RestoreLog{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RestoreLog) GetItems() []RestoreLogByPod {
	if o == nil || o.Items == nil {
		var ret []RestoreLogByPod
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreLog) GetItemsOk() (*[]RestoreLogByPod, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RestoreLog) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []RestoreLogByPod and assigns it to the Items field.
func (o *RestoreLog) SetItems(v []RestoreLogByPod) {
	o.Items = v
}

// GetRestoreId returns the RestoreId field value if set, zero value otherwise.
func (o *RestoreLog) GetRestoreId() string {
	if o == nil || o.RestoreId == nil {
		var ret string
		return ret
	}
	return *o.RestoreId
}

// GetRestoreIdOk returns a tuple with the RestoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreLog) GetRestoreIdOk() (*string, bool) {
	if o == nil || o.RestoreId == nil {
		return nil, false
	}
	return o.RestoreId, true
}

// HasRestoreId returns a boolean if a field has been set.
func (o *RestoreLog) HasRestoreId() bool {
	return o != nil && o.RestoreId != nil
}

// SetRestoreId gets a reference to the given string and assigns it to the RestoreId field.
func (o *RestoreLog) SetRestoreId(v string) {
	o.RestoreId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.RestoreId != nil {
		toSerialize["restoreId"] = o.RestoreId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items     []RestoreLogByPod `json:"items,omitempty"`
		RestoreId *string           `json:"restoreId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "restoreId"})
	} else {
		return err
	}
	o.Items = all.Items
	o.RestoreId = all.RestoreId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
