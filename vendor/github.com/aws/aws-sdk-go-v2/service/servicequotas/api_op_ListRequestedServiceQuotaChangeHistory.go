// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the quota increase requests for the specified Amazon Web Service.
func (c *Client) ListRequestedServiceQuotaChangeHistory(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error) {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRequestedServiceQuotaChangeHistory", params, optFns, c.addOperationListRequestedServiceQuotaChangeHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRequestedServiceQuotaChangeHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRequestedServiceQuotaChangeHistoryInput struct {

	// Specifies the maximum number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results. An API operation
	// can return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies a value for receiving additional results after you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	// Specifies at which level within the Amazon Web Services account the quota
	// request applies to.
	QuotaRequestedAtLevel types.AppliedLevelEnum

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServices operation.
	ServiceCode *string

	// Specifies that you want to filter the results to only the requests with the
	// matching status.
	Status types.RequestStatus

	noSmithyDocumentSerde
}

type ListRequestedServiceQuotaChangeHistoryOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Information about the quota increase requests.
	RequestedQuotas []types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRequestedServiceQuotaChangeHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRequestedServiceQuotaChangeHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRequestedServiceQuotaChangeHistory{}, middleware.After)
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
	if err = addListRequestedServiceQuotaChangeHistoryResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistory(options.Region), middleware.Before); err != nil {
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

// ListRequestedServiceQuotaChangeHistoryAPIClient is a client that implements the
// ListRequestedServiceQuotaChangeHistory operation.
type ListRequestedServiceQuotaChangeHistoryAPIClient interface {
	ListRequestedServiceQuotaChangeHistory(context.Context, *ListRequestedServiceQuotaChangeHistoryInput, ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error)
}

var _ ListRequestedServiceQuotaChangeHistoryAPIClient = (*Client)(nil)

// ListRequestedServiceQuotaChangeHistoryPaginatorOptions is the paginator options
// for ListRequestedServiceQuotaChangeHistory
type ListRequestedServiceQuotaChangeHistoryPaginatorOptions struct {
	// Specifies the maximum number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results. An API operation
	// can return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRequestedServiceQuotaChangeHistoryPaginator is a paginator for
// ListRequestedServiceQuotaChangeHistory
type ListRequestedServiceQuotaChangeHistoryPaginator struct {
	options   ListRequestedServiceQuotaChangeHistoryPaginatorOptions
	client    ListRequestedServiceQuotaChangeHistoryAPIClient
	params    *ListRequestedServiceQuotaChangeHistoryInput
	nextToken *string
	firstPage bool
}

// NewListRequestedServiceQuotaChangeHistoryPaginator returns a new
// ListRequestedServiceQuotaChangeHistoryPaginator
func NewListRequestedServiceQuotaChangeHistoryPaginator(client ListRequestedServiceQuotaChangeHistoryAPIClient, params *ListRequestedServiceQuotaChangeHistoryInput, optFns ...func(*ListRequestedServiceQuotaChangeHistoryPaginatorOptions)) *ListRequestedServiceQuotaChangeHistoryPaginator {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryInput{}
	}

	options := ListRequestedServiceQuotaChangeHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRequestedServiceQuotaChangeHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRequestedServiceQuotaChangeHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRequestedServiceQuotaChangeHistory page.
func (p *ListRequestedServiceQuotaChangeHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRequestedServiceQuotaChangeHistory(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListRequestedServiceQuotaChangeHistory",
	}
}

type opListRequestedServiceQuotaChangeHistoryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListRequestedServiceQuotaChangeHistoryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListRequestedServiceQuotaChangeHistoryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "servicequotas"
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
				signingName = "servicequotas"
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
				v4aScheme.SigningName = aws.String("servicequotas")
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

func addListRequestedServiceQuotaChangeHistoryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListRequestedServiceQuotaChangeHistoryResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
