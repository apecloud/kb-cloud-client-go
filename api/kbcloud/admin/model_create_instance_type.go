// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

type CreateInstanceType struct {
	// Any characters, no more than 255 characters long.
	Description *string `json:"description,omitempty"`
	// Any characters, no more than 63 characters long.
	DisplayName    *string         `json:"displayName,omitempty"`
	OccupationType *OccupationType `json:"occupationType,omitempty"`
	// Consists of lowercase English letters and numbers, 1-8 characters long. Used to add to the specification code.
	Name                *string              `json:"name,omitempty"`
	ServerConfiguration *ServerConfiguration `json:"serverConfiguration,omitempty"`
	// List of engine names supported by this instance type.
	Engines []string `json:"engines,omitempty"`
	// The Names of the instance types to be copied. All specifications of these instances will be copied to the new instance type.
	SourceInstanceTypeName []string `json:"sourceInstanceTypeName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateInstanceType instantiates a new CreateInstanceType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateInstanceType() *CreateInstanceType {
	this := CreateInstanceType{}
	return &this
}

// NewCreateInstanceTypeWithDefaults instantiates a new CreateInstanceType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateInstanceTypeWithDefaults() *CreateInstanceType {
	this := CreateInstanceType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateInstanceType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateInstanceType) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateInstanceType) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateInstanceType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateInstanceType) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateInstanceType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOccupationType returns the OccupationType field value if set, zero value otherwise.
func (o *CreateInstanceType) GetOccupationType() OccupationType {
	if o == nil || o.OccupationType == nil {
		var ret OccupationType
		return ret
	}
	return *o.OccupationType
}

// GetOccupationTypeOk returns a tuple with the OccupationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetOccupationTypeOk() (*OccupationType, bool) {
	if o == nil || o.OccupationType == nil {
		return nil, false
	}
	return o.OccupationType, true
}

// HasOccupationType returns a boolean if a field has been set.
func (o *CreateInstanceType) HasOccupationType() bool {
	return o != nil && o.OccupationType != nil
}

// SetOccupationType gets a reference to the given OccupationType and assigns it to the OccupationType field.
func (o *CreateInstanceType) SetOccupationType(v OccupationType) {
	o.OccupationType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateInstanceType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateInstanceType) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateInstanceType) SetName(v string) {
	o.Name = &v
}

// GetServerConfiguration returns the ServerConfiguration field value if set, zero value otherwise.
func (o *CreateInstanceType) GetServerConfiguration() ServerConfiguration {
	if o == nil || o.ServerConfiguration == nil {
		var ret ServerConfiguration
		return ret
	}
	return *o.ServerConfiguration
}

// GetServerConfigurationOk returns a tuple with the ServerConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetServerConfigurationOk() (*ServerConfiguration, bool) {
	if o == nil || o.ServerConfiguration == nil {
		return nil, false
	}
	return o.ServerConfiguration, true
}

// HasServerConfiguration returns a boolean if a field has been set.
func (o *CreateInstanceType) HasServerConfiguration() bool {
	return o != nil && o.ServerConfiguration != nil
}

// SetServerConfiguration gets a reference to the given ServerConfiguration and assigns it to the ServerConfiguration field.
func (o *CreateInstanceType) SetServerConfiguration(v ServerConfiguration) {
	o.ServerConfiguration = &v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *CreateInstanceType) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *CreateInstanceType) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *CreateInstanceType) SetEngines(v []string) {
	o.Engines = v
}

// GetSourceInstanceTypeName returns the SourceInstanceTypeName field value if set, zero value otherwise.
func (o *CreateInstanceType) GetSourceInstanceTypeName() []string {
	if o == nil || o.SourceInstanceTypeName == nil {
		var ret []string
		return ret
	}
	return o.SourceInstanceTypeName
}

// GetSourceInstanceTypeNameOk returns a tuple with the SourceInstanceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceType) GetSourceInstanceTypeNameOk() (*[]string, bool) {
	if o == nil || o.SourceInstanceTypeName == nil {
		return nil, false
	}
	return &o.SourceInstanceTypeName, true
}

// HasSourceInstanceTypeName returns a boolean if a field has been set.
func (o *CreateInstanceType) HasSourceInstanceTypeName() bool {
	return o != nil && o.SourceInstanceTypeName != nil
}

// SetSourceInstanceTypeName gets a reference to the given []string and assigns it to the SourceInstanceTypeName field.
func (o *CreateInstanceType) SetSourceInstanceTypeName(v []string) {
	o.SourceInstanceTypeName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateInstanceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.OccupationType != nil {
		toSerialize["occupationType"] = o.OccupationType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServerConfiguration != nil {
		toSerialize["serverConfiguration"] = o.ServerConfiguration
	}
	if o.Engines != nil {
		toSerialize["engines"] = o.Engines
	}
	if o.SourceInstanceTypeName != nil {
		toSerialize["sourceInstanceTypeName"] = o.SourceInstanceTypeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateInstanceType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description            *string              `json:"description,omitempty"`
		DisplayName            *string              `json:"displayName,omitempty"`
		OccupationType         *OccupationType      `json:"occupationType,omitempty"`
		Name                   *string              `json:"name,omitempty"`
		ServerConfiguration    *ServerConfiguration `json:"serverConfiguration,omitempty"`
		Engines                []string             `json:"engines,omitempty"`
		SourceInstanceTypeName []string             `json:"sourceInstanceTypeName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "displayName", "occupationType", "name", "serverConfiguration", "engines", "sourceInstanceTypeName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	if all.OccupationType != nil && !all.OccupationType.IsValid() {
		hasInvalidField = true
	} else {
		o.OccupationType = all.OccupationType
	}
	o.Name = all.Name
	if all.ServerConfiguration != nil && all.ServerConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServerConfiguration = all.ServerConfiguration
	o.Engines = all.Engines
	o.SourceInstanceTypeName = all.SourceInstanceTypeName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
