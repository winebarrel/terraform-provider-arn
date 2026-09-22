# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/segment-definitions/segment-definition-name
output "profile_segment_definitions" {
  value = provider::arn::profile_segment_definitions("domain-name", "segment-definition-name")
}
