// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkteamInput struct {
	_ struct{} `type:"structure"`

	// The name of the work team to return a description of.
	//
	// WorkteamName is a required field
	WorkteamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeWorkteamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkteamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkteamInput"}

	if s.WorkteamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkteamName"))
	}
	if s.WorkteamName != nil && len(*s.WorkteamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkteamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkteamOutput struct {
	_ struct{} `type:"structure"`

	// A Workteam instance that contains information about the work team.
	//
	// Workteam is a required field
	Workteam *Workteam `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkteamOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkteam = "DescribeWorkteam"

// DescribeWorkteamRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets information about a specific work team. You can see information such
// as the create date, the last updated date, membership information, and the
// work team's Amazon Resource Name (ARN).
//
//    // Example sending a request using DescribeWorkteamRequest.
//    req := client.DescribeWorkteamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeWorkteam
func (c *Client) DescribeWorkteamRequest(input *DescribeWorkteamInput) DescribeWorkteamRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkteam,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeWorkteamInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkteamOutput{})
	return DescribeWorkteamRequest{Request: req, Input: input, Copy: c.DescribeWorkteamRequest}
}

// DescribeWorkteamRequest is the request type for the
// DescribeWorkteam API operation.
type DescribeWorkteamRequest struct {
	*aws.Request
	Input *DescribeWorkteamInput
	Copy  func(*DescribeWorkteamInput) DescribeWorkteamRequest
}

// Send marshals and sends the DescribeWorkteam API request.
func (r DescribeWorkteamRequest) Send(ctx context.Context) (*DescribeWorkteamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkteamResponse{
		DescribeWorkteamOutput: r.Request.Data.(*DescribeWorkteamOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkteamResponse is the response type for the
// DescribeWorkteam API operation.
type DescribeWorkteamResponse struct {
	*DescribeWorkteamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkteam request.
func (r *DescribeWorkteamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}