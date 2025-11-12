// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseItem Cluster database information
type DatabaseItem struct {
	// Specify the name of database, which must be unique.
	Name string `json:"name"`
	// Specify the options of database.
	Options map[string]interface{} `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseItem instantiates a new DatabaseItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseItem(name string) *DatabaseItem {
	this := DatabaseItem{}
	this.Name = name
	return &this
}

// NewDatabaseItemWithDefaults instantiates a new DatabaseItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseItemWithDefaults() *DatabaseItem {
	this := DatabaseItem{}
	return &this
}

// GetName returns the Name field value.
func (o *DatabaseItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatabaseItem) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DatabaseItem) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseItem) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DatabaseItem) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *DatabaseItem) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabaseItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string                `json:"name"`
		Options map[string]interface{} `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "options"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
