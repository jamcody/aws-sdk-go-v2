// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBClusterParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB cluster parameter group to return parameter details
	// for.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A value that indicates to return only parameters for a specific source. Parameter
	// sources can be engine, service, or customer.
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClusterParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClusterParametersInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDBClusterParametersOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string `type:"string"`

	// Provides a list of parameters for the DB cluster parameter group.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`
}

// String returns the string representation
func (s DescribeDBClusterParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBClusterParameters = "DescribeDBClusterParameters"

// DescribeDBClusterParametersRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Returns the detailed parameter list for a particular DB cluster parameter
// group.
//
//    // Example sending a request using DescribeDBClusterParametersRequest.
//    req := client.DescribeDBClusterParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBClusterParameters
func (c *Client) DescribeDBClusterParametersRequest(input *DescribeDBClusterParametersInput) DescribeDBClusterParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBClusterParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeDBClusterParametersOutput{})
	return DescribeDBClusterParametersRequest{Request: req, Input: input, Copy: c.DescribeDBClusterParametersRequest}
}

// DescribeDBClusterParametersRequest is the request type for the
// DescribeDBClusterParameters API operation.
type DescribeDBClusterParametersRequest struct {
	*aws.Request
	Input *DescribeDBClusterParametersInput
	Copy  func(*DescribeDBClusterParametersInput) DescribeDBClusterParametersRequest
}

// Send marshals and sends the DescribeDBClusterParameters API request.
func (r DescribeDBClusterParametersRequest) Send(ctx context.Context) (*DescribeDBClusterParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterParametersResponse{
		DescribeDBClusterParametersOutput: r.Request.Data.(*DescribeDBClusterParametersOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterParametersResponse is the response type for the
// DescribeDBClusterParameters API operation.
type DescribeDBClusterParametersResponse struct {
	*DescribeDBClusterParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterParameters request.
func (r *DescribeDBClusterParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}