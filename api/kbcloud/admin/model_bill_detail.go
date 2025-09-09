// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BillDetail Task information
type BillDetail struct {
	// The timestamp of data
	DataTime *int64 `json:"dataTime,omitempty"`
	// The total price
	Price *string `json:"price,omitempty"`
	// The ID of cluster
	ClusterId common.NullableString `json:"clusterID,omitempty"`
	// The Name of cluster
	ClusterName common.NullableString `json:"clusterName,omitempty"`
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

// NewBillDetail instantiates a new BillDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBillDetail() *BillDetail {
	this := BillDetail{}
	return &this
}

// NewBillDetailWithDefaults instantiates a new BillDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBillDetailWithDefaults() *BillDetail {
	this := BillDetail{}
	return &this
}

// GetDataTime returns the DataTime field value if set, zero value otherwise.
func (o *BillDetail) GetDataTime() int64 {
	if o == nil || o.DataTime == nil {
		var ret int64
		return ret
	}
	return *o.DataTime
}

// GetDataTimeOk returns a tuple with the DataTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillDetail) GetDataTimeOk() (*int64, bool) {
	if o == nil || o.DataTime == nil {
		return nil, false
	}
	return o.DataTime, true
}

// HasDataTime returns a boolean if a field has been set.
func (o *BillDetail) HasDataTime() bool {
	return o != nil && o.DataTime != nil
}

// SetDataTime gets a reference to the given int64 and assigns it to the DataTime field.
func (o *BillDetail) SetDataTime(v int64) {
	o.DataTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillDetail) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillDetail) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillDetail) HasPrice() bool {
	return o != nil && o.Price != nil
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *BillDetail) SetPrice(v string) {
	o.Price = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetClusterId() string {
	if o == nil || o.ClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *BillDetail) HasClusterId() bool {
	return o != nil && o.ClusterId.IsSet()
}

// SetClusterId gets a reference to the given common.NullableString and assigns it to the ClusterId field.
func (o *BillDetail) SetClusterId(v string) {
	o.ClusterId.Set(&v)
}

// SetClusterIdNil sets the value for ClusterId to be an explicit nil.
func (o *BillDetail) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil.
func (o *BillDetail) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *BillDetail) HasClusterName() bool {
	return o != nil && o.ClusterName.IsSet()
}

// SetClusterName gets a reference to the given common.NullableString and assigns it to the ClusterName field.
func (o *BillDetail) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// SetClusterNameNil sets the value for ClusterName to be an explicit nil.
func (o *BillDetail) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil.
func (o *BillDetail) UnsetClusterName() {
	o.ClusterName.Unset()
}

// GetEngine returns the Engine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetEngine() string {
	if o == nil || o.Engine.Get() == nil {
		var ret string
		return ret
	}
	return *o.Engine.Get()
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engine.Get(), o.Engine.IsSet()
}

// HasEngine returns a boolean if a field has been set.
func (o *BillDetail) HasEngine() bool {
	return o != nil && o.Engine.IsSet()
}

// SetEngine gets a reference to the given common.NullableString and assigns it to the Engine field.
func (o *BillDetail) SetEngine(v string) {
	o.Engine.Set(&v)
}

// SetEngineNil sets the value for Engine to be an explicit nil.
func (o *BillDetail) SetEngineNil() {
	o.Engine.Set(nil)
}

// UnsetEngine ensures that no value is present for Engine, not even an explicit nil.
func (o *BillDetail) UnsetEngine() {
	o.Engine.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *BillDetail) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given common.NullableString and assigns it to the Version field.
func (o *BillDetail) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *BillDetail) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *BillDetail) UnsetVersion() {
	o.Version.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *BillDetail) HasMode() bool {
	return o != nil && o.Mode.IsSet()
}

// SetMode gets a reference to the given common.NullableString and assigns it to the Mode field.
func (o *BillDetail) SetMode(v string) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil.
func (o *BillDetail) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil.
func (o *BillDetail) UnsetMode() {
	o.Mode.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetOrgName() string {
	if o == nil || o.OrgName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgName.Get()
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgName.Get(), o.OrgName.IsSet()
}

// HasOrgName returns a boolean if a field has been set.
func (o *BillDetail) HasOrgName() bool {
	return o != nil && o.OrgName.IsSet()
}

// SetOrgName gets a reference to the given common.NullableString and assigns it to the OrgName field.
func (o *BillDetail) SetOrgName(v string) {
	o.OrgName.Set(&v)
}

// SetOrgNameNil sets the value for OrgName to be an explicit nil.
func (o *BillDetail) SetOrgNameNil() {
	o.OrgName.Set(nil)
}

// UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil.
func (o *BillDetail) UnsetOrgName() {
	o.OrgName.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillDetail) GetProjectName() string {
	if o == nil || o.ProjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BillDetail) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *BillDetail) HasProjectName() bool {
	return o != nil && o.ProjectName.IsSet()
}

// SetProjectName gets a reference to the given common.NullableString and assigns it to the ProjectName field.
func (o *BillDetail) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil.
func (o *BillDetail) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil.
func (o *BillDetail) UnsetProjectName() {
	o.ProjectName.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o BillDetail) MarshalJSON() ([]byte, error) {
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
func (o *BillDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataTime    *int64                `json:"dataTime,omitempty"`
		Price       *string               `json:"price,omitempty"`
		ClusterId   common.NullableString `json:"clusterID,omitempty"`
		ClusterName common.NullableString `json:"clusterName,omitempty"`
		Engine      common.NullableString `json:"engine,omitempty"`
		Version     common.NullableString `json:"version,omitempty"`
		Mode        common.NullableString `json:"mode,omitempty"`
		OrgName     common.NullableString `json:"orgName,omitempty"`
		ProjectName common.NullableString `json:"projectName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"dataTime", "price", "clusterID", "clusterName", "engine", "version", "mode", "orgName", "projectName"})
	} else {
		return err
	}
	o.DataTime = all.DataTime
	o.Price = all.Price
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
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
