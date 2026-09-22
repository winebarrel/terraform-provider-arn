# arn:aws:eks:ap-northeast-1:111111111111:dashboard/dashboard-name
output "eks_dashboard" {
  value = provider::arn::eks_dashboard("dashboard-name")
}
