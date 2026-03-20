// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DnsBindRequest Request to bind external DNS to a cluster component
type DnsBindRequest struct {
	// Component name to bind DNS for
	Component string `json:"component"`
	// K8s Service name (e.g. vpc, internet)
	ServiceName string `json:"serviceName"`
	// IP or hostname from LoadBalancer to point the DNS record to
	TargetValue string `json:"targetValue"`
	// Root domain (e.g. example.com)
	Domain string `json:"domain"`
	// Subdomain/RR for the record. If not set, defaults to {cluster}-{component}
	Subdomain *string `json:"subdomain,omitempty"`
	// Cloud provider access key (e.g. Alibaba Cloud AccessKeyId)
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	// Cloud provider access secret
	AccessSecret *string `json:"accessSecret,omitempty"`
	// Cloud provider region (e.g. cn-hangzhou)
	Region *string `json:"region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDnsBindRequest instantiates a new DnsBindRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDnsBindRequest(component string, serviceName string, targetValue string, domain string) *DnsBindRequest {
	this := DnsBindRequest{}
	this.Component = component
	this.ServiceName = serviceName
	this.TargetValue = targetValue
	this.Domain = domain
	return &this
}

// NewDnsBindRequestWithDefaults instantiates a new DnsBindRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDnsBindRequestWithDefaults() *DnsBindRequest {
	this := DnsBindRequest{}
	return &this
}

// GetComponent returns the Component field value.
func (o *DnsBindRequest) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *DnsBindRequest) SetComponent(v string) {
	o.Component = v
}

// GetServiceName returns the ServiceName field value.
func (o *DnsBindRequest) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *DnsBindRequest) SetServiceName(v string) {
	o.ServiceName = v
}

// GetTargetValue returns the TargetValue field value.
func (o *DnsBindRequest) GetTargetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetValue
}

// GetTargetValueOk returns a tuple with the TargetValue field value
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetTargetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetValue, true
}

// SetTargetValue sets field value.
func (o *DnsBindRequest) SetTargetValue(v string) {
	o.TargetValue = v
}

// GetDomain returns the Domain field value.
func (o *DnsBindRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *DnsBindRequest) SetDomain(v string) {
	o.Domain = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DnsBindRequest) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DnsBindRequest) HasSubdomain() bool {
	return o != nil && o.Subdomain != nil
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DnsBindRequest) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *DnsBindRequest) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *DnsBindRequest) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *DnsBindRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetAccessSecret returns the AccessSecret field value if set, zero value otherwise.
func (o *DnsBindRequest) GetAccessSecret() string {
	if o == nil || o.AccessSecret == nil {
		var ret string
		return ret
	}
	return *o.AccessSecret
}

// GetAccessSecretOk returns a tuple with the AccessSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetAccessSecretOk() (*string, bool) {
	if o == nil || o.AccessSecret == nil {
		return nil, false
	}
	return o.AccessSecret, true
}

// HasAccessSecret returns a boolean if a field has been set.
func (o *DnsBindRequest) HasAccessSecret() bool {
	return o != nil && o.AccessSecret != nil
}

// SetAccessSecret gets a reference to the given string and assigns it to the AccessSecret field.
func (o *DnsBindRequest) SetAccessSecret(v string) {
	o.AccessSecret = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DnsBindRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsBindRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DnsBindRequest) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DnsBindRequest) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DnsBindRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["targetValue"] = o.TargetValue
	toSerialize["domain"] = o.Domain
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if o.AccessSecret != nil {
		toSerialize["accessSecret"] = o.AccessSecret
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DnsBindRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component    *string `json:"component"`
		ServiceName  *string `json:"serviceName"`
		TargetValue  *string `json:"targetValue"`
		Domain       *string `json:"domain"`
		Subdomain    *string `json:"subdomain,omitempty"`
		AccessKeyId  *string `json:"accessKeyId,omitempty"`
		AccessSecret *string `json:"accessSecret,omitempty"`
		Region       *string `json:"region,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field serviceName missing")
	}
	if all.TargetValue == nil {
		return fmt.Errorf("required field targetValue missing")
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "serviceName", "targetValue", "domain", "subdomain", "accessKeyId", "accessSecret", "region"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.ServiceName = *all.ServiceName
	o.TargetValue = *all.TargetValue
	o.Domain = *all.Domain
	o.Subdomain = all.Subdomain
	o.AccessKeyId = all.AccessKeyId
	o.AccessSecret = all.AccessSecret
	o.Region = all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
