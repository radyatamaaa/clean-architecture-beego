package logger


type LoggingObj struct {
	Feature string
	Message string
	PathFile string
	Host string
	RequestId string
	Data LoggingData
}
type LoggingData struct {
	Method string `json:"method"`
	HttpCode int `json:"code"`
	HttpCodeDesc string `json:"code_desc"`
	Request interface{} `json:"request"`
	Response interface{} `json:"response"`
}

func (l LoggingObj)MapError(message,pathFile,feature string) LoggingObj {
	l.Message = message
	if pathFile != ""{
		l.PathFile = pathFile
	}
	l.Feature = feature
	return l
}
