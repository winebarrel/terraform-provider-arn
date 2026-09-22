# arn:aws:databrew:ap-northeast-1:111111111111:ruleset/resource-id
output "databrew_ruleset" {
  value = provider::arn::databrew_ruleset("resource-id")
}
