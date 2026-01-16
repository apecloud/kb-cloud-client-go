// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RefItem a reference of the vulnerability
type RefItem struct {
	// the type of the reference
	RType string `json:"rType"`
	// the URL of the reference
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRefItem instantiates a new RefItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRefItem(rType string, url string) *RefItem {
	this := RefItem{}
	this.RType = rType
	this.Url = url
	return &this
}

// NewRefItemWithDefaults instantiates a new RefItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRefItemWithDefaults() *RefItem {
	this := RefItem{}
	return &this
}

// GetRType returns the RType field value.
func (o *RefItem) GetRType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RType
}

// GetRTypeOk returns a tuple with the RType field value
// and a boolean to check if the value has been set.
func (o *RefItem) GetRTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RType, true
}

// SetRType sets field value.
func (o *RefItem) SetRType(v string) {
	o.RType = v
}

// GetUrl returns the Url field value.
func (o *RefItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RefItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *RefItem) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RefItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["rType"] = o.RType
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RefItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RType *string `json:"rType"`
		Url   *string `json:"url"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RType == nil {
		return fmt.Errorf("required field rType missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"rType", "url"})
	} else {
		return err
	}
	o.RType = *all.RType
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
