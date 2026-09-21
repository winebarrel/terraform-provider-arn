// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: honeycode
// Source: https://servicereference.us-east-1.amazonaws.com/v1/honeycode/honeycode.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "honeycode_screen", Service: "honeycode", Resource: "screen", Template: "arn:${Partition}:honeycode:${Region}:${Account}:screen:workbook/${WorkbookId}/app/${AppId}/screen/${ScreenId}"},
		{Name: "honeycode_screen_automation", Service: "honeycode", Resource: "screen-automation", Template: "arn:${Partition}:honeycode:${Region}:${Account}:screen-automation:workbook/${WorkbookId}/app/${AppId}/screen/${ScreenId}/automation/${AutomationId}"},
		{Name: "honeycode_table", Service: "honeycode", Resource: "table", Template: "arn:${Partition}:honeycode:${Region}:${Account}:table:workbook/${WorkbookId}/table/${TableId}"},
		{Name: "honeycode_workbook", Service: "honeycode", Resource: "workbook", Template: "arn:${Partition}:honeycode:${Region}:${Account}:workbook:workbook/${WorkbookId}"},
	})
}
