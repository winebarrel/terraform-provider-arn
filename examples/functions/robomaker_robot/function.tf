# arn:aws:robomaker:ap-northeast-1:111111111111:robot/robot-name/created-on-epoch
output "robomaker_robot" {
  value = provider::arn::robomaker_robot("robot-name", "created-on-epoch")
}
