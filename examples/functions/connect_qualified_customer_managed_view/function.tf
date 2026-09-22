# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/view/view-id:view-qualifier
output "connect_qualified_customer_managed_view" {
  value = provider::arn::connect_qualified_customer_managed_view("instance-id", "view-id", "view-qualifier")
}
