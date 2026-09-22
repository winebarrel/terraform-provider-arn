# arn:aws:connect:ap-northeast-1:aws:view/view-id
output "connect_aws_managed_view" {
  value = provider::arn::connect_aws_managed_view("view-id")
}
