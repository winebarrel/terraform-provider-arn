# arn:aws:robomaker:ap-northeast-1:111111111111:world-template/world-template-job-id
output "robomaker_world_template" {
  value = provider::arn::robomaker_world_template("world-template-job-id")
}
