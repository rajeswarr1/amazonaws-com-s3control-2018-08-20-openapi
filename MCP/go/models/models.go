package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteBucketTaggingRequest represents the DeleteBucketTaggingRequest schema from the OpenAPI specification
type DeleteBucketTaggingRequest struct {
}

// EncryptionConfiguration represents the EncryptionConfiguration schema from the OpenAPI specification
type EncryptionConfiguration struct {
	Replicakmskeyid interface{} `json:"ReplicaKmsKeyID,omitempty"`
}

// ListAccessPointsForObjectLambdaRequest represents the ListAccessPointsForObjectLambdaRequest schema from the OpenAPI specification
type ListAccessPointsForObjectLambdaRequest struct {
}

// S3InitiateRestoreObjectOperation represents the S3InitiateRestoreObjectOperation schema from the OpenAPI specification
type S3InitiateRestoreObjectOperation struct {
	Expirationindays interface{} `json:"ExpirationInDays,omitempty"`
	Glacierjobtier interface{} `json:"GlacierJobTier,omitempty"`
}

// S3ManifestOutputLocation represents the S3ManifestOutputLocation schema from the OpenAPI specification
type S3ManifestOutputLocation struct {
	Expectedmanifestbucketowner interface{} `json:"ExpectedManifestBucketOwner,omitempty"`
	Manifestencryption interface{} `json:"ManifestEncryption,omitempty"`
	Manifestformat interface{} `json:"ManifestFormat"`
	Manifestprefix interface{} `json:"ManifestPrefix,omitempty"`
	Bucket interface{} `json:"Bucket"`
}

// Destination represents the Destination schema from the OpenAPI specification
type Destination struct {
	Encryptionconfiguration interface{} `json:"EncryptionConfiguration,omitempty"`
	Metrics interface{} `json:"Metrics,omitempty"`
	Replicationtime interface{} `json:"ReplicationTime,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Accesscontroltranslation interface{} `json:"AccessControlTranslation,omitempty"`
	Account interface{} `json:"Account,omitempty"`
	Bucket interface{} `json:"Bucket"`
}

// AccountLevel represents the AccountLevel schema from the OpenAPI specification
type AccountLevel struct {
	Advancedcostoptimizationmetrics interface{} `json:"AdvancedCostOptimizationMetrics,omitempty"`
	Advanceddataprotectionmetrics interface{} `json:"AdvancedDataProtectionMetrics,omitempty"`
	Bucketlevel interface{} `json:"BucketLevel"`
	Detailedstatuscodesmetrics interface{} `json:"DetailedStatusCodesMetrics,omitempty"`
	Activitymetrics interface{} `json:"ActivityMetrics,omitempty"`
}

// S3ReplicateObjectOperation represents the S3ReplicateObjectOperation schema from the OpenAPI specification
type S3ReplicateObjectOperation struct {
}

// ObjectLambdaAccessPointAlias represents the ObjectLambdaAccessPointAlias schema from the OpenAPI specification
type ObjectLambdaAccessPointAlias struct {
	Value interface{} `json:"Value,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetAccessPointForObjectLambdaRequest represents the GetAccessPointForObjectLambdaRequest schema from the OpenAPI specification
type GetAccessPointForObjectLambdaRequest struct {
}

// JobManifest represents the JobManifest schema from the OpenAPI specification
type JobManifest struct {
	Location interface{} `json:"Location"`
	Spec interface{} `json:"Spec"`
}

// CreateMultiRegionAccessPointResult represents the CreateMultiRegionAccessPointResult schema from the OpenAPI specification
type CreateMultiRegionAccessPointResult struct {
	Requesttokenarn interface{} `json:"RequestTokenARN,omitempty"`
}

// CloudWatchMetrics represents the CloudWatchMetrics schema from the OpenAPI specification
type CloudWatchMetrics struct {
	Isenabled interface{} `json:"IsEnabled"`
}

// JobFailure represents the JobFailure schema from the OpenAPI specification
type JobFailure struct {
	Failurecode interface{} `json:"FailureCode,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// ListMultiRegionAccessPointsRequest represents the ListMultiRegionAccessPointsRequest schema from the OpenAPI specification
type ListMultiRegionAccessPointsRequest struct {
}

// S3DeleteObjectTaggingOperation represents the S3DeleteObjectTaggingOperation schema from the OpenAPI specification
type S3DeleteObjectTaggingOperation struct {
}

// MultiRegionAccessPointsAsyncResponse represents the MultiRegionAccessPointsAsyncResponse schema from the OpenAPI specification
type MultiRegionAccessPointsAsyncResponse struct {
	Regions interface{} `json:"Regions,omitempty"`
}

// CreateJobRequest represents the CreateJobRequest schema from the OpenAPI specification
type CreateJobRequest struct {
	Confirmationrequired interface{} `json:"ConfirmationRequired,omitempty"`
	Operation interface{} `json:"Operation"`
	Priority interface{} `json:"Priority"`
	Tags interface{} `json:"Tags,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
	Description interface{} `json:"Description,omitempty"`
	Manifest interface{} `json:"Manifest,omitempty"`
	Manifestgenerator interface{} `json:"ManifestGenerator,omitempty"`
	Report interface{} `json:"Report"`
	Rolearn interface{} `json:"RoleArn"`
}

// SSES3 represents the SSES3 schema from the OpenAPI specification
type SSES3 struct {
}

// S3ObjectOwner represents the S3ObjectOwner schema from the OpenAPI specification
type S3ObjectOwner struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Id interface{} `json:"ID,omitempty"`
}

// ListAccessPointsResult represents the ListAccessPointsResult schema from the OpenAPI specification
type ListAccessPointsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Accesspointlist interface{} `json:"AccessPointList,omitempty"`
}

// JobManifestSpec represents the JobManifestSpec schema from the OpenAPI specification
type JobManifestSpec struct {
	Fields interface{} `json:"Fields,omitempty"`
	Format interface{} `json:"Format"`
}

// GetBucketTaggingResult represents the GetBucketTaggingResult schema from the OpenAPI specification
type GetBucketTaggingResult struct {
	Tagset interface{} `json:"TagSet"`
}

// ExistingObjectReplication represents the ExistingObjectReplication schema from the OpenAPI specification
type ExistingObjectReplication struct {
	Status interface{} `json:"Status"`
}

// GetAccessPointPolicyForObjectLambdaRequest represents the GetAccessPointPolicyForObjectLambdaRequest schema from the OpenAPI specification
type GetAccessPointPolicyForObjectLambdaRequest struct {
}

// S3AccessControlPolicy represents the S3AccessControlPolicy schema from the OpenAPI specification
type S3AccessControlPolicy struct {
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Cannedaccesscontrollist interface{} `json:"CannedAccessControlList,omitempty"`
}

// SelectionCriteria represents the SelectionCriteria schema from the OpenAPI specification
type SelectionCriteria struct {
	Maxdepth interface{} `json:"MaxDepth,omitempty"`
	Minstoragebytespercentage interface{} `json:"MinStorageBytesPercentage,omitempty"`
	Delimiter interface{} `json:"Delimiter,omitempty"`
}

// PutMultiRegionAccessPointPolicyResult represents the PutMultiRegionAccessPointPolicyResult schema from the OpenAPI specification
type PutMultiRegionAccessPointPolicyResult struct {
	Requesttokenarn interface{} `json:"RequestTokenARN,omitempty"`
}

