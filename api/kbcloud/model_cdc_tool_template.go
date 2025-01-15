// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CdcToolTemplate struct {
	Image              *string                `json:"image,omitempty"`
	Command            []string               `json:"command,omitempty"`
	Args               []string               `json:"args,omitempty"`
	Replicas           common.NullableInt32   `json:"replicas,omitempty"`
	ConfigFileName     *string                `json:"configFileName,omitempty"`
	ConfigFilePath     *string                `json:"configFilePath,omitempty"`
	ConfigResourceType *CdcConfigResourceType `json:"configResourceType,omitempty"`
	ConfigFormat       *CdcConfigFormatType   `json:"configFormat,omitempty"`
	// whether to enable local recovery feature
	LocalRecoveryEnabled      *bool                 `json:"localRecoveryEnabled,omitempty"`
	LocalRecoveryStorageClass common.NullableString `json:"localRecoveryStorageClass,omitempty"`
	// size of local recovery storage, the unit is MiB
	LocalRecoveryStorageSize common.NullableInt32 `json:"localRecoveryStorageSize,omitempty"`
	LocalRecoveryPath        *string              `json:"localRecoveryPath,omitempty"`
	LogPath                  *string              `json:"LogPath,omitempty"`
	StartupTimeoutMinutes    common.NullableInt32 `json:"startupTimeoutMinutes,omitempty"`
	Properties               map[string]string    `json:"properties,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcToolTemplate instantiates a new CdcToolTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcToolTemplate() *CdcToolTemplate {
	this := CdcToolTemplate{}
	return &this
}

// NewCdcToolTemplateWithDefaults instantiates a new CdcToolTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcToolTemplateWithDefaults() *CdcToolTemplate {
	this := CdcToolTemplate{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasImage() bool {
	return o != nil && o.Image != nil
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CdcToolTemplate) SetImage(v string) {
	o.Image = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetCommand() []string {
	if o == nil || o.Command == nil {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetCommandOk() (*[]string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return &o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasCommand() bool {
	return o != nil && o.Command != nil
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *CdcToolTemplate) SetCommand(v []string) {
	o.Command = v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetArgsOk() (*[]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return &o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasArgs() bool {
	return o != nil && o.Args != nil
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *CdcToolTemplate) SetArgs(v []string) {
	o.Args = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdcToolTemplate) GetReplicas() int32 {
	if o == nil || o.Replicas.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Replicas.Get()
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CdcToolTemplate) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas.Get(), o.Replicas.IsSet()
}

// HasReplicas returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasReplicas() bool {
	return o != nil && o.Replicas.IsSet()
}

// SetReplicas gets a reference to the given common.NullableInt32 and assigns it to the Replicas field.
func (o *CdcToolTemplate) SetReplicas(v int32) {
	o.Replicas.Set(&v)
}

// SetReplicasNil sets the value for Replicas to be an explicit nil.
func (o *CdcToolTemplate) SetReplicasNil() {
	o.Replicas.Set(nil)
}

// UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil.
func (o *CdcToolTemplate) UnsetReplicas() {
	o.Replicas.Unset()
}

// GetConfigFileName returns the ConfigFileName field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetConfigFileName() string {
	if o == nil || o.ConfigFileName == nil {
		var ret string
		return ret
	}
	return *o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetConfigFileNameOk() (*string, bool) {
	if o == nil || o.ConfigFileName == nil {
		return nil, false
	}
	return o.ConfigFileName, true
}

// HasConfigFileName returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasConfigFileName() bool {
	return o != nil && o.ConfigFileName != nil
}

// SetConfigFileName gets a reference to the given string and assigns it to the ConfigFileName field.
func (o *CdcToolTemplate) SetConfigFileName(v string) {
	o.ConfigFileName = &v
}

// GetConfigFilePath returns the ConfigFilePath field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetConfigFilePath() string {
	if o == nil || o.ConfigFilePath == nil {
		var ret string
		return ret
	}
	return *o.ConfigFilePath
}

// GetConfigFilePathOk returns a tuple with the ConfigFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetConfigFilePathOk() (*string, bool) {
	if o == nil || o.ConfigFilePath == nil {
		return nil, false
	}
	return o.ConfigFilePath, true
}

// HasConfigFilePath returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasConfigFilePath() bool {
	return o != nil && o.ConfigFilePath != nil
}

// SetConfigFilePath gets a reference to the given string and assigns it to the ConfigFilePath field.
func (o *CdcToolTemplate) SetConfigFilePath(v string) {
	o.ConfigFilePath = &v
}

// GetConfigResourceType returns the ConfigResourceType field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetConfigResourceType() CdcConfigResourceType {
	if o == nil || o.ConfigResourceType == nil {
		var ret CdcConfigResourceType
		return ret
	}
	return *o.ConfigResourceType
}

// GetConfigResourceTypeOk returns a tuple with the ConfigResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetConfigResourceTypeOk() (*CdcConfigResourceType, bool) {
	if o == nil || o.ConfigResourceType == nil {
		return nil, false
	}
	return o.ConfigResourceType, true
}

// HasConfigResourceType returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasConfigResourceType() bool {
	return o != nil && o.ConfigResourceType != nil
}

// SetConfigResourceType gets a reference to the given CdcConfigResourceType and assigns it to the ConfigResourceType field.
func (o *CdcToolTemplate) SetConfigResourceType(v CdcConfigResourceType) {
	o.ConfigResourceType = &v
}

// GetConfigFormat returns the ConfigFormat field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetConfigFormat() CdcConfigFormatType {
	if o == nil || o.ConfigFormat == nil {
		var ret CdcConfigFormatType
		return ret
	}
	return *o.ConfigFormat
}

// GetConfigFormatOk returns a tuple with the ConfigFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetConfigFormatOk() (*CdcConfigFormatType, bool) {
	if o == nil || o.ConfigFormat == nil {
		return nil, false
	}
	return o.ConfigFormat, true
}

// HasConfigFormat returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasConfigFormat() bool {
	return o != nil && o.ConfigFormat != nil
}

// SetConfigFormat gets a reference to the given CdcConfigFormatType and assigns it to the ConfigFormat field.
func (o *CdcToolTemplate) SetConfigFormat(v CdcConfigFormatType) {
	o.ConfigFormat = &v
}

// GetLocalRecoveryEnabled returns the LocalRecoveryEnabled field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetLocalRecoveryEnabled() bool {
	if o == nil || o.LocalRecoveryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LocalRecoveryEnabled
}

// GetLocalRecoveryEnabledOk returns a tuple with the LocalRecoveryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetLocalRecoveryEnabledOk() (*bool, bool) {
	if o == nil || o.LocalRecoveryEnabled == nil {
		return nil, false
	}
	return o.LocalRecoveryEnabled, true
}

// HasLocalRecoveryEnabled returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasLocalRecoveryEnabled() bool {
	return o != nil && o.LocalRecoveryEnabled != nil
}

// SetLocalRecoveryEnabled gets a reference to the given bool and assigns it to the LocalRecoveryEnabled field.
func (o *CdcToolTemplate) SetLocalRecoveryEnabled(v bool) {
	o.LocalRecoveryEnabled = &v
}

// GetLocalRecoveryStorageClass returns the LocalRecoveryStorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdcToolTemplate) GetLocalRecoveryStorageClass() string {
	if o == nil || o.LocalRecoveryStorageClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalRecoveryStorageClass.Get()
}

// GetLocalRecoveryStorageClassOk returns a tuple with the LocalRecoveryStorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CdcToolTemplate) GetLocalRecoveryStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalRecoveryStorageClass.Get(), o.LocalRecoveryStorageClass.IsSet()
}

// HasLocalRecoveryStorageClass returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasLocalRecoveryStorageClass() bool {
	return o != nil && o.LocalRecoveryStorageClass.IsSet()
}

// SetLocalRecoveryStorageClass gets a reference to the given common.NullableString and assigns it to the LocalRecoveryStorageClass field.
func (o *CdcToolTemplate) SetLocalRecoveryStorageClass(v string) {
	o.LocalRecoveryStorageClass.Set(&v)
}

// SetLocalRecoveryStorageClassNil sets the value for LocalRecoveryStorageClass to be an explicit nil.
func (o *CdcToolTemplate) SetLocalRecoveryStorageClassNil() {
	o.LocalRecoveryStorageClass.Set(nil)
}

// UnsetLocalRecoveryStorageClass ensures that no value is present for LocalRecoveryStorageClass, not even an explicit nil.
func (o *CdcToolTemplate) UnsetLocalRecoveryStorageClass() {
	o.LocalRecoveryStorageClass.Unset()
}

// GetLocalRecoveryStorageSize returns the LocalRecoveryStorageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdcToolTemplate) GetLocalRecoveryStorageSize() int32 {
	if o == nil || o.LocalRecoveryStorageSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.LocalRecoveryStorageSize.Get()
}

// GetLocalRecoveryStorageSizeOk returns a tuple with the LocalRecoveryStorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CdcToolTemplate) GetLocalRecoveryStorageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalRecoveryStorageSize.Get(), o.LocalRecoveryStorageSize.IsSet()
}

// HasLocalRecoveryStorageSize returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasLocalRecoveryStorageSize() bool {
	return o != nil && o.LocalRecoveryStorageSize.IsSet()
}

// SetLocalRecoveryStorageSize gets a reference to the given common.NullableInt32 and assigns it to the LocalRecoveryStorageSize field.
func (o *CdcToolTemplate) SetLocalRecoveryStorageSize(v int32) {
	o.LocalRecoveryStorageSize.Set(&v)
}

// SetLocalRecoveryStorageSizeNil sets the value for LocalRecoveryStorageSize to be an explicit nil.
func (o *CdcToolTemplate) SetLocalRecoveryStorageSizeNil() {
	o.LocalRecoveryStorageSize.Set(nil)
}

// UnsetLocalRecoveryStorageSize ensures that no value is present for LocalRecoveryStorageSize, not even an explicit nil.
func (o *CdcToolTemplate) UnsetLocalRecoveryStorageSize() {
	o.LocalRecoveryStorageSize.Unset()
}

// GetLocalRecoveryPath returns the LocalRecoveryPath field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetLocalRecoveryPath() string {
	if o == nil || o.LocalRecoveryPath == nil {
		var ret string
		return ret
	}
	return *o.LocalRecoveryPath
}

// GetLocalRecoveryPathOk returns a tuple with the LocalRecoveryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetLocalRecoveryPathOk() (*string, bool) {
	if o == nil || o.LocalRecoveryPath == nil {
		return nil, false
	}
	return o.LocalRecoveryPath, true
}

// HasLocalRecoveryPath returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasLocalRecoveryPath() bool {
	return o != nil && o.LocalRecoveryPath != nil
}

// SetLocalRecoveryPath gets a reference to the given string and assigns it to the LocalRecoveryPath field.
func (o *CdcToolTemplate) SetLocalRecoveryPath(v string) {
	o.LocalRecoveryPath = &v
}

// GetLogPath returns the LogPath field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetLogPath() string {
	if o == nil || o.LogPath == nil {
		var ret string
		return ret
	}
	return *o.LogPath
}

// GetLogPathOk returns a tuple with the LogPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetLogPathOk() (*string, bool) {
	if o == nil || o.LogPath == nil {
		return nil, false
	}
	return o.LogPath, true
}

// HasLogPath returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasLogPath() bool {
	return o != nil && o.LogPath != nil
}

// SetLogPath gets a reference to the given string and assigns it to the LogPath field.
func (o *CdcToolTemplate) SetLogPath(v string) {
	o.LogPath = &v
}

// GetStartupTimeoutMinutes returns the StartupTimeoutMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdcToolTemplate) GetStartupTimeoutMinutes() int32 {
	if o == nil || o.StartupTimeoutMinutes.Get() == nil {
		var ret int32
		return ret
	}
	return *o.StartupTimeoutMinutes.Get()
}

// GetStartupTimeoutMinutesOk returns a tuple with the StartupTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CdcToolTemplate) GetStartupTimeoutMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartupTimeoutMinutes.Get(), o.StartupTimeoutMinutes.IsSet()
}

// HasStartupTimeoutMinutes returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasStartupTimeoutMinutes() bool {
	return o != nil && o.StartupTimeoutMinutes.IsSet()
}

// SetStartupTimeoutMinutes gets a reference to the given common.NullableInt32 and assigns it to the StartupTimeoutMinutes field.
func (o *CdcToolTemplate) SetStartupTimeoutMinutes(v int32) {
	o.StartupTimeoutMinutes.Set(&v)
}

// SetStartupTimeoutMinutesNil sets the value for StartupTimeoutMinutes to be an explicit nil.
func (o *CdcToolTemplate) SetStartupTimeoutMinutesNil() {
	o.StartupTimeoutMinutes.Set(nil)
}

// UnsetStartupTimeoutMinutes ensures that no value is present for StartupTimeoutMinutes, not even an explicit nil.
func (o *CdcToolTemplate) UnsetStartupTimeoutMinutes() {
	o.StartupTimeoutMinutes.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CdcToolTemplate) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcToolTemplate) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CdcToolTemplate) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *CdcToolTemplate) SetProperties(v map[string]string) {
	o.Properties = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcToolTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Replicas.IsSet() {
		toSerialize["replicas"] = o.Replicas.Get()
	}
	if o.ConfigFileName != nil {
		toSerialize["configFileName"] = o.ConfigFileName
	}
	if o.ConfigFilePath != nil {
		toSerialize["configFilePath"] = o.ConfigFilePath
	}
	if o.ConfigResourceType != nil {
		toSerialize["configResourceType"] = o.ConfigResourceType
	}
	if o.ConfigFormat != nil {
		toSerialize["configFormat"] = o.ConfigFormat
	}
	if o.LocalRecoveryEnabled != nil {
		toSerialize["localRecoveryEnabled"] = o.LocalRecoveryEnabled
	}
	if o.LocalRecoveryStorageClass.IsSet() {
		toSerialize["localRecoveryStorageClass"] = o.LocalRecoveryStorageClass.Get()
	}
	if o.LocalRecoveryStorageSize.IsSet() {
		toSerialize["localRecoveryStorageSize"] = o.LocalRecoveryStorageSize.Get()
	}
	if o.LocalRecoveryPath != nil {
		toSerialize["localRecoveryPath"] = o.LocalRecoveryPath
	}
	if o.LogPath != nil {
		toSerialize["LogPath"] = o.LogPath
	}
	if o.StartupTimeoutMinutes.IsSet() {
		toSerialize["startupTimeoutMinutes"] = o.StartupTimeoutMinutes.Get()
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcToolTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Image                     *string                `json:"image,omitempty"`
		Command                   []string               `json:"command,omitempty"`
		Args                      []string               `json:"args,omitempty"`
		Replicas                  common.NullableInt32   `json:"replicas,omitempty"`
		ConfigFileName            *string                `json:"configFileName,omitempty"`
		ConfigFilePath            *string                `json:"configFilePath,omitempty"`
		ConfigResourceType        *CdcConfigResourceType `json:"configResourceType,omitempty"`
		ConfigFormat              *CdcConfigFormatType   `json:"configFormat,omitempty"`
		LocalRecoveryEnabled      *bool                  `json:"localRecoveryEnabled,omitempty"`
		LocalRecoveryStorageClass common.NullableString  `json:"localRecoveryStorageClass,omitempty"`
		LocalRecoveryStorageSize  common.NullableInt32   `json:"localRecoveryStorageSize,omitempty"`
		LocalRecoveryPath         *string                `json:"localRecoveryPath,omitempty"`
		LogPath                   *string                `json:"LogPath,omitempty"`
		StartupTimeoutMinutes     common.NullableInt32   `json:"startupTimeoutMinutes,omitempty"`
		Properties                map[string]string      `json:"properties,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"image", "command", "args", "replicas", "configFileName", "configFilePath", "configResourceType", "configFormat", "localRecoveryEnabled", "localRecoveryStorageClass", "localRecoveryStorageSize", "localRecoveryPath", "LogPath", "startupTimeoutMinutes", "properties"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Image = all.Image
	o.Command = all.Command
	o.Args = all.Args
	o.Replicas = all.Replicas
	o.ConfigFileName = all.ConfigFileName
	o.ConfigFilePath = all.ConfigFilePath
	if all.ConfigResourceType != nil && !all.ConfigResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ConfigResourceType = all.ConfigResourceType
	}
	if all.ConfigFormat != nil && !all.ConfigFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ConfigFormat = all.ConfigFormat
	}
	o.LocalRecoveryEnabled = all.LocalRecoveryEnabled
	o.LocalRecoveryStorageClass = all.LocalRecoveryStorageClass
	o.LocalRecoveryStorageSize = all.LocalRecoveryStorageSize
	o.LocalRecoveryPath = all.LocalRecoveryPath
	o.LogPath = all.LogPath
	o.StartupTimeoutMinutes = all.StartupTimeoutMinutes
	o.Properties = all.Properties

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
