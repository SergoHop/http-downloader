package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main(){
	fmt.Println("HTTP загрузчик")

	if len(os.Args) < 2 {
		fmt.Println("ссылка на url")
		return
	}
	url := os.Args[1]
	var output string

	if len(os.Args) > 2 {
		output = os.Args[2]
	}else{
		output = filepath.Base(url)
	}

	err := DownloadFile(url, output)

	if err != nil{
		fmt.Println("ошибка файла")
		os.Exit(1)
	}

	fmt.Println("файл загружен")

}

func DownloadFile(url, output string) error{
	out ,err := os.Create(output)
	if err != nil {
		return err

	}
	defer out.Close()

	resp, err := http.Get(url)

	if err != nil{
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return fmt.Errorf("status: %s", resp.Status)
	}

	_ , err = io.Copy(out, resp.Body)
	if err != nil{
		return err
	}
	return nil
	
}