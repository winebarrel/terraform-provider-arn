# arn:aws:iotsitewise:ap-northeast-1:111111111111:workspace/workspace-name/pipeline/pipeline-name
output "iotsitewise_pipeline" {
  value = provider::arn::iotsitewise_pipeline("workspace-name", "pipeline-name")
}
