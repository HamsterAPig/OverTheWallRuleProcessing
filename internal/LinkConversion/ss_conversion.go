package LinkConversion

// SS链接转换
import (
	. "OverTheWallRuleProcessing/internal"
	"OverTheWallRuleProcessing/internal/ContentParse"
	"errors"
	"regexp"
	"strings"
)

func ParseShadowsocks(s string) (*ProxyInfoNode, error) {
	re := regexp.MustCompile(`^ss:\/\/(.+)@(.+):(\d+)#(.+)$`)
	matches := re.FindStringSubmatch(s)

	if len(matches) != 5 {
		return nil, errors.New("invalid input string")
	}

	baseDecode, _ := ContentParse.ParseB64String(matches[1])
	baseSplit := strings.Split(baseDecode, ":")

	return &ProxyInfoNode{
		Type:      "ss",
		Password:  baseSplit[1],
		Server:    matches[2],
		Port:      matches[3],
		Chipter:   baseSplit[0],
		OrderName: strings.Trim(matches[4], "\r"),
	}, nil
}
