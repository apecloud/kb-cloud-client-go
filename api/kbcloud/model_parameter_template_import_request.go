// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterTemplateImportRequest Import a YAML file as a new custom parameter template in the target organization.
type ParameterTemplateImportRequest struct {
	// Name to give the imported template. It must be unique within the target organization.
	Name string `json:"name"`
	// Description to save with the new template.
	Description string `json:"description"`
	// Target engine used to validate the component, major versions and configuration in the file.
	Engine ParameterTemplateImportEngine `json:"engine"`
	// Upload one parameter template YAML file of at most 1 MiB (1,048,576 bytes). Its name, engine and description are optional source information and are ignored during import.
	File _io.Reader `json:"file"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterTemplateImportRequest instantiates a new ParameterTemplateImportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterTemplateImportRequest(name string, description string, engine ParameterTemplateImportEngine, file _io.Reader) *ParameterTemplateImportRequest {
	this := ParameterTemplateImportRequest{}
	this.Name = name
	this.Description = description
	this.Engine = engine
	this.File = file
	return &this
}

// NewParameterTemplateImportRequestWithDefaults instantiates a new ParameterTemplateImportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterTemplateImportRequestWithDefaults() *ParameterTemplateImportRequest {
	this := ParameterTemplateImportRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *ParameterTemplateImportRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateImportRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParameterTemplateImportRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value.
func (o *ParameterTemplateImportRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateImportRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ParameterTemplateImportRequest) SetDescription(v string) {
	o.Description = v
}

// GetEngine returns the Engine field value.
func (o *ParameterTemplateImportRequest) GetEngine() ParameterTemplateImportEngine {
	if o == nil {
		var ret ParameterTemplateImportEngine
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateImportRequest) GetEngineOk() (*ParameterTemplateImportEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *ParameterTemplateImportRequest) SetEngine(v ParameterTemplateImportEngine) {
	o.Engine = v
}

// GetFile returns the File field value.
func (o *ParameterTemplateImportRequest) GetFile() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateImportRequest) GetFileOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value.
func (o *ParameterTemplateImportRequest) SetFile(v _io.Reader) {
	o.File = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterTemplateImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["engine"] = o.Engine
	toSerialize["file"] = o.File

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterTemplateImportRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                        `json:"name"`
		Description *string                        `json:"description"`
		Engine      *ParameterTemplateImportEngine `json:"engine"`
		File        *_io.Reader                    `json:"file"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.File == nil {
		return fmt.Errorf("required field file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "engine", "file"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Description = *all.Description
	if !all.Engine.IsValid() {
		hasInvalidField = true
	} else {
		o.Engine = *all.Engine
	}
	o.File = *all.File

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
