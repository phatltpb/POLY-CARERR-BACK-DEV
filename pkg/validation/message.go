package validationMessage

var (
	// field message
	Missing = "Vui lòng điền '<field>'"
	IsValid = "'<field>' không đúng định dạng!"
	IsExits = "'<field>' Đã tồn tại!"

	// error message
	TokenMissing = "Bạn chưa đăng nhập, vui lòng đăng nhập"
	TokenInvalid = "Bạn chưa đăng nhập, vui lòng đăng nhập!"
	TokenExpired = "Phiên làm việc đã hết hạn, vui lòng đăng nhập lại!"
	Unauthorized = "Bạn không có quyền truy cập vào đường dẫn này!"
	HaveSomeErr  = "Đã có lỗi xảy ra vui lòng thử lại sau"
	ApiNotFound  = "Api not Found! Please check again!"
	// ApiNotFound  = "Server intver"

	// message error
	LoginFalse         = "Sai tài khoản hoặc mật khẩu!"
	RequireCompany     = "Bạn chưa có công ty, vui lòng tạo công ty!"
	CompanyIsExits     = "Bạn đã có công ty"
	IdIsNotExits       = "Id không tồn tại"
	UserIsNotExits     = "Người dùng không tồn tại"
	UnActive           = " Tài khoản chưa xác thực!"
	UploadError        = "không thể tải file"
	TimeExpired        = "Bạn đã hết thời gian thực hiện thao tác này!"
	UpdateError        = "Đã có lỗi xảy ra khi bạn thay đổi thông tin, vui lòng thử lại"
	ApplyJobIsExit     = "Bạn chỉ có thể ứng tuyển vào vị trí này 1 lần!"
	EmailIsExit        = "Email đã tồn tại!"
	PasswordWrong      = "Mật khẩu hiện tại không đúng!"
	LoginFalseManyTime = "Bạn đã nhập sai tài khoản hoặc mật khẩu quá nhiều lần. Vui lòng quay lại sau một vài phút!"
)
