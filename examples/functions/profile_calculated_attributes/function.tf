# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/calculated-attributes/calculated-attribute-name
output "profile_calculated_attributes" {
  value = provider::arn::profile_calculated_attributes("domain-name", "calculated-attribute-name")
}
