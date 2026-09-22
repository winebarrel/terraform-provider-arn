# arn:aws:outposts:ap-northeast-1:111111111111:site/site-id
output "outposts_site" {
  value = provider::arn::outposts_site("site-id")
}
