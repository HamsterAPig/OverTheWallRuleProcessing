package LinkConversion

import (
	. "OverTheWallRuleProcessing/internal"
	"log"
	"strings"
)

func RunConversionSub(subStrings []string) ([]ProxyInfoNode, error) {
	result := make([]ProxyInfoNode, len(subStrings))
	for i, item := range subStrings {
		if strings.HasPrefix(item, "ss://") {
			ss, err := ParseShadowsocks(item)
			if err != nil {
				log.Printf("failed to parse %s: %s", item, err)
				continue
			}
			result[i] = *ss
		}
	}
	return result, nil
}
