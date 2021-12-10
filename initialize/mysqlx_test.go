package initialize

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestInitMysql(t *testing.T) {
	//copy配置文件到指定位置
	os.Mkdir("./conf", os.ModePerm)
	defer os.RemoveAll("./conf")
	copyFile("./conf/database.yaml", "../conf/database.yaml")

	InitMysql()
	defer MysqlDB.Close()
	first := &MysqlDB
	InitMysql()
	fmt.Println(first, &MysqlDB)
	if first != &MysqlDB {
		t.Log("MySql连接单例失败！")
		t.FailNow()
	}
}

func copyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("打开源文件错误，错误信息=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("打开目标文件错误，错误信息=%v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
