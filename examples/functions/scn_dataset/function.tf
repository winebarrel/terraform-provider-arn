# arn:aws:scn:ap-northeast-1:111111111111:instance/instance-id/namespaces/namespace/datasets/dataset-name
output "scn_dataset" {
  value = provider::arn::scn_dataset("instance-id", "namespace", "dataset-name")
}
