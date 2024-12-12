// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

type BackupRepoView struct {
	// the router to show in backup repo
	Filepath *string `json:"filepath,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoView instantiates a new BackupRepoView object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoView() *BackupRepoView {
	this := BackupRepoView{}
	return &this
}

// NewBackupRepoViewWithDefaults instantiates a new BackupRepoView object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoViewWithDefaults() *BackupRepoView {
	this := BackupRepoView{}
	return &this
}

// GetFilepath returns the Filepath field value if set, zero value otherwise.
func (o *BackupRepoView) GetFilepath() string {
	if o == nil || o.Filepath == nil {
		var ret string
		return ret
	}
	return *o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoView) GetFilepathOk() (*string, bool) {
	if o == nil || o.Filepath == nil {
		return nil, false
	}
	return o.Filepath, true
}

// HasFilepath returns a boolean if a field has been set.
func (o *BackupRepoView) HasFilepath() bool {
	return o != nil && o.Filepath != nil
}

// SetFilepath gets a reference to the given string and assigns it to the Filepath field.
func (o *BackupRepoView) SetFilepath(v string) {
	o.Filepath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filepath != nil {
		toSerialize["filepath"] = o.Filepath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoView) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filepath *string `json:"filepath,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filepath"})
	} else {
		return err
	}
	o.Filepath = all.Filepath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
