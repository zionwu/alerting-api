package v1

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Alert).DeepCopyInto(out.(*Alert))
			return nil
		}, InType: reflect.TypeOf(&Alert{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AlertList).DeepCopyInto(out.(*AlertList))
			return nil
		}, InType: reflect.TypeOf(&AlertList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*EmailConfig).DeepCopyInto(out.(*EmailConfig))
			return nil
		}, InType: reflect.TypeOf(&EmailConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeRule).DeepCopyInto(out.(*NodeRule))
			return nil
		}, InType: reflect.TypeOf(&NodeRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Notifier).DeepCopyInto(out.(*Notifier))
			return nil
		}, InType: reflect.TypeOf(&Notifier{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NotifierList).DeepCopyInto(out.(*NotifierList))
			return nil
		}, InType: reflect.TypeOf(&NotifierList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PagerdutyConfig).DeepCopyInto(out.(*PagerdutyConfig))
			return nil
		}, InType: reflect.TypeOf(&PagerdutyConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodRule).DeepCopyInto(out.(*PodRule))
			return nil
		}, InType: reflect.TypeOf(&PodRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SlackConfig).DeepCopyInto(out.(*SlackConfig))
			return nil
		}, InType: reflect.TypeOf(&SlackConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WebhookConfig).DeepCopyInto(out.(*WebhookConfig))
			return nil
		}, InType: reflect.TypeOf(&WebhookConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkloadRule).DeepCopyInto(out.(*WorkloadRule))
			return nil
		}, InType: reflect.TypeOf(&WorkloadRule{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alert) DeepCopyInto(out *Alert) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.TargetSelector != nil {
		in, out := &in.TargetSelector, &out.TargetSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.NodeRule = in.NodeRule
	out.PodRule = in.PodRule
	out.WorkloadRule = in.WorkloadRule
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alert.
func (in *Alert) DeepCopy() *Alert {
	if in == nil {
		return nil
	}
	out := new(Alert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertList) DeepCopyInto(out *AlertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertList.
func (in *AlertList) DeepCopy() *AlertList {
	if in == nil {
		return nil
	}
	out := new(AlertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfig) DeepCopyInto(out *EmailConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfig.
func (in *EmailConfig) DeepCopy() *EmailConfig {
	if in == nil {
		return nil
	}
	out := new(EmailConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRule) DeepCopyInto(out *NodeRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRule.
func (in *NodeRule) DeepCopy() *NodeRule {
	if in == nil {
		return nil
	}
	out := new(NodeRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notifier) DeepCopyInto(out *Notifier) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.EmailConfig != nil {
		in, out := &in.EmailConfig, &out.EmailConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(EmailConfig)
			**out = **in
		}
	}
	if in.SlackConfig != nil {
		in, out := &in.SlackConfig, &out.SlackConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(SlackConfig)
			**out = **in
		}
	}
	if in.PagerdutyConfig != nil {
		in, out := &in.PagerdutyConfig, &out.PagerdutyConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(PagerdutyConfig)
			**out = **in
		}
	}
	if in.WebhookConfig != nil {
		in, out := &in.WebhookConfig, &out.WebhookConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebhookConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notifier.
func (in *Notifier) DeepCopy() *Notifier {
	if in == nil {
		return nil
	}
	out := new(Notifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Notifier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotifierList) DeepCopyInto(out *NotifierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Notifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotifierList.
func (in *NotifierList) DeepCopy() *NotifierList {
	if in == nil {
		return nil
	}
	out := new(NotifierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotifierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerdutyConfig) DeepCopyInto(out *PagerdutyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerdutyConfig.
func (in *PagerdutyConfig) DeepCopy() *PagerdutyConfig {
	if in == nil {
		return nil
	}
	out := new(PagerdutyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodRule) DeepCopyInto(out *PodRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodRule.
func (in *PodRule) DeepCopy() *PodRule {
	if in == nil {
		return nil
	}
	out := new(PodRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfig) DeepCopyInto(out *SlackConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfig.
func (in *SlackConfig) DeepCopy() *SlackConfig {
	if in == nil {
		return nil
	}
	out := new(SlackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRule) DeepCopyInto(out *WorkloadRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRule.
func (in *WorkloadRule) DeepCopy() *WorkloadRule {
	if in == nil {
		return nil
	}
	out := new(WorkloadRule)
	in.DeepCopyInto(out)
	return out
}
