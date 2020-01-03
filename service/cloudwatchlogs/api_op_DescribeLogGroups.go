// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLogGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The prefix to match.
	LogGroupNamePrefix *string `locationName:"logGroupNamePrefix" min:"1" type:"string"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeLogGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLogGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLogGroupsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.LogGroupNamePrefix != nil && len(*s.LogGroupNamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupNamePrefix", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLogGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The log groups.
	LogGroups []LogGroup `locationName:"logGroups" type:"list"`

	// The token for the next set of items to return. The token expires after 24
	// hours.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeLogGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLogGroups = "DescribeLogGroups"

// DescribeLogGroupsRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists the specified log groups. You can list all your log groups or filter
// the results by prefix. The results are ASCII-sorted by log group name.
//
//    // Example sending a request using DescribeLogGroupsRequest.
//    req := client.DescribeLogGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeLogGroups
func (c *Client) DescribeLogGroupsRequest(input *DescribeLogGroupsInput) DescribeLogGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeLogGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeLogGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeLogGroupsOutput{})
	return DescribeLogGroupsRequest{Request: req, Input: input, Copy: c.DescribeLogGroupsRequest}
}

// DescribeLogGroupsRequest is the request type for the
// DescribeLogGroups API operation.
type DescribeLogGroupsRequest struct {
	*aws.Request
	Input *DescribeLogGroupsInput
	Copy  func(*DescribeLogGroupsInput) DescribeLogGroupsRequest
}

// Send marshals and sends the DescribeLogGroups API request.
func (r DescribeLogGroupsRequest) Send(ctx context.Context) (*DescribeLogGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLogGroupsResponse{
		DescribeLogGroupsOutput: r.Request.Data.(*DescribeLogGroupsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeLogGroupsRequestPaginator returns a paginator for DescribeLogGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeLogGroupsRequest(input)
//   p := cloudwatchlogs.NewDescribeLogGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeLogGroupsPaginator(req DescribeLogGroupsRequest) DescribeLogGroupsPaginator {
	return DescribeLogGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeLogGroupsInput
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

// DescribeLogGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeLogGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeLogGroupsPaginator) CurrentPage() *DescribeLogGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeLogGroupsOutput)
}

// DescribeLogGroupsResponse is the response type for the
// DescribeLogGroups API operation.
type DescribeLogGroupsResponse struct {
	*DescribeLogGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLogGroups request.
func (r *DescribeLogGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}