# arn:aws:sagemaker:ap-northeast-1:111111111111:lineage-group/lineage-group-name
output "sagemaker_lineage_group" {
  value = provider::arn::sagemaker_lineage_group("lineage-group-name")
}
