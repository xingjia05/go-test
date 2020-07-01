// Package myexcel ...
package myexcel

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"

	_ "github.com/go-sql-driver/mysql" //仅仅是是希望它执行init()函数而已
)

// 仅仅是是希望它执行init()函数而已

type ora struct {
	name       string
	userName   string
	phone      string
	email      string
	gender     string
	mobile     string
	cardid     string
	position   string
	structCode string
}

type ret struct {
	data    retData
	code    int
	status  int
	message string
}

type retData struct {
	isSuccess int
	list      interface{}
}

// MyExcel ...
func MyExcel() {
	//test()
	//var name string
	//fmt.Println("please enter your name:")
	//fmt.Scanf("%s", &name)

	//excelData := readExcel(name)
	//writeExcel(excelData)
	//writeExcel1(excelData)

	structMap := getStructMap()

	sendFile(structMap)
}

// sendFile ...
func sendFile(structMap map[string]int) {
	dir, err := ioutil.ReadDir("/tmp/file1/")
	if err != nil {
		fmt.Println("read dir fail,error:", err)
	}
	for _, file := range dir {
		structCode := strings.TrimSuffix(file.Name(), ".xlsx")
		if structMap[structCode] == 0 {
			continue
		}

		bodyBuf := &bytes.Buffer{}
		sendWriter := multipart.NewWriter(bodyBuf)                                      //创建新的写入
		fileWriter, err := sendWriter.CreateFormFile("file", "/tmp/file1/"+file.Name()) //创建form上传文件
		checkErr(err)
		fmt.Println("/tmp/file1/" + file.Name())

		f, err := os.Open("/tmp/file1/" + file.Name()) //打开要发送的文件
		checkErr(err)
		defer f.Close()

		_, err = io.Copy(fileWriter, f)
		checkErr(err)

		sendWriter.WriteField("structId", strconv.Itoa(structMap[structCode]))

		//formType := sendWriter.FormDataContentType()
		sendWriter.Close()
		client := &http.Client{}
		//req, _ := http.NewRequest("POST", "http://rprprprp_local.com:8997/api/staff/import", bodyBuf)
		//req.Header.Add("Cookie", "zlenv=dev; UM_distinctid=1719f9af9c6606-051abe39c3633f-30657607-fa000-1719f9af9c7832; awt=ST-175643-SwWmvb9bEJJCvDzq8b34lAlJ7Xk-avataruuap-695f4b788-jt4bw; st=ST-196-3vFgx2mwYDa6IXKcOPCnOxWplIA-localhost; PHPSESSID=9ljk8lv9gkelhhpqu8u10po7nd; structCompanyCode=RPX0001; brandId=13")
		//req.Header.Add("Content-Type", "Content-Type: multipart/form-data; boundary=----WebKitFormBoundary178X77fTKeSTjOhy")
		//req, _ := http.NewRequest("POST", "https://acp-test.rp-field.com/api/staff/import", bodyBuf)
		//req.Header.Add("Cookie", "zlenv=test; UM_distinctid=1720d10ec7823-06c13141a56458-30657607-fa000-1720d10ec7b1e5; st=ST-175704-d-68dofIhVWmpg3-rLhvt4WXIsg-avataruuap-695f4b788-jt4bw; PHPSESSID=qt1991kggdqoaj8kl7ugja4gsr; structCompanyCode=RPX0001; brandId=13")
		req, _ := http.NewRequest("POST", "https:///api/staff/import", bodyBuf)
		req.Header.Add("Cookie", "UM_distinctid=1720d10ec7823-06c13141a56458-30657607-fa000-1720d10ec7b1e5; PHPSESSID=6qd0h1a6mmdianh3tjapce2qn5; st=ST-16860-s7WagKKgMgI5rXS1F1RuNtoWpWM-avataruuap-5ff4f7df5c-7xmkn; structCompanyCode=RPX0001; brandId=13")
		req.Header.Add("Content-Type", sendWriter.FormDataContentType())
		resp, err := client.Do(req)
		checkErr(err)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		//var res ret
		var res interface{}
		json.Unmarshal(body, &res)
		if res == nil {
			fmt.Println("fail:", file.Name())
			continue
		}
		resM := res.(map[string]interface{})
		for k, v := range resM {
			if k != "data" {
				continue
			}
			resV := v.(map[string]interface{})
			for kk, vv := range resV {
				if kk == "isSuccess" && vv.(float64) == 0 {
					fmt.Println(v)
				}
			}
		}
	}
}

