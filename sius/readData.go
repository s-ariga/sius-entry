package sius

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 射座割りの情報をExcelから読み込む
func readData(input string) ([]InputData, error) {
	f, err := excelize.OpenFile(input)
	if err != nil {
		log.Fatal(err, "入力ファイルが読み込めませんでした: ", input)
	}
	defer f.Close()

	wsList := f.GetSheetList()

	data := []InputData{}

	for _, ws := range wsList {
		var d InputData
		d.Position = ws

		// 変なやり方だとは思うけど
		cols := []string{"A%d", "B%d", "C%d", "D%d", "E%d"}
		for _, col := range cols {
			topCell := fmt.Sprintf(col, 1)
			cellValue, err := f.GetCellValue(ws, topCell)
			if err != nil {
				log.Fatal(err, "1行目が変")
			}

			// データは2行目から
			row := 2
			for {
				cell, err := f.GetCellValue(ws, fmt.Sprintf(col, row))
				if err != nil || cell == "" {
					break
				}

				switch cellValue {
				case "Team":
					d.Team = append(d.Team, cell)
				case "Name":
					d.Name = append(d.Name, cell)
				case "Relay":
					d.Relay = append(d.Relay, cell)
				case "Target Number":
					d.Target = append(d.Target, cell)
				case "Nation":
					d.Nation = append(d.Nation, cell)
				default:
					continue
				}
				row++
			}
		}

		data = append(data, d)
	}

	return data, nil
}

// Id情報を射手一覧から読み込む
func readIdData(input string) ([]IdData, error) {
	f, err := excelize.OpenFile(input)
	if err != nil {
		log.Fatal(err, "日ラID情報が読み込めませんでした: ", input)
	}
	defer f.Close()

	var data []IdData

	row := 2
	for {
		id, err1 := f.GetCellValue("Sheet1", fmt.Sprintf("F%d", row))
		fName, err2 := f.GetCellValue("Sheet1", fmt.Sprintf("C%d", row))
		lName, err3 := f.GetCellValue("Sheet1", fmt.Sprintf("B%d", row))
		roma, err4 := f.GetCellValue("Sheet1", fmt.Sprintf("E%d", row))

		// idについては、いろんな書き方をするチームがあるので、ここできれいにしておく
		id = sanitize(id)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("何か変", row)
		}
		// 姓が空になったら終了
		if lName == "" {
			break
		}

		data = append(data, IdData{
			Id:     id,
			FName:  fName,
			LName:  lName,
			Romaji: roma,
		})

		row++
	}

	return data, nil
}

// IDに入っていると喜ばしくない文字を消去
// lettersによく含まれる文字を列挙
func sanitize(id string) string {
	letters := []string{" ", "\u3000", "_", "-"}

	for _, l := range letters {
		id = strings.ReplaceAll(id, l, "")
	}
	return id
}
