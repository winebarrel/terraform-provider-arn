# arn:aws:geo:ap-northeast-1:111111111111:route-calculator/calculator-name
output "geo_route_calculator" {
  value = provider::arn::geo_route_calculator("calculator-name")
}
