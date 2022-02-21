package cryptService

//
//const SaltSize = 16
//
//func hashWithSalted(plain string) string {
//	buf := make([]byte, SaltSize, SaltSize+sha1.Size)
//	_, err := io.ReadFull(, buf)
//	if err != nil {
//		fmt.Println("random read failed ->", err)
//	}
//	fmt.Println(string(buf))
//
//	h := sha1.New()
//	h.Write(buf)
//	h.Write([]byte(plain))
//
//	return base64.URLEncoding.EncodeToString(h.Sum(buf))
//}
//
//func match(secret, plain string) bool {
//	data, _ := base64.URLEncoding.DecodeString(secret)
//	if len(data) != SaltSize+sha1.Size {
//		fmt.Println("wrong length of data")
//		return false
//	}
//	h := sha1.New()
//	h.Write(data[:SaltSize])
//	h.Write([]byte(plain))
//	return bytes.Equal(h.Sum(nil), data[SaltSize:])
//}
//
//func main() {
//	h := hashWithSalted("888888")
//	fmt.Println(len(h), h)
//
//	fmt.Println(match(h, "888888"))
//}
