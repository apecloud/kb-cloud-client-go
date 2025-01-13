// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Component Component info in environment
type Component struct {
	// the name of component
	Name string `json:"name"`
	// the version of component
	Version string `json:"version"`
	// the dependencies of component
	Dependencies []string `json:"dependencies,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponent instantiates a new Component object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponent(name string, version string) *Component {
	this := Component{}
	this.Name = name
	this.Version = version
	return &this
}

// NewComponentWithDefaults instantiates a new Component object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentWithDefaults() *Component {
	this := Component{}
	return &this
}

// GetName returns the Name field value.
func (o *Component) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Component) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Component) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value.
func (o *Component) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Component) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *Component) SetVersion(v string) {
	o.Version = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *Component) GetDependencies() []string {
	if o == nil || o.Dependencies == nil {
		var ret []string
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetDependenciesOk() (*[]string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return &o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *Component) HasDependencies() bool {
	return o != nil && o.Dependencies != nil
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *Component) SetDependencies(v []string) {
	o.Dependencies = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Component) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Component) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string  `json:"name"`
		Version      *string  `json:"version"`
		Dependencies []string `json:"dependencies,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "version", "dependencies"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Version = *all.Version
	o.Dependencies = all.Dependencies

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
