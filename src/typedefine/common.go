package typedefine

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ResponseHttpErrRsp(message string) HttpErrRsp {
	httpRsp := HttpErrRsp{}
	httpRsp.Code = 40020
	httpRsp.Message = message
	return httpRsp
}

func ResponseHttpSucRsp(data interface{}) HttpSucRsp {
	httpRsp := HttpSucRsp{}
	httpRsp.Code = 200
	httpRsp.Message = "ok"
	httpRsp.Data = data
	return httpRsp
}

func ReadJSONFileToStruct(filePath string, target interface{}) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// 将 JSON 数据解析到目标结构体
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return nil
}
