# arn:aws:redshift:ap-northeast-1:111111111111:datashare:producer-cluster-namespace/data-share-name
output "redshift_datashare" {
  value = provider::arn::redshift_datashare("producer-cluster-namespace", "data-share-name")
}
