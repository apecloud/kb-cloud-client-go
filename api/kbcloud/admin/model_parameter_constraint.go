// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterConstraint struct {
	// a alias with major version, used to distinguish between different parameter templates
	Family string `json:"family"`
	// match the major version set in the component
	MajorVersion string `json:"majorVersion"`
	// deprecated
	Versions []string `json:"versions,omitempty"`
	// name of config spec, included in configSpecs, equal to componentDefinition.configs[x].name
	ConfigSpecName string `json:"configSpecName"`
	// used in KB version < 1.0, constraint of the configConstraint, equal to componentDefinition.configs[x].constraintRef
	ConstraintRef *string `json:"constraintRef,omitempty"`
	// regular expression of the parameters, mainly used by the frontend for parameter parsing.
	Regex string `json:"regex"`
	// name of configuration file
	ConfigFileName string `json:"configFileName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterConstraint instantiates a new ParameterConstraint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterConstraint(family string, majorVersion string, configSpecName string, regex string, configFileName string) *ParameterConstraint {
	this := ParameterConstraint{}
	this.Family = family
	this.MajorVersion = majorVersion
	this.ConfigSpecName = configSpecName
	this.Regex = regex
	this.ConfigFileName = configFileName
	return &this
}

// NewParameterConstraintWithDefaults instantiates a new ParameterConstraint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterConstraintWithDefaults() *ParameterConstraint {
	this := ParameterConstraint{}
	return &this
}

// GetFamily returns the Family field value.
func (o *ParameterConstraint) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value.
func (o *ParameterConstraint) SetFamily(v string) {
	o.Family = v
}

// GetMajorVersion returns the MajorVersion field value.
func (o *ParameterConstraint) GetMajorVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetMajorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MajorVersion, true
}

// SetMajorVersion sets field value.
func (o *ParameterConstraint) SetMajorVersion(v string) {
	o.MajorVersion = v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ParameterConstraint) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ParameterConstraint) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *ParameterConstraint) SetVersions(v []string) {
	o.Versions = v
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *ParameterConstraint) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *ParameterConstraint) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetConstraintRef returns the ConstraintRef field value if set, zero value otherwise.
func (o *ParameterConstraint) GetConstraintRef() string {
	if o == nil || o.ConstraintRef == nil {
		var ret string
		return ret
	}
	return *o.ConstraintRef
}

// GetConstraintRefOk returns a tuple with the ConstraintRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetConstraintRefOk() (*string, bool) {
	if o == nil || o.ConstraintRef == nil {
		return nil, false
	}
	return o.ConstraintRef, true
}

// HasConstraintRef returns a boolean if a field has been set.
func (o *ParameterConstraint) HasConstraintRef() bool {
	return o != nil && o.ConstraintRef != nil
}

// SetConstraintRef gets a reference to the given string and assigns it to the ConstraintRef field.
func (o *ParameterConstraint) SetConstraintRef(v string) {
	o.ConstraintRef = &v
}

// GetRegex returns the Regex field value.
func (o *ParameterConstraint) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value.
func (o *ParameterConstraint) SetRegex(v string) {
	o.Regex = v
}

// GetConfigFileName returns the ConfigFileName field value.
func (o *ParameterConstraint) GetConfigFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value
// and a boolean to check if the value has been set.
func (o *ParameterConstraint) GetConfigFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigFileName, true
}

// SetConfigFileName sets field value.
func (o *ParameterConstraint) SetConfigFileName(v string) {
	o.ConfigFileName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["family"] = o.Family
	toSerialize["majorVersion"] = o.MajorVersion
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	toSerialize["configSpecName"] = o.ConfigSpecName
	if o.ConstraintRef != nil {
		toSerialize["constraintRef"] = o.ConstraintRef
	}
	toSerialize["regex"] = o.Regex
	toSerialize["configFileName"] = o.ConfigFileName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterConstraint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family         *string  `json:"family"`
		MajorVersion   *string  `json:"majorVersion"`
		Versions       []string `json:"versions,omitempty"`
		ConfigSpecName *string  `json:"configSpecName"`
		ConstraintRef  *string  `json:"constraintRef,omitempty"`
		Regex          *string  `json:"regex"`
		ConfigFileName *string  `json:"configFileName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Family == nil {
		return fmt.Errorf("required field family missing")
	}
	if all.MajorVersion == nil {
		return fmt.Errorf("required field majorVersion missing")
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.Regex == nil {
		return fmt.Errorf("required field regex missing")
	}
	if all.ConfigFileName == nil {
		return fmt.Errorf("required field configFileName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"family", "majorVersion", "versions", "configSpecName", "constraintRef", "regex", "configFileName"})
	} else {
		return err
	}
	o.Family = *all.Family
	o.MajorVersion = *all.MajorVersion
	o.Versions = all.Versions
	o.ConfigSpecName = *all.ConfigSpecName
	o.ConstraintRef = all.ConstraintRef
	o.Regex = *all.Regex
	o.ConfigFileName = *all.ConfigFileName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
