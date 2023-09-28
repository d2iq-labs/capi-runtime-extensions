//go:build !ignore_autogenerated

// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterConfig) DeepCopyInto(out *AWSClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterConfig.
func (in *AWSClusterConfig) DeepCopy() *AWSClusterConfig {
	if in == nil {
		return nil
	}
	out := new(AWSClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterConfigSpec) DeepCopyInto(out *AWSClusterConfigSpec) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(Region)
		**out = **in
	}
	in.GenericClusterConfig.DeepCopyInto(&out.GenericClusterConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterConfigSpec.
func (in *AWSClusterConfigSpec) DeepCopy() *AWSClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AWSClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
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
func (in *DockerClusterConfig) DeepCopyInto(out *DockerClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerClusterConfig.
func (in *DockerClusterConfig) DeepCopy() *DockerClusterConfig {
	if in == nil {
		return nil
	}
	out := new(DockerClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerClusterConfigSpec) DeepCopyInto(out *DockerClusterConfigSpec) {
	*out = *in
	in.GenericClusterConfig.DeepCopyInto(&out.GenericClusterConfig)
	if in.CustomImage != nil {
		in, out := &in.CustomImage, &out.CustomImage
		*out = new(OCIImage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerClusterConfigSpec.
func (in *DockerClusterConfigSpec) DeepCopy() *DockerClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DockerClusterConfigSpec)
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
func (in *GenericClusterConfig) DeepCopyInto(out *GenericClusterConfig) {
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
	in.ImageRegistries.DeepCopyInto(&out.ImageRegistries)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new(Addons)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterConfig.
func (in *GenericClusterConfig) DeepCopy() *GenericClusterConfig {
	if in == nil {
		return nil
	}
	out := new(GenericClusterConfig)
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
func (in *ImageRegistries) DeepCopyInto(out *ImageRegistries) {
	*out = *in
	if in.ImageRegistryCredentials != nil {
		in, out := &in.ImageRegistryCredentials, &out.ImageRegistryCredentials
		*out = make(ImageRegistryCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistries.
func (in *ImageRegistries) DeepCopy() *ImageRegistries {
	if in == nil {
		return nil
	}
	out := new(ImageRegistries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ImageRegistryCredentials) DeepCopyInto(out *ImageRegistryCredentials) {
	{
		in := &in
		*out = make(ImageRegistryCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryCredentials.
func (in ImageRegistryCredentials) DeepCopy() ImageRegistryCredentials {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryCredentials)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryCredentialsResource) DeepCopyInto(out *ImageRegistryCredentialsResource) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryCredentialsResource.
func (in *ImageRegistryCredentialsResource) DeepCopy() *ImageRegistryCredentialsResource {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryCredentialsResource)
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
