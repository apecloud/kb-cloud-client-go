// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentRunMetadata struct {
	// Actual model name used for this run. The value reuses the existing /llm/models string contract.
	Model                *string `json:"model,omitempty"`
	PlaybookVersion      *string `json:"playbookVersion,omitempty"`
	ToolsetVersion       *string `json:"toolsetVersion,omitempty"`
	PromptOrSkillVersion *string `json:"promptOrSkillVersion,omitempty"`
	// Sanitized model strategy/version metadata for audit and replay. Must not include llm.config, provider credentials, or secrets.
	ModelConfig     map[string]interface{} `json:"modelConfig,omitempty"`
	BuilderVersions map[string]string      `json:"builderVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRunMetadata instantiates a new AiAgentRunMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRunMetadata() *AiAgentRunMetadata {
	this := AiAgentRunMetadata{}
	return &this
}

// NewAiAgentRunMetadataWithDefaults instantiates a new AiAgentRunMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRunMetadataWithDefaults() *AiAgentRunMetadata {
	this := AiAgentRunMetadata{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiAgentRunMetadata) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunMetadata) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiAgentRunMetadata) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiAgentRunMetadata) SetModel(v string) {
	o.Model = &v
}

// GetPlaybookVersion returns the PlaybookVersion field value if set, zero value otherwise.
func (o *AiAgentRunMetadata) GetPlaybookVersion() string {
	if o == nil || o.PlaybookVersion == nil {
		var ret string
		return ret
	}
	return *o.PlaybookVersion
}

// GetPlaybookVersionOk returns a tuple with the PlaybookVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunMetadata) GetPlaybookVersionOk() (*string, bool) {
	if o == nil || o.PlaybookVersion == nil {
		return nil, false
	}
	return o.PlaybookVersion, true
}

// HasPlaybookVersion returns a boolean if a field has been set.
func (o *AiAgentRunMetadata) HasPlaybookVersion() bool {
	return o != nil && o.PlaybookVersion != nil
}

// SetPlaybookVersion gets a reference to the given string and assigns it to the PlaybookVersion field.
func (o *AiAgentRunMetadata) SetPlaybookVersion(v string) {
	o.PlaybookVersion = &v
}

// GetToolsetVersion returns the ToolsetVersion field value if set, zero value otherwise.
func (o *AiAgentRunMetadata) GetToolsetVersion() string {
	if o == nil || o.ToolsetVersion == nil {
		var ret string
		return ret
	}
	return *o.ToolsetVersion
}

// GetToolsetVersionOk returns a tuple with the ToolsetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunMetadata) GetToolsetVersionOk() (*string, bool) {
	if o == nil || o.ToolsetVersion == nil {
		return nil, false
	}
	return o.ToolsetVersion, true
}

// HasToolsetVersion returns a boolean if a field has been set.
func (o *AiAgentRunMetadata) HasToolsetVersion() bool {
	return o != nil && o.ToolsetVersion != nil
}

// SetToolsetVersion gets a reference to the given string and assigns it to the ToolsetVersion field.
func (o *AiAgentRunMetadata) SetToolsetVersion(v string) {
	o.ToolsetVersion = &v
}

// GetPromptOrSkillVersion returns the PromptOrSkillVersion field value if set, zero value otherwise.
func (o *AiAgentRunMetadata) GetPromptOrSkillVersion() string {
	if o == nil || o.PromptOrSkillVersion == nil {
		var ret string
		return ret
	}
	return *o.PromptOrSkillVersion
}

// GetPromptOrSkillVersionOk returns a tuple with the PromptOrSkillVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunMetadata) GetPromptOrSkillVersionOk() (*string, bool) {
	if o == nil || o.PromptOrSkillVersion == nil {
		return nil, false
	}
	return o.PromptOrSkillVersion, true
}

// HasPromptOrSkillVersion returns a boolean if a field has been set.
func (o *AiAgentRunMetadata) HasPromptOrSkillVersion() bool {
	return o != nil && o.PromptOrSkillVersion != nil
}

// SetPromptOrSkillVersion gets a reference to the given string and assigns it to the PromptOrSkillVersion field.
func (o *AiAgentRunMetadata) SetPromptOrSkillVersion(v string) {
	o.PromptOrSkillVersion = &v
}

// GetModelConfig returns the ModelConfig field value if set, zero value otherwise.
func (o *AiAgentRunMetadata) GetModelConfig() map[string]interface{} {
	if o == nil || o.ModelConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ModelConfig
}

// GetModelConfigOk returns a tuple with the ModelConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunMetadata) GetModelConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.ModelConfig == nil {
		return nil, false
	}
	return &o.ModelConfig, true
}

// HasModelConfig returns a boolean if a field has been set.
func (o *AiAgentRunMetadata) HasModelConfig() bool {
	return o != nil && o.ModelConfig != nil
}

// SetModelConfig gets a reference to the given map[string]interface{} and assigns it to the ModelConfig field.
func (o *AiAgentRunMetadata) SetModelConfig(v map[string]interface{}) {
	o.ModelConfig = v
}

// GetBuilderVersions returns the BuilderVersions field value if set, zero value otherwise.
func (o *AiAgentRunMetadata) GetBuilderVersions() map[string]string {
	if o == nil || o.BuilderVersions == nil {
		var ret map[string]string
		return ret
	}
	return o.BuilderVersions
}

// GetBuilderVersionsOk returns a tuple with the BuilderVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunMetadata) GetBuilderVersionsOk() (*map[string]string, bool) {
	if o == nil || o.BuilderVersions == nil {
		return nil, false
	}
	return &o.BuilderVersions, true
}

// HasBuilderVersions returns a boolean if a field has been set.
func (o *AiAgentRunMetadata) HasBuilderVersions() bool {
	return o != nil && o.BuilderVersions != nil
}

// SetBuilderVersions gets a reference to the given map[string]string and assigns it to the BuilderVersions field.
func (o *AiAgentRunMetadata) SetBuilderVersions(v map[string]string) {
	o.BuilderVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRunMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.PlaybookVersion != nil {
		toSerialize["playbookVersion"] = o.PlaybookVersion
	}
	if o.ToolsetVersion != nil {
		toSerialize["toolsetVersion"] = o.ToolsetVersion
	}
	if o.PromptOrSkillVersion != nil {
		toSerialize["promptOrSkillVersion"] = o.PromptOrSkillVersion
	}
	if o.ModelConfig != nil {
		toSerialize["modelConfig"] = o.ModelConfig
	}
	if o.BuilderVersions != nil {
		toSerialize["builderVersions"] = o.BuilderVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRunMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Model                *string                `json:"model,omitempty"`
		PlaybookVersion      *string                `json:"playbookVersion,omitempty"`
		ToolsetVersion       *string                `json:"toolsetVersion,omitempty"`
		PromptOrSkillVersion *string                `json:"promptOrSkillVersion,omitempty"`
		ModelConfig          map[string]interface{} `json:"modelConfig,omitempty"`
		BuilderVersions      map[string]string      `json:"builderVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"model", "playbookVersion", "toolsetVersion", "promptOrSkillVersion", "modelConfig", "builderVersions"})
	} else {
		return err
	}
	o.Model = all.Model
	o.PlaybookVersion = all.PlaybookVersion
	o.ToolsetVersion = all.ToolsetVersion
	o.PromptOrSkillVersion = all.PromptOrSkillVersion
	o.ModelConfig = all.ModelConfig
	o.BuilderVersions = all.BuilderVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
