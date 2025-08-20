// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmTablespace struct {
	// tablespace name
	Name string `json:"name"`
	// damengdb users belong to this tablespace
	Users *string `json:"users,omitempty"`
	// the physical files of this tablespace
	Files []DmFile `json:"files"`
	// the total size of this tablespace
	TotalSize *string `json:"totalSize,omitempty"`
	// the used size of this tablespace
	UsedSize *string `json:"usedSize,omitempty"`
	// the used ratio of this tablespace
	UsedRatio *string `json:"usedRatio,omitempty"`
	// whether this tablespace auto extend
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// the max size of this tablespace
	MaxSize *string `json:"maxSize,omitempty"`
	// the used ratio in max size of this tablespace
	UsedRatioInMax *string `json:"usedRatioInMax,omitempty"`
	// the status of this tablespace
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmTablespace instantiates a new DmTablespace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmTablespace(name string, files []DmFile) *DmTablespace {
	this := DmTablespace{}
	this.Name = name
	this.Files = files
	return &this
}

// NewDmTablespaceWithDefaults instantiates a new DmTablespace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmTablespaceWithDefaults() *DmTablespace {
	this := DmTablespace{}
	return &this
}

// GetName returns the Name field value.
func (o *DmTablespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DmTablespace) SetName(v string) {
	o.Name = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *DmTablespace) GetUsers() string {
	if o == nil || o.Users == nil {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetUsersOk() (*string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *DmTablespace) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *DmTablespace) SetUsers(v string) {
	o.Users = &v
}

// GetFiles returns the Files field value.
func (o *DmTablespace) GetFiles() []DmFile {
	if o == nil {
		var ret []DmFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetFilesOk() (*[]DmFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *DmTablespace) SetFiles(v []DmFile) {
	o.Files = v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *DmTablespace) GetTotalSize() string {
	if o == nil || o.TotalSize == nil {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetTotalSizeOk() (*string, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *DmTablespace) HasTotalSize() bool {
	return o != nil && o.TotalSize != nil
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *DmTablespace) SetTotalSize(v string) {
	o.TotalSize = &v
}

// GetUsedSize returns the UsedSize field value if set, zero value otherwise.
func (o *DmTablespace) GetUsedSize() string {
	if o == nil || o.UsedSize == nil {
		var ret string
		return ret
	}
	return *o.UsedSize
}

// GetUsedSizeOk returns a tuple with the UsedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetUsedSizeOk() (*string, bool) {
	if o == nil || o.UsedSize == nil {
		return nil, false
	}
	return o.UsedSize, true
}

// HasUsedSize returns a boolean if a field has been set.
func (o *DmTablespace) HasUsedSize() bool {
	return o != nil && o.UsedSize != nil
}

// SetUsedSize gets a reference to the given string and assigns it to the UsedSize field.
func (o *DmTablespace) SetUsedSize(v string) {
	o.UsedSize = &v
}

// GetUsedRatio returns the UsedRatio field value if set, zero value otherwise.
func (o *DmTablespace) GetUsedRatio() string {
	if o == nil || o.UsedRatio == nil {
		var ret string
		return ret
	}
	return *o.UsedRatio
}

// GetUsedRatioOk returns a tuple with the UsedRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetUsedRatioOk() (*string, bool) {
	if o == nil || o.UsedRatio == nil {
		return nil, false
	}
	return o.UsedRatio, true
}

// HasUsedRatio returns a boolean if a field has been set.
func (o *DmTablespace) HasUsedRatio() bool {
	return o != nil && o.UsedRatio != nil
}

// SetUsedRatio gets a reference to the given string and assigns it to the UsedRatio field.
func (o *DmTablespace) SetUsedRatio(v string) {
	o.UsedRatio = &v
}

// GetAutoExtend returns the AutoExtend field value if set, zero value otherwise.
func (o *DmTablespace) GetAutoExtend() bool {
	if o == nil || o.AutoExtend == nil {
		var ret bool
		return ret
	}
	return *o.AutoExtend
}

// GetAutoExtendOk returns a tuple with the AutoExtend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetAutoExtendOk() (*bool, bool) {
	if o == nil || o.AutoExtend == nil {
		return nil, false
	}
	return o.AutoExtend, true
}

// HasAutoExtend returns a boolean if a field has been set.
func (o *DmTablespace) HasAutoExtend() bool {
	return o != nil && o.AutoExtend != nil
}

// SetAutoExtend gets a reference to the given bool and assigns it to the AutoExtend field.
func (o *DmTablespace) SetAutoExtend(v bool) {
	o.AutoExtend = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *DmTablespace) GetMaxSize() string {
	if o == nil || o.MaxSize == nil {
		var ret string
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetMaxSizeOk() (*string, bool) {
	if o == nil || o.MaxSize == nil {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *DmTablespace) HasMaxSize() bool {
	return o != nil && o.MaxSize != nil
}

// SetMaxSize gets a reference to the given string and assigns it to the MaxSize field.
func (o *DmTablespace) SetMaxSize(v string) {
	o.MaxSize = &v
}

// GetUsedRatioInMax returns the UsedRatioInMax field value if set, zero value otherwise.
func (o *DmTablespace) GetUsedRatioInMax() string {
	if o == nil || o.UsedRatioInMax == nil {
		var ret string
		return ret
	}
	return *o.UsedRatioInMax
}

// GetUsedRatioInMaxOk returns a tuple with the UsedRatioInMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetUsedRatioInMaxOk() (*string, bool) {
	if o == nil || o.UsedRatioInMax == nil {
		return nil, false
	}
	return o.UsedRatioInMax, true
}

// HasUsedRatioInMax returns a boolean if a field has been set.
func (o *DmTablespace) HasUsedRatioInMax() bool {
	return o != nil && o.UsedRatioInMax != nil
}

// SetUsedRatioInMax gets a reference to the given string and assigns it to the UsedRatioInMax field.
func (o *DmTablespace) SetUsedRatioInMax(v string) {
	o.UsedRatioInMax = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DmTablespace) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DmTablespace) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DmTablespace) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmTablespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	toSerialize["files"] = o.Files
	if o.TotalSize != nil {
		toSerialize["totalSize"] = o.TotalSize
	}
	if o.UsedSize != nil {
		toSerialize["usedSize"] = o.UsedSize
	}
	if o.UsedRatio != nil {
		toSerialize["usedRatio"] = o.UsedRatio
	}
	if o.AutoExtend != nil {
		toSerialize["autoExtend"] = o.AutoExtend
	}
	if o.MaxSize != nil {
		toSerialize["maxSize"] = o.MaxSize
	}
	if o.UsedRatioInMax != nil {
		toSerialize["usedRatioInMax"] = o.UsedRatioInMax
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
func (o *DmTablespace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string   `json:"name"`
		Users          *string   `json:"users,omitempty"`
		Files          *[]DmFile `json:"files"`
		TotalSize      *string   `json:"totalSize,omitempty"`
		UsedSize       *string   `json:"usedSize,omitempty"`
		UsedRatio      *string   `json:"usedRatio,omitempty"`
		AutoExtend     *bool     `json:"autoExtend,omitempty"`
		MaxSize        *string   `json:"maxSize,omitempty"`
		UsedRatioInMax *string   `json:"usedRatioInMax,omitempty"`
		Status         *string   `json:"status,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "users", "files", "totalSize", "usedSize", "usedRatio", "autoExtend", "maxSize", "usedRatioInMax", "status"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Users = all.Users
	o.Files = *all.Files
	o.TotalSize = all.TotalSize
	o.UsedSize = all.UsedSize
	o.UsedRatio = all.UsedRatio
	o.AutoExtend = all.AutoExtend
	o.MaxSize = all.MaxSize
	o.UsedRatioInMax = all.UsedRatioInMax
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
