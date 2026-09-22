# arn:aws:observabilityadmin:ap-northeast-1:111111111111:organization-centralization-rule/centralization-rule-name
output "observabilityadmin_organization_centralization_rule" {
  value = provider::arn::observabilityadmin_organization_centralization_rule("centralization-rule-name")
}
