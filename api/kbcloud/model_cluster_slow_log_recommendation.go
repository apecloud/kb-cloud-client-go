// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogRecommendation A rule-based recommendation for a slow log diagnosis
type ClusterSlowLogRecommendation struct {
	// Recommendation type. Current values are create_index, rewrite_sql, reduce_result_set, and investigate_lock.
	Type *string `json:"type,omitempty"`
	Sql  *string `json:"sql,omitempty"`
	// Localized slow log diagnosis text with dynamic arguments
	Description *ClusterSlowLogLocalizedText `json:"description,omitempty"`
	// Evidence source used by the local rule
	Source *string `json:"source,omitempty"`
	// Recommendation safety level. Current values are manual_review_required and safe_hint_only.
	Safety *string `json:"safety,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogRecommendation instantiates a new ClusterSlowLogRecommendation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogRecommendation() *ClusterSlowLogRecommendation {
	this := ClusterSlowLogRecommendation{}
	return &this
}

// NewClusterSlowLogRecommendationWithDefaults instantiates a new ClusterSlowLogRecommendation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogRecommendationWithDefaults() *ClusterSlowLogRecommendation {
	this := ClusterSlowLogRecommendation{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterSlowLogRecommendation) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogRecommendation) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterSlowLogRecommendation) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterSlowLogRecommendation) SetType(v string) {
	o.Type = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *ClusterSlowLogRecommendation) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogRecommendation) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *ClusterSlowLogRecommendation) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *ClusterSlowLogRecommendation) SetSql(v string) {
	o.Sql = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterSlowLogRecommendation) GetDescription() ClusterSlowLogLocalizedText {
	if o == nil || o.Description == nil {
		var ret ClusterSlowLogLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogRecommendation) GetDescriptionOk() (*ClusterSlowLogLocalizedText, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterSlowLogRecommendation) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given ClusterSlowLogLocalizedText and assigns it to the Description field.
func (o *ClusterSlowLogRecommendation) SetDescription(v ClusterSlowLogLocalizedText) {
	o.Description = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ClusterSlowLogRecommendation) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogRecommendation) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ClusterSlowLogRecommendation) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ClusterSlowLogRecommendation) SetSource(v string) {
	o.Source = &v
}

// GetSafety returns the Safety field value if set, zero value otherwise.
func (o *ClusterSlowLogRecommendation) GetSafety() string {
	if o == nil || o.Safety == nil {
		var ret string
		return ret
	}
	return *o.Safety
}

// GetSafetyOk returns a tuple with the Safety field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogRecommendation) GetSafetyOk() (*string, bool) {
	if o == nil || o.Safety == nil {
		return nil, false
	}
	return o.Safety, true
}

// HasSafety returns a boolean if a field has been set.
func (o *ClusterSlowLogRecommendation) HasSafety() bool {
	return o != nil && o.Safety != nil
}

// SetSafety gets a reference to the given string and assigns it to the Safety field.
func (o *ClusterSlowLogRecommendation) SetSafety(v string) {
	o.Safety = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Safety != nil {
		toSerialize["safety"] = o.Safety
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogRecommendation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type        *string                      `json:"type,omitempty"`
		Sql         *string                      `json:"sql,omitempty"`
		Description *ClusterSlowLogLocalizedText `json:"description,omitempty"`
		Source      *string                      `json:"source,omitempty"`
		Safety      *string                      `json:"safety,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "sql", "description", "source", "safety"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Type = all.Type
	o.Sql = all.Sql
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.Source = all.Source
	o.Safety = all.Safety

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
