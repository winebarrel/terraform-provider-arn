# arn:aws:cloudformation:ap-northeast-1:111111111111:generatedTemplate/id
output "cloudformation_generatedtemplate" {
  value = provider::arn::cloudformation_generatedtemplate("id")
}
