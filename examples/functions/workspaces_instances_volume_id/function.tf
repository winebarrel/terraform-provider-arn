# arn:aws:ec2:ap-northeast-1:111111111111:volume/volume-id
output "workspaces_instances_volume_id" {
  value = provider::arn::workspaces_instances_volume_id("volume-id")
}
