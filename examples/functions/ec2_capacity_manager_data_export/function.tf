# arn:aws:ec2:ap-northeast-1:111111111111:capacity-manager-data-export/capacity-manager-data-export-id
output "ec2_capacity_manager_data_export" {
  value = provider::arn::ec2_capacity_manager_data_export("capacity-manager-data-export-id")
}
