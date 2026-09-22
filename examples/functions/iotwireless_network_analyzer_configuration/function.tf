# arn:aws:iotwireless:ap-northeast-1:111111111111:NetworkAnalyzerConfiguration/network-analyzer-configuration-name
output "iotwireless_network_analyzer_configuration" {
  value = provider::arn::iotwireless_network_analyzer_configuration("network-analyzer-configuration-name")
}
