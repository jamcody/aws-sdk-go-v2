// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateSkillFromSkillGroupInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of a skill. Required.
	SkillGroupArn *string `type:"string"`

	// The ARN of a skill group to associate to a skill.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateSkillFromSkillGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateSkillFromSkillGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateSkillFromSkillGroupInput"}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateSkillFromSkillGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateSkillFromSkillGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateSkillFromSkillGroup = "DisassociateSkillFromSkillGroup"

// DisassociateSkillFromSkillGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Disassociates a skill from a skill group.
//
//    // Example sending a request using DisassociateSkillFromSkillGroupRequest.
//    req := client.DisassociateSkillFromSkillGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DisassociateSkillFromSkillGroup
func (c *Client) DisassociateSkillFromSkillGroupRequest(input *DisassociateSkillFromSkillGroupInput) DisassociateSkillFromSkillGroupRequest {
	op := &aws.Operation{
		Name:       opDisassociateSkillFromSkillGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateSkillFromSkillGroupInput{}
	}

	req := c.newRequest(op, input, &DisassociateSkillFromSkillGroupOutput{})
	return DisassociateSkillFromSkillGroupRequest{Request: req, Input: input, Copy: c.DisassociateSkillFromSkillGroupRequest}
}

// DisassociateSkillFromSkillGroupRequest is the request type for the
// DisassociateSkillFromSkillGroup API operation.
type DisassociateSkillFromSkillGroupRequest struct {
	*aws.Request
	Input *DisassociateSkillFromSkillGroupInput
	Copy  func(*DisassociateSkillFromSkillGroupInput) DisassociateSkillFromSkillGroupRequest
}

// Send marshals and sends the DisassociateSkillFromSkillGroup API request.
func (r DisassociateSkillFromSkillGroupRequest) Send(ctx context.Context) (*DisassociateSkillFromSkillGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateSkillFromSkillGroupResponse{
		DisassociateSkillFromSkillGroupOutput: r.Request.Data.(*DisassociateSkillFromSkillGroupOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateSkillFromSkillGroupResponse is the response type for the
// DisassociateSkillFromSkillGroup API operation.
type DisassociateSkillFromSkillGroupResponse struct {
	*DisassociateSkillFromSkillGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateSkillFromSkillGroup request.
func (r *DisassociateSkillFromSkillGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}