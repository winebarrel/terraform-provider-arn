# arn:aws:apigateway:ap-northeast-1:111111111111:/portalproducts/portal-product-id/productrestendpointpages/product-rest-endpoint-page-id
output "apigateway_product_rest_endpoint_page" {
  value = provider::arn::apigateway_product_rest_endpoint_page("portal-product-id", "product-rest-endpoint-page-id")
}
