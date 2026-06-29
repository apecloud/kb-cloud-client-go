// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisDataCapabilities struct {
	Analysis  *RedisCapability `json:"analysis,omitempty"`
	Browser   *RedisCapability `json:"browser,omitempty"`
	Database  *RedisCapability `json:"database,omitempty"`
	Workbench *RedisCapability `json:"workbench,omitempty"`
	Pubsub    *RedisCapability `json:"pubsub,omitempty"`
	SlowLog   *RedisCapability `json:"slowLog,omitempty"`
	Info      *RedisCapability `json:"info,omitempty"`
	Cluster   *RedisCapability `json:"cluster,omitempty"`
	Write     *RedisCapability `json:"write,omitempty"`
	Delete    *RedisCapability `json:"delete,omitempty"`
	Json      *RedisCapability `json:"json,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisDataCapabilities instantiates a new RedisDataCapabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisDataCapabilities() *RedisDataCapabilities {
	this := RedisDataCapabilities{}
	return &this
}

// NewRedisDataCapabilitiesWithDefaults instantiates a new RedisDataCapabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisDataCapabilitiesWithDefaults() *RedisDataCapabilities {
	this := RedisDataCapabilities{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetAnalysis() RedisCapability {
	if o == nil || o.Analysis == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetAnalysisOk() (*RedisCapability, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasAnalysis() bool {
	return o != nil && o.Analysis != nil
}

// SetAnalysis gets a reference to the given RedisCapability and assigns it to the Analysis field.
func (o *RedisDataCapabilities) SetAnalysis(v RedisCapability) {
	o.Analysis = &v
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetBrowser() RedisCapability {
	if o == nil || o.Browser == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetBrowserOk() (*RedisCapability, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasBrowser() bool {
	return o != nil && o.Browser != nil
}

// SetBrowser gets a reference to the given RedisCapability and assigns it to the Browser field.
func (o *RedisDataCapabilities) SetBrowser(v RedisCapability) {
	o.Browser = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetDatabase() RedisCapability {
	if o == nil || o.Database == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetDatabaseOk() (*RedisCapability, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given RedisCapability and assigns it to the Database field.
func (o *RedisDataCapabilities) SetDatabase(v RedisCapability) {
	o.Database = &v
}

// GetWorkbench returns the Workbench field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetWorkbench() RedisCapability {
	if o == nil || o.Workbench == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Workbench
}

// GetWorkbenchOk returns a tuple with the Workbench field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetWorkbenchOk() (*RedisCapability, bool) {
	if o == nil || o.Workbench == nil {
		return nil, false
	}
	return o.Workbench, true
}

// HasWorkbench returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasWorkbench() bool {
	return o != nil && o.Workbench != nil
}

// SetWorkbench gets a reference to the given RedisCapability and assigns it to the Workbench field.
func (o *RedisDataCapabilities) SetWorkbench(v RedisCapability) {
	o.Workbench = &v
}

// GetPubsub returns the Pubsub field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetPubsub() RedisCapability {
	if o == nil || o.Pubsub == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Pubsub
}

// GetPubsubOk returns a tuple with the Pubsub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetPubsubOk() (*RedisCapability, bool) {
	if o == nil || o.Pubsub == nil {
		return nil, false
	}
	return o.Pubsub, true
}

// HasPubsub returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasPubsub() bool {
	return o != nil && o.Pubsub != nil
}

// SetPubsub gets a reference to the given RedisCapability and assigns it to the Pubsub field.
func (o *RedisDataCapabilities) SetPubsub(v RedisCapability) {
	o.Pubsub = &v
}

// GetSlowLog returns the SlowLog field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetSlowLog() RedisCapability {
	if o == nil || o.SlowLog == nil {
		var ret RedisCapability
		return ret
	}
	return *o.SlowLog
}

// GetSlowLogOk returns a tuple with the SlowLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetSlowLogOk() (*RedisCapability, bool) {
	if o == nil || o.SlowLog == nil {
		return nil, false
	}
	return o.SlowLog, true
}

// HasSlowLog returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasSlowLog() bool {
	return o != nil && o.SlowLog != nil
}

// SetSlowLog gets a reference to the given RedisCapability and assigns it to the SlowLog field.
func (o *RedisDataCapabilities) SetSlowLog(v RedisCapability) {
	o.SlowLog = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetInfo() RedisCapability {
	if o == nil || o.Info == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetInfoOk() (*RedisCapability, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasInfo() bool {
	return o != nil && o.Info != nil
}

// SetInfo gets a reference to the given RedisCapability and assigns it to the Info field.
func (o *RedisDataCapabilities) SetInfo(v RedisCapability) {
	o.Info = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetCluster() RedisCapability {
	if o == nil || o.Cluster == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetClusterOk() (*RedisCapability, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given RedisCapability and assigns it to the Cluster field.
func (o *RedisDataCapabilities) SetCluster(v RedisCapability) {
	o.Cluster = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetWrite() RedisCapability {
	if o == nil || o.Write == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetWriteOk() (*RedisCapability, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasWrite() bool {
	return o != nil && o.Write != nil
}

// SetWrite gets a reference to the given RedisCapability and assigns it to the Write field.
func (o *RedisDataCapabilities) SetWrite(v RedisCapability) {
	o.Write = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetDelete() RedisCapability {
	if o == nil || o.Delete == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetDeleteOk() (*RedisCapability, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasDelete() bool {
	return o != nil && o.Delete != nil
}

// SetDelete gets a reference to the given RedisCapability and assigns it to the Delete field.
func (o *RedisDataCapabilities) SetDelete(v RedisCapability) {
	o.Delete = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RedisDataCapabilities) GetJson() RedisCapability {
	if o == nil || o.Json == nil {
		var ret RedisCapability
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDataCapabilities) GetJsonOk() (*RedisCapability, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RedisDataCapabilities) HasJson() bool {
	return o != nil && o.Json != nil
}

// SetJson gets a reference to the given RedisCapability and assigns it to the Json field.
func (o *RedisDataCapabilities) SetJson(v RedisCapability) {
	o.Json = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisDataCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Browser != nil {
		toSerialize["browser"] = o.Browser
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Workbench != nil {
		toSerialize["workbench"] = o.Workbench
	}
	if o.Pubsub != nil {
		toSerialize["pubsub"] = o.Pubsub
	}
	if o.SlowLog != nil {
		toSerialize["slowLog"] = o.SlowLog
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisDataCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Analysis  *RedisCapability `json:"analysis,omitempty"`
		Browser   *RedisCapability `json:"browser,omitempty"`
		Database  *RedisCapability `json:"database,omitempty"`
		Workbench *RedisCapability `json:"workbench,omitempty"`
		Pubsub    *RedisCapability `json:"pubsub,omitempty"`
		SlowLog   *RedisCapability `json:"slowLog,omitempty"`
		Info      *RedisCapability `json:"info,omitempty"`
		Cluster   *RedisCapability `json:"cluster,omitempty"`
		Write     *RedisCapability `json:"write,omitempty"`
		Delete    *RedisCapability `json:"delete,omitempty"`
		Json      *RedisCapability `json:"json,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"analysis", "browser", "database", "workbench", "pubsub", "slowLog", "info", "cluster", "write", "delete", "json"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Analysis != nil && all.Analysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Analysis = all.Analysis
	if all.Browser != nil && all.Browser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Browser = all.Browser
	if all.Database != nil && all.Database.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Database = all.Database
	if all.Workbench != nil && all.Workbench.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Workbench = all.Workbench
	if all.Pubsub != nil && all.Pubsub.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pubsub = all.Pubsub
	if all.SlowLog != nil && all.SlowLog.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SlowLog = all.SlowLog
	if all.Info != nil && all.Info.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Info = all.Info
	if all.Cluster != nil && all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = all.Cluster
	if all.Write != nil && all.Write.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Write = all.Write
	if all.Delete != nil && all.Delete.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Delete = all.Delete
	if all.Json != nil && all.Json.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Json = all.Json

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
