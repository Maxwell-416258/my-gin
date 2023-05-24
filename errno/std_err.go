package errno

var (
	OK                  = &Errno{0, "success!"}
	InternalServerError = &Errno{1001, "Internal server error"}
	Errbind             = &Errno{1002, "Error occurred while binding the request body struct"}
	ErrEncrypt          = &Errno{Code: 20101, Message: "Error occurred while encrypting the "}
	ErrDatabase         = &Errno{Code: 20002, Message: "Database error"}
	ErrArgsNotComplete  = &Errno{Code: 20101, Message: "Incomplete parameters"}
)
