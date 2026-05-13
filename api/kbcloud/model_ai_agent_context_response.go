// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContextResponse struct {
	// Organization-scoped single-cluster Agent execution scope.
	Scope             AiAgentResourceScope     `json:"scope"`
	Cluster           *AiAgentClusterContext   `json:"cluster,omitempty"`
	DefaultTimeRanges []AiAgentTimeRange       `json:"defaultTimeRanges,omitempty"`
	Playbooks         []AiAgentPlaybookSummary `json:"playbooks,omitempty"`
	Capabilities      map[string]bool          `json:"capabilities"`
	Suggestions       []AiAgentSuggestion      `json:"suggestions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentContextResponse instantiates a new AiAgentContextResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentContextResponse(scope AiAgentResourceScope, capabilities map[string]bool) *AiAgentContextResponse {
	this := AiAgentContextResponse{}
	this.Scope = scope
	this.Capabilities = capabilities
	return &this
}

// NewAiAgentContextResponseWithDefaults instantiates a new AiAgentContextResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentContextResponseWithDefaults() *AiAgentContextResponse {
	this := AiAgentContextResponse{}
	return &this
}

// GetScope returns the Scope field value.
func (o *AiAgentContextResponse) GetScope() AiAgentResourceScope {
	if o == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiAgentContextResponse) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiAgentContextResponse) SetScope(v AiAgentResourceScope) {
	o.Scope = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *AiAgentContextResponse) GetCluster() AiAgentClusterContext {
	if o == nil || o.Cluster == nil {
		var ret AiAgentClusterContext
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContextResponse) GetClusterOk() (*AiAgentClusterContext, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *AiAgentContextResponse) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given AiAgentClusterContext and assigns it to the Cluster field.
func (o *AiAgentContextResponse) SetCluster(v AiAgentClusterContext) {
	o.Cluster = &v
}

// GetDefaultTimeRanges returns the DefaultTimeRanges field value if set, zero value otherwise.
func (o *AiAgentContextResponse) GetDefaultTimeRanges() []AiAgentTimeRange {
	if o == nil || o.DefaultTimeRanges == nil {
		var ret []AiAgentTimeRange
		return ret
	}
	return o.DefaultTimeRanges
}

// GetDefaultTimeRangesOk returns a tuple with the DefaultTimeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContextResponse) GetDefaultTimeRangesOk() (*[]AiAgentTimeRange, bool) {
	if o == nil || o.DefaultTimeRanges == nil {
		return nil, false
	}
	return &o.DefaultTimeRanges, true
}

// HasDefaultTimeRanges returns a boolean if a field has been set.
func (o *AiAgentContextResponse) HasDefaultTimeRanges() bool {
	return o != nil && o.DefaultTimeRanges != nil
}

// SetDefaultTimeRanges gets a reference to the given []AiAgentTimeRange and assigns it to the DefaultTimeRanges field.
func (o *AiAgentContextResponse) SetDefaultTimeRanges(v []AiAgentTimeRange) {
	o.DefaultTimeRanges = v
}

// GetPlaybooks returns the Playbooks field value if set, zero value otherwise.
func (o *AiAgentContextResponse) GetPlaybooks() []AiAgentPlaybookSummary {
	if o == nil || o.Playbooks == nil {
		var ret []AiAgentPlaybookSummary
		return ret
	}
	return o.Playbooks
}

// GetPlaybooksOk returns a tuple with the Playbooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContextResponse) GetPlaybooksOk() (*[]AiAgentPlaybookSummary, bool) {
	if o == nil || o.Playbooks == nil {
		return nil, false
	}
	return &o.Playbooks, true
}

// HasPlaybooks returns a boolean if a field has been set.
func (o *AiAgentContextResponse) HasPlaybooks() bool {
	return o != nil && o.Playbooks != nil
}

// SetPlaybooks gets a reference to the given []AiAgentPlaybookSummary and assigns it to the Playbooks field.
func (o *AiAgentContextResponse) SetPlaybooks(v []AiAgentPlaybookSummary) {
	o.Playbooks = v
}

// GetCapabilities returns the Capabilities field value.
func (o *AiAgentContextResponse) GetCapabilities() map[string]bool {
	if o == nil {
		var ret map[string]bool
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *AiAgentContextResponse) GetCapabilitiesOk() (*map[string]bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value.
func (o *AiAgentContextResponse) SetCapabilities(v map[string]bool) {
	o.Capabilities = v
}

// GetSuggestions returns the Suggestions field value if set, zero value otherwise.
func (o *AiAgentContextResponse) GetSuggestions() []AiAgentSuggestion {
	if o == nil || o.Suggestions == nil {
		var ret []AiAgentSuggestion
		return ret
	}
	return o.Suggestions
}

// GetSuggestionsOk returns a tuple with the Suggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContextResponse) GetSuggestionsOk() (*[]AiAgentSuggestion, bool) {
	if o == nil || o.Suggestions == nil {
		return nil, false
	}
	return &o.Suggestions, true
}

// HasSuggestions returns a boolean if a field has been set.
func (o *AiAgentContextResponse) HasSuggestions() bool {
	return o != nil && o.Suggestions != nil
}

// SetSuggestions gets a reference to the given []AiAgentSuggestion and assigns it to the Suggestions field.
func (o *AiAgentContextResponse) SetSuggestions(v []AiAgentSuggestion) {
	o.Suggestions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentContextResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["scope"] = o.Scope
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.DefaultTimeRanges != nil {
		toSerialize["defaultTimeRanges"] = o.DefaultTimeRanges
	}
	if o.Playbooks != nil {
		toSerialize["playbooks"] = o.Playbooks
	}
	toSerialize["capabilities"] = o.Capabilities
	if o.Suggestions != nil {
		toSerialize["suggestions"] = o.Suggestions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentContextResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Scope             *AiAgentResourceScope    `json:"scope"`
		Cluster           *AiAgentClusterContext   `json:"cluster,omitempty"`
		DefaultTimeRanges []AiAgentTimeRange       `json:"defaultTimeRanges,omitempty"`
		Playbooks         []AiAgentPlaybookSummary `json:"playbooks,omitempty"`
		Capabilities      *map[string]bool         `json:"capabilities"`
		Suggestions       []AiAgentSuggestion      `json:"suggestions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Capabilities == nil {
		return fmt.Errorf("required field capabilities missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scope", "cluster", "defaultTimeRanges", "playbooks", "capabilities", "suggestions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = *all.Scope
	if all.Cluster != nil && all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = all.Cluster
	o.DefaultTimeRanges = all.DefaultTimeRanges
	o.Playbooks = all.Playbooks
	o.Capabilities = *all.Capabilities
	o.Suggestions = all.Suggestions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
