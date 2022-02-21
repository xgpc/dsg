package cryptService

var RSAKey struct {
	Public  []byte
	Private []byte
}

func SetRsaKey() {
	publicKey := ReadRSAKey("public.pem")
	privateKey := ReadRSAKey("private.pem")
	RSAKey.Public = publicKey
	RSAKey.Private = privateKey
}
