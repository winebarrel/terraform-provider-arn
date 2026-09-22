# arn:aws:apigateway:ap-northeast-1:111111111111:/portalproducts/portal-product-id/productpages/product-page-id
output "apigateway_product_page" {
  value = provider::arn::apigateway_product_page("portal-product-id", "product-page-id")
}
