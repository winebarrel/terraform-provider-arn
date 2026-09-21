// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: purchase-orders
// Source: https://servicereference.us-east-1.amazonaws.com/v1/purchase-orders/purchase-orders.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "purchase_orders_purchase_order", Service: "purchase-orders", Resource: "purchase-order", Template: "arn:${Partition}:purchase-orders::${Account}:purchase-order/${ResourceName}"},
	})
}
