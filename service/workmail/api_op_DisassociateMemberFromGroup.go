// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateMemberFromGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the group from which members are removed.
	//
	// GroupId is a required field
	GroupId *string `min:"12" type:"string" required:"true"`

	// The identifier for the member to be removed to the group.
	//
	// MemberId is a required field
	MemberId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the group exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateMemberFromGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateMemberFromGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateMemberFromGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}
	if s.GroupId != nil && len(*s.GroupId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupId", 12))
	}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateMemberFromGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateMemberFromGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateMemberFromGroup = "DisassociateMemberFromGroup"

// DisassociateMemberFromGroupRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Removes a member from a group.
//
//    // Example sending a request using DisassociateMemberFromGroupRequest.
//    req := client.DisassociateMemberFromGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DisassociateMemberFromGroup
func (c *Client) DisassociateMemberFromGroupRequest(input *DisassociateMemberFromGroupInput) DisassociateMemberFromGroupRequest {
	op := &aws.Operation{
		Name:       opDisassociateMemberFromGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateMemberFromGroupInput{}
	}

	req := c.newRequest(op, input, &DisassociateMemberFromGroupOutput{})
	return DisassociateMemberFromGroupRequest{Request: req, Input: input, Copy: c.DisassociateMemberFromGroupRequest}
}

// DisassociateMemberFromGroupRequest is the request type for the
// DisassociateMemberFromGroup API operation.
type DisassociateMemberFromGroupRequest struct {
	*aws.Request
	Input *DisassociateMemberFromGroupInput
	Copy  func(*DisassociateMemberFromGroupInput) DisassociateMemberFromGroupRequest
}

// Send marshals and sends the DisassociateMemberFromGroup API request.
func (r DisassociateMemberFromGroupRequest) Send(ctx context.Context) (*DisassociateMemberFromGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateMemberFromGroupResponse{
		DisassociateMemberFromGroupOutput: r.Request.Data.(*DisassociateMemberFromGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateMemberFromGroupResponse is the response type for the
// DisassociateMemberFromGroup API operation.
type DisassociateMemberFromGroupResponse struct {
	*DisassociateMemberFromGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateMemberFromGroup request.
func (r *DisassociateMemberFromGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}