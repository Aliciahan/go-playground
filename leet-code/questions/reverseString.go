package questions

func ReverseStringRecursive(str string) string{

	len_str := len(str)
	if len_str == 1 {
		return str
	} else {
		return ReverseStringRecursive(str[1:])+string(str[0])
	}

}

func ReverseStringBytes(s []byte) []byte  {
	str_len := len(s)
	if str_len ==1 {
		return s
	}else{
		return append(ReverseStringBytes(s[1:]), s[0])
	}

}
