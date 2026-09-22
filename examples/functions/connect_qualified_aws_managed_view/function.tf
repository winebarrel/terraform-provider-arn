# arn:aws:connect:ap-northeast-1:aws:view/view-id:view-qualifier
output "connect_qualified_aws_managed_view" {
  value = provider::arn::connect_qualified_aws_managed_view("view-id", "view-qualifier")
}
