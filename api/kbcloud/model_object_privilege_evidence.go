// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeEvidence struct {
	Kind   string `json:"kind"`
	Format string `json:"format"`
	// Read-only native evidence; never replay as a privilege mutation.
	Items []string `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeEvidence instantiates a new ObjectPrivilegeEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeEvidence(kind string, format string, items []string) *ObjectPrivilegeEvidence {
	this := ObjectPrivilegeEvidence{}
	this.Kind = kind
	this.Format = format
	this.Items = items
	return &this
}

// NewObjectPrivilegeEvidenceWithDefaults instantiates a new ObjectPrivilegeEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeEvidenceWithDefaults() *ObjectPrivilegeEvidence {
	this := ObjectPrivilegeEvidence{}
	return &this
}

// GetKind returns the Kind field value.
func (o *ObjectPrivilegeEvidence) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeEvidence) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *ObjectPrivilegeEvidence) SetKind(v string) {
	o.Kind = v
}

// GetFormat returns the Format field value.
func (o *ObjectPrivilegeEvidence) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeEvidence) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *ObjectPrivilegeEvidence) SetFormat(v string) {
	o.Format = v
}

// GetItems returns the Items field value.
func (o *ObjectPrivilegeEvidence) GetItems() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeEvidence) GetItemsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ObjectPrivilegeEvidence) SetItems(v []string) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["kind"] = o.Kind
	toSerialize["format"] = o.Format
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind   *string   `json:"kind"`
		Format *string   `json:"format"`
		Items  *[]string `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kind", "format", "items"})
	} else {
		return err
	}
	o.Kind = *all.Kind
	o.Format = *all.Format
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
