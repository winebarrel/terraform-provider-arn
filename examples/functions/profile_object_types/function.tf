# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/object-types/object-type-name
output "profile_object_types" {
  value = provider::arn::profile_object_types("domain-name", "object-type-name")
}
