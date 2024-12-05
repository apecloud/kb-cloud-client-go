// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MonitorDataSinkUpdate External Endpoint update option
type MonitorDataSinkUpdate struct {
	// type of monitor data sink(logs or metrics)
	MonitorDataSinkType common.NullableString `json:"monitorDataSinkType"`
	// Name for relative Environment
	EnvironmentName string `json:"environmentName"`
	// url for exporter(eg. url for ElasticSearch)
	ExporterUrl *string `json:"exporterUrl,omitempty"`
	// Endpoint name
	Name common.NullableString `json:"name,omitempty"`
	// Username
	Username common.NullableString `json:"username,omitempty"`
	// Password
	Password common.NullableString `json:"password,omitempty"`
	// apiKey
	ApiKey common.NullableString `json:"apiKey,omitempty"`
	// indexName for ElasticSearch
	IndexName common.NullableString `json:"indexName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorDataSinkUpdate instantiates a new MonitorDataSinkUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDataSinkUpdate(monitorDataSinkType common.NullableString, environmentName string) *MonitorDataSinkUpdate {
	this := MonitorDataSinkUpdate{}
	this.MonitorDataSinkType = monitorDataSinkType
	this.EnvironmentName = environmentName
	return &this
}

// NewMonitorDataSinkUpdateWithDefaults instantiates a new MonitorDataSinkUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDataSinkUpdateWithDefaults() *MonitorDataSinkUpdate {
	this := MonitorDataSinkUpdate{}
	return &this
}

// GetMonitorDataSinkType returns the MonitorDataSinkType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MonitorDataSinkUpdate) GetMonitorDataSinkType() string {
	if o == nil || o.MonitorDataSinkType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MonitorDataSinkType.Get()
}

// GetMonitorDataSinkTypeOk returns a tuple with the MonitorDataSinkType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkUpdate) GetMonitorDataSinkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitorDataSinkType.Get(), o.MonitorDataSinkType.IsSet()
}

// SetMonitorDataSinkType sets field value.
func (o *MonitorDataSinkUpdate) SetMonitorDataSinkType(v string) {
	o.MonitorDataSinkType.Set(&v)
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *MonitorDataSinkUpdate) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSinkUpdate) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *MonitorDataSinkUpdate) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetExporterUrl returns the ExporterUrl field value if set, zero value otherwise.
func (o *MonitorDataSinkUpdate) GetExporterUrl() string {
	if o == nil || o.ExporterUrl == nil {
		var ret string
		return ret
	}
	return *o.ExporterUrl
}

// GetExporterUrlOk returns a tuple with the ExporterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDataSinkUpdate) GetExporterUrlOk() (*string, bool) {
	if o == nil || o.ExporterUrl == nil {
		return nil, false
	}
	return o.ExporterUrl, true
}

// HasExporterUrl returns a boolean if a field has been set.
func (o *MonitorDataSinkUpdate) HasExporterUrl() bool {
	return o != nil && o.ExporterUrl != nil
}

// SetExporterUrl gets a reference to the given string and assigns it to the ExporterUrl field.
func (o *MonitorDataSinkUpdate) SetExporterUrl(v string) {
	o.ExporterUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkUpdate) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MonitorDataSinkUpdate) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *MonitorDataSinkUpdate) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *MonitorDataSinkUpdate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *MonitorDataSinkUpdate) UnsetName() {
	o.Name.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkUpdate) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkUpdate) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *MonitorDataSinkUpdate) HasUsername() bool {
	return o != nil && o.Username.IsSet()
}

// SetUsername gets a reference to the given common.NullableString and assigns it to the Username field.
func (o *MonitorDataSinkUpdate) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil.
func (o *MonitorDataSinkUpdate) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil.
func (o *MonitorDataSinkUpdate) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkUpdate) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkUpdate) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MonitorDataSinkUpdate) HasPassword() bool {
	return o != nil && o.Password.IsSet()
}

// SetPassword gets a reference to the given common.NullableString and assigns it to the Password field.
func (o *MonitorDataSinkUpdate) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil.
func (o *MonitorDataSinkUpdate) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil.
func (o *MonitorDataSinkUpdate) UnsetPassword() {
	o.Password.Unset()
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkUpdate) GetApiKey() string {
	if o == nil || o.ApiKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkUpdate) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *MonitorDataSinkUpdate) HasApiKey() bool {
	return o != nil && o.ApiKey.IsSet()
}

// SetApiKey gets a reference to the given common.NullableString and assigns it to the ApiKey field.
func (o *MonitorDataSinkUpdate) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}

// SetApiKeyNil sets the value for ApiKey to be an explicit nil.
func (o *MonitorDataSinkUpdate) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil.
func (o *MonitorDataSinkUpdate) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetIndexName returns the IndexName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkUpdate) GetIndexName() string {
	if o == nil || o.IndexName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IndexName.Get()
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkUpdate) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexName.Get(), o.IndexName.IsSet()
}

// HasIndexName returns a boolean if a field has been set.
func (o *MonitorDataSinkUpdate) HasIndexName() bool {
	return o != nil && o.IndexName.IsSet()
}

// SetIndexName gets a reference to the given common.NullableString and assigns it to the IndexName field.
func (o *MonitorDataSinkUpdate) SetIndexName(v string) {
	o.IndexName.Set(&v)
}

// SetIndexNameNil sets the value for IndexName to be an explicit nil.
func (o *MonitorDataSinkUpdate) SetIndexNameNil() {
	o.IndexName.Set(nil)
}

// UnsetIndexName ensures that no value is present for IndexName, not even an explicit nil.
func (o *MonitorDataSinkUpdate) UnsetIndexName() {
	o.IndexName.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorDataSinkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["monitorDataSinkType"] = o.MonitorDataSinkType.Get()
	toSerialize["environmentName"] = o.EnvironmentName
	if o.ExporterUrl != nil {
		toSerialize["exporterUrl"] = o.ExporterUrl
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.ApiKey.IsSet() {
		toSerialize["apiKey"] = o.ApiKey.Get()
	}
	if o.IndexName.IsSet() {
		toSerialize["indexName"] = o.IndexName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorDataSinkUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MonitorDataSinkType common.NullableString `json:"monitorDataSinkType"`
		EnvironmentName     *string               `json:"environmentName"`
		ExporterUrl         *string               `json:"exporterUrl,omitempty"`
		Name                common.NullableString `json:"name,omitempty"`
		Username            common.NullableString `json:"username,omitempty"`
		Password            common.NullableString `json:"password,omitempty"`
		ApiKey              common.NullableString `json:"apiKey,omitempty"`
		IndexName           common.NullableString `json:"indexName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.MonitorDataSinkType.IsSet() {
		return fmt.Errorf("required field monitorDataSinkType missing")
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"monitorDataSinkType", "environmentName", "exporterUrl", "name", "username", "password", "apiKey", "indexName"})
	} else {
		return err
	}
	o.MonitorDataSinkType = all.MonitorDataSinkType
	o.EnvironmentName = *all.EnvironmentName
	o.ExporterUrl = all.ExporterUrl
	o.Name = all.Name
	o.Username = all.Username
	o.Password = all.Password
	o.ApiKey = all.ApiKey
	o.IndexName = all.IndexName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
