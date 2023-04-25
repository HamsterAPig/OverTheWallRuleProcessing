package ConfigGenerator

import (
	. "OverTheWallRuleProcessing/internal"
	"fmt"
)

func GenerateConfigClash(nodeInfo []ProxyInfoNode, filepath string) string {
	strTmp := "proxies:\n"
	for _, item := range nodeInfo {
		if item.Type == "ss" {
			strTmp += fmt.Sprintf("    - { name: %s, type: ss, server: %s, port: %s, cipher: %s, password: %s, udp: true }\n", item.OrderName, item.Server, item.Port, item.Chipter, item.Password)
		}
	}
	return strTmp
}
