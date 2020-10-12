// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a cache policy, including the following metadata:
//
//     * The policy’s
// identifier.
//
//     * The date and time when the policy was last modified.
//
// To get
// a cache policy, you must provide the policy’s identifier. If the cache policy is
// attached to a distribution’s cache behavior, you can get the policy’s identifier
// using ListDistributions or GetDistribution. If the cache policy is not attached
// to a cache behavior, you can get the identifier using ListCachePolicies.
func (c *Client) GetCachePolicy(ctx context.Context, params *GetCachePolicyInput, optFns ...func(*Options)) (*GetCachePolicyOutput, error) {
	if params == nil {
		params = &GetCachePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCachePolicy", params, optFns, addOperationGetCachePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCachePolicyInput struct {

	// The unique identifier for the cache policy. If the cache policy is attached to a
	// distribution’s cache behavior, you can get the policy’s identifier using
	// ListDistributions or GetDistribution. If the cache policy is not attached to a
	// cache behavior, you can get the identifier using ListCachePolicies.
	//
	// This member is required.
	Id *string
}

type GetCachePolicyOutput struct {

	// The cache policy.
	CachePolicy *types.CachePolicy

	// The current version of the cache policy.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCachePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetCachePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetCachePolicy{}, middleware.After)
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
	addOpGetCachePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCachePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCachePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetCachePolicy",
	}
}