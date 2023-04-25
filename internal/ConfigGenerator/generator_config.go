package ConfigGenerator

import (
	. "OverTheWallRuleProcessing/internal"
	"os"
)

func RunGenerateConfig(nodeinfo []ProxyInfoNode, filepath string, target_type string) error {
	var res string
	if "clash" == target_type {
		res = GenerateConfigClash(nodeinfo, filepath)
	}
	WriteFile(filepath, res)
	return nil
}

func WriteFile(filename, content string) error {
	var err error

	// 打开文件，如果不存在则创建新文件
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
