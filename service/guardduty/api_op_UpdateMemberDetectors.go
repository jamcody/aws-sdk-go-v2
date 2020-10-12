// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Contains information on member accounts to be updated.
func (c *Client) UpdateMemberDetectors(ctx context.Context, params *UpdateMemberDetectorsInput, optFns ...func(*Options)) (*UpdateMemberDetectorsOutput, error) {
	if params == nil {
		params = &UpdateMemberDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMemberDetectors", params, optFns, addOperationUpdateMemberDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMemberDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMemberDetectorsInput struct {

	// A list of member account IDs to be updated.
	//
	// This member is required.
	AccountIds []*string

	// The detector ID of the master account.
	//
	// This member is required.
	DetectorId *string

	// An object describes which data sources will be updated.
	DataSources *types.DataSourceConfigurations
}

type UpdateMemberDetectorsOutput struct {

	// A list of member account IDs that were unable to be processed along with an
	// explanation for why they were not processed.
	//
	// This member is required.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMemberDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMemberDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMemberDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateMemberDetectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMemberDetectors(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateMemberDetectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "UpdateMemberDetectors",
	}
}