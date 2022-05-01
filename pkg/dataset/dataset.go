package dataset

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"yslinear/go-covid19/pkg/setting"
	"yslinear/go-covid19/service/fst_stock_service"

	"golang.org/x/text/width"
)

func Setup() {
	fileUrl := "https://data.nhi.gov.tw/resource/Nhi_Fst/Fstdata.csv"
	err := DownloadFile(GetDatasetSavePath()+"Fstdata.csv", fileUrl)
	if err != nil {
		panic(err)
	}
	log.Printf("[info] Downloaded %s", fileUrl)

	csvData, err := LoadCsv(GetDatasetSavePath() + "Fstdata.csv")
	if err != nil {
		panic(err)
	}

	for _, row := range csvData {
		code, _ := strconv.Atoi(row[0])
		lng, _ := strconv.ParseFloat(row[3], 32)
		lat, _ := strconv.ParseFloat(row[4], 32)

		loc, _ := time.LoadLocation("Asia/Taipei")
		t, _ := time.ParseInLocation("2006/01/02 15:04:05", row[8], loc)

		address := width.Narrow.String(row[2])
		r, _ := regexp.Compile(`(\W+?[縣市])`)
		city := r.FindString(address)
		address = strings.Replace(address, city, "", 1)
		r, _ = regexp.Compile(`(\W+?市區|\W+?鎮區|\W+?[鄉鎮市區])`)
		district := r.FindString(address)
		address = strings.Replace(address, district, "", 1)

		fstStockService := fst_stock_service.FstStock{
			Hospital: fst_stock_service.Hospital{
				Code:     code,
				Name:     row[1],
				City:     city,
				District: district,
				Address:  address,
				Lng:      float32(lng),
				Lat:      float32(lat),
				Phone:    row[5],
			},
			Brand:     row[6],
			Amount:    row[7],
			Remark:    row[9],
			CreatedAt: t,
		}
		if err := fstStockService.Add(); err != nil {
			panic(err)
		}
	}
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

func LoadCsv(filepath string) ([][]string, error) {
	f, err := os.Open(filepath)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
