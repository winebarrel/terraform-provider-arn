# arn:aws:catalog:ap-northeast-1:111111111111:portfolio/portfolio-id
output "servicecatalog_portfolio" {
  value = provider::arn::servicecatalog_portfolio("portfolio-id")
}
