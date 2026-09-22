# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/*
output "mobiletargeting_apps" {
  value = provider::arn::mobiletargeting_apps()
}
