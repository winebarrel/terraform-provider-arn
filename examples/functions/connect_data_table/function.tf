# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/data-table/data-table-id
output "connect_data_table" {
  value = provider::arn::connect_data_table("instance-id", "data-table-id")
}
