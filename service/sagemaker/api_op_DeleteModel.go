// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteModelInput struct {
	_ struct{} `type:"structure"`

	// The name of the model to delete.
	//
	// ModelName is a required field
	ModelName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteModelInput"}

	if s.ModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteModelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteModelOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteModel = "DeleteModel"

// DeleteModelRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes a model. The DeleteModel API deletes only the model entry that was
// created in Amazon SageMaker when you called the CreateModel (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateModel.html)
// API. It does not delete model artifacts, inference code, or the IAM role
// that you specified when creating the model.
//
//    // Example sending a request using DeleteModelRequest.
//    req := client.DeleteModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteModel
func (c *Client) DeleteModelRequest(input *DeleteModelInput) DeleteModelRequest {
	op := &aws.Operation{
		Name:       opDeleteModel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteModelInput{}
	}

	req := c.newRequest(op, input, &DeleteModelOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteModelRequest{Request: req, Input: input, Copy: c.DeleteModelRequest}
}

// DeleteModelRequest is the request type for the
// DeleteModel API operation.
type DeleteModelRequest struct {
	*aws.Request
	Input *DeleteModelInput
	Copy  func(*DeleteModelInput) DeleteModelRequest
}

// Send marshals and sends the DeleteModel API request.
func (r DeleteModelRequest) Send(ctx context.Context) (*DeleteModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteModelResponse{
		DeleteModelOutput: r.Request.Data.(*DeleteModelOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteModelResponse is the response type for the
// DeleteModel API operation.
type DeleteModelResponse struct {
	*DeleteModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteModel request.
func (r *DeleteModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}