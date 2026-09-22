# arn:aws:catalog:ap-northeast-1:111111111111:product/product-id
output "servicecatalog_product" {
  value = provider::arn::servicecatalog_product("product-id")
}
