# arn:aws:sagemaker:ap-northeast-1:111111111111:cluster/cluster-id
output "sagemaker_cluster" {
  value = provider::arn::sagemaker_cluster("cluster-id")
}
