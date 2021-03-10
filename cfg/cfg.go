package cfg

import "fmt"

// SetCfgFile 设置配置文件
func SetCfgFile(filePath string) error {
	fmt.Printf("[%s] Set Done!\n", filePath)
	return nil
}
