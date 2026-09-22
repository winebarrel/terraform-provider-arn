# arn:aws:emr-serverless:ap-northeast-1:111111111111:/applications/application-id
output "emr_serverless_application" {
  value = provider::arn::emr_serverless_application("application-id")
}
