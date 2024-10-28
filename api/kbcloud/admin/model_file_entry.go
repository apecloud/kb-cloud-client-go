// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// FileEntry the entry of files

type FileEntry struct {
	// description the entry type
	IsDir *bool `json:"IsDir,omitempty"`
	// the full path of file
	FullPath *string `json:"fullPath,omitempty"`
	// the filename
	Filename *string `json:"filename,omitempty"`
	// the size of entry
	Size *int64 `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFileEntry instantiates a new FileEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFileEntry() *FileEntry {
	this := FileEntry{}
	return &this
}

// NewFileEntryWithDefaults instantiates a new FileEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFileEntryWithDefaults() *FileEntry {
	this := FileEntry{}
	return &this
}

// GetIsDir returns the IsDir field value if set, zero value otherwise.
func (o *FileEntry) GetIsDir() bool {
	if o == nil || o.IsDir == nil {
		var ret bool
		return ret
	}
	return *o.IsDir
}

// GetIsDirOk returns a tuple with the IsDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetIsDirOk() (*bool, bool) {
	if o == nil || o.IsDir == nil {
		return nil, false
	}
	return o.IsDir, true
}

// HasIsDir returns a boolean if a field has been set.
func (o *FileEntry) HasIsDir() bool {
	return o != nil && o.IsDir != nil
}

// SetIsDir gets a reference to the given bool and assigns it to the IsDir field.
func (o *FileEntry) SetIsDir(v bool) {
	o.IsDir = &v
}

// GetFullPath returns the FullPath field value if set, zero value otherwise.
func (o *FileEntry) GetFullPath() string {
	if o == nil || o.FullPath == nil {
		var ret string
		return ret
	}
	return *o.FullPath
}

// GetFullPathOk returns a tuple with the FullPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetFullPathOk() (*string, bool) {
	if o == nil || o.FullPath == nil {
		return nil, false
	}
	return o.FullPath, true
}

// HasFullPath returns a boolean if a field has been set.
func (o *FileEntry) HasFullPath() bool {
	return o != nil && o.FullPath != nil
}

// SetFullPath gets a reference to the given string and assigns it to the FullPath field.
func (o *FileEntry) SetFullPath(v string) {
	o.FullPath = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *FileEntry) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *FileEntry) HasFilename() bool {
	return o != nil && o.Filename != nil
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *FileEntry) SetFilename(v string) {
	o.Filename = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileEntry) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntry) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileEntry) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *FileEntry) SetSize(v int64) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FileEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.IsDir != nil {
		toSerialize["IsDir"] = o.IsDir
	}
	if o.FullPath != nil {
		toSerialize["fullPath"] = o.FullPath
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FileEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsDir    *bool   `json:"IsDir,omitempty"`
		FullPath *string `json:"fullPath,omitempty"`
		Filename *string `json:"filename,omitempty"`
		Size     *int64  `json:"size,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"IsDir", "fullPath", "filename", "size"})
	} else {
		return err
	}
	o.IsDir = all.IsDir
	o.FullPath = all.FullPath
	o.Filename = all.Filename
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
