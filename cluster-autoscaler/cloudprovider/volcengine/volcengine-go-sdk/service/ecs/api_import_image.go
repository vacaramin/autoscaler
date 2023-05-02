// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opImportImageCommon = "ImportImage"

// ImportImageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ImportImageCommon operation. The "output" return
// value will be populated with the ImportImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ImportImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after ImportImageCommon Send returns without error.
//
// See ImportImageCommon for more information on using the ImportImageCommon
// API call, and error handling.
//
//	// Example sending a request using the ImportImageCommonRequest method.
//	req, resp := client.ImportImageCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) ImportImageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opImportImageCommon,
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

// ImportImageCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ImportImageCommon for usage and error information.
func (c *ECS) ImportImageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ImportImageCommonRequest(input)
	return out, req.Send()
}

// ImportImageCommonWithContext is the same as ImportImageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ImportImageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ImportImageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ImportImageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opImportImage = "ImportImage"

// ImportImageRequest generates a "volcengine/request.Request" representing the
// client's request for the ImportImage operation. The "output" return
// value will be populated with the ImportImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ImportImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after ImportImageCommon Send returns without error.
//
// See ImportImage for more information on using the ImportImage
// API call, and error handling.
//
//	// Example sending a request using the ImportImageRequest method.
//	req, resp := client.ImportImageRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) ImportImageRequest(input *ImportImageInput) (req *request.Request, output *ImportImageOutput) {
	op := &request.Operation{
		Name:       opImportImage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportImageInput{}
	}

	output = &ImportImageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ImportImage API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ImportImage for usage and error information.
func (c *ECS) ImportImage(input *ImportImageInput) (*ImportImageOutput, error) {
	req, out := c.ImportImageRequest(input)
	return out, req.Send()
}

// ImportImageWithContext is the same as ImportImage with the addition of
// the ability to pass a context and additional request options.
//
// See ImportImage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ImportImageWithContext(ctx volcengine.Context, input *ImportImageInput, opts ...request.Option) (*ImportImageOutput, error) {
	req, out := c.ImportImageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ImportImageInput struct {
	_ struct{} `type:"structure"`

	Architecture *string `type:"string"`

	BootMode *string `type:"string"`

	Description *string `type:"string"`

	ImageName *string `type:"string"`

	OsType *string `type:"string"`

	Platform *string `type:"string"`

	PlatformVersion *string `type:"string"`

	ProjectName *string `type:"string"`

	Url *string `type:"string"`
}

// String returns the string representation
func (s ImportImageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ImportImageInput) GoString() string {
	return s.String()
}

// SetArchitecture sets the Architecture field's value.
func (s *ImportImageInput) SetArchitecture(v string) *ImportImageInput {
	s.Architecture = &v
	return s
}

// SetBootMode sets the BootMode field's value.
func (s *ImportImageInput) SetBootMode(v string) *ImportImageInput {
	s.BootMode = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ImportImageInput) SetDescription(v string) *ImportImageInput {
	s.Description = &v
	return s
}

// SetImageName sets the ImageName field's value.
func (s *ImportImageInput) SetImageName(v string) *ImportImageInput {
	s.ImageName = &v
	return s
}

// SetOsType sets the OsType field's value.
func (s *ImportImageInput) SetOsType(v string) *ImportImageInput {
	s.OsType = &v
	return s
}

// SetPlatform sets the Platform field's value.
func (s *ImportImageInput) SetPlatform(v string) *ImportImageInput {
	s.Platform = &v
	return s
}

// SetPlatformVersion sets the PlatformVersion field's value.
func (s *ImportImageInput) SetPlatformVersion(v string) *ImportImageInput {
	s.PlatformVersion = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ImportImageInput) SetProjectName(v string) *ImportImageInput {
	s.ProjectName = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *ImportImageInput) SetUrl(v string) *ImportImageInput {
	s.Url = &v
	return s
}

type ImportImageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ImageId *string `type:"string"`
}

// String returns the string representation
func (s ImportImageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ImportImageOutput) GoString() string {
	return s.String()
}

// SetImageId sets the ImageId field's value.
func (s *ImportImageOutput) SetImageId(v string) *ImportImageOutput {
	s.ImageId = &v
	return s
}
