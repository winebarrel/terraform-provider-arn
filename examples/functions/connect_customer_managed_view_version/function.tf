# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/view/view-id:view-version
output "connect_customer_managed_view_version" {
  value = provider::arn::connect_customer_managed_view_version("instance-id", "view-id", "view-version")
}