// SSEKMSEncryption represents the SSEKMSEncryption schema from the OpenAPI specification
type SSEKMSEncryption struct {
	Keyid interface{} `json:"KeyId"`
}

// AdvancedDataProtectionMetrics represents the AdvancedDataProtectionMetrics schema from the OpenAPI specification
type AdvancedDataProtectionMetrics struct {
	Isenabled interface{} `json:"IsEnabled,omitempty"`
}

// GetStorageLensConfigurationTaggingResult represents the GetStorageLensConfigurationTaggingResult schema from the OpenAPI specification
type GetStorageLensConfigurationTaggingResult struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// ListStorageLensConfigurationsResult represents the ListStorageLensConfigurationsResult schema from the OpenAPI specification
type ListStorageLensConfigurationsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Storagelensconfigurationlist interface{} `json:"StorageLensConfigurationList,omitempty"`
}

// Exclude represents the Exclude schema from the OpenAPI specification
type Exclude struct {
	Buckets interface{} `json:"Buckets,omitempty"`
	Regions interface{} `json:"Regions,omitempty"`
}

// PutBucketPolicyRequest represents the PutBucketPolicyRequest schema from the OpenAPI specification
type PutBucketPolicyRequest struct {
	Policy interface{} `json:"Policy"`
}

// DeleteAccessPointRequest represents the DeleteAccessPointRequest schema from the OpenAPI specification
type DeleteAccessPointRequest struct {
}

// DeleteMarkerReplication represents the DeleteMarkerReplication schema from the OpenAPI specification
type DeleteMarkerReplication struct {
	Status interface{} `json:"Status"`
}

// JobManifestLocation represents the JobManifestLocation schema from the OpenAPI specification
type JobManifestLocation struct {
	Objectversionid interface{} `json:"ObjectVersionId,omitempty"`
	Etag interface{} `json:"ETag"`
	Objectarn interface{} `json:"ObjectArn"`
}

