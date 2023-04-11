# black-it
Blackout the text

Just like:
```
You can see me, You can't see me.

to:

You can see me, ████████████████.
```

## functions
### Mark(text string, level int) string
Mark text as blackable, but it doesn't actually blackout. It's useful for saving document or storing database.
```go
fmt.Println(Mark("English", 2))
// [BI:█#2]English[:BI]
```

### MarkWithFill(text string, level int, fill string) string
Custom fill string
```go
var (
str1 = "English"
str2 = "简体中文"
str3 = "Русский"
)
fmt.Printf("Language: %s\n", MarkWithFill(str1, 3, "🇺🇸"))
// Language: [BI:🇺🇸#3]English[:BI]
fmt.Printf("Language: %s\n", MarkWithFill(str2, 3, "🇨🇳"))
// Language: [BI:🇨🇳#3]简体中文[:BI]
fmt.Printf("Language: %s\n", MarkWithFill(str3, 3, "🇷🇺"))
// Language: [BI:🇷🇺#3]Русский[:BI]
```

### Unmark(text string, level int) string
Unmark text
```go
fmt.Println(Unmark("[BI:█#2]English[:BI]", 2))
// English
fmt.Println(Unmark("[BI:█#2]English[:BI]", 1))
// [BI:█#2]English[:BI]
```

### BlackIt(text string, level int) string
Blackout text
```go
var (
	str1 = "English"
	str2 = "简体中文"
	str3 = "Русский"
	str  = "Language:\n[BI:🇺🇸#1]" + str1 + "[:BI]\n[BI:🇨🇳#2]" + str2 + "[:BI]\n[BI:🇷🇺#3]" + str3 + "[:BI]"
)
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
```