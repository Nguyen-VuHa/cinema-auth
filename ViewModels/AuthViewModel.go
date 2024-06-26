package viewmodels

import "service-auth/DTO"

// type response đăng ký tài khoản
type SigUpViewModel struct {
	DTO.BaseReponseDTO
	Data DTO.AuthService_SignUp_Response `json:"data"`
}

// type response đăng nhập tài khoản
type SignInViewModel struct {
	DTO.BaseReponseDTO
	Data DTO.AuthService_SignIn_Response `json:"data"`
}

// type response đăng nhập tài khoản
type SignInFacebookViewModel struct {
	DTO.BaseReponseDTO
	Data DTO.AuthService_SignInFacebook_Response `json:"data"`
}

// type response đăng nhập tài khoản
type CallBackFacebookViewModel struct {
	DTO.BaseReponseDTO
	Data DTO.AuthService_Callback_Facebook_Response `json:"data"`
}
