package dataset

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"yslinear/go-covid19/pkg/setting"
	"yslinear/go-covid19/service/fst_stock_service"
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
		fstStockService := fst_stock_service.FstStock{
			Hospital: fst_stock_service.Hospital{
				Code:    code,
				Name:    row[1],
				Address: row[2],
				Lng:     float32(lng),
				Lat:     float32(lat),
				Phone:   row[5],
			},
			Brand:  row[6],
			Amount: row[7],
			Remark: row[9],
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
