// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Metadb_backup struct {
	// the path of storaging backups
	Path *string `json:"path,omitempty"`
	// Date/time when the backup finished being processed.
	CompletionTimestamp string `json:"completionTimestamp"`
	// Date/time when the backup was created.
	CreationTimestamp string `json:"creationTimestamp"`
	// The duration time of backup execution. When converted to a string, the form is "1h2m0.5s".
	Duration string `json:"duration"`
	// name of the backup
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetadb_backup instantiates a new Metadb_backup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetadb_backup(completionTimestamp string, creationTimestamp string, duration string, name string) *Metadb_backup {
	this := Metadb_backup{}
	this.CompletionTimestamp = completionTimestamp
	this.CreationTimestamp = creationTimestamp
	this.Duration = duration
	this.Name = name
	return &this
}

// NewMetadb_backupWithDefaults instantiates a new Metadb_backup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetadb_backupWithDefaults() *Metadb_backup {
	this := Metadb_backup{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Metadb_backup) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_backup) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Metadb_backup) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Metadb_backup) SetPath(v string) {
	o.Path = &v
}

// GetCompletionTimestamp returns the CompletionTimestamp field value.
func (o *Metadb_backup) GetCompletionTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CompletionTimestamp
}

// GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field value
// and a boolean to check if the value has been set.
func (o *Metadb_backup) GetCompletionTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionTimestamp, true
}

// SetCompletionTimestamp sets field value.
func (o *Metadb_backup) SetCompletionTimestamp(v string) {
	o.CompletionTimestamp = v
}

// GetCreationTimestamp returns the CreationTimestamp field value.
func (o *Metadb_backup) GetCreationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
func (o *Metadb_backup) GetCreationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value.
func (o *Metadb_backup) SetCreationTimestamp(v string) {
	o.CreationTimestamp = v
}

// GetDuration returns the Duration field value.
func (o *Metadb_backup) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *Metadb_backup) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *Metadb_backup) SetDuration(v string) {
	o.Duration = v
}

// GetName returns the Name field value.
func (o *Metadb_backup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metadb_backup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Metadb_backup) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Metadb_backup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	toSerialize["completionTimestamp"] = o.CompletionTimestamp
	toSerialize["creationTimestamp"] = o.CreationTimestamp
	toSerialize["duration"] = o.Duration
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Metadb_backup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Path                *string `json:"path,omitempty"`
		CompletionTimestamp *string `json:"completionTimestamp"`
		CreationTimestamp   *string `json:"creationTimestamp"`
		Duration            *string `json:"duration"`
		Name                *string `json:"name"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CompletionTimestamp == nil {
		return fmt.Errorf("required field completionTimestamp missing")
	}
	if all.CreationTimestamp == nil {
		return fmt.Errorf("required field creationTimestamp missing")
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"path", "completionTimestamp", "creationTimestamp", "duration", "name"})
	} else {
		return err
	}
	o.Path = all.Path
	o.CompletionTimestamp = *all.CompletionTimestamp
	o.CreationTimestamp = *all.CreationTimestamp
	o.Duration = *all.Duration
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
