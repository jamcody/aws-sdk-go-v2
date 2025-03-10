// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Composes an email message and immediately queues it for sending. This operation
// is more flexible than the SendEmail API operation. When you use the SendRawEmail
// operation, you can specify the headers of the message as well as its content.
// This flexibility is useful, for example, when you want to send a multipart MIME
// email (such a message that contains both a text and an HTML version). You can
// also use this operation to send messages that include attachments. The
// SendRawEmail operation has the following requirements:
//   - You can only send email from verified email addresses or domains (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html)
//     . If you try to send email from an address that isn't verified, the operation
//     results in an "Email address not verified" error.
//   - If your account is still in the Amazon SES sandbox (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/request-production-access.html)
//     , you can only send email to other verified addresses in your account, or to
//     addresses that are associated with the Amazon SES mailbox simulator (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mailbox-simulator.html)
//     .
//   - The maximum message size, including attachments, is 10 MB.
//   - Each message has to include at least one recipient address. A recipient
//     address includes any address on the To:, CC:, or BCC: lines.
//   - If you send a single message to more than one recipient address, and one of
//     the recipient addresses isn't in a valid format (that is, it's not in the format
//     UserName@[SubDomain.]Domain.TopLevelDomain), Amazon SES rejects the entire
//     message, even if the other addresses are valid.
//   - Each message can include up to 50 recipient addresses across the To:, CC:,
//     or BCC: lines. If you need to send a single message to more than 50 recipients,
//     you have to split the list of recipient addresses into groups of less than 50
//     recipients, and send separate messages to each group.
//   - Amazon SES allows you to specify 8-bit Content-Transfer-Encoding for MIME
//     message parts. However, if Amazon SES has to modify the contents of your message
//     (for example, if you use open and click tracking), 8-bit content isn't
//     preserved. For this reason, we highly recommend that you encode all content that
//     isn't 7-bit ASCII. For more information, see MIME Encoding (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html#send-email-mime-encoding)
//     in the Amazon SES Developer Guide.
//
// Additionally, keep the following considerations in mind when using the
// SendRawEmail operation:
//   - Although you can customize the message headers when using the SendRawEmail
//     operation, Amazon SES will automatically apply its own Message-ID and Date
//     headers; if you passed these headers when creating the message, they will be
//     overwritten by the values that Amazon SES provides.
//   - If you are using sending authorization to send on behalf of another user,
//     SendRawEmail enables you to specify the cross-account identity for the email's
//     Source, From, and Return-Path parameters in one of two ways: you can pass
//     optional parameters SourceArn , FromArn , and/or ReturnPathArn to the API, or
//     you can include the following X-headers in the header of your raw email:
//   - X-SES-SOURCE-ARN
//   - X-SES-FROM-ARN
//   - X-SES-RETURN-PATH-ARN Don't include these X-headers in the DKIM signature.
//     Amazon SES removes these before it sends the email. If you only specify the
//     SourceIdentityArn parameter, Amazon SES sets the From and Return-Path
//     addresses to the same identity that you specified. For more information about
//     sending authorization, see the Using Sending Authorization with Amazon SES (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html)
//     in the Amazon SES Developer Guide.
//   - For every message that you send, the total number of recipients (including
//     each recipient in the To:, CC: and BCC: fields) is counted against the maximum
//     number of emails you can send in a 24-hour period (your sending quota). For more
//     information about sending quotas in Amazon SES, see Managing Your Amazon SES
//     Sending Limits (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/manage-sending-limits.html)
//     in the Amazon SES Developer Guide.
func (c *Client) SendRawEmail(ctx context.Context, params *SendRawEmailInput, optFns ...func(*Options)) (*SendRawEmailOutput, error) {
	if params == nil {
		params = &SendRawEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendRawEmail", params, optFns, c.addOperationSendRawEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendRawEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a single raw email using Amazon SES. For more
// information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html)
// .
type SendRawEmailInput struct {

	// The raw email message itself. The message has to meet the following criteria:
	//   - The message has to contain a header and a body, separated by a blank line.
	//   - All of the required header fields must be present in the message.
	//   - Each part of a multipart MIME message must be formatted properly.
	//   - Attachments must be of a content type that Amazon SES supports. For a list
	//   on unsupported content types, see Unsupported Attachment Types (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mime-types.html)
	//   in the Amazon SES Developer Guide.
	//   - The entire message must be base64-encoded.
	//   - If any of the MIME parts in your message contain content that is outside of
	//   the 7-bit ASCII character range, we highly recommend that you encode that
	//   content. For more information, see Sending Raw Email (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html)
	//   in the Amazon SES Developer Guide.
	//   - Per RFC 5321 (https://tools.ietf.org/html/rfc5321#section-4.5.3.1.6) , the
	//   maximum length of each line of text, including the , must not exceed 1,000
	//   characters.
	//
	// This member is required.
	RawMessage *types.RawMessage

	// The name of the configuration set to use when you send an email using
	// SendRawEmail .
	ConfigurationSetName *string

	// A list of destinations for the message, consisting of To:, CC:, and BCC:
	// addresses.
	Destinations []string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to specify a particular "From" address in the header of the raw email.
	// Instead of using this parameter, you can use the X-header X-SES-FROM-ARN in the
	// raw message of the email. If you use both the FromArn parameter and the
	// corresponding X-header, Amazon SES uses the value of the FromArn parameter. For
	// information about when to use this parameter, see the description of
	// SendRawEmail in this guide, or see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-delegate-sender-tasks-email.html)
	// .
	FromArn *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the ReturnPath parameter. For
	// example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com ) attaches a policy to
	// it that authorizes you to use feedback@example.com , then you would specify the
	// ReturnPathArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com ,
	// and the ReturnPath to be feedback@example.com . Instead of using this parameter,
	// you can use the X-header X-SES-RETURN-PATH-ARN in the raw message of the email.
	// If you use both the ReturnPathArn parameter and the corresponding X-header,
	// Amazon SES uses the value of the ReturnPathArn parameter. For information about
	// when to use this parameter, see the description of SendRawEmail in this guide,
	// or see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-delegate-sender-tasks-email.html)
	// .
	ReturnPathArn *string

	// The identity's email address. If you do not provide a value for this parameter,
	// you must specify a "From" address in the raw text of the message. (You can also
	// specify both.) Amazon SES does not support the SMTPUTF8 extension, as described
	// in RFC6531 (https://tools.ietf.org/html/rfc6531) . For this reason, the local
	// part of a source email address (the part of the email address that precedes the
	// @ sign) may only contain 7-bit ASCII characters (https://en.wikipedia.org/wiki/Email_address#Local-part)
	// . If the domain part of an address (the part after the @ sign) contains
	// non-ASCII characters, they must be encoded using Punycode, as described in
	// RFC3492 (https://tools.ietf.org/html/rfc3492.html) . The sender name (also known
	// as the friendly name) may contain non-ASCII characters. These characters must be
	// encoded using MIME encoded-word syntax, as described in RFC 2047 (https://tools.ietf.org/html/rfc2047)
	// . MIME encoded-word syntax uses the following form:
	// =?charset?encoding?encoded-text?= . If you specify the Source parameter and
	// have feedback forwarding enabled, then bounces and complaints will be sent to
	// this email address. This takes precedence over any Return-Path header that you
	// might include in the raw text of the message.
	Source *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to send for the email address specified in the Source parameter. For
	// example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com ) attaches a policy to
	// it that authorizes you to send from user@example.com , then you would specify
	// the SourceArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com ,
	// and the Source to be user@example.com . Instead of using this parameter, you can
	// use the X-header X-SES-SOURCE-ARN in the raw message of the email. If you use
	// both the SourceArn parameter and the corresponding X-header, Amazon SES uses
	// the value of the SourceArn parameter. For information about when to use this
	// parameter, see the description of SendRawEmail in this guide, or see the Amazon
	// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-delegate-sender-tasks-email.html)
	// .
	SourceArn *string

	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send using SendRawEmail . Tags correspond to characteristics of the email that
	// you define, so that you can publish email sending events.
	Tags []types.MessageTag

	noSmithyDocumentSerde
}

// Represents a unique message ID.
type SendRawEmailOutput struct {

	// The unique message identifier returned from the SendRawEmail action.
	//
	// This member is required.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendRawEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendRawEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendRawEmail{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSendRawEmailResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendRawEmailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendRawEmail(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSendRawEmail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendRawEmail",
	}
}

type opSendRawEmailResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSendRawEmailResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSendRawEmailResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ses"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ses"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ses")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addSendRawEmailResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSendRawEmailResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
