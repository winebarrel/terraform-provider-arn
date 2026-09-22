# arn:aws:application-signals-mcp:ap-northeast-1:111111111111:mcp-server/*
output "application_signals_mcp_mcp_server" {
  value = provider::arn::application_signals_mcp_mcp_server()
}
