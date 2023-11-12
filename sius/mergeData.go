// Seiichi Ariga <seiichi.ariga@gmail.com>

package sius

import (
	"fmt"
	"strings"
)

// 読み込んだ、射座割りの情報と日ラIDの情報をくっつけていく

func mergeData(data []InputData, idData []IdData) ([]OutputData, error) {

	var output []OutputData

	for _, d := range data {

		var out OutputData
		// 姿勢はそのまま
		out.Position = d.Position

		// Start Numberを姿勢によって作る(最初の番号)
		startNum := 0
		switch out.Position {
		case "ARM":
			startNum = ARM
		case "ARW":
			startNum = ARW
		case "R3PM":
			startNum = R3PM
		case "R3PW":
			startNum = R3PW
		case "RPRM":
			startNum = RPRM
		case "RPRW":
			startNum = RPRW
		case "ARMT":
			startNum = ARMT
		case "ARPR":
			startNum = ARPR
		default:
			continue
		}

		for i := range d.Name {
			// 分かっている部分はコピーしていく
			out.StartNum = append(out.StartNum, fmt.Sprintf("%d", startNum+i))
			out.Name = append(out.Name, d.Name[i])
			out.Team = append(out.Team, d.Team[i])
			out.Nation = append(out.Nation, d.Nation[i])
			out.Relay = append(out.Relay, d.Relay[i])
			out.Target = append(out.Target, d.Target[i])

			// 二つのデータの名前を比較して、Idとローマ字をコピー
			name := replaceSpace(d.Name[i])
			// 名前の検索が失敗する可能性があるので、データだけ作っておく
			// ここでメモリー確保 ↓へ
			out.Id = append(out.Id, "")
			out.DispName = append(out.DispName, "")
			for _, id := range idData {
				idName := catName(id.LName, id.FName)
				// 万一名前が一致しないと、ここのappendが実行されない可能性がある
				if name == idName {
					// こっちではappendせずに代入
					out.Id[i] = id.Id
					out.DispName[i] = id.Romaji
				}
			}

		}

		output = append(output, out)
	}

	return output, nil
}

// 全角あるいは半角のスペースを含んでいるかもしれない姓、名を
// 結合してスペースを消去して返す
func catName(lName, fName string) string {
	name := lName + fName

	name = strings.ReplaceAll(name, " ", "")      // 半角スペース
	name = strings.ReplaceAll(name, "\u3000", "") // 全角スペース

	return name
}

func replaceSpace(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ReplaceAll(name, "\u3000", "")

	return name
}
