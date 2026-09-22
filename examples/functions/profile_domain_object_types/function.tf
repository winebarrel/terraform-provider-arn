# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/domain-object-types/object-type-name
output "profile_domain_object_types" {
  value = provider::arn::profile_domain_object_types("domain-name", "object-type-name")
}
