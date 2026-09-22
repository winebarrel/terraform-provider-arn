# arn:aws:securityhub:ap-northeast-1:111111111111:product/company/product-id
output "securityhub_product" {
  value = provider::arn::securityhub_product("company", "product-id")
}
