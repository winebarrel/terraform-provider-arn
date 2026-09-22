# arn:aws:outposts:ap-northeast-1:111111111111:outpost/outpost-id
output "outposts_outpost" {
  value = provider::arn::outposts_outpost("outpost-id")
}
