// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowTest) DeepCopyInto(out *FlowTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowTest.
func (in *FlowTest) DeepCopy() *FlowTest {
	if in == nil {
		return nil
	}
	out := new(FlowTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowTestList) DeepCopyInto(out *FlowTestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlowTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowTestList.
func (in *FlowTestList) DeepCopy() *FlowTestList {
	if in == nil {
		return nil
	}
	out := new(FlowTestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowTestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowTestSpec) DeepCopyInto(out *FlowTestSpec) {
	*out = *in
	out.ReferencePod = in.ReferencePod
	out.ReferenceFlow = in.ReferenceFlow
	if in.SentMessages != nil {
		in, out := &in.SentMessages, &out.SentMessages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowTestSpec.
func (in *FlowTestSpec) DeepCopy() *FlowTestSpec {
	if in == nil {
		return nil
	}
	out := new(FlowTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowTestStatus) DeepCopyInto(out *FlowTestStatus) {
	*out = *in
	in.FailedMatch.DeepCopyInto(&out.FailedMatch)
	in.FailedFilter.DeepCopyInto(&out.FailedFilter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowTestStatus.
func (in *FlowTestStatus) DeepCopy() *FlowTestStatus {
	if in == nil {
		return nil
	}
	out := new(FlowTestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferenceObject) DeepCopyInto(out *ReferenceObject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferenceObject.
func (in *ReferenceObject) DeepCopy() *ReferenceObject {
	if in == nil {
		return nil
	}
	out := new(ReferenceObject)
	in.DeepCopyInto(out)
	return out
}
