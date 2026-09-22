# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/analysistemplate/analysis-template-id
output "cleanrooms_analysistemplate" {
  value = provider::arn::cleanrooms_analysistemplate("membership-id", "analysis-template-id")
}
