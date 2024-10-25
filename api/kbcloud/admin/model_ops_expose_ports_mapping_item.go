// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION OpsExposePortsMappingItem
type OpsExposePortsMappingItem struct {
	// NODESCRIPTION Old
	Old *int32 `json:"old,omitempty"`
	// NODESCRIPTION New
	New *int32 `json:"new,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsExposePortsMappingItem instantiates a new OpsExposePortsMappingItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsExposePortsMappingItem() *OpsExposePortsMappingItem {
	this := OpsExposePortsMappingItem{}
	return &this
}

// NewOpsExposePortsMappingItemWithDefaults instantiates a new OpsExposePortsMappingItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsExposePortsMappingItemWithDefaults() *OpsExposePortsMappingItem {
	this := OpsExposePortsMappingItem{}
	return &this
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *OpsExposePortsMappingItem) GetOld() int32 {
	if o == nil || o.Old == nil {
		var ret int32
		return ret
	}
	return *o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsExposePortsMappingItem) GetOldOk() (*int32, bool) {
	if o == nil || o.Old == nil {
		return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *OpsExposePortsMappingItem) HasOld() bool {
	return o != nil && o.Old != nil
}

// SetOld gets a reference to the given int32 and assigns it to the Old field.
func (o *OpsExposePortsMappingItem) SetOld(v int32) {
	o.Old = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *OpsExposePortsMappingItem) GetNew() int32 {
	if o == nil || o.New == nil {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsExposePortsMappingItem) GetNewOk() (*int32, bool) {
	if o == nil || o.New == nil {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *OpsExposePortsMappingItem) HasNew() bool {
	return o != nil && o.New != nil
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *OpsExposePortsMappingItem) SetNew(v int32) {
	o.New = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsExposePortsMappingItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Old != nil {
		toSerialize["old"] = o.Old
	}
	if o.New != nil {
		toSerialize["new"] = o.New
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsExposePortsMappingItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Old *int32 `json:"old,omitempty"`
		New *int32 `json:"new,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"old", "new"})
	} else {
		return err
	}
	o.Old = all.Old
	o.New = all.New

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
