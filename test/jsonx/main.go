package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"tpmt-zt/common"
)

type A struct {
	ApiCollection []AA `json:"apiCollection"`
}

type AA struct {
	Name  string `json:"name"`
	Items []AA   `json:"items"`
	Api   ApiNow `json:"api"`
}

type ApiNow struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func main() {

	path := common.GetCurrentDirectory()

	path = filepath.Join(path, "tpmt-http.apifox.json")

	fmt.Println("path:", path)

	cc, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return

	}

	var aa A
	json.Unmarshal(cc, &aa)

	var data []*Redata
	for _, v := range aa.ApiCollection {
		for _, item := range v.Items {
			datas := DG(item)
			for _, redata := range datas {
				data = append(data, &Redata{
					Name:               redata.Name,
					Method:             redata.Method,
					Path:               redata.Path,
					InterfaceGroupName: item.Name,
				})
			}
		}

	}

	str, _ := os.Getwd()

	cccc, _ := getAllExcel(str)
	fmt.Println(cccc)

	for _, itemKK := range cccc {
		// 获取当前路径
		f := excelize.NewFile()

		index, _ := f.NewSheet("Sheet1")
		for i, v := range data {
			SetExcelData(f, 0, int64(i+1), v.Name)
			SetExcelData(f, 1, int64(i+1), v.Method)
			SetExcelData(f, 2, int64(i+1), v.Path)
			SetExcelData(f, 3, int64(i+1), v.InterfaceGroupName)
		}

		f.SetActiveSheet(index)
		// Save spreadsheet by the given path.
		if err := f.SaveAs(filepath.Join(itemKK + ".xlsx")); err != nil {
			fmt.Println(err)
		}

	}

}

func getAllExcel(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, 100)
	for _, filename := range files {
		allPath := path + "\\" + filename.Name()
		if IsDir(allPath) {
			nextFiles, _ := getAllExcel(allPath)
			res = append(res, nextFiles...)
		} else {
			if strings.HasSuffix(filename.Name(), ".json") {
				res = append(res, allPath)
			}
		}
	}
	return res, nil
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func DG(data AA) (res []*Redata) {
	if data.Api.Method != "" {

		if data.Api.Method == "delete" {
			data.Api.Path = strings.Replace(data.Api.Path, "/{id}", "", -10)
		}
		res = append(res, &Redata{
			Name:   data.Name,
			Method: data.Api.Method,
			Path:   data.Api.Path,
		})
	}
	if data.Items == nil {
		return res
	} else {
		for _, item := range data.Items {
			reDatas := DG(item)
			for _, reData := range reDatas {
				res = append(res, reData)
			}

		}
	}
	return res
}

type Redata struct {
	Name               string `json:"name"`
	Method             string `json:"method"`
	Path               string `json:"path"`
	InterfaceGroupName string `json:"interface_group_name"`
}

func SetExcelData(f *excelize.File, num int64, key int64, data string) int64 {

	weizhi := fmt.Sprintf("%s%v", numExchangeLetter(num), key)

	f.SetCellValue("Sheet1", weizhi, data)
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle("Sheet1", weizhi, weizhi, style)
	return num + 1

}

func numExchangeLetter(number int64) string {
	var res string
	pp := (number + 1) / 26
	if pp > 26 {
		return ""
	}

	if pp > 0 {
		resA := ExcelLetter[pp-1]
		resB := ExcelLetter[(number+1)%26]
		res = fmt.Sprintf("%s%s", resA, resB)
	} else {
		res = fmt.Sprintf("%s", ExcelLetter[number])
	}

	return res

}

var ExcelLetter = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "k", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
