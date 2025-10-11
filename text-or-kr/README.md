# Go

## 시작 (Start)

- 초기화 (initialize)

```bash
    # Create new go module `example/hello`
    # make project directory
    mkdir ~/example/hello # (windows) `md "%HOMEPATH%\example/hello"
    go mod init example/hello # project mod init

    cd ~/example/hello # move the project directory
    touch hello.go # create go demo file
    code . # open editor vscode
    go version # 설치 후 확인
    go env # 환경설정 보기
```

- 실행 (Run)

```bash

    go run .
    ./hello # excute hello

```

---

## Code Examples

## - 입력받고 출력하기

```go
package main

import ("bufio" "fmt" "log" "os")

// 별칭 설정 가능, aliases
var pl = fmt.Println
// Entry point -> main function
func main() {

    pl("What is your name?")

    reader := bufio.NewReader(os.Stdin)
    name, err := reader.ReadString('\n')

    if err == nil {
        pl("Hello", name)
    } else {
        log.Fatal(err)
    }
}
```

## 변수 선언 및 할당 형식

```go
package main

import ( "fmt" )

// 별칭 설정 가능, aliases
var pl = fmt.Println

/* Shift +Alt + A */

// main : 프로그램 시작점, 하나만 있어야 함
func main() {

 // Variables
 // 변수 선언 기본 형식 : `var name type`

 var nName string = "Viv" // 기본 형식
 var v1, v2 = 1.2, 3.4    // 순서대로 할당 선언
 var v3 = "hello"         // 타입을 생략
 v4 := 2.4                // var 를 제거하여 형식 유추 (float64)
 v4 = 5.4                 // 다시 할당

 pl(nName, v1, v2, v3, v4)
}
```

## Default Types

```go
package main

import (
 "fmt"
 "reflect"
)

var pl = fmt.Println

func main() {

 // Default type : int, float64, bool, string, rune
 pl(reflect.TypeOf(25)) // int
 pl(reflect.TypeOf(3.14)) // float64
 pl(reflect.TypeOf(true)) // bool
 pl(reflect.TypeOf("Hello, Wrold")) // string
 pl(reflect.TypeOf('😀')) // int32
}
```

## Convert

```go
package main

import (
 "fmt"
 "reflect"
 "strconv"
)

var pl = fmt.Println

func main() {

 // Convert
 cV1 := 1.5
 cV2 := int(cV1)
 pl(cV2)

 cV3 := "50000000"
 // string to int
 cV4, err := strconv.Atoi(cV3)
 pl(cV4, err, reflect.TypeOf(cV4))

 cV5 := 50000000
 // int to ascii string
 cV6 := strconv.Itoa(cV5)
 pl(cV6, reflect.TypeOf(cV6))

 cV7 := "3.14"
 // string to float64
 if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
  pl(cV8, err, reflect.TypeOf(cV8))
 }

 cV9 := fmt.Sprintf("%f", 3.141592)
 pl(cV9)
}
```
