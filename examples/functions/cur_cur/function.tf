# arn:aws:cur:ap-northeast-1:111111111111:definition/report-name
output "cur_cur" {
  value = provider::arn::cur_cur("report-name")
}
