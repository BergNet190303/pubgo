package publice

type ResultDef struct {
	ErrCode int
	ErrMsg  string
	Code    string
}

type ResultDesign struct {
	ErrCode int
	Msg     string
	Data    interface{}
	Code    string
}

type ResultDesignOverRall struct {
	Code    string
	ErrCode int
	Msg     string
	Data    interface{}
	Num     int
}

func SuccessDef(msg string) ResultDef {
	return ResultDef{
		ErrCode: 0,
		ErrMsg:  msg,
		Code:    "1",
	}
}

func FiledDef(msg string) ResultDef {
	return ResultDef{
		ErrMsg:  msg,
		ErrCode: 1,
		Code:    "0",
	}
}

func Result(msg string, res interface{}, ErrCode int, Code string) ResultDesign {
	result := ResultDesign{
		Msg:     msg,
		Data:    res,
		ErrCode: ErrCode,
		Code:    Code,
	}
	return result
}

func OverRallResult(msg string, res interface{}, ErrCode int, Code string) ResultDesignOverRall {
	result := ResultDesignOverRall{
		Msg:     msg,
		Data:    res,
		ErrCode: ErrCode,
		Code:    Code,
	}
	return result
}
