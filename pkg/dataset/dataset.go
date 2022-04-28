package dataset

import (
	"io"
	"log"
	"net/http"
	"os"
	"yslinear/go-covid19/pkg/setting"
)

func Setup() {
	fileUrl := "https://data.nhi.gov.tw/resource/Nhi_Fst/Fstdata.csv"
	err := DownloadFile(GetDatasetSavePath()+"Fstdata.csv", fileUrl)
	if err != nil {
		panic(err)
	}
	log.Printf("[info] Downloaded %s", fileUrl)
}

func GetDatasetSavePath() string {
	return setting.AppSetting.DatasetSavePath
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
