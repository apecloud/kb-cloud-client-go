// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchShardRecovery struct {
	Index           common.NullableString          `json:"index,omitempty"`
	Shard           common.NullableInt64           `json:"shard,omitempty"`
	Type            common.NullableString          `json:"type,omitempty"`
	Stage           common.NullableString          `json:"stage,omitempty"`
	Primary         common.NullableBool            `json:"primary,omitempty"`
	StartTimeMillis common.NullableInt64           `json:"startTimeMillis,omitempty"`
	TotalTimeMillis common.NullableInt64           `json:"totalTimeMillis,omitempty"`
	Source          *ElasticsearchShardNode        `json:"source,omitempty"`
	Target          *ElasticsearchShardNode        `json:"target,omitempty"`
	Translog        *ElasticsearchRecoveryTranslog `json:"translog,omitempty"`
	Size            *ElasticsearchRecoverySize     `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchShardRecovery instantiates a new ElasticsearchShardRecovery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchShardRecovery() *ElasticsearchShardRecovery {
	this := ElasticsearchShardRecovery{}
	return &this
}

// NewElasticsearchShardRecoveryWithDefaults instantiates a new ElasticsearchShardRecovery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchShardRecoveryWithDefaults() *ElasticsearchShardRecovery {
	this := ElasticsearchShardRecovery{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetIndex() string {
	if o == nil || o.Index.Get() == nil {
		var ret string
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasIndex() bool {
	return o != nil && o.Index.IsSet()
}

// SetIndex gets a reference to the given common.NullableString and assigns it to the Index field.
func (o *ElasticsearchShardRecovery) SetIndex(v string) {
	o.Index.Set(&v)
}

// SetIndexNil sets the value for Index to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetIndex() {
	o.Index.Unset()
}

// GetShard returns the Shard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetShard() int64 {
	if o == nil || o.Shard.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Shard.Get()
}

// GetShardOk returns a tuple with the Shard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetShardOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shard.Get(), o.Shard.IsSet()
}

// HasShard returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasShard() bool {
	return o != nil && o.Shard.IsSet()
}

// SetShard gets a reference to the given common.NullableInt64 and assigns it to the Shard field.
func (o *ElasticsearchShardRecovery) SetShard(v int64) {
	o.Shard.Set(&v)
}

// SetShardNil sets the value for Shard to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetShardNil() {
	o.Shard.Set(nil)
}

// UnsetShard ensures that no value is present for Shard, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetShard() {
	o.Shard.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasType() bool {
	return o != nil && o.Type.IsSet()
}

// SetType gets a reference to the given common.NullableString and assigns it to the Type field.
func (o *ElasticsearchShardRecovery) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetType() {
	o.Type.Unset()
}

// GetStage returns the Stage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetStage() string {
	if o == nil || o.Stage.Get() == nil {
		var ret string
		return ret
	}
	return *o.Stage.Get()
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stage.Get(), o.Stage.IsSet()
}

// HasStage returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasStage() bool {
	return o != nil && o.Stage.IsSet()
}

// SetStage gets a reference to the given common.NullableString and assigns it to the Stage field.
func (o *ElasticsearchShardRecovery) SetStage(v string) {
	o.Stage.Set(&v)
}

// SetStageNil sets the value for Stage to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetStageNil() {
	o.Stage.Set(nil)
}

// UnsetStage ensures that no value is present for Stage, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetStage() {
	o.Stage.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetPrimary() bool {
	if o == nil || o.Primary.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasPrimary() bool {
	return o != nil && o.Primary.IsSet()
}

// SetPrimary gets a reference to the given common.NullableBool and assigns it to the Primary field.
func (o *ElasticsearchShardRecovery) SetPrimary(v bool) {
	o.Primary.Set(&v)
}

// SetPrimaryNil sets the value for Primary to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetPrimary() {
	o.Primary.Unset()
}

// GetStartTimeMillis returns the StartTimeMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetStartTimeMillis() int64 {
	if o == nil || o.StartTimeMillis.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeMillis.Get()
}

// GetStartTimeMillisOk returns a tuple with the StartTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetStartTimeMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimeMillis.Get(), o.StartTimeMillis.IsSet()
}

// HasStartTimeMillis returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasStartTimeMillis() bool {
	return o != nil && o.StartTimeMillis.IsSet()
}

// SetStartTimeMillis gets a reference to the given common.NullableInt64 and assigns it to the StartTimeMillis field.
func (o *ElasticsearchShardRecovery) SetStartTimeMillis(v int64) {
	o.StartTimeMillis.Set(&v)
}

// SetStartTimeMillisNil sets the value for StartTimeMillis to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetStartTimeMillisNil() {
	o.StartTimeMillis.Set(nil)
}

// UnsetStartTimeMillis ensures that no value is present for StartTimeMillis, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetStartTimeMillis() {
	o.StartTimeMillis.Unset()
}

// GetTotalTimeMillis returns the TotalTimeMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchShardRecovery) GetTotalTimeMillis() int64 {
	if o == nil || o.TotalTimeMillis.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalTimeMillis.Get()
}

// GetTotalTimeMillisOk returns a tuple with the TotalTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchShardRecovery) GetTotalTimeMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeMillis.Get(), o.TotalTimeMillis.IsSet()
}

// HasTotalTimeMillis returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasTotalTimeMillis() bool {
	return o != nil && o.TotalTimeMillis.IsSet()
}

// SetTotalTimeMillis gets a reference to the given common.NullableInt64 and assigns it to the TotalTimeMillis field.
func (o *ElasticsearchShardRecovery) SetTotalTimeMillis(v int64) {
	o.TotalTimeMillis.Set(&v)
}

// SetTotalTimeMillisNil sets the value for TotalTimeMillis to be an explicit nil.
func (o *ElasticsearchShardRecovery) SetTotalTimeMillisNil() {
	o.TotalTimeMillis.Set(nil)
}

// UnsetTotalTimeMillis ensures that no value is present for TotalTimeMillis, not even an explicit nil.
func (o *ElasticsearchShardRecovery) UnsetTotalTimeMillis() {
	o.TotalTimeMillis.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ElasticsearchShardRecovery) GetSource() ElasticsearchShardNode {
	if o == nil || o.Source == nil {
		var ret ElasticsearchShardNode
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchShardRecovery) GetSourceOk() (*ElasticsearchShardNode, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given ElasticsearchShardNode and assigns it to the Source field.
func (o *ElasticsearchShardRecovery) SetSource(v ElasticsearchShardNode) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ElasticsearchShardRecovery) GetTarget() ElasticsearchShardNode {
	if o == nil || o.Target == nil {
		var ret ElasticsearchShardNode
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchShardRecovery) GetTargetOk() (*ElasticsearchShardNode, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given ElasticsearchShardNode and assigns it to the Target field.
func (o *ElasticsearchShardRecovery) SetTarget(v ElasticsearchShardNode) {
	o.Target = &v
}

// GetTranslog returns the Translog field value if set, zero value otherwise.
func (o *ElasticsearchShardRecovery) GetTranslog() ElasticsearchRecoveryTranslog {
	if o == nil || o.Translog == nil {
		var ret ElasticsearchRecoveryTranslog
		return ret
	}
	return *o.Translog
}

// GetTranslogOk returns a tuple with the Translog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchShardRecovery) GetTranslogOk() (*ElasticsearchRecoveryTranslog, bool) {
	if o == nil || o.Translog == nil {
		return nil, false
	}
	return o.Translog, true
}

// HasTranslog returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasTranslog() bool {
	return o != nil && o.Translog != nil
}

// SetTranslog gets a reference to the given ElasticsearchRecoveryTranslog and assigns it to the Translog field.
func (o *ElasticsearchShardRecovery) SetTranslog(v ElasticsearchRecoveryTranslog) {
	o.Translog = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ElasticsearchShardRecovery) GetSize() ElasticsearchRecoverySize {
	if o == nil || o.Size == nil {
		var ret ElasticsearchRecoverySize
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchShardRecovery) GetSizeOk() (*ElasticsearchRecoverySize, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ElasticsearchShardRecovery) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given ElasticsearchRecoverySize and assigns it to the Size field.
func (o *ElasticsearchShardRecovery) SetSize(v ElasticsearchRecoverySize) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchShardRecovery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	if o.Shard.IsSet() {
		toSerialize["shard"] = o.Shard.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Stage.IsSet() {
		toSerialize["stage"] = o.Stage.Get()
	}
	if o.Primary.IsSet() {
		toSerialize["primary"] = o.Primary.Get()
	}
	if o.StartTimeMillis.IsSet() {
		toSerialize["startTimeMillis"] = o.StartTimeMillis.Get()
	}
	if o.TotalTimeMillis.IsSet() {
		toSerialize["totalTimeMillis"] = o.TotalTimeMillis.Get()
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Translog != nil {
		toSerialize["translog"] = o.Translog
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchShardRecovery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Index           common.NullableString          `json:"index,omitempty"`
		Shard           common.NullableInt64           `json:"shard,omitempty"`
		Type            common.NullableString          `json:"type,omitempty"`
		Stage           common.NullableString          `json:"stage,omitempty"`
		Primary         common.NullableBool            `json:"primary,omitempty"`
		StartTimeMillis common.NullableInt64           `json:"startTimeMillis,omitempty"`
		TotalTimeMillis common.NullableInt64           `json:"totalTimeMillis,omitempty"`
		Source          *ElasticsearchShardNode        `json:"source,omitempty"`
		Target          *ElasticsearchShardNode        `json:"target,omitempty"`
		Translog        *ElasticsearchRecoveryTranslog `json:"translog,omitempty"`
		Size            *ElasticsearchRecoverySize     `json:"size,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"index", "shard", "type", "stage", "primary", "startTimeMillis", "totalTimeMillis", "source", "target", "translog", "size"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Index = all.Index
	o.Shard = all.Shard
	o.Type = all.Type
	o.Stage = all.Stage
	o.Primary = all.Primary
	o.StartTimeMillis = all.StartTimeMillis
	o.TotalTimeMillis = all.TotalTimeMillis
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	if all.Translog != nil && all.Translog.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Translog = all.Translog
	if all.Size != nil && all.Size.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
