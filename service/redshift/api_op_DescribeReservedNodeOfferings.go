// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeReservedNodeOfferingsInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeReservedNodeOfferings request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The unique identifier for the offering.
	ReservedNodeOfferingId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReservedNodeOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeReservedNodeOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// A list of ReservedNodeOffering objects.
	ReservedNodeOfferings []ReservedNodeOffering `locationNameList:"ReservedNodeOffering" type:"list"`
}

// String returns the string representation
func (s DescribeReservedNodeOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReservedNodeOfferings = "DescribeReservedNodeOfferings"

// DescribeReservedNodeOfferingsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns a list of the available reserved node offerings by Amazon Redshift
// with their descriptions including the node type, the fixed and recurring
// costs of reserving the node and duration the node will be reserved for you.
// These descriptions help you determine which reserve node offering you want
// to purchase. You then use the unique offering ID in you call to PurchaseReservedNodeOffering
// to reserve one or more nodes for your Amazon Redshift cluster.
//
// For more information about reserved node offerings, go to Purchasing Reserved
// Nodes (https://docs.aws.amazon.com/redshift/latest/mgmt/purchase-reserved-node-instance.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using DescribeReservedNodeOfferingsRequest.
//    req := client.DescribeReservedNodeOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeReservedNodeOfferings
func (c *Client) DescribeReservedNodeOfferingsRequest(input *DescribeReservedNodeOfferingsInput) DescribeReservedNodeOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedNodeOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReservedNodeOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedNodeOfferingsOutput{})
	return DescribeReservedNodeOfferingsRequest{Request: req, Input: input, Copy: c.DescribeReservedNodeOfferingsRequest}
}

// DescribeReservedNodeOfferingsRequest is the request type for the
// DescribeReservedNodeOfferings API operation.
type DescribeReservedNodeOfferingsRequest struct {
	*aws.Request
	Input *DescribeReservedNodeOfferingsInput
	Copy  func(*DescribeReservedNodeOfferingsInput) DescribeReservedNodeOfferingsRequest
}

// Send marshals and sends the DescribeReservedNodeOfferings API request.
func (r DescribeReservedNodeOfferingsRequest) Send(ctx context.Context) (*DescribeReservedNodeOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedNodeOfferingsResponse{
		DescribeReservedNodeOfferingsOutput: r.Request.Data.(*DescribeReservedNodeOfferingsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedNodeOfferingsRequestPaginator returns a paginator for DescribeReservedNodeOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedNodeOfferingsRequest(input)
//   p := redshift.NewDescribeReservedNodeOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedNodeOfferingsPaginator(req DescribeReservedNodeOfferingsRequest) DescribeReservedNodeOfferingsPaginator {
	return DescribeReservedNodeOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedNodeOfferingsInput
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

// DescribeReservedNodeOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedNodeOfferingsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedNodeOfferingsPaginator) CurrentPage() *DescribeReservedNodeOfferingsOutput {
	return p.Pager.CurrentPage().(*DescribeReservedNodeOfferingsOutput)
}

// DescribeReservedNodeOfferingsResponse is the response type for the
// DescribeReservedNodeOfferings API operation.
type DescribeReservedNodeOfferingsResponse struct {
	*DescribeReservedNodeOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedNodeOfferings request.
func (r *DescribeReservedNodeOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}