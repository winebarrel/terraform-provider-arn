# arn:aws:resiliencehub:ap-northeast-1:111111111111:app-assessment/app-assessment-id
output "resiliencehub_app_assessment" {
  value = provider::arn::resiliencehub_app_assessment("app-assessment-id")
}
