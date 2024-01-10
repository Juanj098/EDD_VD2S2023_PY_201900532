package AMerkle

type NodoMerkle struct {
	Left   *NodoMerkle
	Right  *NodoMerkle
	B_info *NBloque
	Valor  string
}

type InfoBloque struct {
	Date  string
	State string
	Name  string
	Tutor int
}

type NBloque struct {
	Sig   *NBloque
	Ant   *NBloque
	Value *InfoBloque
}
