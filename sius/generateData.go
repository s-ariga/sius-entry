// seiichi.ariga@gmail.com

package sius

import (
	"log"
)

// Excelize を使ってデータを作成する

// InputDataは、射座割りから取得したデータ
// 種目名、クラブ名、射手名(スペースを含む)、射群、射座番号
type InputData struct {
	// TODO: 入力データの型を定義する
	Position string
	Team     []string // Team
	Nation   []string // Nation
	Name     []string // Name
	Relay    []string // Relay
	Target   []string // Target
}

// Idデータは射手一覧から取得したIDと名前
// 名前は、姓、名、とローマ字読みの２つ
type IdData struct {
	// TODO: IDデータの型を定義する
	Id     string // F列
	FName  string // C列
	LName  string // B列
	Romaji string // E列
}

// OutputDataはSIUS出力するデータ
type OutputData struct {
	// TODO: 出力データの型を定義する
	Position string
	StartNum []string // Start Number
	Name     []string // Name
	DispName []string // Display Name
	Nation   []string // Nation (実はチーム名)
	Team     []string // Team (団体)
	Relay    []string // Relay
	Target   []string // Target Number
	Id       []string // IssfId
}

// ワークシート名から取得すればいいかも
var Positions = []string{
	"ARM", // 1001 StartNum
	"ARW",
	"R3PM",
	"R3PW",
	"RPRM",
	"RPRW",
	"ARMT",
	"ARPR",
}

const (
	ARM  = 1001
	R3PM = 2001
	RPRM = 3001
	ARW  = 4001
	R3PW = 5001
	RPRW = 6001
	ARMT = 7001
	ARPR = 8001
)

func GenerateSIUSData(inputPath, idInputPath, outputPath string) error {
	// Open input Excel file
	data, err := readData(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	idData, err := readIdData(idInputPath)
	if err != nil {
		log.Fatal(err)
	}

	outputData, err := mergeData(data, idData)
	if err != nil {
		log.Fatal(err)
	}

	err = writeData(outputPath, outputData)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
