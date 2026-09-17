// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayDataSourceSpec struct {
	Name   string                        `json:"name"`
	Type   string                        `json:"type"`
	Engine AiDataGatewayDataSourceEngine `json:"engine"`
	// Defaults to enabled on creation; omitted on update preserves the current status.
	Status           *AiDataGatewayDataSourceStatus `json:"status,omitempty"`
	CloudClusterId   *int64                         `json:"cloudClusterId,omitempty"`
	CloudClusterName *string                        `json:"cloudClusterName,omitempty"`
	EnvironmentName  *string                        `json:"environmentName,omitempty"`
	// Structured connection fields (host, numeric port, database/databaseName, username/user, password, sslMode, sslRootCert, sslCert, sslKey, serverName, tlsEnabled) or encrypted connection-string fields (dsn/url/uri/connectionString). Omitted and masked secret values on update preserve the stored value. Cloud bindings require an explicit database username and use the authoritative managed endpoint. All datasource connection information. For external databases this comes from API input; for Cloud/KubeBlocks managed datasources Cloud resolves and fills the connection information. Sensitive keys such as password, token, secret, privateKey, accessKey, and credential are encrypted at rest and masked in user-facing responses. Internal Runtime config resolves the decrypted view only inside the trusted server boundary.
	ConnectionConfig map[string]interface{} `json:"connectionConfig,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayDataSourceSpec instantiates a new AiDataGatewayDataSourceSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayDataSourceSpec(name string, typeVar string, engine AiDataGatewayDataSourceEngine) *AiDataGatewayDataSourceSpec {
	this := AiDataGatewayDataSourceSpec{}
	this.Name = name
	this.Type = typeVar
	this.Engine = engine
	return &this
}

// NewAiDataGatewayDataSourceSpecWithDefaults instantiates a new AiDataGatewayDataSourceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewayDataSourceSpecWithDefaults() *AiDataGatewayDataSourceSpec {
	this := AiDataGatewayDataSourceSpec{}
	return &this
}

// GetName returns the Name field value.
func (o *AiDataGatewayDataSourceSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiDataGatewayDataSourceSpec) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *AiDataGatewayDataSourceSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiDataGatewayDataSourceSpec) SetType(v string) {
	o.Type = v
}

// GetEngine returns the Engine field value.
func (o *AiDataGatewayDataSourceSpec) GetEngine() AiDataGatewayDataSourceEngine {
	if o == nil {
		var ret AiDataGatewayDataSourceEngine
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetEngineOk() (*AiDataGatewayDataSourceEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *AiDataGatewayDataSourceSpec) SetEngine(v AiDataGatewayDataSourceEngine) {
	o.Engine = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayDataSourceSpec) GetStatus() AiDataGatewayDataSourceStatus {
	if o == nil || o.Status == nil {
		var ret AiDataGatewayDataSourceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetStatusOk() (*AiDataGatewayDataSourceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayDataSourceSpec) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiDataGatewayDataSourceStatus and assigns it to the Status field.
func (o *AiDataGatewayDataSourceSpec) SetStatus(v AiDataGatewayDataSourceStatus) {
	o.Status = &v
}

// GetCloudClusterId returns the CloudClusterId field value if set, zero value otherwise.
func (o *AiDataGatewayDataSourceSpec) GetCloudClusterId() int64 {
	if o == nil || o.CloudClusterId == nil {
		var ret int64
		return ret
	}
	return *o.CloudClusterId
}

// GetCloudClusterIdOk returns a tuple with the CloudClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetCloudClusterIdOk() (*int64, bool) {
	if o == nil || o.CloudClusterId == nil {
		return nil, false
	}
	return o.CloudClusterId, true
}

// HasCloudClusterId returns a boolean if a field has been set.
func (o *AiDataGatewayDataSourceSpec) HasCloudClusterId() bool {
	return o != nil && o.CloudClusterId != nil
}

// SetCloudClusterId gets a reference to the given int64 and assigns it to the CloudClusterId field.
func (o *AiDataGatewayDataSourceSpec) SetCloudClusterId(v int64) {
	o.CloudClusterId = &v
}

// GetCloudClusterName returns the CloudClusterName field value if set, zero value otherwise.
func (o *AiDataGatewayDataSourceSpec) GetCloudClusterName() string {
	if o == nil || o.CloudClusterName == nil {
		var ret string
		return ret
	}
	return *o.CloudClusterName
}

// GetCloudClusterNameOk returns a tuple with the CloudClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetCloudClusterNameOk() (*string, bool) {
	if o == nil || o.CloudClusterName == nil {
		return nil, false
	}
	return o.CloudClusterName, true
}

// HasCloudClusterName returns a boolean if a field has been set.
func (o *AiDataGatewayDataSourceSpec) HasCloudClusterName() bool {
	return o != nil && o.CloudClusterName != nil
}

// SetCloudClusterName gets a reference to the given string and assigns it to the CloudClusterName field.
func (o *AiDataGatewayDataSourceSpec) SetCloudClusterName(v string) {
	o.CloudClusterName = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *AiDataGatewayDataSourceSpec) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *AiDataGatewayDataSourceSpec) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *AiDataGatewayDataSourceSpec) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetConnectionConfig returns the ConnectionConfig field value if set, zero value otherwise.
func (o *AiDataGatewayDataSourceSpec) GetConnectionConfig() map[string]interface{} {
	if o == nil || o.ConnectionConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectionConfig
}

// GetConnectionConfigOk returns a tuple with the ConnectionConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSourceSpec) GetConnectionConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.ConnectionConfig == nil {
		return nil, false
	}
	return &o.ConnectionConfig, true
}

// HasConnectionConfig returns a boolean if a field has been set.
func (o *AiDataGatewayDataSourceSpec) HasConnectionConfig() bool {
	return o != nil && o.ConnectionConfig != nil
}

// SetConnectionConfig gets a reference to the given map[string]interface{} and assigns it to the ConnectionConfig field.
func (o *AiDataGatewayDataSourceSpec) SetConnectionConfig(v map[string]interface{}) {
	o.ConnectionConfig = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayDataSourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["engine"] = o.Engine
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CloudClusterId != nil {
		toSerialize["cloudClusterId"] = o.CloudClusterId
	}
	if o.CloudClusterName != nil {
		toSerialize["cloudClusterName"] = o.CloudClusterName
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.ConnectionConfig != nil {
		toSerialize["connectionConfig"] = o.ConnectionConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayDataSourceSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string                        `json:"name"`
		Type             *string                        `json:"type"`
		Engine           *AiDataGatewayDataSourceEngine `json:"engine"`
		Status           *AiDataGatewayDataSourceStatus `json:"status,omitempty"`
		CloudClusterId   *int64                         `json:"cloudClusterId,omitempty"`
		CloudClusterName *string                        `json:"cloudClusterName,omitempty"`
		EnvironmentName  *string                        `json:"environmentName,omitempty"`
		ConnectionConfig map[string]interface{}         `json:"connectionConfig,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "engine", "status", "cloudClusterId", "cloudClusterName", "environmentName", "connectionConfig"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Type = *all.Type
	if !all.Engine.IsValid() {
		hasInvalidField = true
	} else {
		o.Engine = *all.Engine
	}
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.CloudClusterId = all.CloudClusterId
	o.CloudClusterName = all.CloudClusterName
	o.EnvironmentName = all.EnvironmentName
	o.ConnectionConfig = all.ConnectionConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
