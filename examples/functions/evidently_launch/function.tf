# arn:aws:evidently:ap-northeast-1:111111111111:project/project-name/launch/launch-name
output "evidently_launch" {
  value = provider::arn::evidently_launch("project-name", "launch-name")
}
