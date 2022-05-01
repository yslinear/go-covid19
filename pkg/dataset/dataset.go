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
	"yslinear/go-covid19/service/fst_service"
	"yslinear/go-covid19/service/hospital_service"

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
		hospitalCode, _ := strconv.Atoi(row[0])
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
		amount, _ := strconv.Atoi(row[7])

		hospitalService := hospital_service.Hospital{
			Code: hospitalCode,
		}
		if count, err := hospitalService.Count(); err != nil {
			panic(err)
		} else if count == 0 {
			hospitalService.Name = row[1]
			hospitalService.City = city
			hospitalService.District = district
			hospitalService.Address = address
			hospitalService.Lng = lng
			hospitalService.Lat = lat
			hospitalService.Phone = row[5]
			hospitalService.Add()
		}

		fstService := fst_service.Fst{
			HospitalCode: hospitalCode,
			CreatedAt:    t,
		}

		if count, err := fstService.Count(); err != nil {
			panic(err)
		} else if count == 0 {
			fstService.Brand = row[6]
			fstService.Amount = amount
			fstService.Remark = row[9]
			fstService.Add()
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
