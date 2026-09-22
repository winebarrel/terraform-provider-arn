# arn:aws:apigateway:ap-northeast-1:111111111111:/portalproducts/portal-product-id
output "apigateway_portal_product" {
  value = provider::arn::apigateway_portal_product("portal-product-id")
}
