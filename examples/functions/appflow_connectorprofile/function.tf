# arn:aws:appflow:ap-northeast-1:111111111111:connectorprofile/profile-name
output "appflow_connectorprofile" {
  value = provider::arn::appflow_connectorprofile("profile-name")
}
