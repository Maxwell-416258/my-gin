package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

// 实现Error方法，才属于error接口类型

type Err struct {
	Code    int
	Message string
	Err     error
}

func (err Err) Error() string {
	return fmt.Sprintf("Err -code:%d,message:%s,error:%s", err.Code, err.Message, err.Err)
}

func (err Errno) Error() string {
	return err.Message
}

// 创建错误
func New(errno *Errno, err error) *Err {
	return &Err{errno.Code, errno.Message, err}
}

// 解析定制错误，主要为了标准化返回
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch err_type := err.(type) {
	case *Err:
		return err_type.Code, err_type.Error()
	case *Errno:
		return err_type.Code, err_type.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}
