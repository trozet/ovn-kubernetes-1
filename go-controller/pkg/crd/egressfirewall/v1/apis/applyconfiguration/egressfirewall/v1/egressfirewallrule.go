/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/egressfirewall/v1"
)

// EgressFirewallRuleApplyConfiguration represents a declarative configuration of the EgressFirewallRule type for use
// with apply.
type EgressFirewallRuleApplyConfiguration struct {
	Type  *v1.EgressFirewallRuleType                   `json:"type,omitempty"`
	Ports []EgressFirewallPortApplyConfiguration       `json:"ports,omitempty"`
	To    *EgressFirewallDestinationApplyConfiguration `json:"to,omitempty"`
}

// EgressFirewallRuleApplyConfiguration constructs a declarative configuration of the EgressFirewallRule type for use with
// apply.
func EgressFirewallRule() *EgressFirewallRuleApplyConfiguration {
	return &EgressFirewallRuleApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *EgressFirewallRuleApplyConfiguration) WithType(value v1.EgressFirewallRuleType) *EgressFirewallRuleApplyConfiguration {
	b.Type = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *EgressFirewallRuleApplyConfiguration) WithPorts(values ...*EgressFirewallPortApplyConfiguration) *EgressFirewallRuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithTo sets the To field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the To field is set to the value of the last call.
func (b *EgressFirewallRuleApplyConfiguration) WithTo(value *EgressFirewallDestinationApplyConfiguration) *EgressFirewallRuleApplyConfiguration {
	b.To = value
	return b
}