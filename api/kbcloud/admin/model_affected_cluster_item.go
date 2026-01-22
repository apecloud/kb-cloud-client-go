// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// AffectedClusterItem a cluster which is affected by the vulnerability
type AffectedClusterItem struct {
	// the ID of the cluster
	Id *string `json:"id,omitempty"`
	// the name of the cluster
	Name *string `json:"name,omitempty"`
	// the display name of the cluster
	DisplayName *string `json:"displayName,omitempty"`
	// the type of the cluster
	ClusterType *string `json:"clusterType,omitempty"`
	// the engine of the cluster
	Engine *string `json:"engine,omitempty"`
	// the name of the organization
	OrgName *string `json:"orgName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAffectedClusterItem instantiates a new AffectedClusterItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAffectedClusterItem() *AffectedClusterItem {
	this := AffectedClusterItem{}
	return &this
}

// NewAffectedClusterItemWithDefaults instantiates a new AffectedClusterItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAffectedClusterItemWithDefaults() *AffectedClusterItem {
	this := AffectedClusterItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffectedClusterItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedClusterItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffectedClusterItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffectedClusterItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AffectedClusterItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedClusterItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AffectedClusterItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AffectedClusterItem) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AffectedClusterItem) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedClusterItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AffectedClusterItem) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AffectedClusterItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise.
func (o *AffectedClusterItem) GetClusterType() string {
	if o == nil || o.ClusterType == nil {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedClusterItem) GetClusterTypeOk() (*string, bool) {
	if o == nil || o.ClusterType == nil {
		return nil, false
	}
	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *AffectedClusterItem) HasClusterType() bool {
	return o != nil && o.ClusterType != nil
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *AffectedClusterItem) SetClusterType(v string) {
	o.ClusterType = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *AffectedClusterItem) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedClusterItem) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *AffectedClusterItem) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *AffectedClusterItem) SetEngine(v string) {
	o.Engine = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AffectedClusterItem) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedClusterItem) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AffectedClusterItem) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AffectedClusterItem) SetOrgName(v string) {
	o.OrgName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AffectedClusterItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ClusterType != nil {
		toSerialize["clusterType"] = o.ClusterType
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AffectedClusterItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		DisplayName *string `json:"displayName,omitempty"`
		ClusterType *string `json:"clusterType,omitempty"`
		Engine      *string `json:"engine,omitempty"`
		OrgName     *string `json:"orgName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "displayName", "clusterType", "engine", "orgName"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.DisplayName = all.DisplayName
	o.ClusterType = all.ClusterType
	o.Engine = all.Engine
	o.OrgName = all.OrgName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
