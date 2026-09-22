# arn:aws:robomaker:ap-northeast-1:111111111111:simulation-application/application-name/created-on-epoch
output "robomaker_simulation_application" {
  value = provider::arn::robomaker_simulation_application("application-name", "created-on-epoch")
}
