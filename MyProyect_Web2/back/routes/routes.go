package routes

import (
	"back/config"
	"back/controllers"
	"back/middleware"

	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Ruta para obtener todos los usuarios
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Ruta para crear un nuevo usuario
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r)
	}).Methods("POST")

	// Ruta para eliminar un usuario por ID (protegida con token)
	r.Handle("/users/{id}", middleware.ValidateToken(http.HandlerFunc(controllers.DeleteUser))).Methods("DELETE")

	// Rutas para el menú
	r.HandleFunc("/menu", controllers.GetMenu).Methods("GET")

	// Rutas para pedidos
	r.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")

	// Rutas para historial de ventas
	r.HandleFunc("/sales-history", controllers.GetSalesHistory).Methods("GET")

	// Ruta para verificar la conexión a la base de datos
	r.HandleFunc("/ping-db", func(w http.ResponseWriter, r *http.Request) {
		if err := config.DB.Ping(); err != nil {
			http.Error(w, "No se pudo conectar a la base de datos", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Conexión exitosa a la base de datos"))
	}).Methods("GET")

	// Ruta para ver reportes
	r.HandleFunc("/view-reports", controllers.ViewReports).Methods("GET")

	// Ruta para crear un reporte
	r.HandleFunc("/reports", controllers.CreateReport).Methods("POST")

	// Ruta para obtener todos los reportes
	r.HandleFunc("/reports", controllers.GetReports).Methods("GET")

	// Ruta para autorizar descuentos (protegida con token)
	r.Handle("/authorize-discounts", middleware.ValidateToken(http.HandlerFunc(controllers.AuthorizeDiscounts))).Methods("POST")
	r.Handle("/authorize-discounts", middleware.ValidateToken(http.HandlerFunc(controllers.GetAuthorizedDiscounts))).Methods("GET")

	// Ruta para autorizar cambios de contraseña (protegida con token)
	r.Handle("/authorize-password-change", middleware.ValidateToken(http.HandlerFunc(controllers.AuthorizePasswordChange))).Methods("POST")

	// Ruta para obtener los registros de permisos (protegida con token)
	r.Handle("/permissions-logs", middleware.ValidateToken(http.HandlerFunc(controllers.GetPermissionsLogs))).Methods("GET")

	// Ruta para generar token
	r.HandleFunc("/generate-token", controllers.GenerateToken).Methods("POST")
}
