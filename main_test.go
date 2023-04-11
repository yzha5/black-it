package black_it

import (
	"fmt"
	"testing"
)

var (
	str1 = "English"
	str2 = "简体中文"
	str3 = "Русский"
	str  = "Language:\n[BI:🇺🇸#1]" + str1 + "[:BI]\n[BI:🇨🇳#2]" + str2 + "[:BI]\n[BI:🇷🇺#3]" + str3 + "[:BI]"
)

func TestMark(t *testing.T) {
	fmt.Println(Mark(str1, 2))
	// [BI:█#2]English[:BI]
}

func TestMarkWithFill(t *testing.T) {
	fmt.Printf("Language: %s\n", MarkWithFill(str1, 3, "🇺🇸"))
	// Language: [BI:🇺🇸#3]English[:BI]
	fmt.Printf("Language: %s\n", MarkWithFill(str2, 3, "🇨🇳"))
	// Language: [BI:🇨🇳#3]简体中文[:BI]
	fmt.Printf("Language: %s\n", MarkWithFill(str3, 3, "🇷🇺"))
	// Language: [BI:🇷🇺#3]Русский[:BI]
}

func TestUnmark(t *testing.T) {
	fmt.Println(Unmark(str, 2))
	/*
		Language:
		English
		简体中文
		[BI:🇷🇺#3]Русский[:BI]
	*/
}

func TestGetBiObj(t *testing.T) {
	o1 := getBiObj("abc[BI:█#2]level2 text[:BI]")
	fmt.Println(o1) // <nil>
	o2 := getBiObj("[BI:█#2]level2 text[:BI]")
	fmt.Println(o2) // &{level2 text 2 █}
}
func TestBlackIt(t *testing.T) {
	fmt.Println(BlackIt(str, 2))
	/*
		Language:
		English
		简体中文
		🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺🇷🇺
	*/
	fmt.Println(BlackIt(str, 3))
	/*
		Language:
		English
		简体中文
		Русский
	*/
}
