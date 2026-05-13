// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticExportManifest Manifest for a redacted diagnostic package.
type DiagnosticExportManifest struct {
	// Diagnostic task ID.
	TaskId string `json:"taskId"`
	// Generated package name.
	PackageName string `json:"packageName"`
	// Package generation time.
	GeneratedAt time.Time `json:"generatedAt"`
	// Whether the package was produced from redacted files only.
	Redacted bool                   `json:"redacted"`
	Files    []DiagnosticExportFile `json:"files"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticExportManifest instantiates a new DiagnosticExportManifest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticExportManifest(taskId string, packageName string, generatedAt time.Time, redacted bool, files []DiagnosticExportFile) *DiagnosticExportManifest {
	this := DiagnosticExportManifest{}
	this.TaskId = taskId
	this.PackageName = packageName
	this.GeneratedAt = generatedAt
	this.Redacted = redacted
	this.Files = files
	return &this
}

// NewDiagnosticExportManifestWithDefaults instantiates a new DiagnosticExportManifest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticExportManifestWithDefaults() *DiagnosticExportManifest {
	this := DiagnosticExportManifest{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *DiagnosticExportManifest) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportManifest) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *DiagnosticExportManifest) SetTaskId(v string) {
	o.TaskId = v
}

// GetPackageName returns the PackageName field value.
func (o *DiagnosticExportManifest) GetPackageName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportManifest) GetPackageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageName, true
}

// SetPackageName sets field value.
func (o *DiagnosticExportManifest) SetPackageName(v string) {
	o.PackageName = v
}

// GetGeneratedAt returns the GeneratedAt field value.
func (o *DiagnosticExportManifest) GetGeneratedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.GeneratedAt
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportManifest) GetGeneratedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneratedAt, true
}

// SetGeneratedAt sets field value.
func (o *DiagnosticExportManifest) SetGeneratedAt(v time.Time) {
	o.GeneratedAt = v
}

// GetRedacted returns the Redacted field value.
func (o *DiagnosticExportManifest) GetRedacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redacted
}

// GetRedactedOk returns a tuple with the Redacted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportManifest) GetRedactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redacted, true
}

// SetRedacted sets field value.
func (o *DiagnosticExportManifest) SetRedacted(v bool) {
	o.Redacted = v
}

// GetFiles returns the Files field value.
func (o *DiagnosticExportManifest) GetFiles() []DiagnosticExportFile {
	if o == nil {
		var ret []DiagnosticExportFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportManifest) GetFilesOk() (*[]DiagnosticExportFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *DiagnosticExportManifest) SetFiles(v []DiagnosticExportFile) {
	o.Files = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticExportManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	toSerialize["packageName"] = o.PackageName
	if o.GeneratedAt.Nanosecond() == 0 {
		toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["redacted"] = o.Redacted
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticExportManifest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId      *string                 `json:"taskId"`
		PackageName *string                 `json:"packageName"`
		GeneratedAt *time.Time              `json:"generatedAt"`
		Redacted    *bool                   `json:"redacted"`
		Files       *[]DiagnosticExportFile `json:"files"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.PackageName == nil {
		return fmt.Errorf("required field packageName missing")
	}
	if all.GeneratedAt == nil {
		return fmt.Errorf("required field generatedAt missing")
	}
	if all.Redacted == nil {
		return fmt.Errorf("required field redacted missing")
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "packageName", "generatedAt", "redacted", "files"})
	} else {
		return err
	}
	o.TaskId = *all.TaskId
	o.PackageName = *all.PackageName
	o.GeneratedAt = *all.GeneratedAt
	o.Redacted = *all.Redacted
	o.Files = *all.Files

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
