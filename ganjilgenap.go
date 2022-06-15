package edspert_go1_lib

func CekGanjilGenap(num int) string {

	if num%2 == 0 {
		return "genap"
	}else{
		return "ganjil"
	}

}