// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// GetProviderConfigReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) GetProviderConfigReference() xpv1.ProviderConfigReference {
	return p.ProviderConfigReference
}

// GetResourceReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) GetResourceReference() xpv1.TypedReference {
	return p.ResourceReference
}

// SetProviderConfigReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) SetProviderConfigReference(r xpv1.ProviderConfigReference) {
	p.ProviderConfigReference = r
}

// SetResourceReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) SetResourceReference(r xpv1.TypedReference) {
	p.ResourceReference = r
}

// GetObjectKind returns the GroupVersionKind of the ProviderConfigUsage
func (p *ProviderConfigUsage) GetObjectKind() schema.GroupVersionKind {
	return ProviderConfigUsageGroupVersionKind
}

// GetItems returns the list of ProviderConfigUsage as runtime.Object
func (l *ProviderConfigUsageList) GetItems() []runtime.Object {
	res := make([]runtime.Object, len(l.Items))
	for i := range l.Items {
		res[i] = &l.Items[i]
	}
	return res
}

// GetObjectKind returns the GroupVersionKind of the ProviderConfigUsageList
func (l *ProviderConfigUsageList) GetObjectKind() schema.GroupVersionKind {
	return ProviderConfigUsageListGroupVersionKind
}
