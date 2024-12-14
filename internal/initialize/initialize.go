package initialize

import (
	_ "embed"
	"fmt"
	"os"
)

// このディレクティブは、指定されたファイルをバイナリデータとして埋め込む
//
//go:embed templates/.code-forge.yml
var configTemplate []byte

// 設定ファイルとテンプレートファイル配置先のディレクトリを作成
func Initialize() {
	templateDir := "forge"
	outputPath := ".code-forge.yml"

	err := os.MkdirAll(templateDir, os.ModePerm)
	if err != nil {
		fmt.Printf("failed to create config directory: %v", err)
		return
	}

	_, err = os.Stat(outputPath)
	if err == nil {
		fmt.Println(".code-forge.yml already exists")
		return
	}

	// 設定ファイルを書き込む
	err = os.WriteFile(outputPath, configTemplate, 0644)
	if err != nil {
		fmt.Printf("failed to create code-forge.yml: %v", err)
		return
	}

	fmt.Println(".code-forge.yml has been created successfully!")
}
