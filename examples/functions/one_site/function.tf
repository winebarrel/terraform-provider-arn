# arn:aws:one:ap-northeast-1:111111111111:site/site-id
output "one_site" {
  value = provider::arn::one_site("site-id")
}
