package main

import (
	"OverTheWallRuleProcessing/internal/ConfigGenerator"
	"OverTheWallRuleProcessing/internal/ContentParse"
	"OverTheWallRuleProcessing/internal/LinkConversion"
)

func main() {
	baseRaw, _ := ContentParse.GetUrlResp("http://cat.cat77.org/api/v1/client/subscribe?token=09085361e443acc95169baf36dc642fc")
	decodeByte, _ := ContentParse.ParseB64String(baseRaw)
	decodeStrings, _ := ContentParse.SplitAndDecode(decodeByte)
	ans, _ := LinkConversion.RunConversionSub(decodeStrings)
	ConfigGenerator.RunGenerateConfig(ans, "/home/wwwr/w/wp-content/uploads/5ra15aKo6L2756yZ55qE6YWN572u5paH5Lu2/Y2F0LTc3.ym", "clash")
}
