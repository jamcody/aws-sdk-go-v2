// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the group.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The organization under which the group is to be created.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGroupOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the group.
	GroupId *string `min:"12" type:"string"`
}

// String returns the string representation
func (s CreateGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateGroup = "CreateGroup"

// CreateGroupRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Creates a group that can be used in Amazon WorkMail by calling the RegisterToWorkMail
// operation.
//
//    // Example sending a request using CreateGroupRequest.
//    req := client.CreateGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/CreateGroup
func (c *Client) CreateGroupRequest(input *CreateGroupInput) CreateGroupRequest {
	op := &aws.Operation{
		Name:       opCreateGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGroupInput{}
	}

	req := c.newRequest(op, input, &CreateGroupOutput{})
	return CreateGroupRequest{Request: req, Input: input, Copy: c.CreateGroupRequest}
}

// CreateGroupRequest is the request type for the
// CreateGroup API operation.
type CreateGroupRequest struct {
	*aws.Request
	Input *CreateGroupInput
	Copy  func(*CreateGroupInput) CreateGroupRequest
}

// Send marshals and sends the CreateGroup API request.
func (r CreateGroupRequest) Send(ctx context.Context) (*CreateGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupResponse{
		CreateGroupOutput: r.Request.Data.(*CreateGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupResponse is the response type for the
// CreateGroup API operation.
type CreateGroupResponse struct {
	*CreateGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroup request.
func (r *CreateGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}