// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This test case serializes XML lists for the following cases for both input and
// output:
//
//     * Normal XML lists.
//
//     * Normal XML sets.
//
//     * XML lists of
// lists.
//
//     * XML lists with @xmlName on its members
//
//     * Flattened XML
// lists.
//
//     * Flattened XML lists with @xmlName.
//
//     * Lists of structures.
func (c *Client) XmlLists(ctx context.Context, params *XmlListsInput, optFns ...func(*Options)) (*XmlListsOutput, error) {
	if params == nil {
		params = &XmlListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlLists", params, optFns, addOperationXmlListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlListsInput struct {
}

type XmlListsOutput struct {
	BooleanList []*bool

	EnumList []types.FooEnum

	FlattenedList []*string

	FlattenedList2 []*string

	IntegerList []*int32

	// A list of lists of strings.
	NestedStringList [][]*string

	RenamedListMembers []*string

	StringList []*string

	StringSet []*string

	StructureList []*types.StructureListMember

	TimestampList []*time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpXmlLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpXmlLists{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlLists(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opXmlLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlLists",
	}
}