# arn:aws:robomaker:ap-northeast-1:111111111111:world-export-job/world-export-job-id
output "robomaker_world_export_job" {
  value = provider::arn::robomaker_world_export_job("world-export-job-id")
}
