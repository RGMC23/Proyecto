package controllers

import (
	"back/config"
	"back/middleware"
	models "back/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var rolePermissions = map[string][]string{
	"Dueño":    {"gestionar_usuarios", "gestionar_negocio", "ver_reportes"},
	"Admin":    {"gestionar_usuarios", "ver_reportes"},
	"Employee": {"ver_tareas", "registrar_ventas"},
}

// Función genérica para verificar permisos y ejecutar lógica
func HandleWithPermission(w http.ResponseWriter, r *http.Request, userRole string, permission string, action func()) {
	if !models.CheckPermission(userRole, permission) {
		http.Error(w, "No tienes permiso para realizar esta acción", http.StatusForbidden)
		return
	}
	action()
}

// Obtener todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, email, username, role FROM users")
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Role); err != nil {
			http.Error(w, "Error scanning user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Name     string
		Email    string
		Username string
		Password string
		Role     string
	}

	// Decode the JSON body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Respond with a success message (database logic should be added here)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// Crear un nuevo usuario con contraseña aleatoria
func CreateUserWithRandomPassword(w http.ResponseWriter, r *http.Request) {
	// Obtener el rol del usuario autenticado desde el encabezado
	rolUsuario := r.Header.Get("Role")

	// Validar que el usuario sea Dueño o Admin
	if rolUsuario != "Dueño" && rolUsuario != "Admin" {
		http.Error(w, "No tienes permiso para crear usuarios", http.StatusForbidden)
		return
	}

	var user struct {
		Name     string
		Email    string
		Username string
		Role     string
	}

	// Decodificar el cuerpo JSON en la estructura del usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Validar que los campos requeridos no estén vacíos
	if user.Name == "" || user.Email == "" || user.Username == "" || user.Role == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	// Generar una contraseña aleatoria
	randomPassword := models.GenerateRandomPassword(12)

	// Encriptar la contraseña
	hashedPassword, err := models.EncryptPassword(randomPassword)
	if err != nil {
		http.Error(w, "Error al encriptar la contraseña", http.StatusInternalServerError)
		return
	}

	// Guardar el usuario en la base de datos
	query := `INSERT INTO users (name, email, username, password, role) VALUES ($1, $2, $3, $4, $5)`
	result, err := config.DB.Exec(query, user.Name, user.Email, user.Username, hashedPassword, user.Role)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta SQL:", err)
		http.Error(w, "Error al guardar el usuario en la base de datos", http.StatusInternalServerError)
		return
	}

	// Verificar si se insertó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error al verificar las filas afectadas:", err)
		http.Error(w, "Error al verificar la inserción", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		fmt.Println("No se insertó ningún registro en la base de datos")
		http.Error(w, "No se pudo crear el usuario", http.StatusInternalServerError)
		return
	}

	fmt.Println("Usuario creado correctamente en la base de datos")

	// Responder con éxito y mostrar la contraseña generada
	response := map[string]interface{}{
		"message":  "Usuario creado exitosamente",
		"username": user.Username,
		"password": randomPassword, // Mostrar la contraseña generada
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Función para gestionar el negocio
func ManageBusiness(w http.ResponseWriter, r *http.Request) {
	userRole := "Dueño" // Esto debería obtenerse del usuario autenticado
	HandleWithPermission(w, r, userRole, "gestionar_negocio", func() {
		w.Write([]byte("Negocio gestionado con éxito"))
	})
}

// Función para ver reportes (compartida entre Dueño y Admin)
func ViewReports(w http.ResponseWriter, r *http.Request) {
	userRole := "Admin" // Esto debería obtenerse del usuario autenticado
	if userRole != "Dueño" && userRole != "Admin" {
		http.Error(w, "No tienes permiso para ver reportes", http.StatusForbidden)
		return
	}
	HandleWithPermission(w, r, userRole, "ver_reportes", func() {
		w.Write([]byte("Reportes mostrados"))
	})
}

// Función para gestionar empleados (compartida entre Dueño y Admin)
func ManageEmployees(w http.ResponseWriter, r *http.Request) {
	userRole := "Admin" // Esto debería obtenerse del usuario autenticado
	if userRole != "Dueño" && userRole != "Admin" {
		http.Error(w, "No tienes permiso para gestionar empleados", http.StatusForbidden)
		return
	}
	HandleWithPermission(w, r, userRole, "gestionar_empleados", func() {
		w.Write([]byte("Empleados gestionados"))
	})
}

// Función para ver el historial de ventas (compartida entre Dueño, Admin y Employee)
func ViewSalesHistory(w http.ResponseWriter, r *http.Request) {
	userRole := "Admin" // Esto debería obtenerse del usuario autenticado
	if userRole != "Dueño" && userRole != "Admin" && userRole != "Employee" {
		http.Error(w, "No tienes permiso para ver el historial de ventas", http.StatusForbidden)
		return
	}
	HandleWithPermission(w, r, userRole, "ver_historial_ventas", func() {
		w.Write([]byte("Historial de ventas mostrado"))
	})
}

// Función para gestionar usuarios (compartida entre Dueño y Admin)
func ManageUsers(w http.ResponseWriter, r *http.Request) {
	userRole := "Admin" // Esto debería obtenerse del usuario autenticado
	if userRole != "Dueño" && userRole != "Admin" {
		http.Error(w, "No tienes permiso para gestionar usuarios", http.StatusForbidden)
		return
	}
	HandleWithPermission(w, r, userRole, "gestionar_usuarios", func() {
		w.Write([]byte("Usuarios gestionados"))
	})
}

// Función para autorizar descuentos
func AuthorizeDiscounts(w http.ResponseWriter, r *http.Request) {
	userRole := "Dueño" // Esto debería obtenerse del usuario autenticado
	userID := 1         // ID del usuario autenticado (esto debería venir de un token o sesión)

	HandleWithPermission(w, r, userRole, "autorizar_descuentos", func() {
		var requestBody struct {
			DiscountID int    `json:"discount_id"`
			Status     string `json:"status"` // "aprobado" o "rechazado"
		}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Solicitud inválida", http.StatusBadRequest)
			return
		}

		// Registrar la acción en la base de datos
		query := `INSERT INTO permissions_logs (user_id, action, status) VALUES ($1, $2, $3)`
		_, dbErr := config.DB.Exec(query, userID, "autorizar_descuentos", requestBody.Status)
		if dbErr != nil {
			http.Error(w, "Error al registrar la acción", http.StatusInternalServerError)
			return
		}

		// Responder con éxito
		response := map[string]interface{}{
			"message":    "Descuento autorizado exitosamente",
			"discountID": requestBody.DiscountID,
			"status":     requestBody.Status,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}

// Función para autorizar cambios de contraseña
func AuthorizePasswordChange(w http.ResponseWriter, r *http.Request) {
	userRole := "Dueño" // Esto debería obtenerse del usuario autenticado
	userID := 1         // ID del usuario autenticado (esto debería venir de un token o sesión)

	HandleWithPermission(w, r, userRole, "autorizar_cambio_contraseña", func() {
		var requestBody struct {
			UserID int    `json:"user_id"`
			Status string `json:"status"` // "aprobado" o "rechazado"
		}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Solicitud inválida", http.StatusBadRequest)
			return
		}

		// Registrar la acción en la base de datos
		query := `INSERT INTO permissions_logs (user_id, action, status) VALUES ($1, $2, $3)`
		_, dbErr := config.DB.Exec(query, userID, "autorizar_cambio_contraseña", requestBody.Status)
		if dbErr != nil {
			http.Error(w, "Error al registrar la acción", http.StatusInternalServerError)
			return
		}

		// Responder con éxito
		response := map[string]interface{}{
			"message": "Cambio de contraseña autorizado exitosamente",
			"userID":  requestBody.UserID,
			"status":  requestBody.Status,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}

// Función para gestionar el inventario (compartida entre Dueño y Admin)
func ManageInventory(w http.ResponseWriter, r *http.Request) {
	userRole := "Admin" // Esto debería obtenerse del usuario autenticado
	if userRole != "Dueño" && userRole != "Admin" {
		http.Error(w, "No tienes permiso para gestionar el inventario", http.StatusForbidden)
		return
	}
	HandleWithPermission(w, r, userRole, "gestionar_inventario", func() {
		w.Write([]byte("Inventario gestionado"))
	})
}

// Función para ver tareas (solo para Employee)
func ViewTasks(w http.ResponseWriter, r *http.Request) {
	userRole := "Employee" // Esto debería obtenerse del usuario autenticado
	HandleWithPermission(w, r, userRole, "ver_tareas", func() {
		w.Write([]byte("Tareas mostradas"))
	})
}

// Función para registrar ventas (solo para Employee)
func RegisterSales(w http.ResponseWriter, r *http.Request) {
	userRole := "Employee" // Esto debería obtenerse del usuario autenticado
	HandleWithPermission(w, r, userRole, "realizar_ventas", func() {
		w.Write([]byte("Venta registrada con éxito"))
	})
}

// Función para actualizar tareas (solo para Employee)
func UpdateTasks(w http.ResponseWriter, r *http.Request) {
	userRole := "Employee" // Esto debería obtenerse del usuario autenticado
	HandleWithPermission(w, r, userRole, "actualizar_tareas", func() {
		// Lógica para actualizar tareas
		w.Write([]byte("Tarea actualizada con éxito"))
	})
}

// GetPermissionsLogs obtiene los registros de permisos desde la base de datos
func GetPermissionsLogs(w http.ResponseWriter, r *http.Request) {
	// Consulta para obtener los registros de la tabla permissions_logs
	query := `SELECT id, user_id, action, status, created_at FROM permissions_logs`
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, "Error al obtener los registros de permisos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Estructura para almacenar los registros
	var logs []struct {
		ID        int    `json:"id"`
		UserID    int    `json:"user_id"`
		Action    string `json:"action"`
		Status    string `json:"status"`
		CreatedAt string `json:"created_at"`
	}

	// Iterar sobre los resultados y agregarlos a la lista
	for rows.Next() {
		var log struct {
			ID        int    `json:"id"`
			UserID    int    `json:"user_id"`
			Action    string `json:"action"`
			Status    string `json:"status"`
			CreatedAt string `json:"created_at"`
		}
		err := rows.Scan(&log.ID, &log.UserID, &log.Action, &log.Status, &log.CreatedAt)
		if err != nil {
			http.Error(w, "Error al procesar los registros", http.StatusInternalServerError)
			return
		}
		logs = append(logs, log)
	}

	// Responder con los registros en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// GetAuthorizedDiscounts obtiene los descuentos autorizados desde la base de datos
func GetAuthorizedDiscounts(w http.ResponseWriter, r *http.Request) {
	// Consulta para obtener los descuentos autorizados
	query := `SELECT id, user_id, action, status, created_at FROM permissions_logs WHERE action = 'autorizar_descuentos'`
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, "Error al obtener los descuentos autorizados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Estructura para almacenar los registros
	var discounts []struct {
		ID        int    `json:"id"`
		UserID    int    `json:"user_id"`
		Action    string `json:"action"`
		Status    string `json:"status"`
		CreatedAt string `json:"created_at"`
	}

	// Iterar sobre los resultados y agregarlos a la lista
	for rows.Next() {
		var discount struct {
			ID        int    `json:"id"`
			UserID    int    `json:"user_id"`
			Action    string `json:"action"`
			Status    string `json:"status"`
			CreatedAt string `json:"created_at"`
		}
		err := rows.Scan(&discount.ID, &discount.UserID, &discount.Action, &discount.Status, &discount.CreatedAt)
		if err != nil {
			http.Error(w, "Error al procesar los registros", http.StatusInternalServerError)
			return
		}
		discounts = append(discounts, discount)
	}

	// Responder con los registros en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discounts)
}

// Eliminar un usuario por ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Obtener el rol del usuario autenticado desde el encabezado (agregado por el middleware)
	rolUsuario := r.Header.Get("Role")
	fmt.Println("Rol recibido desde el token:", rolUsuario)

	// Validar que el usuario tenga el permiso "gestionar_usuarios"
	if !models.CheckPermission(rolUsuario, "gestionar_usuarios") {
		fmt.Println("Permiso denegado para el rol:", rolUsuario)
		http.Error(w, "No tienes permiso para eliminar usuarios", http.StatusForbidden)
		return
	}

	// Obtener el ID del usuario desde los parámetros de la URL
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Println("ID recibido:", userID)

	// Validar que el ID no esté vacío
	if userID == "" {
		http.Error(w, "El ID del usuario es requerido", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta para eliminar el usuario
	query := `DELETE FROM users WHERE id = $1`
	result, err := config.DB.Exec(query, userID)
	if err != nil {
		http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
		return
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error al verificar la eliminación", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado exitosamente"))
}

// Función para generar un token
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Role string `json:"role"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || (requestBody.Role != "Dueño" && requestBody.Role != "Admin" && requestBody.Role != "Employee") {
		http.Error(w, "Rol inválido", http.StatusBadRequest)
		return
	}

	token, err := middleware.GenerateToken(requestBody.Role)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
