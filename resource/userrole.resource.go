package resource

import (
	// "mime/multipart"
	// "os"

	"gin-starter/entity"
	// "gin-starter/utils"
)

// const (
// 	timeFormat = "2006-01-02 15:04:05"
// )

// GetUserRoleResponse
type GetUserRoleRespone struct {
	List  []*UserRole `json:"list"`
	Meta  *Meta       `json:"meta"`
}

type GetUserRoleRequest struct {
	Search string `form:"search" json:"search"`
	Sort  string `form:"sort" json:"sort"`
	Order string `form:"order" json:"order"`
	Limit int    `form:"limit,default=10" json:"limit"`
	Page  int    `form:"page,default=0" json:"page"`
}

// GetUserRoleByID
type GetUserRoleByID struct {
	ID string `uri:"id" binding:"required"`
}

type CreateUserRoleRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

type UpdateUserRoleRequest struct {
	ID          string `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}

type DeleteUserRoleRequest struct {
	ID string `uri:"id" binding:"required"`
}

type UserRole struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Description string `json:"description"`
}

func NewUserRole(user *entity.UserRole) *UserRole {
	// otpIsNull := false
	// if user.OTP.String != "" {
	// 	otpIsNull = true
	// }

	// dob := "1970-01-01"
	// if user.DOB.Valid {
	// 	dob = user.DOB.Time.Format(timeFormat)
	// }

	return &UserRole{
		ID:          user.ID.String(),
		Name:        user.Name,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.Format(timeFormat),
		UpdatedAt:   user.UpdatedAt.Format(timeFormat),
	}
}

// type ChangePasswordRequest struct {
// 	OldPassword             string `form:"old_password" json:"old_password" binding:"required"`
// 	NewPassword             string `form:"new_password" json:"new_password" binding:"required"`
// 	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" binding:"required"`
// }

// type ForgotPasswordRequest struct {
// 	Email string `form:"email" json:"email" binding:"required"`
// }

// type ForgotPasswordChangeRequest struct {
// 	Token                   string `form:"token" json:"token" binding:"required"`
// 	NewPassword             string `form:"new_password" json:"new_password" binding:"required"`
// 	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" binding:"required"`
// }

// type GetUserByForgotPasswordTokenRequest struct {
// 	Token string `uri:"token" json:"token" binding:"required"`
// }

// type VerifyOTPRequest struct {
// 	Code string `form:"code" json:"code" binding:"required"`
// }
