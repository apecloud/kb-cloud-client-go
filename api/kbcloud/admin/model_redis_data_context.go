// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisDataContext struct {
	DatasourceName *string `json:"datasourceName,omitempty"`
	// Redis ACL user name. Prefer acl.user for new clients.
	AclUser *string          `json:"aclUser,omitempty"`
	Acl     *RedisACLContext `json:"acl,omitempty"`
	// Current logical Redis database index for this request context.
	DbIndex *int64 `json:"dbIndex,omitempty"`
	// Available logical Redis database count. Non-cluster Redis defaults to 16 when the server does not expose a count; Redis Cluster is always 1.
	DbCount *int64 `json:"dbCount,omitempty"`
	// Redis architecture mode, such as standalone, replication, sentinel, or cluster
	Mode         *string                `json:"mode,omitempty"`
	ReadOnly     *bool                  `json:"readOnly,omitempty"`
	Capabilities *RedisDataCapabilities `json:"capabilities,omitempty"`
	Cluster      *RedisClusterState     `json:"cluster,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisDataContext instantiates a new RedisDataContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisDataContext() *RedisDataContext {
	this := RedisDataContext{}
	return &this
}

// NewRedisDataContextWithDefaults instantiates a new RedisDataContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisDataContextWithDefaults() *RedisDataContext {
	this := RedisDataContext{}
	return &this
}

// GetDatasourceName returns the DatasourceName field value if set, zero value otherwise.
func (o *RedisDataContext) GetDatasourceName() string {
	if o == nil || o.DatasourceName == nil {
		var ret string
		return ret
	}
	return *o.DatasourceName
}

// GetDatasourceNameOk returns a tuple with the DatasourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetDatasourceNameOk() (*string, bool) {
	if o == nil || o.DatasourceName == nil {
		return nil, false
	}
	return o.DatasourceName, true
}

// HasDatasourceName returns a boolean if a field has been set.
func (o *RedisDataContext) HasDatasourceName() bool {
	return o != nil && o.DatasourceName != nil
}

// SetDatasourceName gets a reference to the given string and assigns it to the DatasourceName field.
func (o *RedisDataContext) SetDatasourceName(v string) {
	o.DatasourceName = &v
}

// GetAclUser returns the AclUser field value if set, zero value otherwise.
func (o *RedisDataContext) GetAclUser() string {
	if o == nil || o.AclUser == nil {
		var ret string
		return ret
	}
	return *o.AclUser
}

// GetAclUserOk returns a tuple with the AclUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetAclUserOk() (*string, bool) {
	if o == nil || o.AclUser == nil {
		return nil, false
	}
	return o.AclUser, true
}

// HasAclUser returns a boolean if a field has been set.
func (o *RedisDataContext) HasAclUser() bool {
	return o != nil && o.AclUser != nil
}

// SetAclUser gets a reference to the given string and assigns it to the AclUser field.
func (o *RedisDataContext) SetAclUser(v string) {
	o.AclUser = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *RedisDataContext) GetAcl() RedisACLContext {
	if o == nil || o.Acl == nil {
		var ret RedisACLContext
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetAclOk() (*RedisACLContext, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *RedisDataContext) HasAcl() bool {
	return o != nil && o.Acl != nil
}

// SetAcl gets a reference to the given RedisACLContext and assigns it to the Acl field.
func (o *RedisDataContext) SetAcl(v RedisACLContext) {
	o.Acl = &v
}

// GetDbIndex returns the DbIndex field value if set, zero value otherwise.
func (o *RedisDataContext) GetDbIndex() int64 {
	if o == nil || o.DbIndex == nil {
		var ret int64
		return ret
	}
	return *o.DbIndex
}

// GetDbIndexOk returns a tuple with the DbIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetDbIndexOk() (*int64, bool) {
	if o == nil || o.DbIndex == nil {
		return nil, false
	}
	return o.DbIndex, true
}

// HasDbIndex returns a boolean if a field has been set.
func (o *RedisDataContext) HasDbIndex() bool {
	return o != nil && o.DbIndex != nil
}

// SetDbIndex gets a reference to the given int64 and assigns it to the DbIndex field.
func (o *RedisDataContext) SetDbIndex(v int64) {
	o.DbIndex = &v
}

// GetDbCount returns the DbCount field value if set, zero value otherwise.
func (o *RedisDataContext) GetDbCount() int64 {
	if o == nil || o.DbCount == nil {
		var ret int64
		return ret
	}
	return *o.DbCount
}

// GetDbCountOk returns a tuple with the DbCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetDbCountOk() (*int64, bool) {
	if o == nil || o.DbCount == nil {
		return nil, false
	}
	return o.DbCount, true
}

// HasDbCount returns a boolean if a field has been set.
func (o *RedisDataContext) HasDbCount() bool {
	return o != nil && o.DbCount != nil
}

// SetDbCount gets a reference to the given int64 and assigns it to the DbCount field.
func (o *RedisDataContext) SetDbCount(v int64) {
	o.DbCount = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RedisDataContext) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RedisDataContext) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RedisDataContext) SetMode(v string) {
	o.Mode = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *RedisDataContext) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *RedisDataContext) HasReadOnly() bool {
	return o != nil && o.ReadOnly != nil
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *RedisDataContext) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *RedisDataContext) GetCapabilities() RedisDataCapabilities {
	if o == nil || o.Capabilities == nil {
		var ret RedisDataCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetCapabilitiesOk() (*RedisDataCapabilities, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *RedisDataContext) HasCapabilities() bool {
	return o != nil && o.Capabilities != nil
}

// SetCapabilities gets a reference to the given RedisDataCapabilities and assigns it to the Capabilities field.
func (o *RedisDataContext) SetCapabilities(v RedisDataCapabilities) {
	o.Capabilities = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *RedisDataContext) GetCluster() RedisClusterState {
	if o == nil || o.Cluster == nil {
		var ret RedisClusterState
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataContext) GetClusterOk() (*RedisClusterState, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *RedisDataContext) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given RedisClusterState and assigns it to the Cluster field.
func (o *RedisDataContext) SetCluster(v RedisClusterState) {
	o.Cluster = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisDataContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatasourceName != nil {
		toSerialize["datasourceName"] = o.DatasourceName
	}
	if o.AclUser != nil {
		toSerialize["aclUser"] = o.AclUser
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.DbIndex != nil {
		toSerialize["dbIndex"] = o.DbIndex
	}
	if o.DbCount != nil {
		toSerialize["dbCount"] = o.DbCount
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisDataContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasourceName *string                `json:"datasourceName,omitempty"`
		AclUser        *string                `json:"aclUser,omitempty"`
		Acl            *RedisACLContext       `json:"acl,omitempty"`
		DbIndex        *int64                 `json:"dbIndex,omitempty"`
		DbCount        *int64                 `json:"dbCount,omitempty"`
		Mode           *string                `json:"mode,omitempty"`
		ReadOnly       *bool                  `json:"readOnly,omitempty"`
		Capabilities   *RedisDataCapabilities `json:"capabilities,omitempty"`
		Cluster        *RedisClusterState     `json:"cluster,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"datasourceName", "aclUser", "acl", "dbIndex", "dbCount", "mode", "readOnly", "capabilities", "cluster"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DatasourceName = all.DatasourceName
	o.AclUser = all.AclUser
	if all.Acl != nil && all.Acl.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Acl = all.Acl
	o.DbIndex = all.DbIndex
	o.DbCount = all.DbCount
	o.Mode = all.Mode
	o.ReadOnly = all.ReadOnly
	if all.Capabilities != nil && all.Capabilities.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Capabilities = all.Capabilities
	if all.Cluster != nil && all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = all.Cluster

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
