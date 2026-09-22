# arn:aws:mobiletargeting:ap-northeast-1:111111111111:templates
output "mobiletargeting_templates" {
  value = provider::arn::mobiletargeting_templates()
}
