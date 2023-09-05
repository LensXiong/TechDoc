# 将表格中的数据按照格式处理。
格式：	{"GET", "/xxx_api/v1/app/xxx_xxx_page", "", ""}
示例代码：
```
func (this DashboardController) GeTApiData(ctx *gin.Context) {
	inputFile := "/Users/xxxx/api.xlsx" // 输入的 Excel 文件名
	outputFile := "output_api.json"     // 输出的 JSON 文件名

	err := excelToJSON(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Conversion completed. JSON data saved to %s\n", outputFile)
}

func excelToJSON(inputFile string, outputFile string) error {
	// 打开 Excel 文件
	xlFile, err := xlsx.OpenFile(inputFile)
	if err != nil {
		return err
	}

	// 创建一个用于存储结果的切片
	var result []interface{}

	// 遍历 Excel 中的每个工作表
	for _, sheet := range xlFile.Sheets {
		// 获取最大的行和列数
		maxRow, maxCol := sheet.MaxRow, sheet.MaxCol

		// 遍历每一行
		for rowIdx := 0; rowIdx <= maxRow; rowIdx++ {
			// 确保行中至少有2个单元格
			if maxCol < 2 {
				continue
			}

			// 从 Excel 单元格中获取数据
			methodTemp, _ := sheet.Cell(rowIdx, 1)
			method := methodTemp.Value
			urlTemp, _ := sheet.Cell(rowIdx, 2)
			url := urlTemp.Value

			// 将数据添加到结果切片中
			result = append(result, []string{method, url, "", ""})
			result = append(result, "\t")

		}
	}

	// 将结果切片转换为 JSON 格式，并进行格式化
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// 移除末尾的换行符
	jsonString := strings.TrimSuffix(string(jsonResult), "\n")

	// 将生成的 JSON 数据中的方括号 [] 替换为大括号 {}
	jsonString = strings.ReplaceAll(jsonString, "[", "{")
	jsonString = strings.ReplaceAll(jsonString, "]", "}")

	// 去除头部的 '{' 和末尾的 '}'
	jsonString = jsonString[1 : len(jsonString)-1]

	// 将 JSON 数据写入输出文件
	err = writeFile(outputFile, []byte(jsonString))
	if err != nil {
		return err
	}

	return nil
}

func writeFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

```
