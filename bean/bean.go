package bean

import (
	"errors"
	"fmt"
	"github.com/Shivam-nagar23/polling-plugin/bean"
	"io/ioutil"
	"os"
	"strings"
)

const (
	PermissionMode = 0644
	FileName       = "/output/results.json"
)

func ExtractOutRegistryId(hostUrl string) string {
	res := strings.Split(hostUrl, ".")
	return res[0]

}
func CheckFileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); err == nil {
		// exists
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		// not exists
		return false, nil
	} else {
		// Some other error
		return false, err
	}
}
func WriteToFile(file string, fileName string) error {
	fileExist, err := CheckFileExists(bean.FileName)
	if err != nil {
		fmt.Println("error in checking file exist or not", "err", err.Error())
		return err
	}
	if !fileExist {
		if err2 := os.Mkdir("/output", os.ModePerm); err2 != nil {
			return err2
		}
	}
	err = ioutil.WriteFile(fileName, []byte(file), PermissionMode)
	fmt.Println("fileName", fileName)
	fmt.Println("Permission Mode", PermissionMode)
	if err != nil {
		fmt.Println("error in writing results to json file", "err", err)
		return err
	}
	return nil
}

// /445808685819.dkr.ecr.us-east-2.amazonaws.com/devtron/html-ecr:cf50e450-125-588///Sample Image for reference
func GetHostUrlForEcr(registryId, region string) string {
	return fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", registryId, region)
}
