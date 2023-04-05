# Go

## 다운로드 및 설치 ( Download & Installation )

1. [Go](https://go.dev/doc/install)
2. [Git](https://git-scm.com/),  [Github](https://github.com/)
3. [Visual Studio Code](https://code.visualstudio.com/)

## 시작 (Start)

- 초기화 (initialize)

```bash
    # Create new go module `example/hello`
    # make project directory
    mkdir ~/example/hello # (windows) `md "%HOMEPATH%\example/hello"
    go mod init example/hello # project mod init
    
    cd ~/example/hello # move the project directory
    touch hello.go # go file
    code . # open editor vscode
    
    go version # 설치 후 확인
    
    go env # 환경설정 보기
```

- Paste the follwing code into your `hello.go` file and save the file.

```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
```

- 실행 (Run)

```bash
    
    go run . 
    # Hello, World!
    
    go build # created excute file
    
    ./hello # excute hello
    
```
