module github.com/aws/aws-sdk-go-v2/credentials

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.20.0
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.7
	github.com/aws/aws-sdk-go-v2/service/sso v1.13.1
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.15.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.21.1
	github.com/aws/smithy-go v1.14.0
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../

replace github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../feature/ec2/imds/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/sso => ../service/sso/

replace github.com/aws/aws-sdk-go-v2/service/ssooidc => ../service/ssooidc/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../service/sts/
