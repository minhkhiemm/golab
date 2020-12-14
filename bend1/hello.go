package bend1

func ThreeWords() string {
	threeWords := [3]string{"foo", "bar", "baz"}
	return threeWords[1]
}

func TenWords() string {
	tenWords := [10]string{"foo", "bar", "bar", "foo", "bar", "foo", "foo", "foo", "foo", "bar"}
	return tenWords[8]
}
