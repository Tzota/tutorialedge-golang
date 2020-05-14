# tutorialedge-golang

Lessons from and based on https://tutorialedge.net/golang/

# Я познаю мир/go
## Паттерн для случая, когда не можешь расширить строку методом
```go
broken := "G# r#cks!"
replacer := strings.NewReplacer("#",
fixed := replacer.Replace(broken)
fmt.Println(fixed)
```

## Чтение с клавиатуры
```go
fmt.Print("Enter a grade: ")
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
fmt.Println(input)
```

## Конвертирование строки в...
```go
strconv.ParseFloat(input, 64)
```
## Чтение из файла
```go
file, err := os.Open("data.txt")
if err != nil {
    log.Fatal(err)
}
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
err = file.Close()
if err != nil {
    log.Fatal(err)
}
if scanner.Err() != nil {
    log.Fatal(scanner.Err())
}	
```

## Сеттеры
```go
func (d *Date) SetMonth(month int) error {
    if month < 1 || month > 12 {
        return errors.New("invalid month")
    }
    d.Month = month
    return nil
}
```
