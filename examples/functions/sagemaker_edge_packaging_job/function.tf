# arn:aws:sagemaker:ap-northeast-1:111111111111:edge-packaging-job/edge-packaging-job-name
output "sagemaker_edge_packaging_job" {
  value = provider::arn::sagemaker_edge_packaging_job("edge-packaging-job-name")
}
