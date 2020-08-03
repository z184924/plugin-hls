package hls

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type DeviceInfo struct {
	CameraName string `json:"cameraName"`
	SysCode    string `json:"sysCode"`
}

//HOST 此处替换成平台SDK所在服务器IP与端口
const HOST string = "http://111.30.79.44:90"

//APPKEY 此处替换成申请的appkey
const APPKEY string = "6f0a0539"

//SECRET 此处替换成申请的secret
const SECRET string = "6aabd44a437e4f9583b47d1ac090bb6d"

//BASEMETHOD 基本方法
const BASEMETHOD string = "/webapi/service"

//METHOD 请求方法
const METHOD string = "/vss/getPlatCameraResListByUnits"

//HKM3U8URLF 海康m3u8url前半段
// const HKM3U8URLF string = "http://172.16.104.2:6713/mag/hls/"

const HKM3U8URLF string = "http://111.30.79.44:6713/mag/hls/"

//HKM3U8URLB 海康m3u8url后半段
const HKM3U8URLB string = "/1/live.m3u8"

//GetDeviceList 获取设备列表
func GetDeviceList() []*DeviceInfo {
	var mapResult = make(map[string]interface{})
	hkResultJSON := GetPlatEncodeDeviceResListJSON()
	json.Unmarshal([]byte(hkResultJSON), &mapResult)
	deviceInfoMapListInterface := mapResult["data"].(map[string]interface{})["list"]
	jsonTemp, _ := json.Marshal(deviceInfoMapListInterface)
	var deviceInfoList []*DeviceInfo
	json.Unmarshal(jsonTemp, &deviceInfoList)
	// for i := 0; i < len(deviceInfoList); i++ {
	// 	deviceInfoMap := *deviceInfoList[i]
	// 	fmt.Println(deviceInfoMap)
	// }
	return deviceInfoList
}

//GetPlatEncodeDeviceResListJSON 获取摄像头编码信息JSON
func GetPlatEncodeDeviceResListJSON() string {
	var param = make(map[string]interface{})
	param["appkey"] = APPKEY
	param["pageNo"] = 1
	param["pageSize"] = 1000
	param["time"] = time.Now().UnixNano() / 1e6
	paramString, error := json.Marshal(param)
	var resultjson string
	if error == nil {
		var token = BuildToken(string(paramString))
		resultjson = Post(HOST+BASEMETHOD+METHOD+"?token="+token, param, "application/json; charset=UTF-8")
		// fmt.Println(resultjson)
		// fmt.Println("---------------------------------")
	}
	return resultjson
}

//BuildToken 建立token
func BuildToken(paramJSON string) string {
	var result string
	result = CreateMD5(BASEMETHOD + METHOD + paramJSON + SECRET)
	return result
}

//CreateMD5 生成md5
func CreateMD5(sourceString string) string {
	h := md5.New()
	io.WriteString(h, sourceString)
	var md5str = fmt.Sprintf("%x", h.Sum(nil))
	return md5str
}
