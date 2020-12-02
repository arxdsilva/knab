package domains

var operations = map[int]string{
	1: "compra a vista",
	2: "compra parcelada",
	3: "saque",
	4: "pagamento",
}

func IsOperation(op int) bool {
	_, ok := operations[op]
	return ok
}
