// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ChatRequest 聊天消息请求
type ChatRequest struct {
	// 组织名称
	OrgName string `json:"orgName"`
	// 集群名称
	ClusterName string `json:"clusterName"`
	// 会话ID
	SessionId string `json:"sessionId"`
	// 消息ID
	MessageId string `json:"messageID"`
	// 父消息ID
	ParentId string `json:"parentID"`
	// 问题内容
	Query string `json:"query"`
	// 查询类型
	QueryType QueryType `json:"queryType"`
	// LLM模型
	Model *string `json:"model,omitempty"`
	// 创建时间
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatRequest instantiates a new ChatRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatRequest(orgName string, clusterName string, sessionId string, messageId string, parentId string, query string, queryType QueryType) *ChatRequest {
	this := ChatRequest{}
	this.OrgName = orgName
	this.ClusterName = clusterName
	this.SessionId = sessionId
	this.MessageId = messageId
	this.ParentId = parentId
	this.Query = query
	this.QueryType = queryType
	return &this
}

// NewChatRequestWithDefaults instantiates a new ChatRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatRequestWithDefaults() *ChatRequest {
	this := ChatRequest{}
	return &this
}

// GetOrgName returns the OrgName field value.
func (o *ChatRequest) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *ChatRequest) SetOrgName(v string) {
	o.OrgName = v
}

// GetClusterName returns the ClusterName field value.
func (o *ChatRequest) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *ChatRequest) SetClusterName(v string) {
	o.ClusterName = v
}

// GetSessionId returns the SessionId field value.
func (o *ChatRequest) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *ChatRequest) SetSessionId(v string) {
	o.SessionId = v
}

// GetMessageId returns the MessageId field value.
func (o *ChatRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value.
func (o *ChatRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetParentId returns the ParentId field value.
func (o *ChatRequest) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value.
func (o *ChatRequest) SetParentId(v string) {
	o.ParentId = v
}

// GetQuery returns the Query field value.
func (o *ChatRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ChatRequest) SetQuery(v string) {
	o.Query = v
}

// GetQueryType returns the QueryType field value.
func (o *ChatRequest) GetQueryType() QueryType {
	if o == nil {
		var ret QueryType
		return ret
	}
	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetQueryTypeOk() (*QueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value.
func (o *ChatRequest) SetQueryType(v QueryType) {
	o.QueryType = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ChatRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ChatRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ChatRequest) SetModel(v string) {
	o.Model = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChatRequest) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRequest) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChatRequest) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ChatRequest) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgName"] = o.OrgName
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["sessionId"] = o.SessionId
	toSerialize["messageID"] = o.MessageId
	toSerialize["parentID"] = o.ParentId
	toSerialize["query"] = o.Query
	toSerialize["queryType"] = o.QueryType
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName     *string    `json:"orgName"`
		ClusterName *string    `json:"clusterName"`
		SessionId   *string    `json:"sessionId"`
		MessageId   *string    `json:"messageID"`
		ParentId    *string    `json:"parentID"`
		Query       *string    `json:"query"`
		QueryType   *QueryType `json:"queryType"`
		Model       *string    `json:"model,omitempty"`
		CreatedAt   *int64     `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field sessionId missing")
	}
	if all.MessageId == nil {
		return fmt.Errorf("required field messageID missing")
	}
	if all.ParentId == nil {
		return fmt.Errorf("required field parentID missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.QueryType == nil {
		return fmt.Errorf("required field queryType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "clusterName", "sessionId", "messageID", "parentID", "query", "queryType", "model", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = *all.OrgName
	o.ClusterName = *all.ClusterName
	o.SessionId = *all.SessionId
	o.MessageId = *all.MessageId
	o.ParentId = *all.ParentId
	o.Query = *all.Query
	if !all.QueryType.IsValid() {
		hasInvalidField = true
	} else {
		o.QueryType = *all.QueryType
	}
	o.Model = all.Model
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
