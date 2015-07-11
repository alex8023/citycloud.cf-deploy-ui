package entity

type BOSH struct{
	Name string
}

func NewBOSH(name string)(bosh BOSH){
	bosh.Name = name
	return
}