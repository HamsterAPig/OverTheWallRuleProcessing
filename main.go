package main

import (
	"OverTheWallRuleProcessing/config"
	"OverTheWallRuleProcessing/internal/ConfigGenerator"
	"OverTheWallRuleProcessing/internal/ContentParse"
	"OverTheWallRuleProcessing/internal/LinkConversion"
	"log"
)

func main() {

	confLists, _ := config.ReadConf()
	baseRaw, _ := ContentParse.GetUrlResp(confLists[0])
	log.Print(baseRaw)
	decodeByte, _ := ContentParse.ParseB64String(baseRaw)
	decodeStrings, _ := ContentParse.SplitAndDecode(decodeByte)
	ans, _ := LinkConversion.RunConversionSub(decodeStrings)
	log.Print(ans)
	ConfigGenerator.RunGenerateConfig(ans, confLists[1], "clash")
}
