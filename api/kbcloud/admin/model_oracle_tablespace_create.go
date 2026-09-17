// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OracleTablespaceCreate struct {
	// Converted to uppercase. Limited to 30 ASCII bytes for compatibility with Oracle 12.1.
	Name string `json:"name"`
	// New tablespaces default to PERMANENT when omitted.
	Contents *OracleTablespaceContents `json:"contents,omitempty"`
	Bigfile  *bool                     `json:"bigfile,omitempty"`
	// Exactly one file for BIGFILE tablespaces.
	Files []OracleTablespaceFileSpec `json:"files"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOracleTablespaceCreate instantiates a new OracleTablespaceCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOracleTablespaceCreate(name string, files []OracleTablespaceFileSpec) *OracleTablespaceCreate {
	this := OracleTablespaceCreate{}
	this.Name = name
	var bigfile bool = false
	this.Bigfile = &bigfile
	this.Files = files
	return &this
}

// NewOracleTablespaceCreateWithDefaults instantiates a new OracleTablespaceCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOracleTablespaceCreateWithDefaults() *OracleTablespaceCreate {
	this := OracleTablespaceCreate{}
	var bigfile bool = false
	this.Bigfile = &bigfile
	return &this
}

// GetName returns the Name field value.
func (o *OracleTablespaceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OracleTablespaceCreate) SetName(v string) {
	o.Name = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *OracleTablespaceCreate) GetContents() OracleTablespaceContents {
	if o == nil || o.Contents == nil {
		var ret OracleTablespaceContents
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceCreate) GetContentsOk() (*OracleTablespaceContents, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *OracleTablespaceCreate) HasContents() bool {
	return o != nil && o.Contents != nil
}

// SetContents gets a reference to the given OracleTablespaceContents and assigns it to the Contents field.
func (o *OracleTablespaceCreate) SetContents(v OracleTablespaceContents) {
	o.Contents = &v
}

// GetBigfile returns the Bigfile field value if set, zero value otherwise.
func (o *OracleTablespaceCreate) GetBigfile() bool {
	if o == nil || o.Bigfile == nil {
		var ret bool
		return ret
	}
	return *o.Bigfile
}

// GetBigfileOk returns a tuple with the Bigfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceCreate) GetBigfileOk() (*bool, bool) {
	if o == nil || o.Bigfile == nil {
		return nil, false
	}
	return o.Bigfile, true
}

// HasBigfile returns a boolean if a field has been set.
func (o *OracleTablespaceCreate) HasBigfile() bool {
	return o != nil && o.Bigfile != nil
}

// SetBigfile gets a reference to the given bool and assigns it to the Bigfile field.
func (o *OracleTablespaceCreate) SetBigfile(v bool) {
	o.Bigfile = &v
}

// GetFiles returns the Files field value.
func (o *OracleTablespaceCreate) GetFiles() []OracleTablespaceFileSpec {
	if o == nil {
		var ret []OracleTablespaceFileSpec
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceCreate) GetFilesOk() (*[]OracleTablespaceFileSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *OracleTablespaceCreate) SetFiles(v []OracleTablespaceFileSpec) {
	o.Files = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OracleTablespaceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	if o.Bigfile != nil {
		toSerialize["bigfile"] = o.Bigfile
	}
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OracleTablespaceCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string                     `json:"name"`
		Contents *OracleTablespaceContents   `json:"contents,omitempty"`
		Bigfile  *bool                       `json:"bigfile,omitempty"`
		Files    *[]OracleTablespaceFileSpec `json:"files"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "contents", "bigfile", "files"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Contents != nil && !all.Contents.IsValid() {
		hasInvalidField = true
	} else {
		o.Contents = all.Contents
	}
	o.Bigfile = all.Bigfile
	o.Files = *all.Files

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
