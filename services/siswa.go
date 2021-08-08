package services

// <= 40 == c
// > 40 dan <= 70 == b
// > 70 == a

func score (nilai int) string{

	//untuk posisi kondisi c dan b bisa ditukar, namun untuk posisi kondisi a tidak bisa.
	//dikarenakan kondisi a itu selain dari kondisi c dan b.
	if nilai <= 40 {
		return "C"
	} else if nilai >40 && nilai <=70 {
		return "B"
	} else {
		return "A"
	}

}