package beegopro

import (
	"errors"
	"fmt"
	"github.com/beego/bee/internal/pkg/utils"
	beeLogger "github.com/beego/bee/logger"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// write to file
func (c *RenderFile) write(filename string, buf string) (err error) {
	if utils.IsExist(filename) && !isNeedOverwrite(filename) {
		return
	}

	filePath := path.Dir(filename)
	err = createPath(filePath)
	if err != nil {
		err = errors.New("write create path " + err.Error())
		return
	}

	filePathBak := filePath + "/bak"
	err = createPath(filePathBak)
	if err != nil {
		err = errors.New("write create path bak " + err.Error())
		return
	}

	name := path.Base(filename)

	if utils.IsExist(filename) {
		bakName := fmt.Sprintf("%s/%s.%s.bak", filePathBak, name, time.Now().Format("2006.01.02.15.04.05"))
		beeLogger.Log.Infof("bak file '%s'", bakName)
		if err := os.Rename(filename, bakName); err != nil {
			err = errors.New("file is bak error, path is " + bakName)
			return err
		}
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		err = errors.New("write create file " + err.Error())
		return
	}

	output := []byte(buf)

	if c.Option.EnableFormat && filepath.Ext(filename) == ".go" {
		// format code
		var bts []byte
		bts, err = format.Source([]byte(buf))
		if err != nil {
			err = errors.New("format buf error " + err.Error())
			return
		}
		output = bts
	}

	err = ioutil.WriteFile(filename, output, 0644)
	if err != nil {
		err = errors.New("write write file " + err.Error())
		return
	}
	return
}

func isNeedOverwrite(fileName string) (flag bool) {
	seg := "//"
	ext := filepath.Ext(fileName)
	if ext == ".sql" {
		seg = "--"
	}

	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	overwrite := ""
	var contentByte []byte
	contentByte, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	for _, s := range strings.Split(string(contentByte), "\n") {
		s = strings.TrimSpace(strings.TrimPrefix(s, seg))
		if strings.HasPrefix(s, "@BeeOverwrite") {
			overwrite = strings.TrimSpace(s[len("@BeeOverwrite"):])
		}
	}
	if strings.ToLower(overwrite) == "yes" {
		flag = true
		return
	}
	return
}

// createPath ??????os.MkdirAll?????????????????????
func createPath(filePath string) error {
	if !utils.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func getPackagePath() (packagePath string) {
	f, err := os.Open("go.mod")
	if err != nil {
		return
	}
	defer f.Close()
	var contentByte []byte
	contentByte, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	for _, s := range strings.Split(string(contentByte), "\n") {
		packagePath = strings.TrimSpace(strings.TrimPrefix(s, "module"))
		return
	}
	return
}

func getModelType(orm string) (inputType, goType, mysqlType, tag string) {
	kv := strings.SplitN(orm, ",", 2)
	inputType = kv[0]
	switch inputType {
	case "string":
		goType = "string"
		tag = "size(255)"
		// todo use orm data
		mysqlType = "varchar(255) NOT NULL"
	case "text":
		goType = "string"
		tag = "type(longtext)"
		mysqlType = "longtext  NOT NULL"
	case "auto":
		goType = "int"
		tag = "auto"
		mysqlType = "int(11) NOT NULL AUTO_INCREMENT"
	case "pk":
		goType = "int"
		tag = "pk"
		mysqlType = "int(11) NOT NULL"
	case "datetime":
		goType = "time.Time"
		tag = "type(datetime)"
		mysqlType = "datetime NOT NULL"
	case "int", "int8", "int16", "int32", "int64":
		fallthrough
	case "uint", "uint8", "uint16", "uint32", "uint64":
		goType = inputType
		tag = ""
		mysqlType = "int(11) DEFAULT NULL"
	case "bool":
		goType = inputType
		tag = ""
		mysqlType = "int(11) DEFAULT NULL"
	case "float32", "float64":
		goType = inputType
		tag = ""
		mysqlType = "float NOT NULL"
	case "float":
		goType = "float64"
		tag = ""
		mysqlType = "float NOT NULL"
	default:
		beeLogger.Log.Fatalf("not support type: %s", inputType)
	}
	// user set orm tag
	if len(kv) == 2 {
		tag = kv[1]
	}
	return
}
