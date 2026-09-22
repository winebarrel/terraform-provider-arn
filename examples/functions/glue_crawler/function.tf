# arn:aws:glue:ap-northeast-1:111111111111:crawler/crawler-name
output "glue_crawler" {
  value = provider::arn::glue_crawler("crawler-name")
}
