# arn:aws:codedeploy:ap-northeast-1:111111111111:application:application-name
output "codedeploy_application" {
  value = provider::arn::codedeploy_application("application-name")
}
