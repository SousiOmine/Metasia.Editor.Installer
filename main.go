package main

import (
	"fmt"
)

func main() {

	fmt.Println("Metasiaエディタ インストーラーへようこそ！")

	var params InstallParams
	params.SetDefault()

	var installPath string

	fmt.Print("インストール先のパスを入力してください: ")
	fmt.Scanln(&installPath)
	params.Path = installPath

	executor := InstallExecutor{Params: params}
	err := executor.Execute()
	if err != nil {
		fmt.Println("インストールに失敗しました:", err)
		return
	}
	fmt.Println("インストールが完了しました！")
}
