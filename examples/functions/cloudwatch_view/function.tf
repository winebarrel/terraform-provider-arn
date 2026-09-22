# arn:aws:cloudwatch:ap-northeast-1:111111111111:view/view-name
output "cloudwatch_view" {
  value = provider::arn::cloudwatch_view("view-name")
}
