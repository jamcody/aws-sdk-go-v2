// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that was returned when you called the CreateCertificateAuthority
	// action. This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// Use this parameter when paginating results to specify the maximum number
	// of items to return in the response. If additional items exist beyond the
	// number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int64 `min:"1" type:"integer"`

	// Use this parameter when paginating results in a subsequent request after
	// you receive a response with truncated results. Set it to the value of NextToken
	// from the response you just received.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsOutput struct {
	_ struct{} `type:"structure"`

	// When the list is truncated, this value is present and should be used for
	// the NextToken parameter in a subsequent pagination request.
	NextToken *string `min:"1" type:"string"`

	// The tags associated with your private CA.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s ListTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTags = "ListTags"

// ListTagsRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Lists the tags, if any, that are associated with your private CA. Tags are
// labels that you can use to identify and organize your CAs. Each tag consists
// of a key and an optional value. Call the TagCertificateAuthority action to
// add one or more tags to your CA. Call the UntagCertificateAuthority action
// to remove tags.
//
//    // Example sending a request using ListTagsRequest.
//    req := client.ListTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ListTags
func (c *Client) ListTagsRequest(input *ListTagsInput) ListTagsRequest {
	op := &aws.Operation{
		Name:       opListTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTagsInput{}
	}

	req := c.newRequest(op, input, &ListTagsOutput{})
	return ListTagsRequest{Request: req, Input: input, Copy: c.ListTagsRequest}
}

// ListTagsRequest is the request type for the
// ListTags API operation.
type ListTagsRequest struct {
	*aws.Request
	Input *ListTagsInput
	Copy  func(*ListTagsInput) ListTagsRequest
}

// Send marshals and sends the ListTags API request.
func (r ListTagsRequest) Send(ctx context.Context) (*ListTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsResponse{
		ListTagsOutput: r.Request.Data.(*ListTagsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTagsRequestPaginator returns a paginator for ListTags.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTagsRequest(input)
//   p := acmpca.NewListTagsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTagsPaginator(req ListTagsRequest) ListTagsPaginator {
	return ListTagsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTagsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListTagsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTagsPaginator struct {
	aws.Pager
}

func (p *ListTagsPaginator) CurrentPage() *ListTagsOutput {
	return p.Pager.CurrentPage().(*ListTagsOutput)
}

// ListTagsResponse is the response type for the
// ListTags API operation.
type ListTagsResponse struct {
	*ListTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTags request.
func (r *ListTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}