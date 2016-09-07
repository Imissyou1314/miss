package models

import (
    "miss/utils"
)

/**
 *  交互的数据格式
 * 定了了交互的数据格式
 */

type ResultMsg struct {
    //获取信息
    ResultInfo string
    //获取错误的信息
    ResultError string
    //获取是否成功
    isSuccess bool
    //获取数据
    ResultData interface{}
}

/**
 *  获取Result对象
 *  @param: isSucess 是否获取数据成功
 *  @param: ResultError 获取出现错误的错误数据
 *  @param: ResultInfo  获取信息的接送
 *  @parm:  ResultData 获取的信息
 *  @resutl:    ResultMsg对象
 */
func GetResultMsg(isSuccess bool,
                ResultError string,
                ResultInfo string,
                ResultData interface{}) (resutl ResultMsg) {
                var result ResultMsg
                result.ResultInfo = ResultInfo
                result.ResultError = ResultError
                result.ResultData = ResultData
                result.isSuccess = false
                result.isSuccess = isSuccess
                return
            }

/**
 *  获取Json数据的交互对象
 *  @result:    []byte  转换候的对象
 *  @result:     error  错误信息
 */
func GetReusltJsonStr(isSuccess bool,
                ResultError string,
                ResultInfo string,
                ResultData interface{}) ([]byte, error) {
                return utils.ToJson(
                    GetResultMsg(isSuccess,
                        ResultError,
                        ResultInfo,
                        ResultData))
            }
