package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFile(n string) {
	// ファイル名をパース
	dir, n := parseName(n)

	// .forgeディレクトリを探す
	_, err := os.Stat(".forge")
	if err != nil {
		fmt.Println("not found .forge directory")
		return
	}

	// ディレクトリを作成
	if dir != "" {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("failed to create directory")
			return
		}
	}

	// テンプレートファイルを探す
	_, err = os.Stat(".forge/template.txt")
	if err != nil {
		fmt.Println("not found template file")
		return
	}

	// テンプレートファイルを読み込む
	file, err := os.ReadFile(".forge/template.txt")
	if err != nil {
		fmt.Println("failed to read template file")
		return
	}

	// ファイルパスを組み立てる
	fullPath := filepath.Join(dir, n+".txt")

	// 新しいファイルを生成する
	err = os.WriteFile(fullPath, file, 0644)
	if err != nil {
		fmt.Println("failed to generate file")
		return
	}

	// ファイル生成成功
	fmt.Println("generate file successfully")
}

// ファイル名をパースしてディレクトリ名とファイル名を返す
func parseName(path string) (dif, filename string) {
	// パスをスラッシュで分割
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return "", ""
	}

	// 最後の要素をファイル名として取得
	n := parts[len(parts)-1]

	// 最後以外を結合してディレクトリ名として取得
	dir := strings.Join(parts[:len(parts)-1], string(os.PathSeparator))

	// ディレクトリ名とファイル名を返す
	return dir, n
}
