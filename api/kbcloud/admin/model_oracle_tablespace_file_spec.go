// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OracleTablespaceFileSpec struct {
	// Absolute database-server file path or ASM destination. Omit to use Oracle Managed Files (DB_CREATE_FILE_DEST must be configured). Existing files are never reused.
	Name       *string `json:"name,omitempty"`
	SizeMb     int64   `json:"sizeMB"`
	AutoExtend *bool   `json:"autoExtend,omitempty"`
	// Required positive increment when autoExtend is true; otherwise omit.
	NextSizeMb *int64 `json:"nextSizeMB,omitempty"`
	// Maximum size when autoExtend is true; zero or omitted means UNLIMITED. Otherwise omit.
	MaxSizeMb *int64 `json:"maxSizeMB,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOracleTablespaceFileSpec instantiates a new OracleTablespaceFileSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOracleTablespaceFileSpec(sizeMb int64) *OracleTablespaceFileSpec {
	this := OracleTablespaceFileSpec{}
	this.SizeMb = sizeMb
	var autoExtend bool = false
	this.AutoExtend = &autoExtend
	return &this
}

// NewOracleTablespaceFileSpecWithDefaults instantiates a new OracleTablespaceFileSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOracleTablespaceFileSpecWithDefaults() *OracleTablespaceFileSpec {
	this := OracleTablespaceFileSpec{}
	var autoExtend bool = false
	this.AutoExtend = &autoExtend
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OracleTablespaceFileSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFileSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OracleTablespaceFileSpec) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OracleTablespaceFileSpec) SetName(v string) {
	o.Name = &v
}

// GetSizeMb returns the SizeMb field value.
func (o *OracleTablespaceFileSpec) GetSizeMb() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFileSpec) GetSizeMbOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeMb, true
}

// SetSizeMb sets field value.
func (o *OracleTablespaceFileSpec) SetSizeMb(v int64) {
	o.SizeMb = v
}

// GetAutoExtend returns the AutoExtend field value if set, zero value otherwise.
func (o *OracleTablespaceFileSpec) GetAutoExtend() bool {
	if o == nil || o.AutoExtend == nil {
		var ret bool
		return ret
	}
	return *o.AutoExtend
}

// GetAutoExtendOk returns a tuple with the AutoExtend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFileSpec) GetAutoExtendOk() (*bool, bool) {
	if o == nil || o.AutoExtend == nil {
		return nil, false
	}
	return o.AutoExtend, true
}

// HasAutoExtend returns a boolean if a field has been set.
func (o *OracleTablespaceFileSpec) HasAutoExtend() bool {
	return o != nil && o.AutoExtend != nil
}

// SetAutoExtend gets a reference to the given bool and assigns it to the AutoExtend field.
func (o *OracleTablespaceFileSpec) SetAutoExtend(v bool) {
	o.AutoExtend = &v
}

// GetNextSizeMb returns the NextSizeMb field value if set, zero value otherwise.
func (o *OracleTablespaceFileSpec) GetNextSizeMb() int64 {
	if o == nil || o.NextSizeMb == nil {
		var ret int64
		return ret
	}
	return *o.NextSizeMb
}

// GetNextSizeMbOk returns a tuple with the NextSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFileSpec) GetNextSizeMbOk() (*int64, bool) {
	if o == nil || o.NextSizeMb == nil {
		return nil, false
	}
	return o.NextSizeMb, true
}

// HasNextSizeMb returns a boolean if a field has been set.
func (o *OracleTablespaceFileSpec) HasNextSizeMb() bool {
	return o != nil && o.NextSizeMb != nil
}

// SetNextSizeMb gets a reference to the given int64 and assigns it to the NextSizeMb field.
func (o *OracleTablespaceFileSpec) SetNextSizeMb(v int64) {
	o.NextSizeMb = &v
}

// GetMaxSizeMb returns the MaxSizeMb field value if set, zero value otherwise.
func (o *OracleTablespaceFileSpec) GetMaxSizeMb() int64 {
	if o == nil || o.MaxSizeMb == nil {
		var ret int64
		return ret
	}
	return *o.MaxSizeMb
}

// GetMaxSizeMbOk returns a tuple with the MaxSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFileSpec) GetMaxSizeMbOk() (*int64, bool) {
	if o == nil || o.MaxSizeMb == nil {
		return nil, false
	}
	return o.MaxSizeMb, true
}

// HasMaxSizeMb returns a boolean if a field has been set.
func (o *OracleTablespaceFileSpec) HasMaxSizeMb() bool {
	return o != nil && o.MaxSizeMb != nil
}

// SetMaxSizeMb gets a reference to the given int64 and assigns it to the MaxSizeMb field.
func (o *OracleTablespaceFileSpec) SetMaxSizeMb(v int64) {
	o.MaxSizeMb = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OracleTablespaceFileSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["sizeMB"] = o.SizeMb
	if o.AutoExtend != nil {
		toSerialize["autoExtend"] = o.AutoExtend
	}
	if o.NextSizeMb != nil {
		toSerialize["nextSizeMB"] = o.NextSizeMb
	}
	if o.MaxSizeMb != nil {
		toSerialize["maxSizeMB"] = o.MaxSizeMb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OracleTablespaceFileSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name,omitempty"`
		SizeMb     *int64  `json:"sizeMB"`
		AutoExtend *bool   `json:"autoExtend,omitempty"`
		NextSizeMb *int64  `json:"nextSizeMB,omitempty"`
		MaxSizeMb  *int64  `json:"maxSizeMB,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SizeMb == nil {
		return fmt.Errorf("required field sizeMB missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "sizeMB", "autoExtend", "nextSizeMB", "maxSizeMB"})
	} else {
		return err
	}
	o.Name = all.Name
	o.SizeMb = *all.SizeMb
	o.AutoExtend = all.AutoExtend
	o.NextSizeMb = all.NextSizeMb
	o.MaxSizeMb = all.MaxSizeMb

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