// getStructmap ...
func getStructMap() map[string]int {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.5:3306)/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	rows, err := db.Query("select * from structMap")
	checkErr(err)

	structMap := make(map[string]int)
	for rows.Next() {
		var structOuterCode string
		var id, structID int
		rows.Scan(&id, &structOuterCode, &structID)
		structMap[structOuterCode] = structID
	}
	return structMap
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func test() {
	name := "111"
	file := excelize.NewFile()
	file.SetSheetRow("Sheet1", "A1", &[]interface{}{1, 2, 3, 4, 5, 6, 7, 8})
	if err := file.SaveAs("/tmp/file/" + name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

// writeExcel1 ...
func writeExcel1(data []ora) {
	groupMap := make(map[string][]ora)
	for _, row := range data {
		groupMap[row.structCode] = append(groupMap[row.structCode], row)
	}
	for name, listMap := range groupMap {
		name = strings.Trim(name, " \r\n\"\t\r\n ")
		name = strings.TrimRight(name, " \r\n\"\t ")
		name = strings.ToUpper(name)
		file := excelize.NewFile()
		file.SetSheetRow("Sheet1", "A1", &[]interface{}{1, 2, 3, 4, 5, 6, 7, 8})
		file.SetSheetRow("Sheet1", "A2", &[]interface{}{1, 2, 3, 4, 5, 6, 7, 8})
		for k, item := range listMap {
			t := k + 3
			tStr := strconv.Itoa(t)
			file.SetSheetRow("Sheet1", "A"+tStr, &[]interface{}{string(item.name), string(item.userName), string(item.phone), string(item.email), string(item.gender), string(item.mobile), string(item.position), string(item.cardid)})
		}
		if err := file.SaveAs("/tmp/file/" + name + ".xlsx"); err != nil {
			fmt.Println(err)
		}
	}
}

// writeExcel ...
func writeExcel(data []ora) {
	groupMap := make(map[string][]ora)
	for _, row := range data {
		groupMap[row.structCode] = append(groupMap[row.structCode], row)
	}
	for name, listMap := range groupMap {
		name = strings.Trim(name, " \r\n\"\t\r\n ")
		name = strings.TrimRight(name, " \r\n\"\t ")
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("sheet1")
		if err != nil {
			fmt.Println(err.Error())
		}
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"
		cell = row.AddCell()
		cell.Value = "1111"

		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"
		cell = row.AddCell()
		cell.Value = "22222"

		for _, item := range listMap {
			row := sheet.AddRow()
			cell = row.AddCell()
			cell.Value = item.name

			cell = row.AddCell()
			cell.Value = item.userName
			cell = row.AddCell()
			cell.Value = item.phone

			cell = row.AddCell()
			cell.Value = item.email

			cell = row.AddCell()
			cell.Value = item.gender

			cell = row.AddCell()
			cell.Value = item.mobile

			cell = row.AddCell()
			cell.Value = item.position

			cell = row.AddCell()
			cell.Value = item.cardid
		}
		err = file.Save("/tmp/file/" + name + ".xlsx")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// readExcel ...
func readExcel(name string) []ora {
	var listOra []ora
	xlFile, err := xlsx.OpenFile(name)
	if err != nil {
		fmt.Println("failed to open xlsx,reason:", err)
	}
	for k, sheet := range xlFile.Sheets {
		if k != 0 {
			continue
		}
		fmt.Println(k, sheet.Name)
		tmpOra := ora{}
		for k, row := range sheet.Rows {
			if k < 3 {
				continue
			}
			var strs []string

			for _, cell := range row.Cells {
				text := cell.String()
				strs = append(strs, text)
			}

			tmpOra.name = strs[0]
			tmpOra.userName = strs[1]
			tmpOra.phone = strs[2]
			tmpOra.email = strs[3]
			tmpOra.gender = strs[4]
			tmpOra.mobile = strs[5]
			tmpOra.cardid = strs[7]
			tmpOra.position = strs[6]
			tmpOra.structCode = strs[8]
			listOra = append(listOra, tmpOra)
		}
	}
	return listOra
}
