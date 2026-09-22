# arn:aws:resiliencehub:ap-northeast-1:111111111111:app/app-id
output "resiliencehub_application" {
  value = provider::arn::resiliencehub_application("app-id")
}
