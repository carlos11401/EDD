package Estructura



type NodoEzdo struct{
	Id int
	Siguiente *NodoEzdo
	Anterior *NodoEzdo
	Acceso *Nodo
}

// LISTA PARA LOS ENCABEZADOS
type Encabezado struct{
	Primero *NodoEzdo
	Ultimo *NodoEzdo
	Size int 
}


func New_NodoEzdo(val int, nuevoNodo *Nodo) *NodoEzdo{
	return &NodoEzdo{val,nil,nil,nuevoNodo}
}
func New_Encabezado() *Encabezado{
	return &Encabezado{nil,nil,0}
}

func InsertarEncabezado(nuevo *NodoEzdo, lista *Encabezado)  {

	if lista.Primero == nil { 
		lista.Primero = nuevo
		lista.Ultimo = nuevo
		lista.Size += 1
	}else{
		if nuevo.Id < lista.Primero.Id { //insercion al inicio
			nuevo.Siguiente = lista.Primero
			lista.Primero.Anterior = nuevo
			lista.Primero = nuevo
			lista.Size += 1
		} else if nuevo.Id > lista.Ultimo.Id { //insercion al final
			lista.Ultimo.Siguiente = nuevo
			nuevo.Anterior = lista.Ultimo
			lista.Ultimo = nuevo
			lista.Size += 1
		} else{ //busco la posicion correcta
			tmp := lista.Primero
			for tmp != nil{
				if nuevo.Id < tmp.Id { //busco la posicion adecuada para insertar
					//trabajando con el que esta antes de tmp
					tmp.Anterior.Siguiente = nuevo
					nuevo.Anterior = tmp.Anterior
					// trabajando con tmp
					nuevo.Siguiente = tmp
					tmp.Anterior = nuevo
					//aumento el contador 
					lista.Size += 1
					break; // rompo el ciclo for para que no siga buscando
				}
				tmp = tmp.Siguiente
			}
		}
	}
	
}


func BuscarEncabezado(valor int, enc *Encabezado) *NodoEzdo{
	if enc != nil { //si existe el encabezado 
		tmp := enc.Primero
		for tmp != nil {
			if tmp.Id == valor{
				return tmp
			}
			tmp = tmp.Siguiente
		}
	}
	return nil //no existe el encabezado
}