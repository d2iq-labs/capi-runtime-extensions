//go:build !ignore_autogenerated

// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/api/external/sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMILookup) DeepCopyInto(out *AMILookup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMILookup.
func (in *AMILookup) DeepCopy() *AMILookup {
	if in == nil {
		return nil
	}
	out := new(AMILookup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMISpec) DeepCopyInto(out *AMISpec) {
	*out = *in
	if in.Lookup != nil {
		in, out := &in.Lookup, &out.Lookup
		*out = new(AMILookup)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMISpec.
func (in *AMISpec) DeepCopy() *AMISpec {
	if in == nil {
		return nil
	}
	out := new(AMISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSLoadBalancerSpec) DeepCopyInto(out *AWSLoadBalancerSpec) {
	*out = *in
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(v1beta2.ELBScheme)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSLoadBalancerSpec.
func (in *AWSLoadBalancerSpec) DeepCopy() *AWSLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(AWSLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSNetwork) DeepCopyInto(out *AWSNetwork) {
	*out = *in
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(VPC)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSNetwork.
func (in *AWSNetwork) DeepCopy() *AWSNetwork {
	if in == nil {
		return nil
	}
	out := new(AWSNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSNodeSpec) DeepCopyInto(out *AWSNodeSpec) {
	*out = *in
	if in.IAMInstanceProfile != nil {
		in, out := &in.IAMInstanceProfile, &out.IAMInstanceProfile
		*out = new(IAMInstanceProfile)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(InstanceType)
		**out = **in
	}
	if in.AMISpec != nil {
		in, out := &in.AMISpec, &out.AMISpec
		*out = new(AMISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalSecurityGroups != nil {
		in, out := &in.AdditionalSecurityGroups, &out.AdditionalSecurityGroups
		*out = make(AdditionalSecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSNodeSpec.
func (in *AWSNodeSpec) DeepCopy() *AWSNodeSpec {
	if in == nil {
		return nil
	}
	out := new(AWSNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSpec) DeepCopyInto(out *AWSSpec) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(Region)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(AWSNetwork)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneLoadBalancer != nil {
		in, out := &in.ControlPlaneLoadBalancer, &out.ControlPlaneLoadBalancer
		*out = new(AWSLoadBalancerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSpec.
func (in *AWSSpec) DeepCopy() *AWSSpec {
	if in == nil {
		return nil
	}
	out := new(AWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AdditionalSecurityGroup) DeepCopyInto(out *AdditionalSecurityGroup) {
	{
		in := &in
		*out = make(AdditionalSecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalSecurityGroup.
func (in AdditionalSecurityGroup) DeepCopy() AdditionalSecurityGroup {
	if in == nil {
		return nil
	}
	out := new(AdditionalSecurityGroup)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addons) DeepCopyInto(out *Addons) {
	*out = *in
	if in.CNI != nil {
		in, out := &in.CNI, &out.CNI
		*out = new(CNI)
		**out = **in
	}
	if in.NFD != nil {
		in, out := &in.NFD, &out.NFD
		*out = new(NFD)
		**out = **in
	}
	if in.ClusterAutoscaler != nil {
		in, out := &in.ClusterAutoscaler, &out.ClusterAutoscaler
		*out = new(ClusterAutoscaler)
		**out = **in
	}
	if in.CPI != nil {
		in, out := &in.CPI, &out.CPI
		*out = new(CPI)
		**out = **in
	}
	if in.CSIProviders != nil {
		in, out := &in.CSIProviders, &out.CSIProviders
		*out = new(CSIProviders)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addons.
func (in *Addons) DeepCopy() *Addons {
	if in == nil {
		return nil
	}
	out := new(Addons)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllProvidersSpec) DeepCopyInto(out *AllProvidersSpec) {
	*out = *in
	if in.KubernetesImageRepository != nil {
		in, out := &in.KubernetesImageRepository, &out.KubernetesImageRepository
		*out = new(KubernetesImageRepository)
		**out = **in
	}
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(Etcd)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(HTTPProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraAPIServerCertSANs != nil {
		in, out := &in.ExtraAPIServerCertSANs, &out.ExtraAPIServerCertSANs
		*out = make(ExtraAPIServerCertSANs, len(*in))
		copy(*out, *in)
	}
	if in.ImageRegistries != nil {
		in, out := &in.ImageRegistries, &out.ImageRegistries
		*out = make(ImageRegistries, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GlobalImageRegistryMirror != nil {
		in, out := &in.GlobalImageRegistryMirror, &out.GlobalImageRegistryMirror
		*out = new(GlobalImageRegistryMirror)
		(*in).DeepCopyInto(*out)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new(Addons)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllProvidersSpec.
func (in *AllProvidersSpec) DeepCopy() *AllProvidersSpec {
	if in == nil {
		return nil
	}
	out := new(AllProvidersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNI) DeepCopyInto(out *CNI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNI.
func (in *CNI) DeepCopy() *CNI {
	if in == nil {
		return nil
	}
	out := new(CNI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPI) DeepCopyInto(out *CPI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPI.
func (in *CPI) DeepCopy() *CPI {
	if in == nil {
		return nil
	}
	out := new(CPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIProvider) DeepCopyInto(out *CSIProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIProvider.
func (in *CSIProvider) DeepCopy() *CSIProvider {
	if in == nil {
		return nil
	}
	out := new(CSIProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIProviders) DeepCopyInto(out *CSIProviders) {
	*out = *in
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]CSIProvider, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIProviders.
func (in *CSIProviders) DeepCopy() *CSIProviders {
	if in == nil {
		return nil
	}
	out := new(CSIProviders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscaler) DeepCopyInto(out *ClusterAutoscaler) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscaler.
func (in *ClusterAutoscaler) DeepCopy() *ClusterAutoscaler {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(DockerSpec)
		**out = **in
	}
	in.AllProvidersSpec.DeepCopyInto(&out.AllProvidersSpec)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(NodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerNodeSpec) DeepCopyInto(out *DockerNodeSpec) {
	*out = *in
	if in.CustomImage != nil {
		in, out := &in.CustomImage, &out.CustomImage
		*out = new(OCIImage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerNodeSpec.
func (in *DockerNodeSpec) DeepCopy() *DockerNodeSpec {
	if in == nil {
		return nil
	}
	out := new(DockerNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerSpec) DeepCopyInto(out *DockerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerSpec.
func (in *DockerSpec) DeepCopy() *DockerSpec {
	if in == nil {
		return nil
	}
	out := new(DockerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ExtraAPIServerCertSANs) DeepCopyInto(out *ExtraAPIServerCertSANs) {
	{
		in := &in
		*out = make(ExtraAPIServerCertSANs, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraAPIServerCertSANs.
func (in ExtraAPIServerCertSANs) DeepCopy() ExtraAPIServerCertSANs {
	if in == nil {
		return nil
	}
	out := new(ExtraAPIServerCertSANs)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericNodeConfig) DeepCopyInto(out *GenericNodeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericNodeConfig.
func (in *GenericNodeConfig) DeepCopy() *GenericNodeConfig {
	if in == nil {
		return nil
	}
	out := new(GenericNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalImageRegistryMirror) DeepCopyInto(out *GlobalImageRegistryMirror) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(RegistryCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalImageRegistryMirror.
func (in *GlobalImageRegistryMirror) DeepCopy() *GlobalImageRegistryMirror {
	if in == nil {
		return nil
	}
	out := new(GlobalImageRegistryMirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxy) DeepCopyInto(out *HTTPProxy) {
	*out = *in
	if in.AdditionalNo != nil {
		in, out := &in.AdditionalNo, &out.AdditionalNo
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxy.
func (in *HTTPProxy) DeepCopy() *HTTPProxy {
	if in == nil {
		return nil
	}
	out := new(HTTPProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ImageRegistries) DeepCopyInto(out *ImageRegistries) {
	{
		in := &in
		*out = make(ImageRegistries, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistries.
func (in ImageRegistries) DeepCopy() ImageRegistries {
	if in == nil {
		return nil
	}
	out := new(ImageRegistries)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistry) DeepCopyInto(out *ImageRegistry) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(RegistryCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistry.
func (in *ImageRegistry) DeepCopy() *ImageRegistry {
	if in == nil {
		return nil
	}
	out := new(ImageRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFD) DeepCopyInto(out *NFD) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFD.
func (in *NFD) DeepCopy() *NFD {
	if in == nil {
		return nil
	}
	out := new(NFD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigSpec) DeepCopyInto(out *NodeConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(DockerNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigSpec.
func (in *NodeConfigSpec) DeepCopy() *NodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentials) DeepCopyInto(out *RegistryCredentials) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentials.
func (in *RegistryCredentials) DeepCopy() *RegistryCredentials {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Subnets) DeepCopyInto(out *Subnets) {
	{
		in := &in
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnets.
func (in Subnets) DeepCopy() Subnets {
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}
