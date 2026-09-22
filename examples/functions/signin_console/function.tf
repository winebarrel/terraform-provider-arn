# arn:aws:signin:::console/console-name
output "signin_console" {
  value = provider::arn::signin_console("console-name")
}
