// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

type Application struct {

	// The description of the application.
	Description *string

	// The application ID.
	Id *string

	// The application name.
	Name *string
}

// A summary of a configuration profile.
type ConfigurationProfileSummary struct {

	// The application ID.
	ApplicationId *string

	// The ID of the configuration profile.
	Id *string

	// The URI location of the configuration.
	LocationUri *string

	// The name of the configuration profile.
	Name *string

	// The types of validators in the configuration profile.
	ValidatorTypes []ValidatorType
}

// An object that describes a deployment event.
type DeploymentEvent struct {

	// A description of the deployment event. Descriptions include, but are not limited
	// to, the user account or the CloudWatch alarm ARN that initiated a rollback, the
	// percentage of hosts that received the deployment, or in the case of an internal
	// error, a recommendation to attempt a new deployment.
	Description *string

	// The type of deployment event. Deployment event types include the start, stop, or
	// completion of a deployment; a percentage update; the start or stop of a bake
	// period; the start or completion of a rollback.
	EventType DeploymentEventType

	// The date and time the event occurred.
	OccurredAt *time.Time

	// The entity that triggered the deployment event. Events can be triggered by a
	// user, AWS AppConfig, an Amazon CloudWatch alarm, or an internal error.
	TriggeredBy TriggeredBy
}

type DeploymentStrategy struct {

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int32

	// The description of the deployment strategy.
	Description *string

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32

	// The percentage of targets that received a deployed configuration during each
	// interval.
	GrowthFactor *float32

	// The algorithm used to define how percentage grew over time.
	GrowthType GrowthType

	// The deployment strategy ID.
	Id *string

	// The name of the deployment strategy.
	Name *string

	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo ReplicateTo
}

// Information about the deployment.
type DeploymentSummary struct {

	// Time the deployment completed.
	CompletedAt *time.Time

	// The name of the configuration.
	ConfigurationName *string

	// The version of the configuration.
	ConfigurationVersion *string

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int32

	// The sequence number of the deployment.
	DeploymentNumber *int32

	// The amount of time AppConfig monitors for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor *float32

	// The algorithm used to define how percentage grows over time.
	GrowthType GrowthType

	// The percentage of targets for which the deployment is available.
	PercentageComplete *float32

	// Time the deployment started.
	StartedAt *time.Time

	// The state of the deployment.
	State DeploymentState
}

type Environment struct {

	// The application ID.
	ApplicationId *string

	// The description of the environment.
	Description *string

	// The environment ID.
	Id *string

	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []*Monitor

	// The name of the environment.
	Name *string

	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State EnvironmentState
}

// Information about the configuration.
type HostedConfigurationVersionSummary struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile ID.
	ConfigurationProfileId *string

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string

	// A description of the configuration.
	Description *string

	// The configuration version.
	VersionNumber *int32
}

// Amazon CloudWatch alarms to monitor during the deployment process.
type Monitor struct {

	// ARN of the Amazon CloudWatch alarm.
	AlarmArn *string

	// ARN of an IAM role for AppConfig to monitor AlarmArn.
	AlarmRoleArn *string
}

// A validator provides a syntactic or semantic check to ensure the configuration
// you want to deploy functions as intended. To validate your application
// configuration data, you provide a schema or a Lambda function that runs against
// the configuration. The configuration deployment or update can only proceed when
// the configuration data is valid.
type Validator struct {

	// Either the JSON Schema content or the Amazon Resource Name (ARN) of an AWS
	// Lambda function.
	//
	// This member is required.
	Content *string

	// AppConfig supports validators of type JSON_SCHEMA and LAMBDA
	//
	// This member is required.
	Type ValidatorType
}