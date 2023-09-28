//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Vesoft Inc.

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
	"k8s.io/api/autoscaling/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingPolicySpec) DeepCopyInto(out *AutoscalingPolicySpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(v2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingPolicySpec.
func (in *AutoscalingPolicySpec) DeepCopy() *AutoscalingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AutoscalingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingPolicyStatus) DeepCopyInto(out *AutoscalingPolicyStatus) {
	*out = *in
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		*out = (*in).DeepCopy()
	}
	if in.CurrentMetrics != nil {
		in, out := &in.CurrentMetrics, &out.CurrentMetrics
		*out = make([]v2.MetricStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingPolicyStatus.
func (in *AutoscalingPolicyStatus) DeepCopy() *AutoscalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AutoscalingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NebulaAutoscaler) DeepCopyInto(out *NebulaAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NebulaAutoscaler.
func (in *NebulaAutoscaler) DeepCopy() *NebulaAutoscaler {
	if in == nil {
		return nil
	}
	out := new(NebulaAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NebulaAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NebulaAutoscalerCondition) DeepCopyInto(out *NebulaAutoscalerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NebulaAutoscalerCondition.
func (in *NebulaAutoscalerCondition) DeepCopy() *NebulaAutoscalerCondition {
	if in == nil {
		return nil
	}
	out := new(NebulaAutoscalerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NebulaAutoscalerList) DeepCopyInto(out *NebulaAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NebulaAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NebulaAutoscalerList.
func (in *NebulaAutoscalerList) DeepCopy() *NebulaAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(NebulaAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NebulaAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NebulaAutoscalerSpec) DeepCopyInto(out *NebulaAutoscalerSpec) {
	*out = *in
	out.NebulaClusterRef = in.NebulaClusterRef
	in.GraphdPolicy.DeepCopyInto(&out.GraphdPolicy)
	if in.PollingPeriod != nil {
		in, out := &in.PollingPeriod, &out.PollingPeriod
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NebulaAutoscalerSpec.
func (in *NebulaAutoscalerSpec) DeepCopy() *NebulaAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(NebulaAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NebulaAutoscalerStatus) DeepCopyInto(out *NebulaAutoscalerStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	in.GraphdStatus.DeepCopyInto(&out.GraphdStatus)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]NebulaAutoscalerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NebulaAutoscalerStatus.
func (in *NebulaAutoscalerStatus) DeepCopy() *NebulaAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(NebulaAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NebulaClusterRef) DeepCopyInto(out *NebulaClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NebulaClusterRef.
func (in *NebulaClusterRef) DeepCopy() *NebulaClusterRef {
	if in == nil {
		return nil
	}
	out := new(NebulaClusterRef)
	in.DeepCopyInto(out)
	return out
}