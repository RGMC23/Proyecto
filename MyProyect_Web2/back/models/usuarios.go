package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Estructura base para un usuario
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`    // No se incluye en la respuesta JSON por seguridad
	Role     string `json:"role"` // Puede ser "Dueño", "Admin", o "Employee"
}

// Estructura para el rol de Dueño
type Dueño struct {
	User
	BusinessName string   `json:"business_name"`
	Permissions  []string `json:"permissions"` // Lista de permisos específicos
}

// Estructura para el rol de Admin
type Admin struct {
	User
	Permissions []string `json:"permissions"`
}

// Estructura para el rol de Empleado
type Employee struct {
	User
	Department  string   `json:"department"`
	Permissions []string `json:"permissions"` // Lista de permisos específicos
}

// Método para establecer la contraseña con hash
func (u *User) SetPassword(password string) error {
	// Generar un hash seguro de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Método para verificar la contraseña
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Mapa de roles y permisos
var rolePermissions = map[string][]string{
	"Dueño":    {"gestionar_usuarios", "gestionar_negocio", "ver_reportes"},
	"Admin":    {"gestionar_usuarios", "ver_reportes"},
	"Employee": {"ver_tareas", "registrar_ventas"},
}

// Función para verificar si un rol tiene un permiso específico
func CheckPermission(role string, permission string) bool {
	fmt.Println("Rol:", role, "Permiso solicitado:", permission)
	permissions, exists := rolePermissions[role]
	if !exists {
		return false
	}

	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}
