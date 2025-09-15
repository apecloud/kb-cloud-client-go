// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// Bill Task information
type Bill struct {
	// The timestamp of data
	DataTime *int64 `json:"dataTime,omitempty"`
	// The total price
	Price *string `json:"price,omitempty"`
	// The ID of cluster
	ClusterId common.NullableString `json:"clusterID,omitempty"`
	// The Name of cluster
	ClusterName common.NullableString `json:"clusterName,omitempty"`
	// The Display Name of cluster
	ClusterDisplayName common.NullableString `json:"clusterDisplayName,omitempty"`
	// Cluster Application Engine
	Engine common.NullableString `json:"engine,omitempty"`
	// Cluster Application Version
	Version common.NullableString `json:"version,omitempty"`
	// Cluster topology mode
	Mode common.NullableString `json:"mode,omitempty"`
	// The Name of organization
	OrgName common.NullableString `json:"orgName,omitempty"`
	// The Name of project
	ProjectName common.NullableString `json:"projectName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBill instantiates a new Bill object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBill() *Bill {
	this := Bill{}
	return &this
}

// NewBillWithDefaults instantiates a new Bill object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBillWithDefaults() *Bill {
	this := Bill{}
	return &this
}

// GetDataTime returns the DataTime field value if set, zero value otherwise.
func (o *Bill) GetDataTime() int64 {
	if o == nil || o.DataTime == nil {
		var ret int64
		return ret
	}
	return *o.DataTime
}

// GetDataTimeOk returns a tuple with the DataTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetDataTimeOk() (*int64, bool) {
	if o == nil || o.DataTime == nil {
		return nil, false
	}
	return o.DataTime, true
}

// HasDataTime returns a boolean if a field has been set.
func (o *Bill) HasDataTime() bool {
	return o != nil && o.DataTime != nil
}

// SetDataTime gets a reference to the given int64 and assigns it to the DataTime field.
func (o *Bill) SetDataTime(v int64) {
	o.DataTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Bill) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Bill) HasPrice() bool {
	return o != nil && o.Price != nil
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *Bill) SetPrice(v string) {
	o.Price = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetClusterId() string {
	if o == nil || o.ClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *Bill) HasClusterId() bool {
	return o != nil && o.ClusterId.IsSet()
}

// SetClusterId gets a reference to the given common.NullableString and assigns it to the ClusterId field.
func (o *Bill) SetClusterId(v string) {
	o.ClusterId.Set(&v)
}

// SetClusterIdNil sets the value for ClusterId to be an explicit nil.
func (o *Bill) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil.
func (o *Bill) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *Bill) HasClusterName() bool {
	return o != nil && o.ClusterName.IsSet()
}

// SetClusterName gets a reference to the given common.NullableString and assigns it to the ClusterName field.
func (o *Bill) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// SetClusterNameNil sets the value for ClusterName to be an explicit nil.
func (o *Bill) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil.
func (o *Bill) UnsetClusterName() {
	o.ClusterName.Unset()
}

// GetClusterDisplayName returns the ClusterDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetClusterDisplayName() string {
	if o == nil || o.ClusterDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterDisplayName.Get()
}

// GetClusterDisplayNameOk returns a tuple with the ClusterDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetClusterDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterDisplayName.Get(), o.ClusterDisplayName.IsSet()
}

// HasClusterDisplayName returns a boolean if a field has been set.
func (o *Bill) HasClusterDisplayName() bool {
	return o != nil && o.ClusterDisplayName.IsSet()
}

// SetClusterDisplayName gets a reference to the given common.NullableString and assigns it to the ClusterDisplayName field.
func (o *Bill) SetClusterDisplayName(v string) {
	o.ClusterDisplayName.Set(&v)
}

// SetClusterDisplayNameNil sets the value for ClusterDisplayName to be an explicit nil.
func (o *Bill) SetClusterDisplayNameNil() {
	o.ClusterDisplayName.Set(nil)
}

// UnsetClusterDisplayName ensures that no value is present for ClusterDisplayName, not even an explicit nil.
func (o *Bill) UnsetClusterDisplayName() {
	o.ClusterDisplayName.Unset()
}

// GetEngine returns the Engine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetEngine() string {
	if o == nil || o.Engine.Get() == nil {
		var ret string
		return ret
	}
	return *o.Engine.Get()
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engine.Get(), o.Engine.IsSet()
}

// HasEngine returns a boolean if a field has been set.
func (o *Bill) HasEngine() bool {
	return o != nil && o.Engine.IsSet()
}

// SetEngine gets a reference to the given common.NullableString and assigns it to the Engine field.
func (o *Bill) SetEngine(v string) {
	o.Engine.Set(&v)
}

// SetEngineNil sets the value for Engine to be an explicit nil.
func (o *Bill) SetEngineNil() {
	o.Engine.Set(nil)
}

// UnsetEngine ensures that no value is present for Engine, not even an explicit nil.
func (o *Bill) UnsetEngine() {
	o.Engine.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *Bill) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given common.NullableString and assigns it to the Version field.
func (o *Bill) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *Bill) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *Bill) UnsetVersion() {
	o.Version.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *Bill) HasMode() bool {
	return o != nil && o.Mode.IsSet()
}

// SetMode gets a reference to the given common.NullableString and assigns it to the Mode field.
func (o *Bill) SetMode(v string) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil.
func (o *Bill) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil.
func (o *Bill) UnsetMode() {
	o.Mode.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetOrgName() string {
	if o == nil || o.OrgName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgName.Get()
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgName.Get(), o.OrgName.IsSet()
}

// HasOrgName returns a boolean if a field has been set.
func (o *Bill) HasOrgName() bool {
	return o != nil && o.OrgName.IsSet()
}

// SetOrgName gets a reference to the given common.NullableString and assigns it to the OrgName field.
func (o *Bill) SetOrgName(v string) {
	o.OrgName.Set(&v)
}

// SetOrgNameNil sets the value for OrgName to be an explicit nil.
func (o *Bill) SetOrgNameNil() {
	o.OrgName.Set(nil)
}

// UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil.
func (o *Bill) UnsetOrgName() {
	o.OrgName.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bill) GetProjectName() string {
	if o == nil || o.ProjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Bill) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *Bill) HasProjectName() bool {
	return o != nil && o.ProjectName.IsSet()
}

// SetProjectName gets a reference to the given common.NullableString and assigns it to the ProjectName field.
func (o *Bill) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil.
func (o *Bill) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil.
func (o *Bill) UnsetProjectName() {
	o.ProjectName.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o Bill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DataTime != nil {
		toSerialize["dataTime"] = o.DataTime
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.ClusterId.IsSet() {
		toSerialize["clusterID"] = o.ClusterId.Get()
	}
	if o.ClusterName.IsSet() {
		toSerialize["clusterName"] = o.ClusterName.Get()
	}
	if o.ClusterDisplayName.IsSet() {
		toSerialize["clusterDisplayName"] = o.ClusterDisplayName.Get()
	}
	if o.Engine.IsSet() {
		toSerialize["engine"] = o.Engine.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if o.OrgName.IsSet() {
		toSerialize["orgName"] = o.OrgName.Get()
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Bill) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataTime           *int64                `json:"dataTime,omitempty"`
		Price              *string               `json:"price,omitempty"`
		ClusterId          common.NullableString `json:"clusterID,omitempty"`
		ClusterName        common.NullableString `json:"clusterName,omitempty"`
		ClusterDisplayName common.NullableString `json:"clusterDisplayName,omitempty"`
		Engine             common.NullableString `json:"engine,omitempty"`
		Version            common.NullableString `json:"version,omitempty"`
		Mode               common.NullableString `json:"mode,omitempty"`
		OrgName            common.NullableString `json:"orgName,omitempty"`
		ProjectName        common.NullableString `json:"projectName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"dataTime", "price", "clusterID", "clusterName", "clusterDisplayName", "engine", "version", "mode", "orgName", "projectName"})
	} else {
		return err
	}
	o.DataTime = all.DataTime
	o.Price = all.Price
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.ClusterDisplayName = all.ClusterDisplayName
	o.Engine = all.Engine
	o.Version = all.Version
	o.Mode = all.Mode
	o.OrgName = all.OrgName
	o.ProjectName = all.ProjectName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
