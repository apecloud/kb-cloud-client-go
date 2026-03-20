// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterDNSBinding Cluster DNS binding information
type ClusterDNSBinding struct {
	// Binding ID
	Id *int64 `json:"id,omitempty"`
	// Cluster ID
	ClusterId *int64 `json:"clusterId,omitempty"`
	// Component name
	Component *string `json:"component,omitempty"`
	// K8s Service name
	ServiceName *string `json:"serviceName,omitempty"`
	// Bound DNS hostname
	DnsHostname *string `json:"dnsHostname,omitempty"`
	// Target IP or hostname
	TargetValue *string `json:"targetValue,omitempty"`
	// DNS record type (A, CNAME, etc.)
	RecordType *string `json:"recordType,omitempty"`
	// DNS provider (aliyun, aws, etc.)
	Provider *string `json:"provider,omitempty"`
	// Creation time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Last update time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterDNSBinding instantiates a new ClusterDNSBinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterDNSBinding() *ClusterDNSBinding {
	this := ClusterDNSBinding{}
	return &this
}

// NewClusterDNSBindingWithDefaults instantiates a new ClusterDNSBinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterDNSBindingWithDefaults() *ClusterDNSBinding {
	this := ClusterDNSBinding{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterDNSBinding) SetId(v int64) {
	o.Id = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetClusterId() int64 {
	if o == nil || o.ClusterId == nil {
		var ret int64
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetClusterIdOk() (*int64, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given int64 and assigns it to the ClusterId field.
func (o *ClusterDNSBinding) SetClusterId(v int64) {
	o.ClusterId = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ClusterDNSBinding) SetComponent(v string) {
	o.Component = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ClusterDNSBinding) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetDnsHostname returns the DnsHostname field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetDnsHostname() string {
	if o == nil || o.DnsHostname == nil {
		var ret string
		return ret
	}
	return *o.DnsHostname
}

// GetDnsHostnameOk returns a tuple with the DnsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetDnsHostnameOk() (*string, bool) {
	if o == nil || o.DnsHostname == nil {
		return nil, false
	}
	return o.DnsHostname, true
}

// HasDnsHostname returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasDnsHostname() bool {
	return o != nil && o.DnsHostname != nil
}

// SetDnsHostname gets a reference to the given string and assigns it to the DnsHostname field.
func (o *ClusterDNSBinding) SetDnsHostname(v string) {
	o.DnsHostname = &v
}

// GetTargetValue returns the TargetValue field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetTargetValue() string {
	if o == nil || o.TargetValue == nil {
		var ret string
		return ret
	}
	return *o.TargetValue
}

// GetTargetValueOk returns a tuple with the TargetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetTargetValueOk() (*string, bool) {
	if o == nil || o.TargetValue == nil {
		return nil, false
	}
	return o.TargetValue, true
}

// HasTargetValue returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasTargetValue() bool {
	return o != nil && o.TargetValue != nil
}

// SetTargetValue gets a reference to the given string and assigns it to the TargetValue field.
func (o *ClusterDNSBinding) SetTargetValue(v string) {
	o.TargetValue = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasRecordType() bool {
	return o != nil && o.RecordType != nil
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *ClusterDNSBinding) SetRecordType(v string) {
	o.RecordType = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ClusterDNSBinding) SetProvider(v string) {
	o.Provider = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ClusterDNSBinding) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ClusterDNSBinding) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDNSBinding) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ClusterDNSBinding) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ClusterDNSBinding) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterDNSBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.DnsHostname != nil {
		toSerialize["dnsHostname"] = o.DnsHostname
	}
	if o.TargetValue != nil {
		toSerialize["targetValue"] = o.TargetValue
	}
	if o.RecordType != nil {
		toSerialize["recordType"] = o.RecordType
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterDNSBinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *int64     `json:"id,omitempty"`
		ClusterId   *int64     `json:"clusterId,omitempty"`
		Component   *string    `json:"component,omitempty"`
		ServiceName *string    `json:"serviceName,omitempty"`
		DnsHostname *string    `json:"dnsHostname,omitempty"`
		TargetValue *string    `json:"targetValue,omitempty"`
		RecordType  *string    `json:"recordType,omitempty"`
		Provider    *string    `json:"provider,omitempty"`
		CreatedAt   *time.Time `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "clusterId", "component", "serviceName", "dnsHostname", "targetValue", "recordType", "provider", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Id = all.Id
	o.ClusterId = all.ClusterId
	o.Component = all.Component
	o.ServiceName = all.ServiceName
	o.DnsHostname = all.DnsHostname
	o.TargetValue = all.TargetValue
	o.RecordType = all.RecordType
	o.Provider = all.Provider
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
