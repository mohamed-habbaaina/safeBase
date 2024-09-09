package model

type DatabaseConnection struct {
	Name     string `json:"name" validate:"required"`
	Host     string `json:"host" validate:"required,hostname|ip"`
	Port     int    `json:"port" validate:"required,min=1,max=65535"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:"required,oneof=mysql postgres"`
}
