module github.com/aws/aws-sdk-go-v2/feature/ec2/imds/internal/configtesting

go 1.15

require (
	github.com/aws/aws-sdk-go-v2/config v1.18.32
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.7
)

replace github.com/aws/aws-sdk-go-v2 => ../../../../../

replace github.com/aws/aws-sdk-go-v2/config => ../../../../../config/

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../../../credentials/

replace github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../../../../../feature/ec2/imds/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/internal/ini => ../../../../../internal/ini/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../../../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/sso => ../../../../../service/sso/

replace github.com/aws/aws-sdk-go-v2/service/ssooidc => ../../../../../service/ssooidc/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../../../../../service/sts/
