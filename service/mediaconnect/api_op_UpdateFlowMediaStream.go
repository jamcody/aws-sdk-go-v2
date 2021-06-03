// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing media stream.
func (c *Client) UpdateFlowMediaStream(ctx context.Context, params *UpdateFlowMediaStreamInput, optFns ...func(*Options)) (*UpdateFlowMediaStreamOutput, error) {
	if params == nil {
		params = &UpdateFlowMediaStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowMediaStream", params, optFns, addOperationUpdateFlowMediaStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowMediaStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The fields that you want to update in the media stream.
type UpdateFlowMediaStreamInput struct {

	// The Amazon Resource Name (ARN) of the flow.
	//
	// This member is required.
	FlowArn *string

	// The name of the media stream that you want to update.
	//
	// This member is required.
	MediaStreamName *string

	// The attributes that you want to assign to the media stream.
	Attributes *types.MediaStreamAttributesRequest

	// The sample rate (in Hz) for the stream. If the media stream type is video or
	// ancillary data, set this value to 90000. If the media stream type is audio, set
	// this value to either 48000 or 96000.
	ClockRate int32

	// Description
	Description *string

	// The type of media stream.
	MediaStreamType types.MediaStreamType

	// The resolution of the video.
	VideoFormat *string
}

type UpdateFlowMediaStreamOutput struct {

	// The ARN of the flow that is associated with the media stream that you updated.
	FlowArn *string

	// The media stream that you updated.
	MediaStream *types.MediaStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFlowMediaStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlowMediaStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlowMediaStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateFlowMediaStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowMediaStream(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateFlowMediaStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "UpdateFlowMediaStream",
	}
}