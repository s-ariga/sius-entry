// Seiichi Ariga<seiichi.ariga@gmail.com>

package sius

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// 最後に、Excelとして出力する

func writeData(file string, data []OutputData) error {
	f := excelize.NewFile()
	defer f.Close()

	for i := 0; i < len(data); i++ {
		_, err := f.NewSheet(data[i].Position)
		if err != nil {
			log.Fatal(err)
		}

		err = f.SetCellValue(data[i].Position, "A1", "Start Number")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "B1", "Name")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "C1", "Display Name")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "D1", "Nation")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "E1", "Team")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "F1", "Relay")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "G1", "Target Number")
		if err != nil {
			log.Fatal(err)
		}
		err = f.SetCellValue(data[i].Position, "H1", "IssfId")
		if err != nil {
			log.Fatal(err)
		}

		row := 2
		for j := 0; j < len(data[i].StartNum); j++ {
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("A%d", row+j), data[i].StartNum[j])
			if err != nil {
				log.Fatal(err)
			}
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("B%d", row+j), data[i].Name[j])
			if err != nil {
				log.Fatal(err)
			}
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("C%d", row+j), data[i].DispName[j])
			if err != nil {
				log.Fatal(err)
			}
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("D%d", row+j), data[i].Nation[j])
			if err != nil {
				log.Fatal(err)
			}
			// Teamについては、"団体"とは言っていたらクラブチーム名、"個人"の場合は空白にしておく
			if data[i].Team[j] == "団体" {
				err = f.SetCellValue(data[i].Position, fmt.Sprintf("E%d", row+j), data[i].Nation[j])
				if err != nil {
					log.Fatal(err)
				}
			}
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("F%d", row+j), data[i].Relay[j])
			if err != nil {
				log.Fatal(err)
			}
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("G%d", row+j), data[i].Target[j])
			if err != nil {
				log.Fatal(err)
			}
			err = f.SetCellValue(data[i].Position, fmt.Sprintf("H%d", row+j), data[i].Id[j])
			if err != nil {
				log.Fatal(err)
			}
		}

	}

	// Excelファイルを作成するとSheet1が作られてしまうので、これを消去しておく
	err := f.DeleteSheet("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	err = f.SaveAs(file)
	if err != nil {
		log.Fatal(err)
	}

	return nil

}
