// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OracleTablespace Oracle tablespace in the datasource connection's current container. Sizes are in MiB. Supports Oracle 12c and 19c.
type OracleTablespace struct {
	Name string `json:"name"`
	// PERMANENT, TEMPORARY, or UNDO.
	Contents      string `json:"contents"`
	Status        string `json:"status"`
	Bigfile       bool   `json:"bigfile"`
	BlockSize     int64  `json:"blockSize"`
	ContainerName string `json:"containerName"`
	// Whether this is a default permanent or temporary tablespace, including default temporary group members.
	IsDefault       bool    `json:"isDefault"`
	AllocatedSizeMb float64 `json:"allocatedSizeMB"`
	// Allocated file space minus free extents, or temporary file space currently allocated to temporary segments.
	UsedSizeMb float64 `json:"usedSizeMB"`
	// Configured file growth limit; does not imply available disk capacity.
	MaxSizeMb float64                `json:"maxSizeMB"`
	Files     []OracleTablespaceFile `json:"files"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOracleTablespace instantiates a new OracleTablespace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOracleTablespace(name string, contents string, status string, bigfile bool, blockSize int64, containerName string, isDefault bool, allocatedSizeMb float64, usedSizeMb float64, maxSizeMb float64, files []OracleTablespaceFile) *OracleTablespace {
	this := OracleTablespace{}
	this.Name = name
	this.Contents = contents
	this.Status = status
	this.Bigfile = bigfile
	this.BlockSize = blockSize
	this.ContainerName = containerName
	this.IsDefault = isDefault
	this.AllocatedSizeMb = allocatedSizeMb
	this.UsedSizeMb = usedSizeMb
	this.MaxSizeMb = maxSizeMb
	this.Files = files
	return &this
}

// NewOracleTablespaceWithDefaults instantiates a new OracleTablespace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOracleTablespaceWithDefaults() *OracleTablespace {
	this := OracleTablespace{}
	return &this
}

// GetName returns the Name field value.
func (o *OracleTablespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OracleTablespace) SetName(v string) {
	o.Name = v
}

// GetContents returns the Contents field value.
func (o *OracleTablespace) GetContents() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetContentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value.
func (o *OracleTablespace) SetContents(v string) {
	o.Contents = v
}

// GetStatus returns the Status field value.
func (o *OracleTablespace) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OracleTablespace) SetStatus(v string) {
	o.Status = v
}

// GetBigfile returns the Bigfile field value.
func (o *OracleTablespace) GetBigfile() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Bigfile
}

// GetBigfileOk returns a tuple with the Bigfile field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetBigfileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bigfile, true
}

// SetBigfile sets field value.
func (o *OracleTablespace) SetBigfile(v bool) {
	o.Bigfile = v
}

// GetBlockSize returns the BlockSize field value.
func (o *OracleTablespace) GetBlockSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetBlockSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockSize, true
}

// SetBlockSize sets field value.
func (o *OracleTablespace) SetBlockSize(v int64) {
	o.BlockSize = v
}

// GetContainerName returns the ContainerName field value.
func (o *OracleTablespace) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value.
func (o *OracleTablespace) SetContainerName(v string) {
	o.ContainerName = v
}

// GetIsDefault returns the IsDefault field value.
func (o *OracleTablespace) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value.
func (o *OracleTablespace) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetAllocatedSizeMb returns the AllocatedSizeMb field value.
func (o *OracleTablespace) GetAllocatedSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AllocatedSizeMb
}

// GetAllocatedSizeMbOk returns a tuple with the AllocatedSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetAllocatedSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedSizeMb, true
}

// SetAllocatedSizeMb sets field value.
func (o *OracleTablespace) SetAllocatedSizeMb(v float64) {
	o.AllocatedSizeMb = v
}

// GetUsedSizeMb returns the UsedSizeMb field value.
func (o *OracleTablespace) GetUsedSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.UsedSizeMb
}

// GetUsedSizeMbOk returns a tuple with the UsedSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetUsedSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedSizeMb, true
}

// SetUsedSizeMb sets field value.
func (o *OracleTablespace) SetUsedSizeMb(v float64) {
	o.UsedSizeMb = v
}

// GetMaxSizeMb returns the MaxSizeMb field value.
func (o *OracleTablespace) GetMaxSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxSizeMb
}

// GetMaxSizeMbOk returns a tuple with the MaxSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetMaxSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSizeMb, true
}

// SetMaxSizeMb sets field value.
func (o *OracleTablespace) SetMaxSizeMb(v float64) {
	o.MaxSizeMb = v
}

// GetFiles returns the Files field value.
func (o *OracleTablespace) GetFiles() []OracleTablespaceFile {
	if o == nil {
		var ret []OracleTablespaceFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *OracleTablespace) GetFilesOk() (*[]OracleTablespaceFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *OracleTablespace) SetFiles(v []OracleTablespaceFile) {
	o.Files = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OracleTablespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["contents"] = o.Contents
	toSerialize["status"] = o.Status
	toSerialize["bigfile"] = o.Bigfile
	toSerialize["blockSize"] = o.BlockSize
	toSerialize["containerName"] = o.ContainerName
	toSerialize["isDefault"] = o.IsDefault
	toSerialize["allocatedSizeMB"] = o.AllocatedSizeMb
	toSerialize["usedSizeMB"] = o.UsedSizeMb
	toSerialize["maxSizeMB"] = o.MaxSizeMb
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OracleTablespace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string                 `json:"name"`
		Contents        *string                 `json:"contents"`
		Status          *string                 `json:"status"`
		Bigfile         *bool                   `json:"bigfile"`
		BlockSize       *int64                  `json:"blockSize"`
		ContainerName   *string                 `json:"containerName"`
		IsDefault       *bool                   `json:"isDefault"`
		AllocatedSizeMb *float64                `json:"allocatedSizeMB"`
		UsedSizeMb      *float64                `json:"usedSizeMB"`
		MaxSizeMb       *float64                `json:"maxSizeMB"`
		Files           *[]OracleTablespaceFile `json:"files"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Contents == nil {
		return fmt.Errorf("required field contents missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Bigfile == nil {
		return fmt.Errorf("required field bigfile missing")
	}
	if all.BlockSize == nil {
		return fmt.Errorf("required field blockSize missing")
	}
	if all.ContainerName == nil {
		return fmt.Errorf("required field containerName missing")
	}
	if all.IsDefault == nil {
		return fmt.Errorf("required field isDefault missing")
	}
	if all.AllocatedSizeMb == nil {
		return fmt.Errorf("required field allocatedSizeMB missing")
	}
	if all.UsedSizeMb == nil {
		return fmt.Errorf("required field usedSizeMB missing")
	}
	if all.MaxSizeMb == nil {
		return fmt.Errorf("required field maxSizeMB missing")
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "contents", "status", "bigfile", "blockSize", "containerName", "isDefault", "allocatedSizeMB", "usedSizeMB", "maxSizeMB", "files"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Contents = *all.Contents
	o.Status = *all.Status
	o.Bigfile = *all.Bigfile
	o.BlockSize = *all.BlockSize
	o.ContainerName = *all.ContainerName
	o.IsDefault = *all.IsDefault
	o.AllocatedSizeMb = *all.AllocatedSizeMb
	o.UsedSizeMb = *all.UsedSizeMb
	o.MaxSizeMb = *all.MaxSizeMb
	o.Files = *all.Files

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
