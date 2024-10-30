// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type ComponentsItemVolumesItem struct {
	// volume name
	Name *string `json:"name,omitempty"`
	// Storage size, the unit is Gi.
	Storage *float64 `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewComponentsItemVolumesItem instantiates a new ComponentsItemVolumesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentsItemVolumesItem() *ComponentsItemVolumesItem {
	this := ComponentsItemVolumesItem{}
	return &this
}

// NewComponentsItemVolumesItemWithDefaults instantiates a new ComponentsItemVolumesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentsItemVolumesItemWithDefaults() *ComponentsItemVolumesItem {
	this := ComponentsItemVolumesItem{}
	return &this
}
// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentsItemVolumesItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItemVolumesItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentsItemVolumesItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentsItemVolumesItem) SetName(v string) {
	o.Name = &v
}


// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ComponentsItemVolumesItem) GetStorage() float64 {
	if o == nil || o.Storage == nil {
		var ret float64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItemVolumesItem) GetStorageOk() (*float64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ComponentsItemVolumesItem) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float64 and assigns it to the Storage field.
func (o *ComponentsItemVolumesItem) SetStorage(v float64) {
	o.Storage = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ComponentsItemVolumesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentsItemVolumesItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name,omitempty"`
		Storage *float64 `json:"storage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "name", "storage",  })
	} else {
		return err
	}
	o.Name = all.Name
	o.Storage = all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
