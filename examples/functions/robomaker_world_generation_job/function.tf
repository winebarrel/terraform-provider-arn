# arn:aws:robomaker:ap-northeast-1:111111111111:world-generation-job/world-generation-job-id
output "robomaker_world_generation_job" {
  value = provider::arn::robomaker_world_generation_job("world-generation-job-id")
}
