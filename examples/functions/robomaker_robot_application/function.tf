# arn:aws:robomaker:ap-northeast-1:111111111111:robot-application/application-name/created-on-epoch
output "robomaker_robot_application" {
  value = provider::arn::robomaker_robot_application("application-name", "created-on-epoch")
}
