// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudfront
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudfront/cloudfront.json
// Functions: 19
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudfront_anycast_ip_list", Service: "cloudfront", Resource: "anycast-ip-list", Template: "arn:${Partition}:cloudfront::${Account}:anycast-ip-list/${Id}"},
		{Name: "cloudfront_cache_policy", Service: "cloudfront", Resource: "cache-policy", Template: "arn:${Partition}:cloudfront::${Account}:cache-policy/${Id}"},
		{Name: "cloudfront_connection_function", Service: "cloudfront", Resource: "connection-function", Template: "arn:${Partition}:cloudfront::${Account}:connection-function/${Id}"},
		{Name: "cloudfront_connection_group", Service: "cloudfront", Resource: "connection-group", Template: "arn:${Partition}:cloudfront::${Account}:connection-group/${Id}"},
		{Name: "cloudfront_continuous_deployment_policy", Service: "cloudfront", Resource: "continuous-deployment-policy", Template: "arn:${Partition}:cloudfront::${Account}:continuous-deployment-policy/${Id}"},
		{Name: "cloudfront_distribution", Service: "cloudfront", Resource: "distribution", Template: "arn:${Partition}:cloudfront::${Account}:distribution/${DistributionId}"},
		{Name: "cloudfront_distribution_tenant", Service: "cloudfront", Resource: "distribution-tenant", Template: "arn:${Partition}:cloudfront::${Account}:distribution-tenant/${Id}"},
		{Name: "cloudfront_field_level_encryption_config", Service: "cloudfront", Resource: "field-level-encryption-config", Template: "arn:${Partition}:cloudfront::${Account}:field-level-encryption-config/${Id}"},
		{Name: "cloudfront_field_level_encryption_profile", Service: "cloudfront", Resource: "field-level-encryption-profile", Template: "arn:${Partition}:cloudfront::${Account}:field-level-encryption-profile/${Id}"},
		{Name: "cloudfront_function", Service: "cloudfront", Resource: "function", Template: "arn:${Partition}:cloudfront::${Account}:function/${Name}"},
		{Name: "cloudfront_key_value_store", Service: "cloudfront", Resource: "key-value-store", Template: "arn:${Partition}:cloudfront::${Account}:key-value-store/${Name}"},
		{Name: "cloudfront_origin_access_control", Service: "cloudfront", Resource: "origin-access-control", Template: "arn:${Partition}:cloudfront::${Account}:origin-access-control/${Id}"},
		{Name: "cloudfront_origin_access_identity", Service: "cloudfront", Resource: "origin-access-identity", Template: "arn:${Partition}:cloudfront::${Account}:origin-access-identity/${Id}"},
		{Name: "cloudfront_origin_request_policy", Service: "cloudfront", Resource: "origin-request-policy", Template: "arn:${Partition}:cloudfront::${Account}:origin-request-policy/${Id}"},
		{Name: "cloudfront_realtime_log_config", Service: "cloudfront", Resource: "realtime-log-config", Template: "arn:${Partition}:cloudfront::${Account}:realtime-log-config/${Name}"},
		{Name: "cloudfront_response_headers_policy", Service: "cloudfront", Resource: "response-headers-policy", Template: "arn:${Partition}:cloudfront::${Account}:response-headers-policy/${Id}"},
		{Name: "cloudfront_streaming_distribution", Service: "cloudfront", Resource: "streaming-distribution", Template: "arn:${Partition}:cloudfront::${Account}:streaming-distribution/${DistributionId}"},
		{Name: "cloudfront_trust_store", Service: "cloudfront", Resource: "trust-store", Template: "arn:${Partition}:cloudfront::${Account}:trust-store/${Id}"},
		{Name: "cloudfront_vpcorigin", Service: "cloudfront", Resource: "vpcorigin", Template: "arn:${Partition}:cloudfront::${Account}:vpcorigin/${Id}"},
	})
}