// UpdateJobStatusResult represents the UpdateJobStatusResult schema from the OpenAPI specification
type UpdateJobStatusResult struct {
	Status interface{} `json:"Status,omitempty"`
	Statusupdatereason interface{} `json:"StatusUpdateReason,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// LifecycleRuleFilter represents the LifecycleRuleFilter schema from the OpenAPI specification
type LifecycleRuleFilter struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag S3Tag `json:"Tag,omitempty"` // A container for a key-value name pair.
	And interface{} `json:"And,omitempty"`
	Objectsizegreaterthan interface{} `json:"ObjectSizeGreaterThan,omitempty"`
	Objectsizelessthan interface{} `json:"ObjectSizeLessThan,omitempty"`
}

// S3SetObjectRetentionOperation represents the S3SetObjectRetentionOperation schema from the OpenAPI specification
type S3SetObjectRetentionOperation struct {
	Bypassgovernanceretention interface{} `json:"BypassGovernanceRetention,omitempty"`
	Retention interface{} `json:"Retention"`
}

// EstablishedMultiRegionAccessPointPolicy represents the EstablishedMultiRegionAccessPointPolicy schema from the OpenAPI specification
type EstablishedMultiRegionAccessPointPolicy struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// GetBucketVersioningResult represents the GetBucketVersioningResult schema from the OpenAPI specification
type GetBucketVersioningResult struct {
	Mfadelete interface{} `json:"MFADelete,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GeneratedManifestEncryption represents the GeneratedManifestEncryption schema from the OpenAPI specification
type GeneratedManifestEncryption struct {
	Ssekms interface{} `json:"SSEKMS,omitempty"`
	Sses3 interface{} `json:"SSES3,omitempty"`
}

// SSES3Encryption represents the SSES3Encryption schema from the OpenAPI specification
type SSES3Encryption struct {
}

// GetPublicAccessBlockRequest represents the GetPublicAccessBlockRequest schema from the OpenAPI specification
type GetPublicAccessBlockRequest struct {
}

// DeleteBucketRequest represents the DeleteBucketRequest schema from the OpenAPI specification
type DeleteBucketRequest struct {
}

// DeleteJobTaggingResult represents the DeleteJobTaggingResult schema from the OpenAPI specification
type DeleteJobTaggingResult struct {
}

// ActivityMetrics represents the ActivityMetrics schema from the OpenAPI specification
type ActivityMetrics struct {
	Isenabled interface{} `json:"IsEnabled,omitempty"`
}

// CreateBucketRequest represents the CreateBucketRequest schema from the OpenAPI specification
type CreateBucketRequest struct {
	Createbucketconfiguration interface{} `json:"CreateBucketConfiguration,omitempty"`
}

// SubmitMultiRegionAccessPointRoutesRequest represents the SubmitMultiRegionAccessPointRoutesRequest schema from the OpenAPI specification
type SubmitMultiRegionAccessPointRoutesRequest struct {
	Routeupdates interface{} `json:"RouteUpdates"`
}

// DeleteAccessPointForObjectLambdaRequest represents the DeleteAccessPointForObjectLambdaRequest schema from the OpenAPI specification
type DeleteAccessPointForObjectLambdaRequest struct {
}

// SseKmsEncryptedObjects represents the SseKmsEncryptedObjects schema from the OpenAPI specification
type SseKmsEncryptedObjects struct {
	Status interface{} `json:"Status"`
}

// JobManifestGeneratorFilter represents the JobManifestGeneratorFilter schema from the OpenAPI specification
type JobManifestGeneratorFilter struct {
	Objectreplicationstatuses interface{} `json:"ObjectReplicationStatuses,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Eligibleforreplication interface{} `json:"EligibleForReplication,omitempty"`
}

// PutBucketVersioningRequest represents the PutBucketVersioningRequest schema from the OpenAPI specification
type PutBucketVersioningRequest struct {
	Versioningconfiguration interface{} `json:"VersioningConfiguration"`
}

// ListRegionalBucketsRequest represents the ListRegionalBucketsRequest schema from the OpenAPI specification
type ListRegionalBucketsRequest struct {
}

// GetAccessPointRequest represents the GetAccessPointRequest schema from the OpenAPI specification
type GetAccessPointRequest struct {
}

// S3UserMetadata represents the S3UserMetadata schema from the OpenAPI specification
type S3UserMetadata struct {
}

// GetAccessPointPolicyResult represents the GetAccessPointPolicyResult schema from the OpenAPI specification
type GetAccessPointPolicyResult struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// S3JobManifestGenerator represents the S3JobManifestGenerator schema from the OpenAPI specification
type S3JobManifestGenerator struct {
	Manifestoutputlocation interface{} `json:"ManifestOutputLocation,omitempty"`
	Sourcebucket interface{} `json:"SourceBucket"`
	Enablemanifestoutput interface{} `json:"EnableManifestOutput"`
	Expectedbucketowner interface{} `json:"ExpectedBucketOwner,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
}

// S3ObjectMetadata represents the S3ObjectMetadata schema from the OpenAPI specification
type S3ObjectMetadata struct {
	Httpexpiresdate interface{} `json:"HttpExpiresDate,omitempty"`
	Contentencoding interface{} `json:"ContentEncoding,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
	Ssealgorithm interface{} `json:"SSEAlgorithm,omitempty"`
	Cachecontrol interface{} `json:"CacheControl,omitempty"`
	Contentlanguage interface{} `json:"ContentLanguage,omitempty"`
	Contentlength interface{} `json:"ContentLength,omitempty"`
	Contentmd5 interface{} `json:"ContentMD5,omitempty"`
	Requestercharged interface{} `json:"RequesterCharged,omitempty"`
	Contentdisposition interface{} `json:"ContentDisposition,omitempty"`
	Usermetadata interface{} `json:"UserMetadata,omitempty"`
}

// DescribeJobResult represents the DescribeJobResult schema from the OpenAPI specification
type DescribeJobResult struct {
	Job interface{} `json:"Job,omitempty"`
}

// ListStorageLensConfigurationsRequest represents the ListStorageLensConfigurationsRequest schema from the OpenAPI specification
type ListStorageLensConfigurationsRequest struct {
}

// PutStorageLensConfigurationTaggingResult represents the PutStorageLensConfigurationTaggingResult schema from the OpenAPI specification
type PutStorageLensConfigurationTaggingResult struct {
}

// RegionReport represents the RegionReport schema from the OpenAPI specification
type RegionReport struct {
	Bucket interface{} `json:"Bucket,omitempty"`
	Bucketaccountid interface{} `json:"BucketAccountId,omitempty"`
	Region interface{} `json:"Region,omitempty"`
}

// DeleteMultiRegionAccessPointResult represents the DeleteMultiRegionAccessPointResult schema from the OpenAPI specification
type DeleteMultiRegionAccessPointResult struct {
	Requesttokenarn interface{} `json:"RequestTokenARN,omitempty"`
}

// AccessPoint represents the AccessPoint schema from the OpenAPI specification
type AccessPoint struct {
	Name interface{} `json:"Name"`
	Networkorigin interface{} `json:"NetworkOrigin"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Accesspointarn interface{} `json:"AccessPointArn,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
	Bucket interface{} `json:"Bucket"`
	Bucketaccountid interface{} `json:"BucketAccountId,omitempty"`
}

// GetAccessPointConfigurationForObjectLambdaResult represents the GetAccessPointConfigurationForObjectLambdaResult schema from the OpenAPI specification
type GetAccessPointConfigurationForObjectLambdaResult struct {
	Configuration interface{} `json:"Configuration,omitempty"`
}

// ReplicaModifications represents the ReplicaModifications schema from the OpenAPI specification
type ReplicaModifications struct {
	Status interface{} `json:"Status"`
}

// StorageLensAwsOrg represents the StorageLensAwsOrg schema from the OpenAPI specification
type StorageLensAwsOrg struct {
	Arn interface{} `json:"Arn"`
}

// DeleteMultiRegionAccessPointRequest represents the DeleteMultiRegionAccessPointRequest schema from the OpenAPI specification
type DeleteMultiRegionAccessPointRequest struct {
	Details interface{} `json:"Details"`
	Clienttoken interface{} `json:"ClientToken"`
}

// DeleteBucketPolicyRequest represents the DeleteBucketPolicyRequest schema from the OpenAPI specification
type DeleteBucketPolicyRequest struct {
}

// DeletePublicAccessBlockRequest represents the DeletePublicAccessBlockRequest schema from the OpenAPI specification
type DeletePublicAccessBlockRequest struct {
}

// StorageLensTag represents the StorageLensTag schema from the OpenAPI specification
type StorageLensTag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ObjectLambdaConfiguration represents the ObjectLambdaConfiguration schema from the OpenAPI specification
type ObjectLambdaConfiguration struct {
	Allowedfeatures interface{} `json:"AllowedFeatures,omitempty"`
	Cloudwatchmetricsenabled interface{} `json:"CloudWatchMetricsEnabled,omitempty"`
	Supportingaccesspoint interface{} `json:"SupportingAccessPoint"`
	Transformationconfigurations interface{} `json:"TransformationConfigurations"`
}

// StorageLensDataExportEncryption represents the StorageLensDataExportEncryption schema from the OpenAPI specification
type StorageLensDataExportEncryption struct {
	Ssekms interface{} `json:"SSEKMS,omitempty"`
	Sses3 interface{} `json:"SSES3,omitempty"`
}

// GetBucketVersioningRequest represents the GetBucketVersioningRequest schema from the OpenAPI specification
type GetBucketVersioningRequest struct {
}

// GetAccessPointPolicyStatusForObjectLambdaResult represents the GetAccessPointPolicyStatusForObjectLambdaResult schema from the OpenAPI specification
type GetAccessPointPolicyStatusForObjectLambdaResult struct {
	Policystatus PolicyStatus `json:"PolicyStatus,omitempty"` // Indicates whether this access point policy is public. For more information about how Amazon S3 evaluates policies to determine whether they are public, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The Meaning of "Public"</a> in the <i>Amazon S3 User Guide</i>.
}

// GetMultiRegionAccessPointRoutesRequest represents the GetMultiRegionAccessPointRoutesRequest schema from the OpenAPI specification
type GetMultiRegionAccessPointRoutesRequest struct {
}

// ListAccessPointsRequest represents the ListAccessPointsRequest schema from the OpenAPI specification
type ListAccessPointsRequest struct {
}

// GetJobTaggingResult represents the GetJobTaggingResult schema from the OpenAPI specification
type GetJobTaggingResult struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// S3SetObjectAclOperation represents the S3SetObjectAclOperation schema from the OpenAPI specification
type S3SetObjectAclOperation struct {
	Accesscontrolpolicy interface{} `json:"AccessControlPolicy,omitempty"`
}

// ListJobsResult represents the ListJobsResult schema from the OpenAPI specification
type ListJobsResult struct {
	Jobs interface{} `json:"Jobs,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Region represents the Region schema from the OpenAPI specification
type Region struct {
	Bucket interface{} `json:"Bucket"`
	Bucketaccountid interface{} `json:"BucketAccountId,omitempty"`
}

// GetAccessPointPolicyStatusResult represents the GetAccessPointPolicyStatusResult schema from the OpenAPI specification
type GetAccessPointPolicyStatusResult struct {
	Policystatus interface{} `json:"PolicyStatus,omitempty"`
}

// AsyncErrorDetails represents the AsyncErrorDetails schema from the OpenAPI specification
type AsyncErrorDetails struct {
	Message interface{} `json:"Message,omitempty"`
	Requestid interface{} `json:"RequestId,omitempty"`
	Resource interface{} `json:"Resource,omitempty"`
	Code interface{} `json:"Code,omitempty"`
}

// LifecycleRule represents the LifecycleRule schema from the OpenAPI specification
type LifecycleRule struct {
	Expiration interface{} `json:"Expiration,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
	Id interface{} `json:"ID,omitempty"`
	Noncurrentversionexpiration interface{} `json:"NoncurrentVersionExpiration,omitempty"`
	Noncurrentversiontransitions interface{} `json:"NoncurrentVersionTransitions,omitempty"`
	Status interface{} `json:"Status"`
	Transitions interface{} `json:"Transitions,omitempty"`
	Abortincompletemultipartupload interface{} `json:"AbortIncompleteMultipartUpload,omitempty"`
}

// PutAccessPointConfigurationForObjectLambdaRequest represents the PutAccessPointConfigurationForObjectLambdaRequest schema from the OpenAPI specification
type PutAccessPointConfigurationForObjectLambdaRequest struct {
	Configuration interface{} `json:"Configuration"`
}

// PutMultiRegionAccessPointPolicyInput represents the PutMultiRegionAccessPointPolicyInput schema from the OpenAPI specification
type PutMultiRegionAccessPointPolicyInput struct {
	Policy interface{} `json:"Policy"`
	Name interface{} `json:"Name"`
}

// DescribeMultiRegionAccessPointOperationResult represents the DescribeMultiRegionAccessPointOperationResult schema from the OpenAPI specification
type DescribeMultiRegionAccessPointOperationResult struct {
	Asyncoperation interface{} `json:"AsyncOperation,omitempty"`
}

// GetMultiRegionAccessPointPolicyStatusRequest represents the GetMultiRegionAccessPointPolicyStatusRequest schema from the OpenAPI specification
type GetMultiRegionAccessPointPolicyStatusRequest struct {
}

// MultiRegionAccessPointPolicyDocument represents the MultiRegionAccessPointPolicyDocument schema from the OpenAPI specification
type MultiRegionAccessPointPolicyDocument struct {
	Established interface{} `json:"Established,omitempty"`
	Proposed interface{} `json:"Proposed,omitempty"`
}

// UpdateJobPriorityResult represents the UpdateJobPriorityResult schema from the OpenAPI specification
type UpdateJobPriorityResult struct {
	Jobid interface{} `json:"JobId"`
	Priority interface{} `json:"Priority"`
}

// GetMultiRegionAccessPointPolicyResult represents the GetMultiRegionAccessPointPolicyResult schema from the OpenAPI specification
type GetMultiRegionAccessPointPolicyResult struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// GetJobTaggingRequest represents the GetJobTaggingRequest schema from the OpenAPI specification
type GetJobTaggingRequest struct {
}

// NoncurrentVersionExpiration represents the NoncurrentVersionExpiration schema from the OpenAPI specification
type NoncurrentVersionExpiration struct {
	Newernoncurrentversions interface{} `json:"NewerNoncurrentVersions,omitempty"`
	Noncurrentdays interface{} `json:"NoncurrentDays,omitempty"`
}

// Endpoints represents the Endpoints schema from the OpenAPI specification
type Endpoints struct {
}

// UpdateJobStatusRequest represents the UpdateJobStatusRequest schema from the OpenAPI specification
type UpdateJobStatusRequest struct {
}

// ListRegionalBucketsResult represents the ListRegionalBucketsResult schema from the OpenAPI specification
type ListRegionalBucketsResult struct {
	Regionalbucketlist interface{} `json:"RegionalBucketList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// S3ObjectLockLegalHold represents the S3ObjectLockLegalHold schema from the OpenAPI specification
type S3ObjectLockLegalHold struct {
	Status interface{} `json:"Status"`
}

// GetBucketResult represents the GetBucketResult schema from the OpenAPI specification
type GetBucketResult struct {
	Bucket interface{} `json:"Bucket,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Publicaccessblockenabled interface{} `json:"PublicAccessBlockEnabled,omitempty"`
}

// LifecycleRuleAndOperator represents the LifecycleRuleAndOperator schema from the OpenAPI specification
type LifecycleRuleAndOperator struct {
	Objectsizegreaterthan interface{} `json:"ObjectSizeGreaterThan,omitempty"`
	Objectsizelessthan interface{} `json:"ObjectSizeLessThan,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// JobReport represents the JobReport schema from the OpenAPI specification
type JobReport struct {
	Format interface{} `json:"Format,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Reportscope interface{} `json:"ReportScope,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
	Enabled interface{} `json:"Enabled"`
}

// ReplicationTime represents the ReplicationTime schema from the OpenAPI specification
type ReplicationTime struct {
	Status interface{} `json:"Status"`
	Time interface{} `json:"Time"`
}

// GetAccessPointResult represents the GetAccessPointResult schema from the OpenAPI specification
type GetAccessPointResult struct {
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
	Bucketaccountid interface{} `json:"BucketAccountId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Publicaccessblockconfiguration PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty"` // <p>The <code>PublicAccessBlock</code> configuration that you want to apply to this Amazon S3 account. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The Meaning of "Public"</a> in the <i>Amazon S3 User Guide</i>.</p> <p>This data type is not supported for Amazon S3 on Outposts.</p>
	Accesspointarn interface{} `json:"AccessPointArn,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Networkorigin interface{} `json:"NetworkOrigin,omitempty"`
}

// BucketLevel represents the BucketLevel schema from the OpenAPI specification
type BucketLevel struct {
	Prefixlevel interface{} `json:"PrefixLevel,omitempty"`
	Activitymetrics interface{} `json:"ActivityMetrics,omitempty"`
	Advancedcostoptimizationmetrics interface{} `json:"AdvancedCostOptimizationMetrics,omitempty"`
	Advanceddataprotectionmetrics interface{} `json:"AdvancedDataProtectionMetrics,omitempty"`
	Detailedstatuscodesmetrics interface{} `json:"DetailedStatusCodesMetrics,omitempty"`
}

// PutAccessPointPolicyRequest represents the PutAccessPointPolicyRequest schema from the OpenAPI specification
type PutAccessPointPolicyRequest struct {
	Policy interface{} `json:"Policy"`
}

// AccessControlTranslation represents the AccessControlTranslation schema from the OpenAPI specification
type AccessControlTranslation struct {
	Owner interface{} `json:"Owner"`
}

// S3Grantee represents the S3Grantee schema from the OpenAPI specification
type S3Grantee struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Identifier interface{} `json:"Identifier,omitempty"`
	Typeidentifier interface{} `json:"TypeIdentifier,omitempty"`
}

// GetBucketReplicationResult represents the GetBucketReplicationResult schema from the OpenAPI specification
type GetBucketReplicationResult struct {
	Replicationconfiguration interface{} `json:"ReplicationConfiguration,omitempty"`
}

// JobProgressSummary represents the JobProgressSummary schema from the OpenAPI specification
type JobProgressSummary struct {
	Numberoftasksfailed interface{} `json:"NumberOfTasksFailed,omitempty"`
	Numberoftaskssucceeded interface{} `json:"NumberOfTasksSucceeded,omitempty"`
	Timers interface{} `json:"Timers,omitempty"`
	Totalnumberoftasks interface{} `json:"TotalNumberOfTasks,omitempty"`
}

// AsyncResponseDetails represents the AsyncResponseDetails schema from the OpenAPI specification
type AsyncResponseDetails struct {
	Errordetails interface{} `json:"ErrorDetails,omitempty"`
	Multiregionaccesspointdetails interface{} `json:"MultiRegionAccessPointDetails,omitempty"`
}

// DeleteJobTaggingRequest represents the DeleteJobTaggingRequest schema from the OpenAPI specification
type DeleteJobTaggingRequest struct {
}

// GetMultiRegionAccessPointResult represents the GetMultiRegionAccessPointResult schema from the OpenAPI specification
type GetMultiRegionAccessPointResult struct {
	Accesspoint interface{} `json:"AccessPoint,omitempty"`
}

// MultiRegionAccessPointRoute represents the MultiRegionAccessPointRoute schema from the OpenAPI specification
type MultiRegionAccessPointRoute struct {
	Bucket interface{} `json:"Bucket,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Trafficdialpercentage interface{} `json:"TrafficDialPercentage"`
}

// S3SetObjectTaggingOperation represents the S3SetObjectTaggingOperation schema from the OpenAPI specification
type S3SetObjectTaggingOperation struct {
	Tagset interface{} `json:"TagSet,omitempty"`
}

// ListAccessPointsForObjectLambdaResult represents the ListAccessPointsForObjectLambdaResult schema from the OpenAPI specification
type ListAccessPointsForObjectLambdaResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Objectlambdaaccesspointlist interface{} `json:"ObjectLambdaAccessPointList,omitempty"`
}

// GetStorageLensConfigurationRequest represents the GetStorageLensConfigurationRequest schema from the OpenAPI specification
type GetStorageLensConfigurationRequest struct {
}

// AsyncOperation represents the AsyncOperation schema from the OpenAPI specification
type AsyncOperation struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Requestparameters interface{} `json:"RequestParameters,omitempty"`
	Requeststatus interface{} `json:"RequestStatus,omitempty"`
	Requesttokenarn interface{} `json:"RequestTokenARN,omitempty"`
	Responsedetails interface{} `json:"ResponseDetails,omitempty"`
}

// GetAccessPointForObjectLambdaResult represents the GetAccessPointForObjectLambdaResult schema from the OpenAPI specification
type GetAccessPointForObjectLambdaResult struct {
	Name interface{} `json:"Name,omitempty"`
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
}

// PrefixLevel represents the PrefixLevel schema from the OpenAPI specification
type PrefixLevel struct {
	Storagemetrics interface{} `json:"StorageMetrics"`
}

// GetBucketTaggingRequest represents the GetBucketTaggingRequest schema from the OpenAPI specification
type GetBucketTaggingRequest struct {
}

// Tagging represents the Tagging schema from the OpenAPI specification
type Tagging struct {
	Tagset interface{} `json:"TagSet"`
}

// CreateAccessPointForObjectLambdaResult represents the CreateAccessPointForObjectLambdaResult schema from the OpenAPI specification
type CreateAccessPointForObjectLambdaResult struct {
	Alias interface{} `json:"Alias,omitempty"`
	Objectlambdaaccesspointarn interface{} `json:"ObjectLambdaAccessPointArn,omitempty"`
}

// GetMultiRegionAccessPointPolicyRequest represents the GetMultiRegionAccessPointPolicyRequest schema from the OpenAPI specification
type GetMultiRegionAccessPointPolicyRequest struct {
}

// SSEKMS represents the SSEKMS schema from the OpenAPI specification
type SSEKMS struct {
	Keyid interface{} `json:"KeyId"`
}

// DeleteAccessPointPolicyRequest represents the DeleteAccessPointPolicyRequest schema from the OpenAPI specification
type DeleteAccessPointPolicyRequest struct {
}

// GetStorageLensConfigurationTaggingRequest represents the GetStorageLensConfigurationTaggingRequest schema from the OpenAPI specification
type GetStorageLensConfigurationTaggingRequest struct {
}

// CreateBucketConfiguration represents the CreateBucketConfiguration schema from the OpenAPI specification
type CreateBucketConfiguration struct {
	Locationconstraint interface{} `json:"LocationConstraint,omitempty"`
}

// LifecycleExpiration represents the LifecycleExpiration schema from the OpenAPI specification
type LifecycleExpiration struct {
	Date interface{} `json:"Date,omitempty"`
	Days interface{} `json:"Days,omitempty"`
	Expiredobjectdeletemarker interface{} `json:"ExpiredObjectDeleteMarker,omitempty"`
}

// S3AccessControlList represents the S3AccessControlList schema from the OpenAPI specification
type S3AccessControlList struct {
	Owner interface{} `json:"Owner"`
	Grants interface{} `json:"Grants,omitempty"`
}

// PutBucketTaggingRequest represents the PutBucketTaggingRequest schema from the OpenAPI specification
type PutBucketTaggingRequest struct {
	Tagging interface{} `json:"Tagging"`
}

// AbortIncompleteMultipartUpload represents the AbortIncompleteMultipartUpload schema from the OpenAPI specification
type AbortIncompleteMultipartUpload struct {
	Daysafterinitiation interface{} `json:"DaysAfterInitiation,omitempty"`
}

// JobDescriptor represents the JobDescriptor schema from the OpenAPI specification
type JobDescriptor struct {
	Suspendeddate interface{} `json:"SuspendedDate,omitempty"`
	Confirmationrequired interface{} `json:"ConfirmationRequired,omitempty"`
	Jobarn interface{} `json:"JobArn,omitempty"`
	Progresssummary interface{} `json:"ProgressSummary,omitempty"`
	Generatedmanifestdescriptor interface{} `json:"GeneratedManifestDescriptor,omitempty"`
	Report interface{} `json:"Report,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Suspendedcause interface{} `json:"SuspendedCause,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Terminationdate interface{} `json:"TerminationDate,omitempty"`
	Failurereasons interface{} `json:"FailureReasons,omitempty"`
	Manifest interface{} `json:"Manifest,omitempty"`
	Statusupdatereason interface{} `json:"StatusUpdateReason,omitempty"`
	Manifestgenerator interface{} `json:"ManifestGenerator,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
}

// CreateJobResult represents the CreateJobResult schema from the OpenAPI specification
type CreateJobResult struct {
	Jobid interface{} `json:"JobId,omitempty"`
}

// StorageLensDataExport represents the StorageLensDataExport schema from the OpenAPI specification
type StorageLensDataExport struct {
	Cloudwatchmetrics interface{} `json:"CloudWatchMetrics,omitempty"`
	S3bucketdestination interface{} `json:"S3BucketDestination,omitempty"`
}

// GetBucketLifecycleConfigurationResult represents the GetBucketLifecycleConfigurationResult schema from the OpenAPI specification
type GetBucketLifecycleConfigurationResult struct {
	Rules interface{} `json:"Rules,omitempty"`
}

// PutBucketReplicationRequest represents the PutBucketReplicationRequest schema from the OpenAPI specification
type PutBucketReplicationRequest struct {
	Replicationconfiguration interface{} `json:"ReplicationConfiguration"`
}

// PutStorageLensConfigurationTaggingRequest represents the PutStorageLensConfigurationTaggingRequest schema from the OpenAPI specification
type PutStorageLensConfigurationTaggingRequest struct {
	Tags interface{} `json:"Tags"`
}

// PublicAccessBlockConfiguration represents the PublicAccessBlockConfiguration schema from the OpenAPI specification
type PublicAccessBlockConfiguration struct {
	Blockpublicacls interface{} `json:"BlockPublicAcls,omitempty"`
	Blockpublicpolicy interface{} `json:"BlockPublicPolicy,omitempty"`
	Ignorepublicacls interface{} `json:"IgnorePublicAcls,omitempty"`
	Restrictpublicbuckets interface{} `json:"RestrictPublicBuckets,omitempty"`
}

// ObjectLambdaContentTransformation represents the ObjectLambdaContentTransformation schema from the OpenAPI specification
type ObjectLambdaContentTransformation struct {
	Awslambda interface{} `json:"AwsLambda,omitempty"`
}

// PolicyStatus represents the PolicyStatus schema from the OpenAPI specification
type PolicyStatus struct {
	Ispublic interface{} `json:"IsPublic,omitempty"`
}

// SourceSelectionCriteria represents the SourceSelectionCriteria schema from the OpenAPI specification
type SourceSelectionCriteria struct {
	Replicamodifications interface{} `json:"ReplicaModifications,omitempty"`
	Ssekmsencryptedobjects interface{} `json:"SseKmsEncryptedObjects,omitempty"`
}

// S3Tag represents the S3Tag schema from the OpenAPI specification
type S3Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// GetAccessPointPolicyStatusRequest represents the GetAccessPointPolicyStatusRequest schema from the OpenAPI specification
type GetAccessPointPolicyStatusRequest struct {
}

// MultiRegionAccessPointReport represents the MultiRegionAccessPointReport schema from the OpenAPI specification
type MultiRegionAccessPointReport struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Publicaccessblock PublicAccessBlockConfiguration `json:"PublicAccessBlock,omitempty"` // <p>The <code>PublicAccessBlock</code> configuration that you want to apply to this Amazon S3 account. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The Meaning of "Public"</a> in the <i>Amazon S3 User Guide</i>.</p> <p>This data type is not supported for Amazon S3 on Outposts.</p>
	Regions interface{} `json:"Regions,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
}

// ListStorageLensConfigurationEntry represents the ListStorageLensConfigurationEntry schema from the OpenAPI specification
type ListStorageLensConfigurationEntry struct {
	Storagelensarn interface{} `json:"StorageLensArn"`
	Homeregion interface{} `json:"HomeRegion"`
	Id interface{} `json:"Id"`
	Isenabled interface{} `json:"IsEnabled,omitempty"`
}

// GetMultiRegionAccessPointPolicyStatusResult represents the GetMultiRegionAccessPointPolicyStatusResult schema from the OpenAPI specification
type GetMultiRegionAccessPointPolicyStatusResult struct {
	Established PolicyStatus `json:"Established,omitempty"` // Indicates whether this access point policy is public. For more information about how Amazon S3 evaluates policies to determine whether they are public, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The Meaning of "Public"</a> in the <i>Amazon S3 User Guide</i>.
}

// RegionalBucket represents the RegionalBucket schema from the OpenAPI specification
type RegionalBucket struct {
	Bucketarn interface{} `json:"BucketArn,omitempty"`
	Creationdate interface{} `json:"CreationDate"`
	Outpostid interface{} `json:"OutpostId,omitempty"`
	Publicaccessblockenabled interface{} `json:"PublicAccessBlockEnabled"`
	Bucket interface{} `json:"Bucket"`
}

// ReplicationRuleAndOperator represents the ReplicationRuleAndOperator schema from the OpenAPI specification
type ReplicationRuleAndOperator struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetAccessPointPolicyRequest represents the GetAccessPointPolicyRequest schema from the OpenAPI specification
type GetAccessPointPolicyRequest struct {
}

// ReplicationRuleFilter represents the ReplicationRuleFilter schema from the OpenAPI specification
type ReplicationRuleFilter struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag S3Tag `json:"Tag,omitempty"` // A container for a key-value name pair.
	And interface{} `json:"And,omitempty"`
}

// GetMultiRegionAccessPointRequest represents the GetMultiRegionAccessPointRequest schema from the OpenAPI specification
type GetMultiRegionAccessPointRequest struct {
}

// PutJobTaggingRequest represents the PutJobTaggingRequest schema from the OpenAPI specification
type PutJobTaggingRequest struct {
	Tags interface{} `json:"Tags"`
}

// JobListDescriptor represents the JobListDescriptor schema from the OpenAPI specification
type JobListDescriptor struct {
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Progresssummary interface{} `json:"ProgressSummary,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Terminationdate interface{} `json:"TerminationDate,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// GetBucketPolicyResult represents the GetBucketPolicyResult schema from the OpenAPI specification
type GetBucketPolicyResult struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// GetStorageLensConfigurationResult represents the GetStorageLensConfigurationResult schema from the OpenAPI specification
type GetStorageLensConfigurationResult struct {
	Storagelensconfiguration interface{} `json:"StorageLensConfiguration,omitempty"`
}

// StorageLensConfiguration represents the StorageLensConfiguration schema from the OpenAPI specification
type StorageLensConfiguration struct {
	Storagelensarn interface{} `json:"StorageLensArn,omitempty"`
	Accountlevel interface{} `json:"AccountLevel"`
	Awsorg interface{} `json:"AwsOrg,omitempty"`
	Dataexport interface{} `json:"DataExport,omitempty"`
	Exclude interface{} `json:"Exclude,omitempty"`
	Id interface{} `json:"Id"`
	Include interface{} `json:"Include,omitempty"`
	Isenabled interface{} `json:"IsEnabled"`
}

// GetBucketReplicationRequest represents the GetBucketReplicationRequest schema from the OpenAPI specification
type GetBucketReplicationRequest struct {
}

// ReplicationRule represents the ReplicationRule schema from the OpenAPI specification
type ReplicationRule struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Sourceselectioncriteria interface{} `json:"SourceSelectionCriteria,omitempty"`
	Destination interface{} `json:"Destination"`
	Priority interface{} `json:"Priority,omitempty"`
	Bucket interface{} `json:"Bucket"`
	Existingobjectreplication interface{} `json:"ExistingObjectReplication,omitempty"`
	Status interface{} `json:"Status"`
	Deletemarkerreplication interface{} `json:"DeleteMarkerReplication,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
	Id interface{} `json:"ID,omitempty"`
}

// VersioningConfiguration represents the VersioningConfiguration schema from the OpenAPI specification
type VersioningConfiguration struct {
	Mfadelete interface{} `json:"MFADelete,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// S3Grant represents the S3Grant schema from the OpenAPI specification
type S3Grant struct {
	Permission interface{} `json:"Permission,omitempty"`
	Grantee interface{} `json:"Grantee,omitempty"`
}

// DescribeJobRequest represents the DescribeJobRequest schema from the OpenAPI specification
type DescribeJobRequest struct {
}

// DeleteBucketReplicationRequest represents the DeleteBucketReplicationRequest schema from the OpenAPI specification
type DeleteBucketReplicationRequest struct {
}

// GetPublicAccessBlockOutput represents the GetPublicAccessBlockOutput schema from the OpenAPI specification
type GetPublicAccessBlockOutput struct {
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration,omitempty"`
}

// Include represents the Include schema from the OpenAPI specification
type Include struct {
	Buckets interface{} `json:"Buckets,omitempty"`
	Regions interface{} `json:"Regions,omitempty"`
}

// PutStorageLensConfigurationRequest represents the PutStorageLensConfigurationRequest schema from the OpenAPI specification
type PutStorageLensConfigurationRequest struct {
	Storagelensconfiguration interface{} `json:"StorageLensConfiguration"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateAccessPointResult represents the CreateAccessPointResult schema from the OpenAPI specification
type CreateAccessPointResult struct {
	Accesspointarn interface{} `json:"AccessPointArn,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
}

// ReplicationConfiguration represents the ReplicationConfiguration schema from the OpenAPI specification
type ReplicationConfiguration struct {
	Role interface{} `json:"Role"`
	Rules interface{} `json:"Rules"`
}

// VpcConfiguration represents the VpcConfiguration schema from the OpenAPI specification
type VpcConfiguration struct {
	Vpcid interface{} `json:"VpcId"`
}

// ListMultiRegionAccessPointsResult represents the ListMultiRegionAccessPointsResult schema from the OpenAPI specification
type ListMultiRegionAccessPointsResult struct {
	Accesspoints interface{} `json:"AccessPoints,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteBucketLifecycleConfigurationRequest represents the DeleteBucketLifecycleConfigurationRequest schema from the OpenAPI specification
type DeleteBucketLifecycleConfigurationRequest struct {
}

// ProposedMultiRegionAccessPointPolicy represents the ProposedMultiRegionAccessPointPolicy schema from the OpenAPI specification
type ProposedMultiRegionAccessPointPolicy struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// GetBucketLifecycleConfigurationRequest represents the GetBucketLifecycleConfigurationRequest schema from the OpenAPI specification
type GetBucketLifecycleConfigurationRequest struct {
}

// DeleteStorageLensConfigurationRequest represents the DeleteStorageLensConfigurationRequest schema from the OpenAPI specification
type DeleteStorageLensConfigurationRequest struct {
}

// GetAccessPointConfigurationForObjectLambdaRequest represents the GetAccessPointConfigurationForObjectLambdaRequest schema from the OpenAPI specification
type GetAccessPointConfigurationForObjectLambdaRequest struct {
}

// GetMultiRegionAccessPointRoutesResult represents the GetMultiRegionAccessPointRoutesResult schema from the OpenAPI specification
type GetMultiRegionAccessPointRoutesResult struct {
	Routes interface{} `json:"Routes,omitempty"`
	Mrap interface{} `json:"Mrap,omitempty"`
}

// S3BucketDestination represents the S3BucketDestination schema from the OpenAPI specification
type S3BucketDestination struct {
	Format interface{} `json:"Format"`
	Outputschemaversion interface{} `json:"OutputSchemaVersion"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Accountid interface{} `json:"AccountId"`
	Arn interface{} `json:"Arn"`
	Encryption interface{} `json:"Encryption,omitempty"`
}

// S3SetObjectLegalHoldOperation represents the S3SetObjectLegalHoldOperation schema from the OpenAPI specification
type S3SetObjectLegalHoldOperation struct {
	Legalhold interface{} `json:"LegalHold"`
}

// JobManifestGenerator represents the JobManifestGenerator schema from the OpenAPI specification
type JobManifestGenerator struct {
	S3jobmanifestgenerator interface{} `json:"S3JobManifestGenerator,omitempty"`
}

// DeleteAccessPointPolicyForObjectLambdaRequest represents the DeleteAccessPointPolicyForObjectLambdaRequest schema from the OpenAPI specification
type DeleteAccessPointPolicyForObjectLambdaRequest struct {
}

// DeleteStorageLensConfigurationTaggingResult represents the DeleteStorageLensConfigurationTaggingResult schema from the OpenAPI specification
type DeleteStorageLensConfigurationTaggingResult struct {
}

// DeleteStorageLensConfigurationTaggingRequest represents the DeleteStorageLensConfigurationTaggingRequest schema from the OpenAPI specification
type DeleteStorageLensConfigurationTaggingRequest struct {
}

// PutMultiRegionAccessPointPolicyRequest represents the PutMultiRegionAccessPointPolicyRequest schema from the OpenAPI specification
type PutMultiRegionAccessPointPolicyRequest struct {
	Clienttoken interface{} `json:"ClientToken"`
	Details interface{} `json:"Details"`
}

// DetailedStatusCodesMetrics represents the DetailedStatusCodesMetrics schema from the OpenAPI specification
type DetailedStatusCodesMetrics struct {
	Isenabled interface{} `json:"IsEnabled,omitempty"`
}

// Transition represents the Transition schema from the OpenAPI specification
type Transition struct {
	Date interface{} `json:"Date,omitempty"`
	Days interface{} `json:"Days,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
}

// CreateAccessPointRequest represents the CreateAccessPointRequest schema from the OpenAPI specification
type CreateAccessPointRequest struct {
	Bucket interface{} `json:"Bucket"`
	Bucketaccountid interface{} `json:"BucketAccountId,omitempty"`
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration,omitempty"`
	Vpcconfiguration interface{} `json:"VpcConfiguration,omitempty"`
}

// PutPublicAccessBlockRequest represents the PutPublicAccessBlockRequest schema from the OpenAPI specification
type PutPublicAccessBlockRequest struct {
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration"`
}

// AwsLambdaTransformation represents the AwsLambdaTransformation schema from the OpenAPI specification
type AwsLambdaTransformation struct {
	Functionarn interface{} `json:"FunctionArn"`
	Functionpayload interface{} `json:"FunctionPayload,omitempty"`
}

// GetBucketPolicyRequest represents the GetBucketPolicyRequest schema from the OpenAPI specification
type GetBucketPolicyRequest struct {
}

// ObjectLambdaAccessPoint represents the ObjectLambdaAccessPoint schema from the OpenAPI specification
type ObjectLambdaAccessPoint struct {
	Alias interface{} `json:"Alias,omitempty"`
	Name interface{} `json:"Name"`
	Objectlambdaaccesspointarn interface{} `json:"ObjectLambdaAccessPointArn,omitempty"`
}

// PutJobTaggingResult represents the PutJobTaggingResult schema from the OpenAPI specification
type PutJobTaggingResult struct {
}

// PutAccessPointPolicyForObjectLambdaRequest represents the PutAccessPointPolicyForObjectLambdaRequest schema from the OpenAPI specification
type PutAccessPointPolicyForObjectLambdaRequest struct {
	Policy interface{} `json:"Policy"`
}

// ReplicationTimeValue represents the ReplicationTimeValue schema from the OpenAPI specification
type ReplicationTimeValue struct {
	Minutes interface{} `json:"Minutes,omitempty"`
}

// DeleteMultiRegionAccessPointInput represents the DeleteMultiRegionAccessPointInput schema from the OpenAPI specification
type DeleteMultiRegionAccessPointInput struct {
	Name interface{} `json:"Name"`
}

// CreateBucketResult represents the CreateBucketResult schema from the OpenAPI specification
type CreateBucketResult struct {
	Bucketarn interface{} `json:"BucketArn,omitempty"`
}

// S3CopyObjectOperation represents the S3CopyObjectOperation schema from the OpenAPI specification
type S3CopyObjectOperation struct {
	Newobjectmetadata interface{} `json:"NewObjectMetadata,omitempty"`
	Bucketkeyenabled interface{} `json:"BucketKeyEnabled,omitempty"`
	Sseawskmskeyid interface{} `json:"SSEAwsKmsKeyId,omitempty"`
	Checksumalgorithm interface{} `json:"ChecksumAlgorithm,omitempty"`
	Objectlockmode interface{} `json:"ObjectLockMode,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Targetresource interface{} `json:"TargetResource,omitempty"`
	Unmodifiedsinceconstraint interface{} `json:"UnModifiedSinceConstraint,omitempty"`
	Accesscontrolgrants interface{} `json:"AccessControlGrants,omitempty"`
	Objectlockretainuntildate interface{} `json:"ObjectLockRetainUntilDate,omitempty"`
	Newobjecttagging interface{} `json:"NewObjectTagging,omitempty"`
	Modifiedsinceconstraint interface{} `json:"ModifiedSinceConstraint,omitempty"`
	Redirectlocation interface{} `json:"RedirectLocation,omitempty"`
	Metadatadirective interface{} `json:"MetadataDirective,omitempty"`
	Objectlocklegalholdstatus interface{} `json:"ObjectLockLegalHoldStatus,omitempty"`
	Targetkeyprefix interface{} `json:"TargetKeyPrefix,omitempty"`
	Cannedaccesscontrollist interface{} `json:"CannedAccessControlList,omitempty"`
	Requesterpays interface{} `json:"RequesterPays,omitempty"`
}

// ListJobsRequest represents the ListJobsRequest schema from the OpenAPI specification
type ListJobsRequest struct {
}

// MultiRegionAccessPointRegionalResponse represents the MultiRegionAccessPointRegionalResponse schema from the OpenAPI specification
type MultiRegionAccessPointRegionalResponse struct {
	Name interface{} `json:"Name,omitempty"`
	Requeststatus interface{} `json:"RequestStatus,omitempty"`
}

// CreateAccessPointForObjectLambdaRequest represents the CreateAccessPointForObjectLambdaRequest schema from the OpenAPI specification
type CreateAccessPointForObjectLambdaRequest struct {
	Configuration interface{} `json:"Configuration"`
}

// GetAccessPointPolicyStatusForObjectLambdaRequest represents the GetAccessPointPolicyStatusForObjectLambdaRequest schema from the OpenAPI specification
type GetAccessPointPolicyStatusForObjectLambdaRequest struct {
}

// GetBucketRequest represents the GetBucketRequest schema from the OpenAPI specification
type GetBucketRequest struct {
}

// SubmitMultiRegionAccessPointRoutesResult represents the SubmitMultiRegionAccessPointRoutesResult schema from the OpenAPI specification
type SubmitMultiRegionAccessPointRoutesResult struct {
}

// Metrics represents the Metrics schema from the OpenAPI specification
type Metrics struct {
	Eventthreshold interface{} `json:"EventThreshold,omitempty"`
	Status interface{} `json:"Status"`
}

// JobTimers represents the JobTimers schema from the OpenAPI specification
type JobTimers struct {
	Elapsedtimeinactiveseconds interface{} `json:"ElapsedTimeInActiveSeconds,omitempty"`
}

// CreateMultiRegionAccessPointRequest represents the CreateMultiRegionAccessPointRequest schema from the OpenAPI specification
type CreateMultiRegionAccessPointRequest struct {
	Clienttoken interface{} `json:"ClientToken"`
	Details interface{} `json:"Details"`
}

// ObjectLambdaTransformationConfiguration represents the ObjectLambdaTransformationConfiguration schema from the OpenAPI specification
type ObjectLambdaTransformationConfiguration struct {
	Contenttransformation interface{} `json:"ContentTransformation"`
	Actions interface{} `json:"Actions"`
}

// JobOperation represents the JobOperation schema from the OpenAPI specification
type JobOperation struct {
	S3putobjectretention S3SetObjectRetentionOperation `json:"S3PutObjectRetention,omitempty"` // Contains the configuration parameters for the Object Lock retention action for an S3 Batch Operations job. Batch Operations passes every object to the underlying <code>PutObjectRetention</code> API operation. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-retention-date.html">Using S3 Object Lock retention with S3 Batch Operations</a> in the <i>Amazon S3 User Guide</i>.
	Lambdainvoke interface{} `json:"LambdaInvoke,omitempty"`
	S3initiaterestoreobject interface{} `json:"S3InitiateRestoreObject,omitempty"`
	S3putobjectcopy interface{} `json:"S3PutObjectCopy,omitempty"`
	S3putobjectlegalhold S3SetObjectLegalHoldOperation `json:"S3PutObjectLegalHold,omitempty"` // Contains the configuration for an S3 Object Lock legal hold operation that an S3 Batch Operations job passes to every object to the underlying <code>PutObjectLegalHold</code> API operation. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-legal-hold.html">Using S3 Object Lock legal hold with S3 Batch Operations</a> in the <i>Amazon S3 User Guide</i>.
	S3putobjecttagging interface{} `json:"S3PutObjectTagging,omitempty"`
	S3replicateobject interface{} `json:"S3ReplicateObject,omitempty"`
	S3deleteobjecttagging interface{} `json:"S3DeleteObjectTagging,omitempty"`
	S3putobjectacl interface{} `json:"S3PutObjectAcl,omitempty"`
}

// PutBucketLifecycleConfigurationRequest represents the PutBucketLifecycleConfigurationRequest schema from the OpenAPI specification
type PutBucketLifecycleConfigurationRequest struct {
	Lifecycleconfiguration interface{} `json:"LifecycleConfiguration,omitempty"`
}

// S3Retention represents the S3Retention schema from the OpenAPI specification
type S3Retention struct {
	Mode interface{} `json:"Mode,omitempty"`
	Retainuntildate interface{} `json:"RetainUntilDate,omitempty"`
}

// CreateMultiRegionAccessPointInput represents the CreateMultiRegionAccessPointInput schema from the OpenAPI specification
type CreateMultiRegionAccessPointInput struct {
	Publicaccessblock PublicAccessBlockConfiguration `json:"PublicAccessBlock,omitempty"` // <p>The <code>PublicAccessBlock</code> configuration that you want to apply to this Amazon S3 account. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The Meaning of "Public"</a> in the <i>Amazon S3 User Guide</i>.</p> <p>This data type is not supported for Amazon S3 on Outposts.</p>
	Regions interface{} `json:"Regions"`
	Name interface{} `json:"Name"`
}

// AsyncRequestParameters represents the AsyncRequestParameters schema from the OpenAPI specification
type AsyncRequestParameters struct {
	Deletemultiregionaccesspointrequest interface{} `json:"DeleteMultiRegionAccessPointRequest,omitempty"`
	Putmultiregionaccesspointpolicyrequest interface{} `json:"PutMultiRegionAccessPointPolicyRequest,omitempty"`
	Createmultiregionaccesspointrequest interface{} `json:"CreateMultiRegionAccessPointRequest,omitempty"`
}

// NoncurrentVersionTransition represents the NoncurrentVersionTransition schema from the OpenAPI specification
type NoncurrentVersionTransition struct {
	Noncurrentdays interface{} `json:"NoncurrentDays,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
}

// DescribeMultiRegionAccessPointOperationRequest represents the DescribeMultiRegionAccessPointOperationRequest schema from the OpenAPI specification
type DescribeMultiRegionAccessPointOperationRequest struct {
}

// UpdateJobPriorityRequest represents the UpdateJobPriorityRequest schema from the OpenAPI specification
type UpdateJobPriorityRequest struct {
}

// LifecycleConfiguration represents the LifecycleConfiguration schema from the OpenAPI specification
type LifecycleConfiguration struct {
	Rules interface{} `json:"Rules,omitempty"`
}

// LambdaInvokeOperation represents the LambdaInvokeOperation schema from the OpenAPI specification
type LambdaInvokeOperation struct {
	Functionarn interface{} `json:"FunctionArn,omitempty"`
}

// AdvancedCostOptimizationMetrics represents the AdvancedCostOptimizationMetrics schema from the OpenAPI specification
type AdvancedCostOptimizationMetrics struct {
	Isenabled interface{} `json:"IsEnabled,omitempty"`
}

// S3GeneratedManifestDescriptor represents the S3GeneratedManifestDescriptor schema from the OpenAPI specification
type S3GeneratedManifestDescriptor struct {
	Format interface{} `json:"Format,omitempty"`
	Location JobManifestLocation `json:"Location,omitempty"` // Contains the information required to locate a manifest object.
}

// PrefixLevelStorageMetrics represents the PrefixLevelStorageMetrics schema from the OpenAPI specification
type PrefixLevelStorageMetrics struct {
	Isenabled interface{} `json:"IsEnabled,omitempty"`
	Selectioncriteria SelectionCriteria `json:"SelectionCriteria,omitempty"` // <p/>
}

// GetAccessPointPolicyForObjectLambdaResult represents the GetAccessPointPolicyForObjectLambdaResult schema from the OpenAPI specification
type GetAccessPointPolicyForObjectLambdaResult struct {
	Policy interface{} `json:"Policy,omitempty"`
}
