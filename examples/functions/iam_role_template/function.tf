# arn:aws:iam::aws:role-template/aws-service-principal/role-template-name:role-template-major-version
output "iam_role_template" {
  value = provider::arn::iam_role_template("aws-service-principal", "role-template-name", "role-template-major-version")
}
