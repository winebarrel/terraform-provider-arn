# arn:aws:ce::111111111111:costcategory/identifier
output "ce_costcategory" {
  value = provider::arn::ce_costcategory("identifier")
}
