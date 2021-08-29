package twfxr

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"time"
)

var (
	asiaTaipei = time.FixedZone("UTC+8", 8*60*60)
)

const (
	csvFileURL = "https://rate.bot.com.tw/xrt/flcsv/0/day"
)

var (
	ErrNotFound = errors.New("not found")
)

type Metadata struct {
	QuotedAt time.Time
}

func GetCurrencyExchangeRate(ctx context.Context, currency Currency) (CurrencyExchangeRate, Metadata, error) {
	results, metadata, err := GetCurrencyExchangeRates(ctx)
	if err != nil {
		return CurrencyExchangeRate{}, metadata, err
	}

	v, ok := results[currency]
	if !ok {
		return CurrencyExchangeRate{}, metadata, fmt.Errorf("no such currency: %w", ErrNotFound)
	}

	return v, metadata, nil
}

func GetCurrencyExchangeRates(ctx context.Context) (map[Currency]CurrencyExchangeRate, Metadata, error) {
	filename, data, err := getExchangeRateCSVFile(ctx)
	if err != nil {
		return nil, Metadata{}, err
	}

	metadata, err := parseMetadata(filename)
	if err != nil {
		return nil, metadata, err
	}

	currencies, err := parseCSV(bytes.NewReader(data))

	return currencies, metadata, err
}

func getExchangeRateCSVFile(ctx context.Context) (filename string, data []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, csvFileURL, nil)
	if err != nil {
		return "", nil, err
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	contentDisposition := resp.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return "", nil, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	return params["filename"], data, nil
}

func parseMetadata(filename string) (metadata Metadata, err error) {
	// filename: ExchangeRate@202108280526.csv
	metadata.QuotedAt, err = time.ParseInLocation("200601021504", filename[13:len(filename)-4], asiaTaipei)
	if err != nil {
		return metadata, fmt.Errorf("failed to parse filename: %w", err)
	}

	return metadata, nil
}

func parseCSV(reader io.Reader) (map[Currency]CurrencyExchangeRate, error) {
	r := csv.NewReader(reader)
	r.FieldsPerRecord = -1

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	currencies := make(map[Currency]CurrencyExchangeRate)

	for _, record := range records[1:] {
		data := map[string]json.RawMessage{
			"Currency": json.RawMessage(fmt.Sprintf("%q", record[0])),
			// Buying
			"Buying-Cash":            json.RawMessage(fmt.Sprintf("%s", record[2])),
			"Buying-Spot":            json.RawMessage(fmt.Sprintf("%s", record[3])),
			"Buying-Forward-10Days":  json.RawMessage(fmt.Sprintf("%s", record[4])),
			"Buying-Forward-30Days":  json.RawMessage(fmt.Sprintf("%s", record[5])),
			"Buying-Forward-60Days":  json.RawMessage(fmt.Sprintf("%s", record[6])),
			"Buying-Forward-90Days":  json.RawMessage(fmt.Sprintf("%s", record[7])),
			"Buying-Forward-120Days": json.RawMessage(fmt.Sprintf("%s", record[8])),
			"Buying-Forward-150Days": json.RawMessage(fmt.Sprintf("%s", record[9])),
			"Buying-Forward-180Days": json.RawMessage(fmt.Sprintf("%s", record[10])),
			// Selling
			"Selling-Cash":            json.RawMessage(fmt.Sprintf("%s", record[12])),
			"Selling-Spot":            json.RawMessage(fmt.Sprintf("%s", record[13])),
			"Selling-Forward-10Days":  json.RawMessage(fmt.Sprintf("%s", record[14])),
			"Selling-Forward-30Days":  json.RawMessage(fmt.Sprintf("%s", record[15])),
			"Selling-Forward-60Days":  json.RawMessage(fmt.Sprintf("%s", record[16])),
			"Selling-Forward-90Days":  json.RawMessage(fmt.Sprintf("%s", record[17])),
			"Selling-Forward-120Days": json.RawMessage(fmt.Sprintf("%s", record[18])),
			"Selling-Forward-150Days": json.RawMessage(fmt.Sprintf("%s", record[19])),
			"Selling-Forward-180Days": json.RawMessage(fmt.Sprintf("%s", record[20])),
		}

		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		var exchangeRate CurrencyExchangeRate

		if err := json.Unmarshal(b, &exchangeRate); err != nil {
			return nil, err
		}

		currencies[Currency(record[0])] = exchangeRate
	}

	return currencies, nil
}
