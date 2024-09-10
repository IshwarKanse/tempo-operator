/*
Copyright The Kubernetes Authors.

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

package v1beta1

// VolumeAttachmentSpecApplyConfiguration represents an declarative configuration of the VolumeAttachmentSpec type for use
// with apply.
type VolumeAttachmentSpecApplyConfiguration struct {
	Attacher *string                                   `json:"attacher,omitempty"`
	Source   *VolumeAttachmentSourceApplyConfiguration `json:"source,omitempty"`
	NodeName *string                                   `json:"nodeName,omitempty"`
}

// VolumeAttachmentSpecApplyConfiguration constructs an declarative configuration of the VolumeAttachmentSpec type for use with
// apply.
func VolumeAttachmentSpec() *VolumeAttachmentSpecApplyConfiguration {
	return &VolumeAttachmentSpecApplyConfiguration{}
}

// WithAttacher sets the Attacher field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Attacher field is set to the value of the last call.
func (b *VolumeAttachmentSpecApplyConfiguration) WithAttacher(value string) *VolumeAttachmentSpecApplyConfiguration {
	b.Attacher = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *VolumeAttachmentSpecApplyConfiguration) WithSource(value *VolumeAttachmentSourceApplyConfiguration) *VolumeAttachmentSpecApplyConfiguration {
	b.Source = value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *VolumeAttachmentSpecApplyConfiguration) WithNodeName(value string) *VolumeAttachmentSpecApplyConfiguration {
	b.NodeName = &value
	return b
}
