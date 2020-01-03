// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePresignedNotebookInstanceUrlInput struct {
	_ struct{} `type:"structure"`

	// The name of the notebook instance.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`

	// The duration of the session, in seconds. The default is 12 hours.
	SessionExpirationDurationInSeconds *int64 `min:"1800" type:"integer"`
}

// String returns the string representation
func (s CreatePresignedNotebookInstanceUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePresignedNotebookInstanceUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePresignedNotebookInstanceUrlInput"}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}
	if s.SessionExpirationDurationInSeconds != nil && *s.SessionExpirationDurationInSeconds < 1800 {
		invalidParams.Add(aws.NewErrParamMinValue("SessionExpirationDurationInSeconds", 1800))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePresignedNotebookInstanceUrlOutput struct {
	_ struct{} `type:"structure"`

	// A JSON object that contains the URL string.
	AuthorizedUrl *string `type:"string"`
}

// String returns the string representation
func (s CreatePresignedNotebookInstanceUrlOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePresignedNotebookInstanceUrl = "CreatePresignedNotebookInstanceUrl"

// CreatePresignedNotebookInstanceUrlRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns a URL that you can use to connect to the Jupyter server from a notebook
// instance. In the Amazon SageMaker console, when you choose Open next to a
// notebook instance, Amazon SageMaker opens a new tab showing the Jupyter server
// home page from the notebook instance. The console uses this API to get the
// URL and show the page.
//
// IAM authorization policies for this API are also enforced for every HTTP
// request and WebSocket frame that attempts to connect to the notebook instance.For
// example, you can restrict access to this API and to the URL that it returns
// to a list of IP addresses that you specify. Use the NotIpAddress condition
// operator and the aws:SourceIP condition context key to specify the list of
// IP addresses that you want to have access to the notebook instance. For more
// information, see Limit Access to a Notebook Instance by IP Address (https://docs.aws.amazon.com/sagemaker/latest/dg/security_iam_id-based-policy-examples.html#nbi-ip-filter).
//
// The URL that you get from a call to is valid only for 5 minutes. If you try
// to use the URL after the 5-minute limit expires, you are directed to the
// AWS console sign-in page.
//
//    // Example sending a request using CreatePresignedNotebookInstanceUrlRequest.
//    req := client.CreatePresignedNotebookInstanceUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreatePresignedNotebookInstanceUrl
func (c *Client) CreatePresignedNotebookInstanceUrlRequest(input *CreatePresignedNotebookInstanceUrlInput) CreatePresignedNotebookInstanceUrlRequest {
	op := &aws.Operation{
		Name:       opCreatePresignedNotebookInstanceUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePresignedNotebookInstanceUrlInput{}
	}

	req := c.newRequest(op, input, &CreatePresignedNotebookInstanceUrlOutput{})
	return CreatePresignedNotebookInstanceUrlRequest{Request: req, Input: input, Copy: c.CreatePresignedNotebookInstanceUrlRequest}
}

// CreatePresignedNotebookInstanceUrlRequest is the request type for the
// CreatePresignedNotebookInstanceUrl API operation.
type CreatePresignedNotebookInstanceUrlRequest struct {
	*aws.Request
	Input *CreatePresignedNotebookInstanceUrlInput
	Copy  func(*CreatePresignedNotebookInstanceUrlInput) CreatePresignedNotebookInstanceUrlRequest
}

// Send marshals and sends the CreatePresignedNotebookInstanceUrl API request.
func (r CreatePresignedNotebookInstanceUrlRequest) Send(ctx context.Context) (*CreatePresignedNotebookInstanceUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePresignedNotebookInstanceUrlResponse{
		CreatePresignedNotebookInstanceUrlOutput: r.Request.Data.(*CreatePresignedNotebookInstanceUrlOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePresignedNotebookInstanceUrlResponse is the response type for the
// CreatePresignedNotebookInstanceUrl API operation.
type CreatePresignedNotebookInstanceUrlResponse struct {
	*CreatePresignedNotebookInstanceUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePresignedNotebookInstanceUrl request.
func (r *CreatePresignedNotebookInstanceUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}