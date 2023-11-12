// クラブ対抗戦の射座割りからSIUS用データを生成する

package main

import (
	"log"

	"github.com/s-ariga/sius-entry/sius"
)

// Excelの操作用にはExcelizeを使う
// CSVにいちいち変換するのがめんどくさいので

const (
	INPUT    = "./input/SIUS用射座割り.xlsx"
	ID_INPUT = "./input/全射手一覧.xlsx"
	OUTPUT   = "./output/SIUSデータ.xlsx"
)

func main() {
	err := sius.GenerateSIUSData(INPUT, ID_INPUT, OUTPUT)
	if err != nil {
		// 取り急ぎ
		// ここまでエラーが来ることはないと思うので
		log.Fatal(err)
	}
}
