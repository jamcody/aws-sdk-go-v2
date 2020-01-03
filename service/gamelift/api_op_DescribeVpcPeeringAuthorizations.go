// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcPeeringAuthorizationsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeVpcPeeringAuthorizationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVpcPeeringAuthorizationsOutput struct {
	_ struct{} `type:"structure"`

	// Collection of objects that describe all valid VPC peering operations for
	// the current AWS account.
	VpcPeeringAuthorizations []VpcPeeringAuthorization `type:"list"`
}

// String returns the string representation
func (s DescribeVpcPeeringAuthorizationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcPeeringAuthorizations = "DescribeVpcPeeringAuthorizations"

// DescribeVpcPeeringAuthorizationsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves valid VPC peering authorizations that are pending for the AWS account.
// This operation returns all VPC peering authorizations and requests for peering.
// This includes those initiated and received by this account.
//
//    * CreateVpcPeeringAuthorization
//
//    * DescribeVpcPeeringAuthorizations
//
//    * DeleteVpcPeeringAuthorization
//
//    * CreateVpcPeeringConnection
//
//    * DescribeVpcPeeringConnections
//
//    * DeleteVpcPeeringConnection
//
//    // Example sending a request using DescribeVpcPeeringAuthorizationsRequest.
//    req := client.DescribeVpcPeeringAuthorizationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeVpcPeeringAuthorizations
func (c *Client) DescribeVpcPeeringAuthorizationsRequest(input *DescribeVpcPeeringAuthorizationsInput) DescribeVpcPeeringAuthorizationsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcPeeringAuthorizations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcPeeringAuthorizationsInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcPeeringAuthorizationsOutput{})
	return DescribeVpcPeeringAuthorizationsRequest{Request: req, Input: input, Copy: c.DescribeVpcPeeringAuthorizationsRequest}
}

// DescribeVpcPeeringAuthorizationsRequest is the request type for the
// DescribeVpcPeeringAuthorizations API operation.
type DescribeVpcPeeringAuthorizationsRequest struct {
	*aws.Request
	Input *DescribeVpcPeeringAuthorizationsInput
	Copy  func(*DescribeVpcPeeringAuthorizationsInput) DescribeVpcPeeringAuthorizationsRequest
}

// Send marshals and sends the DescribeVpcPeeringAuthorizations API request.
func (r DescribeVpcPeeringAuthorizationsRequest) Send(ctx context.Context) (*DescribeVpcPeeringAuthorizationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcPeeringAuthorizationsResponse{
		DescribeVpcPeeringAuthorizationsOutput: r.Request.Data.(*DescribeVpcPeeringAuthorizationsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVpcPeeringAuthorizationsResponse is the response type for the
// DescribeVpcPeeringAuthorizations API operation.
type DescribeVpcPeeringAuthorizationsResponse struct {
	*DescribeVpcPeeringAuthorizationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcPeeringAuthorizations request.
func (r *DescribeVpcPeeringAuthorizationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}