# arn:aws:sagemaker:ap-northeast-1:111111111111:cluster-scheduler-config/cluster-scheduler-config-id
output "sagemaker_cluster_scheduler_config" {
  value = provider::arn::sagemaker_cluster_scheduler_config("cluster-scheduler-config-id")
}
