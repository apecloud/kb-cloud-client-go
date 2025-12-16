// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type BackupOptionOfflineBackupOption struct {
	// whether to use the parent directory of the backup path as the actual backup path
	UseParentDirAsBackupPath *bool `json:"useParentDirAsBackupPath,omitempty"`
	// whether offline backup is supported
	Supported *bool `json:"supported,omitempty"`
	// the root user for offline backup
	RootUser *string `json:"rootUser,omitempty"`
	// the backup method for offline backup
	BackupMethod *string `json:"backupMethod,omitempty"`
	// the component name for offline backup
	ComponentName *string                `json:"componentName,omitempty"`
	Status        map[string]interface{} `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupOptionOfflineBackupOption instantiates a new BackupOptionOfflineBackupOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupOptionOfflineBackupOption() *BackupOptionOfflineBackupOption {
	this := BackupOptionOfflineBackupOption{}
	return &this
}

// NewBackupOptionOfflineBackupOptionWithDefaults instantiates a new BackupOptionOfflineBackupOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupOptionOfflineBackupOptionWithDefaults() *BackupOptionOfflineBackupOption {
	this := BackupOptionOfflineBackupOption{}
	return &this
}

// GetUseParentDirAsBackupPath returns the UseParentDirAsBackupPath field value if set, zero value otherwise.
func (o *BackupOptionOfflineBackupOption) GetUseParentDirAsBackupPath() bool {
	if o == nil || o.UseParentDirAsBackupPath == nil {
		var ret bool
		return ret
	}
	return *o.UseParentDirAsBackupPath
}

// GetUseParentDirAsBackupPathOk returns a tuple with the UseParentDirAsBackupPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionOfflineBackupOption) GetUseParentDirAsBackupPathOk() (*bool, bool) {
	if o == nil || o.UseParentDirAsBackupPath == nil {
		return nil, false
	}
	return o.UseParentDirAsBackupPath, true
}

// HasUseParentDirAsBackupPath returns a boolean if a field has been set.
func (o *BackupOptionOfflineBackupOption) HasUseParentDirAsBackupPath() bool {
	return o != nil && o.UseParentDirAsBackupPath != nil
}

// SetUseParentDirAsBackupPath gets a reference to the given bool and assigns it to the UseParentDirAsBackupPath field.
func (o *BackupOptionOfflineBackupOption) SetUseParentDirAsBackupPath(v bool) {
	o.UseParentDirAsBackupPath = &v
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *BackupOptionOfflineBackupOption) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionOfflineBackupOption) GetSupportedOk() (*bool, bool) {
	if o == nil || o.Supported == nil {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *BackupOptionOfflineBackupOption) HasSupported() bool {
	return o != nil && o.Supported != nil
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *BackupOptionOfflineBackupOption) SetSupported(v bool) {
	o.Supported = &v
}

// GetRootUser returns the RootUser field value if set, zero value otherwise.
func (o *BackupOptionOfflineBackupOption) GetRootUser() string {
	if o == nil || o.RootUser == nil {
		var ret string
		return ret
	}
	return *o.RootUser
}

// GetRootUserOk returns a tuple with the RootUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionOfflineBackupOption) GetRootUserOk() (*string, bool) {
	if o == nil || o.RootUser == nil {
		return nil, false
	}
	return o.RootUser, true
}

// HasRootUser returns a boolean if a field has been set.
func (o *BackupOptionOfflineBackupOption) HasRootUser() bool {
	return o != nil && o.RootUser != nil
}

// SetRootUser gets a reference to the given string and assigns it to the RootUser field.
func (o *BackupOptionOfflineBackupOption) SetRootUser(v string) {
	o.RootUser = &v
}

// GetBackupMethod returns the BackupMethod field value if set, zero value otherwise.
func (o *BackupOptionOfflineBackupOption) GetBackupMethod() string {
	if o == nil || o.BackupMethod == nil {
		var ret string
		return ret
	}
	return *o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionOfflineBackupOption) GetBackupMethodOk() (*string, bool) {
	if o == nil || o.BackupMethod == nil {
		return nil, false
	}
	return o.BackupMethod, true
}

// HasBackupMethod returns a boolean if a field has been set.
func (o *BackupOptionOfflineBackupOption) HasBackupMethod() bool {
	return o != nil && o.BackupMethod != nil
}

// SetBackupMethod gets a reference to the given string and assigns it to the BackupMethod field.
func (o *BackupOptionOfflineBackupOption) SetBackupMethod(v string) {
	o.BackupMethod = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *BackupOptionOfflineBackupOption) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionOfflineBackupOption) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *BackupOptionOfflineBackupOption) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *BackupOptionOfflineBackupOption) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupOptionOfflineBackupOption) GetStatus() map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionOfflineBackupOption) GetStatusOk() (*map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupOptionOfflineBackupOption) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *BackupOptionOfflineBackupOption) SetStatus(v map[string]interface{}) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupOptionOfflineBackupOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.UseParentDirAsBackupPath != nil {
		toSerialize["useParentDirAsBackupPath"] = o.UseParentDirAsBackupPath
	}
	if o.Supported != nil {
		toSerialize["supported"] = o.Supported
	}
	if o.RootUser != nil {
		toSerialize["rootUser"] = o.RootUser
	}
	if o.BackupMethod != nil {
		toSerialize["backupMethod"] = o.BackupMethod
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
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
func (o *BackupOptionOfflineBackupOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UseParentDirAsBackupPath *bool                  `json:"useParentDirAsBackupPath,omitempty"`
		Supported                *bool                  `json:"supported,omitempty"`
		RootUser                 *string                `json:"rootUser,omitempty"`
		BackupMethod             *string                `json:"backupMethod,omitempty"`
		ComponentName            *string                `json:"componentName,omitempty"`
		Status                   map[string]interface{} `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"useParentDirAsBackupPath", "supported", "rootUser", "backupMethod", "componentName", "status"})
	} else {
		return err
	}
	o.UseParentDirAsBackupPath = all.UseParentDirAsBackupPath
	o.Supported = all.Supported
	o.RootUser = all.RootUser
	o.BackupMethod = all.BackupMethod
	o.ComponentName = all.ComponentName
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
