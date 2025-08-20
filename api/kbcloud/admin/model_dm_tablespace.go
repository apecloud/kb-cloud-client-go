// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

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
	// the total size already allocated of this tablespace, unit MB
	AllocatedSizeMb *string `json:"allocatedSizeMB,omitempty"`
	// the current used size of this tablespace, unit MB
	UsedSizeMb *string `json:"usedSizeMB,omitempty"`
	// the current used ratio of this tablespace
	UsedRatio *string `json:"usedRatio,omitempty"`
	// whether this tablespace auto extend
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// the max size can be allocated of this tablespace, unit GB
	MaxSizeGb *string `json:"maxSizeGB,omitempty"`
	// the used ratio in the max allocated size of this tablespace
	UsedRatioInMax *string `json:"usedRatioInMax,omitempty"`
	// the status of this tablespace
	Status *TableSpaceStatus `json:"status,omitempty"`
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

// GetAllocatedSizeMb returns the AllocatedSizeMb field value if set, zero value otherwise.
func (o *DmTablespace) GetAllocatedSizeMb() string {
	if o == nil || o.AllocatedSizeMb == nil {
		var ret string
		return ret
	}
	return *o.AllocatedSizeMb
}

// GetAllocatedSizeMbOk returns a tuple with the AllocatedSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetAllocatedSizeMbOk() (*string, bool) {
	if o == nil || o.AllocatedSizeMb == nil {
		return nil, false
	}
	return o.AllocatedSizeMb, true
}

// HasAllocatedSizeMb returns a boolean if a field has been set.
func (o *DmTablespace) HasAllocatedSizeMb() bool {
	return o != nil && o.AllocatedSizeMb != nil
}

// SetAllocatedSizeMb gets a reference to the given string and assigns it to the AllocatedSizeMb field.
func (o *DmTablespace) SetAllocatedSizeMb(v string) {
	o.AllocatedSizeMb = &v
}

// GetUsedSizeMb returns the UsedSizeMb field value if set, zero value otherwise.
func (o *DmTablespace) GetUsedSizeMb() string {
	if o == nil || o.UsedSizeMb == nil {
		var ret string
		return ret
	}
	return *o.UsedSizeMb
}

// GetUsedSizeMbOk returns a tuple with the UsedSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetUsedSizeMbOk() (*string, bool) {
	if o == nil || o.UsedSizeMb == nil {
		return nil, false
	}
	return o.UsedSizeMb, true
}

// HasUsedSizeMb returns a boolean if a field has been set.
func (o *DmTablespace) HasUsedSizeMb() bool {
	return o != nil && o.UsedSizeMb != nil
}

// SetUsedSizeMb gets a reference to the given string and assigns it to the UsedSizeMb field.
func (o *DmTablespace) SetUsedSizeMb(v string) {
	o.UsedSizeMb = &v
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

// GetMaxSizeGb returns the MaxSizeGb field value if set, zero value otherwise.
func (o *DmTablespace) GetMaxSizeGb() string {
	if o == nil || o.MaxSizeGb == nil {
		var ret string
		return ret
	}
	return *o.MaxSizeGb
}

// GetMaxSizeGbOk returns a tuple with the MaxSizeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetMaxSizeGbOk() (*string, bool) {
	if o == nil || o.MaxSizeGb == nil {
		return nil, false
	}
	return o.MaxSizeGb, true
}

// HasMaxSizeGb returns a boolean if a field has been set.
func (o *DmTablespace) HasMaxSizeGb() bool {
	return o != nil && o.MaxSizeGb != nil
}

// SetMaxSizeGb gets a reference to the given string and assigns it to the MaxSizeGb field.
func (o *DmTablespace) SetMaxSizeGb(v string) {
	o.MaxSizeGb = &v
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
func (o *DmTablespace) GetStatus() TableSpaceStatus {
	if o == nil || o.Status == nil {
		var ret TableSpaceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmTablespace) GetStatusOk() (*TableSpaceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DmTablespace) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given TableSpaceStatus and assigns it to the Status field.
func (o *DmTablespace) SetStatus(v TableSpaceStatus) {
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
	if o.AllocatedSizeMb != nil {
		toSerialize["allocatedSizeMB"] = o.AllocatedSizeMb
	}
	if o.UsedSizeMb != nil {
		toSerialize["usedSizeMB"] = o.UsedSizeMb
	}
	if o.UsedRatio != nil {
		toSerialize["usedRatio"] = o.UsedRatio
	}
	if o.AutoExtend != nil {
		toSerialize["autoExtend"] = o.AutoExtend
	}
	if o.MaxSizeGb != nil {
		toSerialize["maxSizeGB"] = o.MaxSizeGb
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
		Name            *string           `json:"name"`
		Users           *string           `json:"users,omitempty"`
		Files           *[]DmFile         `json:"files"`
		AllocatedSizeMb *string           `json:"allocatedSizeMB,omitempty"`
		UsedSizeMb      *string           `json:"usedSizeMB,omitempty"`
		UsedRatio       *string           `json:"usedRatio,omitempty"`
		AutoExtend      *bool             `json:"autoExtend,omitempty"`
		MaxSizeGb       *string           `json:"maxSizeGB,omitempty"`
		UsedRatioInMax  *string           `json:"usedRatioInMax,omitempty"`
		Status          *TableSpaceStatus `json:"status,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "users", "files", "allocatedSizeMB", "usedSizeMB", "usedRatio", "autoExtend", "maxSizeGB", "usedRatioInMax", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Users = all.Users
	o.Files = *all.Files
	o.AllocatedSizeMb = all.AllocatedSizeMb
	o.UsedSizeMb = all.UsedSizeMb
	o.UsedRatio = all.UsedRatio
	o.AutoExtend = all.AutoExtend
	o.MaxSizeGb = all.MaxSizeGb
	o.UsedRatioInMax = all.UsedRatioInMax
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
