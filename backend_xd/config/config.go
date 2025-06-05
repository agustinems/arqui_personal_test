package config

import (
	"log"
	"os"
)

var (
	// jwtSecret contendrá la clave para firmar los tokens
	jwtSecret []byte
	// dsn es la cadena de conexión a MySQL
	dsn string
	// port es el puerto donde arranca el servidor HTTP
	port string
)

func init() {
	// 1) Lectura de la cadena de conexión a MySQL
	//    Ejemplo de export previo en tu terminal:
	//    export DB_DSN="usuario:User@1234@tcp(127.0.0.1:3306)/gimnasio?charset=utf8&parseTime=true&loc=Local"
	dsn = os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN no está configurado. Ejemplo:\n" +
			"  export DB_DSN=\"usuario:User@1234@tcp(127.0.0.1:3306)/gimnasio?charset=utf8&parseTime=true&loc=Local\"")
	}

	// 2) Lectura de la clave secreta para JWT
	//    Ejemplo de export previo:
	//    export JWT_SECRET="una_clave_muy_segura"
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET no está configurado. Ejemplo:\n" +
			"  export JWT_SECRET=\"una_clave_muy_segura\"")
	}
	jwtSecret = []byte(secret)

	// 3) Lectura del puerto (por defecto 8000 si no se define)
	//    Ejemplo de export previo:
	//    export PORT=8000
	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}
func Dsn() string {
	return dsn
}

func JWTSecret() []byte {
	return jwtSecret
}

func Port() string {
	return port
}
