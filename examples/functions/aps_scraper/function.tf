# arn:aws:aps:ap-northeast-1:111111111111:scraper/scraper-id
output "aps_scraper" {
  value = provider::arn::aps_scraper("scraper-id")
}
