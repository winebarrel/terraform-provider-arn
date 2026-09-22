# arn:aws:glue:ap-northeast-1:111111111111:dataQualityRuleset/ruleset-name
output "glue_data_quality_ruleset" {
  value = provider::arn::glue_data_quality_ruleset("ruleset-name")
}
