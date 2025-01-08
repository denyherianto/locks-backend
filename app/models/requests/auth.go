package requests

import "github.com/denyherianto/go-fiber-boilerplate/app/models/entities"

// SignUp
type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Name     string `json:"name" validate:"required,lte=255"`
	Username string `json:"username" validate:"gte=3,lte=20"`
	Password string `json:"password" validate:"required,lte=255"`
}

// SignIn
type SignInRequest struct {
	Identifier string `json:"identifier" validate:"required,lte=255"`
	Password   string `json:"password" validate:"required,lte=255"`
}

type UserRoleResponse struct {
	entities.UserRole
	RoleName string `json:"role_name"`
}

type UserRolePermissionResponse struct {
	ModuleID         string `json:"module_id"`
	ModuleName       string `json:"module_name"`
	PermissionCreate bool   `json:"permission_create"`
	PermissionRead   bool   `json:"permission_read"`
	PermissionUpdate bool   `json:"permission_update"`
	PermissionDelete bool   `json:"permission_delete"`
}

type SignInResponse struct {
	AccessToken  string              `json:"access_token"`
	RefreshToken string              `json:"refresh_token"`
	Name         string              `json:"name"`
	Email        string              `json:"email"`
	Roles        *[]UserRoleResponse `json:"roles"`
}
