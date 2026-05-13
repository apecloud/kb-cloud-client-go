// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentClusterContext struct {
	DisplayName *string `json:"displayName,omitempty"`
	Engine      *string `json:"engine,omitempty"`
	Version     *string `json:"version,omitempty"`
	Status      *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentClusterContext instantiates a new AiAgentClusterContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentClusterContext() *AiAgentClusterContext {
	this := AiAgentClusterContext{}
	return &this
}

// NewAiAgentClusterContextWithDefaults instantiates a new AiAgentClusterContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentClusterContextWithDefaults() *AiAgentClusterContext {
	this := AiAgentClusterContext{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AiAgentClusterContext) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentClusterContext) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AiAgentClusterContext) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AiAgentClusterContext) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *AiAgentClusterContext) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentClusterContext) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *AiAgentClusterContext) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *AiAgentClusterContext) SetEngine(v string) {
	o.Engine = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AiAgentClusterContext) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentClusterContext) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AiAgentClusterContext) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AiAgentClusterContext) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentClusterContext) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentClusterContext) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentClusterContext) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiAgentClusterContext) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentClusterContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
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
func (o *AiAgentClusterContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string `json:"displayName,omitempty"`
		Engine      *string `json:"engine,omitempty"`
		Version     *string `json:"version,omitempty"`
		Status      *string `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"displayName", "engine", "version", "status"})
	} else {
		return err
	}
	o.DisplayName = all.DisplayName
	o.Engine = all.Engine
	o.Version = all.Version
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
