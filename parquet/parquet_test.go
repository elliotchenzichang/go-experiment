package parquet

import (
	"github.com/bxcodec/faker/v3"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
	"log"
	"testing"
	"time"
)

type user struct {
	ID        string  `parquet:"name=id, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	FirstName string  `parquet:"name=firstname, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	LastName  string  `parquet:"name=lastname, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Email     string  `parquet:"name=email, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Phone     string  `parquet:"name=phone, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Blog      string  `parquet:"name=blog, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Username  string  `parquet:"name=username, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Score     float64 `parquet:"name=score, type=DOUBLE"`
	//won't be saved in the parquet file
	CreatedAt time.Time
}

const recordNumber = 10000

func TestWriteSomethingToParquet(t *testing.T) {
	var data []*user
	//create fake data
	for i := 0; i < recordNumber; i++ {
		u := &user{
			ID:        faker.UUIDDigit(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Phone:     faker.Phonenumber(),
			Blog:      faker.URL(),
			Username:  faker.Username(),
			Score:     float64(i),
			CreatedAt: time.Now(),
		}
		data = append(data, u)
	}
	err := generateParquet(data)
	if err != nil {
		log.Fatal(err)
	}
}

func TestParquetForReadingData(t *testing.T) {

}

func generateParquet(data []*user) error {
	log.Println("generating parquet file")
	fw, err := local.NewLocalFileWriter("output.parquet")
	if err != nil {
		return err
	}
	//parameters: writer, type of struct, size
	pw, err := writer.NewParquetWriter(fw, new(user), int64(len(data)))
	if err != nil {
		return err
	}
	//compression type
	pw.CompressionType = parquet.CompressionCodec_GZIP
	defer fw.Close()
	for _, d := range data {
		if err = pw.Write(d); err != nil {
			return err
		}
	}
	if err = pw.WriteStop(); err != nil {
		return err
	}
	return nil
}

func readParquet() ([]*user, error) {
	fr, err := local.NewLocalFileReader("output.parquet")
	if err != nil {
		return nil, err
	}

	pr, err := reader.NewParquetReader(fr, new(user), recordNumber)
	if err != nil {
		return nil, err
	}

	u := make([]*user, recordNumber)
	if err = pr.Read(&u); err != nil {
		return nil, err
	}
	pr.ReadStop()
	fr.Close()
	return u, nil
}

func readPartialParquet(pageSize, page int) ([]*user, error) {
	fr, err := local.NewLocalFileReader("output.parquet")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fr.Close()
	}()

	pr, err := reader.NewParquetReader(fr, new(user), int64(pageSize))
	if err != nil {
		return nil, err
	}
	defer pr.ReadStop()

	//num := pr.GetNumRows()

	pr.SkipRows(int64(pageSize * page))
	u := make([]*user, pageSize)
	if err = pr.Read(&u); err != nil {
		return nil, err
	}

	return u, nil
}
