// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: lex
// Source: https://servicereference.us-east-1.amazonaws.com/v1/lex/lex.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "lex_bot", Service: "lex", Resource: "bot", Template: "arn:${Partition}:lex:${Region}:${Account}:bot/${BotId}"},
		{Name: "lex_bot_2", Service: "lex", Resource: "bot", Template: "arn:${Partition}:lex:${Region}:${Account}:bot:${BotName}"},
		{Name: "lex_bot_alias", Service: "lex", Resource: "bot alias", Template: "arn:${Partition}:lex:${Region}:${Account}:bot-alias/${BotId}/${BotAliasId}"},
		{Name: "lex_bot_alias_2", Service: "lex", Resource: "bot alias", Template: "arn:${Partition}:lex:${Region}:${Account}:bot:${BotName}:${BotAlias}"},
		{Name: "lex_bot_version", Service: "lex", Resource: "bot version", Template: "arn:${Partition}:lex:${Region}:${Account}:bot:${BotName}:${BotVersion}"},
		{Name: "lex_channel", Service: "lex", Resource: "channel", Template: "arn:${Partition}:lex:${Region}:${Account}:bot-channel:${BotName}:${BotAlias}:${ChannelName}"},
		{Name: "lex_intent_version", Service: "lex", Resource: "intent version", Template: "arn:${Partition}:lex:${Region}:${Account}:intent:${IntentName}:${IntentVersion}"},
		{Name: "lex_slottype_version", Service: "lex", Resource: "slottype version", Template: "arn:${Partition}:lex:${Region}:${Account}:slottype:${SlotName}:${SlotVersion}"},
		{Name: "lex_test_set", Service: "lex", Resource: "test set", Template: "arn:${Partition}:lex:${Region}:${Account}:test-set/${TestSetId}"},
	})
}
