// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayDataSource struct {
	Name             *string `json:"name,omitempty"`
	Type             *string `json:"type,omitempty"`
	Engine           *string `json:"engine,omitempty"`
	CloudClusterId   *int64  `json:"cloudClusterId,omitempty"`
	CloudClusterName *string `json:"cloudClusterName,omitempty"`
	EnvironmentName  *string `json:"environmentName,omitempty"`
	// All datasource connection information. For external databases this comes from API input; for Cloud/KubeBlocks managed datasources Cloud resolves and fills the connection information. Sensitive keys such as password, token, secret, privateKey, accessKey, and credential are encrypted at rest and masked in user-facing responses. Internal Runtime config resolves the decrypted view only inside the trusted server boundary.
	ConnectionConfig map[string]interface{} `json:"connectionConfig,omitempty"`
	DatasourceId     *string                `json:"datasourceId,omitempty"`
	GatewayId        *string                `json:"gatewayId,omitempty"`
	OrgName          *string                `json:"orgName,omitempty"`
	Status           *string                `json:"status,omitempty"`
	LastTestStatus   *string                `json:"lastTestStatus,omitempty"`
	LastTestMessage  *string                `json:"lastTestMessage,omitempty"`
	LastTestedAt     *time.Time             `json:"lastTestedAt,omitempty"`
	CreatedAt        *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time             `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayDataSource instantiates a new AiDataGatewayDataSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayDataSource() *AiDataGatewayDataSource {
	this := AiDataGatewayDataSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiDataGatewayDataSource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AiDataGatewayDataSource) SetType(v string) {
	o.Type = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *AiDataGatewayDataSource) SetEngine(v string) {
	o.Engine = &v
}

// GetCloudClusterId returns the CloudClusterId field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetCloudClusterId() int64 {
	if o == nil || o.CloudClusterId == nil {
		var ret int64
		return ret
	}
	return *o.CloudClusterId
}

// GetCloudClusterIdOk returns a tuple with the CloudClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetCloudClusterIdOk() (*int64, bool) {
	if o == nil || o.CloudClusterId == nil {
		return nil, false
	}
	return o.CloudClusterId, true
}

// HasCloudClusterId returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasCloudClusterId() bool {
	return o != nil && o.CloudClusterId != nil
}

// SetCloudClusterId gets a reference to the given int64 and assigns it to the CloudClusterId field.
func (o *AiDataGatewayDataSource) SetCloudClusterId(v int64) {
	o.CloudClusterId = &v
}

// GetCloudClusterName returns the CloudClusterName field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetCloudClusterName() string {
	if o == nil || o.CloudClusterName == nil {
		var ret string
		return ret
	}
	return *o.CloudClusterName
}

// GetCloudClusterNameOk returns a tuple with the CloudClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetCloudClusterNameOk() (*string, bool) {
	if o == nil || o.CloudClusterName == nil {
		return nil, false
	}
	return o.CloudClusterName, true
}

// HasCloudClusterName returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasCloudClusterName() bool {
	return o != nil && o.CloudClusterName != nil
}

// SetCloudClusterName gets a reference to the given string and assigns it to the CloudClusterName field.
func (o *AiDataGatewayDataSource) SetCloudClusterName(v string) {
	o.CloudClusterName = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *AiDataGatewayDataSource) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetConnectionConfig returns the ConnectionConfig field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetConnectionConfig() map[string]interface{} {
	if o == nil || o.ConnectionConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectionConfig
}

// GetConnectionConfigOk returns a tuple with the ConnectionConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetConnectionConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.ConnectionConfig == nil {
		return nil, false
	}
	return &o.ConnectionConfig, true
}

// HasConnectionConfig returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasConnectionConfig() bool {
	return o != nil && o.ConnectionConfig != nil
}

// SetConnectionConfig gets a reference to the given map[string]interface{} and assigns it to the ConnectionConfig field.
func (o *AiDataGatewayDataSource) SetConnectionConfig(v map[string]interface{}) {
	o.ConnectionConfig = v
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetDatasourceId() string {
	if o == nil || o.DatasourceId == nil {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetDatasourceIdOk() (*string, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *AiDataGatewayDataSource) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGatewayDataSource) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGatewayDataSource) SetOrgName(v string) {
	o.OrgName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGatewayDataSource) SetStatus(v string) {
	o.Status = &v
}

// GetLastTestStatus returns the LastTestStatus field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetLastTestStatus() string {
	if o == nil || o.LastTestStatus == nil {
		var ret string
		return ret
	}
	return *o.LastTestStatus
}

// GetLastTestStatusOk returns a tuple with the LastTestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetLastTestStatusOk() (*string, bool) {
	if o == nil || o.LastTestStatus == nil {
		return nil, false
	}
	return o.LastTestStatus, true
}

// HasLastTestStatus returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasLastTestStatus() bool {
	return o != nil && o.LastTestStatus != nil
}

// SetLastTestStatus gets a reference to the given string and assigns it to the LastTestStatus field.
func (o *AiDataGatewayDataSource) SetLastTestStatus(v string) {
	o.LastTestStatus = &v
}

// GetLastTestMessage returns the LastTestMessage field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetLastTestMessage() string {
	if o == nil || o.LastTestMessage == nil {
		var ret string
		return ret
	}
	return *o.LastTestMessage
}

// GetLastTestMessageOk returns a tuple with the LastTestMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetLastTestMessageOk() (*string, bool) {
	if o == nil || o.LastTestMessage == nil {
		return nil, false
	}
	return o.LastTestMessage, true
}

// HasLastTestMessage returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasLastTestMessage() bool {
	return o != nil && o.LastTestMessage != nil
}

// SetLastTestMessage gets a reference to the given string and assigns it to the LastTestMessage field.
func (o *AiDataGatewayDataSource) SetLastTestMessage(v string) {
	o.LastTestMessage = &v
}

// GetLastTestedAt returns the LastTestedAt field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetLastTestedAt() time.Time {
	if o == nil || o.LastTestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTestedAt
}

// GetLastTestedAtOk returns a tuple with the LastTestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetLastTestedAtOk() (*time.Time, bool) {
	if o == nil || o.LastTestedAt == nil {
		return nil, false
	}
	return o.LastTestedAt, true
}

// HasLastTestedAt returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasLastTestedAt() bool {
	return o != nil && o.LastTestedAt != nil
}

// SetLastTestedAt gets a reference to the given time.Time and assigns it to the LastTestedAt field.
func (o *AiDataGatewayDataSource) SetLastTestedAt(v time.Time) {
	o.LastTestedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGatewayDataSource) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayDataSource) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayDataSource) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayDataSource) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiDataGatewayDataSource) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
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
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastTestStatus != nil {
		toSerialize["lastTestStatus"] = o.LastTestStatus
	}
	if o.LastTestMessage != nil {
		toSerialize["lastTestMessage"] = o.LastTestMessage
	}
	if o.LastTestedAt != nil {
		if o.LastTestedAt.Nanosecond() == 0 {
			toSerialize["lastTestedAt"] = o.LastTestedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastTestedAt"] = o.LastTestedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayDataSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string                `json:"name,omitempty"`
		Type             *string                `json:"type,omitempty"`
		Engine           *string                `json:"engine,omitempty"`
		CloudClusterId   *int64                 `json:"cloudClusterId,omitempty"`
		CloudClusterName *string                `json:"cloudClusterName,omitempty"`
		EnvironmentName  *string                `json:"environmentName,omitempty"`
		ConnectionConfig map[string]interface{} `json:"connectionConfig,omitempty"`
		DatasourceId     *string                `json:"datasourceId,omitempty"`
		GatewayId        *string                `json:"gatewayId,omitempty"`
		OrgName          *string                `json:"orgName,omitempty"`
		Status           *string                `json:"status,omitempty"`
		LastTestStatus   *string                `json:"lastTestStatus,omitempty"`
		LastTestMessage  *string                `json:"lastTestMessage,omitempty"`
		LastTestedAt     *time.Time             `json:"lastTestedAt,omitempty"`
		CreatedAt        *time.Time             `json:"createdAt,omitempty"`
		UpdatedAt        *time.Time             `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "engine", "cloudClusterId", "cloudClusterName", "environmentName", "connectionConfig", "datasourceId", "gatewayId", "orgName", "status", "lastTestStatus", "lastTestMessage", "lastTestedAt", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.Engine = all.Engine
	o.CloudClusterId = all.CloudClusterId
	o.CloudClusterName = all.CloudClusterName
	o.EnvironmentName = all.EnvironmentName
	o.ConnectionConfig = all.ConnectionConfig
	o.DatasourceId = all.DatasourceId
	o.GatewayId = all.GatewayId
	o.OrgName = all.OrgName
	o.Status = all.Status
	o.LastTestStatus = all.LastTestStatus
	o.LastTestMessage = all.LastTestMessage
	o.LastTestedAt = all.LastTestedAt
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
