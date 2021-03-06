package middlew

import(
	"net/http"
	"github.com/abotlucas/microblogging/bd"
)
/* ChequeoBD es el middlew que me permite conocer el estado de la BD */
func ChequeoBD (next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		/* next es la funcion recibida, si la conexcion no se perdio, ejecuta la funcion que le enviamos */
		next.ServeHTTP(w,r)
	}
}