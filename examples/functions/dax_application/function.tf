# arn:aws:dax:ap-northeast-1:111111111111:cache/cluster-name
output "dax_application" {
  value = provider::arn::dax_application("cluster-name")
}
