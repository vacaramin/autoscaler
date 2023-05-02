// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisableScalingGroupCommon = "DisableScalingGroup"

// DisableScalingGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableScalingGroupCommon operation. The "output" return
// value will be populated with the DisableScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableScalingGroupCommon Send returns without error.
//
// See DisableScalingGroupCommon for more information on using the DisableScalingGroupCommon
// API call, and error handling.
//
//	// Example sending a request using the DisableScalingGroupCommonRequest method.
//	req, resp := client.DisableScalingGroupCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DisableScalingGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisableScalingGroupCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DisableScalingGroupCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DisableScalingGroupCommon for usage and error information.
func (c *AUTOSCALING) DisableScalingGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisableScalingGroupCommonRequest(input)
	return out, req.Send()
}

// DisableScalingGroupCommonWithContext is the same as DisableScalingGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisableScalingGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DisableScalingGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisableScalingGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisableScalingGroup = "DisableScalingGroup"

// DisableScalingGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableScalingGroup operation. The "output" return
// value will be populated with the DisableScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableScalingGroupCommon Send returns without error.
//
// See DisableScalingGroup for more information on using the DisableScalingGroup
// API call, and error handling.
//
//	// Example sending a request using the DisableScalingGroupRequest method.
//	req, resp := client.DisableScalingGroupRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DisableScalingGroupRequest(input *DisableScalingGroupInput) (req *request.Request, output *DisableScalingGroupOutput) {
	op := &request.Operation{
		Name:       opDisableScalingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableScalingGroupInput{}
	}

	output = &DisableScalingGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisableScalingGroup API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DisableScalingGroup for usage and error information.
func (c *AUTOSCALING) DisableScalingGroup(input *DisableScalingGroupInput) (*DisableScalingGroupOutput, error) {
	req, out := c.DisableScalingGroupRequest(input)
	return out, req.Send()
}

// DisableScalingGroupWithContext is the same as DisableScalingGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DisableScalingGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DisableScalingGroupWithContext(ctx volcengine.Context, input *DisableScalingGroupInput, opts ...request.Option) (*DisableScalingGroupOutput, error) {
	req, out := c.DisableScalingGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisableScalingGroupInput struct {
	_ struct{} `type:"structure"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DisableScalingGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableScalingGroupInput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DisableScalingGroupInput) SetScalingGroupId(v string) *DisableScalingGroupInput {
	s.ScalingGroupId = &v
	return s
}

type DisableScalingGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DisableScalingGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableScalingGroupOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DisableScalingGroupOutput) SetScalingGroupId(v string) *DisableScalingGroupOutput {
	s.ScalingGroupId = &v
	return s
}
