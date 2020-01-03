// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeAliasInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet alias. Specify the alias you want to retrieve.
	//
	// AliasId is a required field
	AliasId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAliasInput"}

	if s.AliasId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeAliasOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains the requested alias.
	Alias *Alias `type:"structure"`
}

// String returns the string representation
func (s DescribeAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAlias = "DescribeAlias"

// DescribeAliasRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves properties for an alias. This operation returns all alias metadata
// and settings. To get an alias's target fleet ID only, use ResolveAlias.
//
// To get alias properties, specify the alias ID. If successful, the requested
// alias record is returned.
//
//    * CreateAlias
//
//    * ListAliases
//
//    * DescribeAlias
//
//    * UpdateAlias
//
//    * DeleteAlias
//
//    * ResolveAlias
//
//    // Example sending a request using DescribeAliasRequest.
//    req := client.DescribeAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeAlias
func (c *Client) DescribeAliasRequest(input *DescribeAliasInput) DescribeAliasRequest {
	op := &aws.Operation{
		Name:       opDescribeAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAliasInput{}
	}

	req := c.newRequest(op, input, &DescribeAliasOutput{})
	return DescribeAliasRequest{Request: req, Input: input, Copy: c.DescribeAliasRequest}
}

// DescribeAliasRequest is the request type for the
// DescribeAlias API operation.
type DescribeAliasRequest struct {
	*aws.Request
	Input *DescribeAliasInput
	Copy  func(*DescribeAliasInput) DescribeAliasRequest
}

// Send marshals and sends the DescribeAlias API request.
func (r DescribeAliasRequest) Send(ctx context.Context) (*DescribeAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAliasResponse{
		DescribeAliasOutput: r.Request.Data.(*DescribeAliasOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAliasResponse is the response type for the
// DescribeAlias API operation.
type DescribeAliasResponse struct {
	*DescribeAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAlias request.
func (r *DescribeAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}