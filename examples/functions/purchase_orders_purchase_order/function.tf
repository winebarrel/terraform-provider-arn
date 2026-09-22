# arn:aws:purchase-orders::111111111111:purchase-order/resource-name
output "purchase_orders_purchase_order" {
  value = provider::arn::purchase_orders_purchase_order("resource-name")
}
