package THash

import (
	"log"
	"strconv"
)

type TablaHash struct {
	Tabla       map[int]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) Insertar(carnet int, name string, pass string, c1 int, c2 int, c3 int) {
	index := t.Calc_Index(carnet)
	var cursos []int
	cursos = append(cursos, c1)
	cursos = append(cursos, c2)
	cursos = append(cursos, c3)
	newNodo := &NodoHash{Llave: index, Dato: &Estudiante{Carnet: carnet, Name: name, Password: pass, Cursos: cursos}}
	if index < t.Capacidad {
		if _, existe := t.Tabla[index]; !existe {
			t.Tabla[index] = *newNodo
			t.Utilizacion++
			t.capacidadT()
		} else {
			contador := 1
			index = t.reIndex(carnet, contador)
			for {
				if _, existe := t.Tabla[index]; existe {
					contador++
					index = t.reIndex(carnet, contador)
				} else {
					newNodo.Llave = index
					t.Tabla[index] = *newNodo
					t.Utilizacion++
					t.capacidadT()
					break
				}
			}
		}
	}
}

func (t *TablaHash) capacidadT() {
	aux := float64(t.Capacidad) * 0.7
	if t.Utilizacion > int(aux) {
		auxPrev := t.Capacidad
		t.Capacidad = t.NewC()
		t.Utilizacion = 0
		t.Reinsertar(auxPrev)
	}
}

func (t *TablaHash) Reinsertar(auxPrev int) {
	auxT := t.Tabla
	t.Tabla = make(map[int]NodoHash)
	for i := 0; i < auxPrev; i++ {
		if user, existe := auxT[i]; existe {
			t.Insertar(user.Dato.Carnet, user.Dato.Name, user.Dato.Password, user.Dato.Cursos[0], user.Dato.Cursos[2], user.Dato.Cursos[2])
		}
	}
}

func (t *TablaHash) NewC() int {
	num := 1
	for num <= 10 {
		c := t.fibonacci(num)
		if c > t.Capacidad {
			return c
		}
		num++
	}
	return -1
}

func (t *TablaHash) fibonacci(x int) int {
	if x == 1 || x == 2 {
		return 1
	} else {
		return t.fibonacci(x-1) + t.fibonacci(x-2)
	}
}

func (t *TablaHash) reIndex(carnet int, cont int) int {
	newI := t.Calc_Index(carnet) + (cont * cont)
	return t.newIndex(newI)
}

func (t *TablaHash) newIndex(newIndex int) int {
	nPos := 0
	if newIndex < t.Capacidad {
		nPos = newIndex
	} else {
		nPos = newIndex - t.Capacidad
		nPos = t.newIndex(nPos)
	}

	return nPos
}

func (t *TablaHash) Calc_Index(carnet int) int {
	var num []int
	for {
		if carnet > 0 {
			digito := carnet % 10
			num = append([]int{digito}, num...)
			carnet = carnet / 10
		} else {
			break
		}
	}
	var nASCII []rune
	for _, numero := range num {
		valor := rune(numero + 48)
		nASCII = append(nASCII, valor)
	}
	final := 0
	for _, nASCII := range nASCII {
		final += int(nASCII)
	}
	index := final % t.Capacidad
	return index
}

func (t *TablaHash) Buscar(carnet string, pass string) *NodoHash {
	temp, err := strconv.Atoi(carnet)
	if err != nil {
		log.Fatal(err)
	}
	index := t.Calc_Index(temp)
	if index < t.Capacidad {
		if user, existe := t.Tabla[index]; existe {
			if user.Dato.Carnet == temp {
				if user.Dato.Password == pass {
					return &user
				}
			} else {
				contador := 1
				index = t.reIndex(temp, contador)
				for {
					if user, existe := t.Tabla[index]; existe {
						if user.Dato.Carnet == temp {
							if user.Dato.Password == pass {
								return &user
							} else {
								return nil
							}
						} else {
							contador++
							index = t.reIndex(temp, contador)
						}
					} else {
						return nil
					}
				}
			}
		}
	}
	return nil
}

func (t *TablaHash) BuscarC(temp int) *NodoHash {
	index := t.Calc_Index(temp)
	if index < t.Capacidad {
		if user, existe := t.Tabla[index]; existe {
			if user.Dato.Carnet == temp {
				return &user

			} else {
				contador := 1
				index = t.reIndex(temp, contador)
				for {
					if user, existe := t.Tabla[index]; existe {
						if user.Dato.Carnet == temp {
							return &user
						} else {
							contador++
							index = t.reIndex(temp, contador)
						}
					} else {
						return nil
					}
				}
			}
		}
	}
	return nil
}

func (t *TablaHash) Arreglo() []NodoHash {
	var array []NodoHash
	if t.Utilizacion > 0 {
		for i := 0; i < t.Capacidad; i++ {
			if user, existe := t.Tabla[i]; existe {
				array = append(array, user)
			}
		}
	}
	return array
}
