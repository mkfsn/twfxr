package twfxr_test

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mkfsn/twfxr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//go:embed testdata/ExchangeRate@202108290526.csv
var ExchangeRatePage string

type twfxrSuite struct {
	suite.Suite
}

func (suite *twfxrSuite) SetupSuite() {
	httpmock.Activate()
	httpmock.RegisterResponder(http.MethodGet, "https://rate.bot.com.tw/xrt/flcsv/0/day",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusOK, ExchangeRatePage)
			resp.Header.Add("Content-Disposition", ` attachment; filename="ExchangeRate@202108290526.csv"`)
			return resp, nil
		},
	)
}

func (suite *twfxrSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (suite *twfxrSuite) TestGetCurrencyExchangeRates() {
	currencies, metadata, err := twfxr.GetCurrencyExchangeRates(context.Background())
	suite.NoError(err)

	expectedMetadata := twfxr.Metadata{
		QuotedAt: time.Date(2021, 8, 29, 5, 26, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
	}

	suite.Equal(expectedMetadata, metadata)

	expectedCurrencies := map[twfxr.Currency]twfxr.CurrencyExchangeRate{
		twfxr.CurrencyUSD: {
			BuyingCash:            27.52000,
			BuyingSpot:            27.84500,
			BuyingForward10Days:   27.86500,
			BuyingForward30Days:   27.86500,
			BuyingForward60Days:   27.86000,
			BuyingForward90Days:   27.85500,
			BuyingForward120Days:  27.85000,
			BuyingForward150Days:  27.84100,
			BuyingForward180Days:  27.83400,
			SellingCash:           28.19000,
			SellingSpot:           27.99500,
			SellingForward10Days:  27.97100,
			SellingForward30Days:  27.97100,
			SellingForward60Days:  27.97100,
			SellingForward90Days:  27.97000,
			SellingForward120Days: 27.97000,
			SellingForward150Days: 27.96900,
			SellingForward180Days: 27.96700,
		},
		twfxr.CurrencyHKD: {
			BuyingCash:            3.43000,
			BuyingSpot:            3.55100,
			BuyingForward10Days:   3.55400,
			BuyingForward30Days:   3.55300,
			BuyingForward60Days:   3.55300,
			BuyingForward90Days:   3.55300,
			BuyingForward120Days:  3.55200,
			BuyingForward150Days:  3.55100,
			BuyingForward180Days:  3.55100,
			SellingCash:           3.63400,
			SellingSpot:           3.62100,
			SellingForward10Days:  3.61500,
			SellingForward30Days:  3.61600,
			SellingForward60Days:  3.61600,
			SellingForward90Days:  3.61600,
			SellingForward120Days: 3.61600,
			SellingForward150Days: 3.61700,
			SellingForward180Days: 3.61700,
		},
		twfxr.CurrencyGBP: {
			BuyingCash:            37.26000,
			BuyingSpot:            38.15500,
			BuyingForward10Days:   38.11600,
			BuyingForward30Days:   38.10800,
			BuyingForward60Days:   38.10900,
			BuyingForward90Days:   38.10800,
			BuyingForward120Days:  38.10100,
			BuyingForward150Days:  38.09500,
			BuyingForward180Days:  38.08800,
			SellingCash:           39.38000,
			SellingSpot:           38.78500,
			SellingForward10Days:  38.52600,
			SellingForward30Days:  38.53700,
			SellingForward60Days:  38.53800,
			SellingForward90Days:  38.53900,
			SellingForward120Days: 38.54500,
			SellingForward150Days: 38.55100,
			SellingForward180Days: 38.55700,
		},
		twfxr.CurrencyAUD: {
			BuyingCash:            20.03000,
			BuyingSpot:            20.24500,
			BuyingForward10Days:   20.14800,
			BuyingForward30Days:   20.14500,
			BuyingForward60Days:   20.14600,
			BuyingForward90Days:   20.14800,
			BuyingForward120Days:  20.14400,
			BuyingForward150Days:  20.14300,
			BuyingForward180Days:  20.14000,
			SellingCash:           20.81000,
			SellingSpot:           20.59000,
			SellingForward10Days:  20.35400,
			SellingForward30Days:  20.36100,
			SellingForward60Days:  20.36400,
			SellingForward90Days:  20.36500,
			SellingForward120Days: 20.36900,
			SellingForward150Days: 20.37700,
			SellingForward180Days: 20.37800,
		},
		twfxr.CurrencyCAD: {
			BuyingCash:            21.65000,
			BuyingSpot:            21.98000,
			BuyingForward10Days:   21.94700,
			BuyingForward30Days:   21.94100,
			BuyingForward60Days:   21.93900,
			BuyingForward90Days:   21.93700,
			BuyingForward120Days:  21.93200,
			BuyingForward150Days:  21.92600,
			BuyingForward180Days:  21.92100,
			SellingCash:           22.56000,
			SellingSpot:           22.31000,
			SellingForward10Days:  22.15300,
			SellingForward30Days:  22.15800,
			SellingForward60Days:  22.15700,
			SellingForward90Days:  22.15500,
			SellingForward120Days: 22.15700,
			SellingForward150Days: 22.15900,
			SellingForward180Days: 22.16100,
		},
		twfxr.CurrencySGD: {
			BuyingCash:            20.17000,
			BuyingSpot:            20.64000,
			BuyingForward10Days:   20.57700,
			BuyingForward30Days:   20.57000,
			BuyingForward60Days:   20.56800,
			BuyingForward90Days:   20.56700,
			BuyingForward120Days:  20.56000,
			BuyingForward150Days:  20.55400,
			BuyingForward180Days:  20.54800,
			SellingCash:           21.08000,
			SellingSpot:           20.86000,
			SellingForward10Days:  20.76200,
			SellingForward30Days:  20.76700,
			SellingForward60Days:  20.76600,
			SellingForward90Days:  20.76400,
			SellingForward120Days: 20.76500,
			SellingForward150Days: 20.76500,
			SellingForward180Days: 20.76600,
		},
		twfxr.CurrencyCHF: {
			BuyingCash:            29.82000,
			BuyingSpot:            30.43000,
			BuyingForward10Days:   30.32200,
			BuyingForward30Days:   30.32600,
			BuyingForward60Days:   30.34900,
			BuyingForward90Days:   30.37200,
			BuyingForward120Days:  30.38900,
			BuyingForward150Days:  30.40600,
			BuyingForward180Days:  30.42400,
			SellingCash:           31.02000,
			SellingSpot:           30.82000,
			SellingForward10Days:  30.58400,
			SellingForward30Days:  30.61200,
			SellingForward60Days:  30.63600,
			SellingForward90Days:  30.65800,
			SellingForward120Days: 30.68800,
			SellingForward150Days: 30.71900,
			SellingForward180Days: 30.74900,
		},
		twfxr.CurrencyJPY: {
			BuyingCash:            0.24490,
			BuyingSpot:            0.25190,
			BuyingForward10Days:   0.25160,
			BuyingForward30Days:   0.25160,
			BuyingForward60Days:   0.25160,
			BuyingForward90Days:   0.25170,
			BuyingForward120Days:  0.25170,
			BuyingForward150Days:  0.25180,
			BuyingForward180Days:  0.25180,
			SellingCash:           0.25770,
			SellingSpot:           0.25650,
			SellingForward10Days:  0.25570,
			SellingForward30Days:  0.25580,
			SellingForward60Days:  0.25580,
			SellingForward90Days:  0.25590,
			SellingForward120Days: 0.25600,
			SellingForward150Days: 0.25620,
			SellingForward180Days: 0.25630,
		},
		twfxr.CurrencyZAR: {
			BuyingCash:            0.00000,
			BuyingSpot:            1.85100,
			BuyingForward10Days:   1.83200,
			BuyingForward30Days:   1.82600,
			BuyingForward60Days:   1.81800,
			BuyingForward90Days:   1.81000,
			BuyingForward120Days:  1.80200,
			BuyingForward150Days:  1.79400,
			BuyingForward180Days:  1.78600,
			SellingCash:           0.00000,
			SellingSpot:           1.94100,
			SellingForward10Days:  1.91300,
			SellingForward30Days:  1.90900,
			SellingForward60Days:  1.90100,
			SellingForward90Days:  1.89400,
			SellingForward120Days: 1.88700,
			SellingForward150Days: 1.87900,
			SellingForward180Days: 1.87200,
		},
		twfxr.CurrencySEK: {
			BuyingCash:            2.85000,
			BuyingSpot:            3.18000,
			BuyingForward10Days:   3.16000,
			BuyingForward30Days:   3.15900,
			BuyingForward60Days:   3.15900,
			BuyingForward90Days:   3.16000,
			BuyingForward120Days:  3.16000,
			BuyingForward150Days:  3.16000,
			BuyingForward180Days:  3.16100,
			SellingCash:           3.37000,
			SellingSpot:           3.30000,
			SellingForward10Days:  3.26100,
			SellingForward30Days:  3.26300,
			SellingForward60Days:  3.26400,
			SellingForward90Days:  3.26400,
			SellingForward120Days: 3.26600,
			SellingForward150Days: 3.26700,
			SellingForward180Days: 3.26800,
		},
		twfxr.CurrencyNZD: {
			BuyingCash:            19.09000,
			BuyingSpot:            19.42000,
			BuyingForward10Days:   19.31700,
			BuyingForward30Days:   19.31000,
			BuyingForward60Days:   19.30500,
			BuyingForward90Days:   19.29900,
			BuyingForward120Days:  19.28600,
			BuyingForward150Days:  19.27200,
			BuyingForward180Days:  19.25900,
			SellingCash:           19.94000,
			SellingSpot:           19.72000,
			SellingForward10Days:  19.52200,
			SellingForward30Days:  19.52700,
			SellingForward60Days:  19.52200,
			SellingForward90Days:  19.51600,
			SellingForward120Days: 19.51000,
			SellingForward150Days: 19.50300,
			SellingForward180Days: 19.49700,
		},
		twfxr.CurrencyTHB: {
			BuyingCash:            0.73030,
			BuyingSpot:            0.83970,
			BuyingForward10Days:   0.00000,
			BuyingForward30Days:   0.00000,
			BuyingForward60Days:   0.00000,
			BuyingForward90Days:   0.00000,
			BuyingForward120Days:  0.00000,
			BuyingForward150Days:  0.00000,
			BuyingForward180Days:  0.00000,
			SellingCash:           0.92030,
			SellingSpot:           0.88570,
			SellingForward10Days:  0.00000,
			SellingForward30Days:  0.00000,
			SellingForward60Days:  0.00000,
			SellingForward90Days:  0.00000,
			SellingForward120Days: 0.00000,
			SellingForward150Days: 0.00000,
			SellingForward180Days: 0.00000,
		},
		twfxr.CurrencyPHP: {
			BuyingCash:            0.48640,
			BuyingSpot:            0.00000,
			BuyingForward10Days:   0.00000,
			BuyingForward30Days:   0.00000,
			BuyingForward60Days:   0.00000,
			BuyingForward90Days:   0.00000,
			BuyingForward120Days:  0.00000,
			BuyingForward150Days:  0.00000,
			BuyingForward180Days:  0.00000,
			SellingCash:           0.61940,
			SellingSpot:           0.00000,
			SellingForward10Days:  0.00000,
			SellingForward30Days:  0.00000,
			SellingForward60Days:  0.00000,
			SellingForward90Days:  0.00000,
			SellingForward120Days: 0.00000,
			SellingForward150Days: 0.00000,
			SellingForward180Days: 0.00000,
		},
		twfxr.CurrencyIDR: {
			BuyingCash:            0.00158,
			BuyingSpot:            0.00000,
			BuyingForward10Days:   0.00000,
			BuyingForward30Days:   0.00000,
			BuyingForward60Days:   0.00000,
			BuyingForward90Days:   0.00000,
			BuyingForward120Days:  0.00000,
			BuyingForward150Days:  0.00000,
			BuyingForward180Days:  0.00000,
			SellingCash:           0.00228,
			SellingSpot:           0.00000,
			SellingForward10Days:  0.00000,
			SellingForward30Days:  0.00000,
			SellingForward60Days:  0.00000,
			SellingForward90Days:  0.00000,
			SellingForward120Days: 0.00000,
			SellingForward150Days: 0.00000,
			SellingForward180Days: 0.00000,
		},
		twfxr.CurrencyEUR: {
			BuyingCash:            32.12000,
			BuyingSpot:            32.63500,
			BuyingForward10Days:   32.64100,
			BuyingForward30Days:   32.64400,
			BuyingForward60Days:   32.66000,
			BuyingForward90Days:   32.67700,
			BuyingForward120Days:  32.69100,
			BuyingForward150Days:  32.70500,
			BuyingForward180Days:  32.71800,
			SellingCash:           33.46000,
			SellingSpot:           33.23500,
			SellingForward10Days:  33.05200,
			SellingForward30Days:  33.07700,
			SellingForward60Days:  33.09700,
			SellingForward90Days:  33.11800,
			SellingForward120Days: 33.14800,
			SellingForward150Days: 33.17800,
			SellingForward180Days: 33.20800,
		},
		twfxr.CurrencyKRW: {
			BuyingCash:            0.02229,
			BuyingSpot:            0.00000,
			BuyingForward10Days:   0.00000,
			BuyingForward30Days:   0.00000,
			BuyingForward60Days:   0.00000,
			BuyingForward90Days:   0.00000,
			BuyingForward120Days:  0.00000,
			BuyingForward150Days:  0.00000,
			BuyingForward180Days:  0.00000,
			SellingCash:           0.02619,
			SellingSpot:           0.00000,
			SellingForward10Days:  0.00000,
			SellingForward30Days:  0.00000,
			SellingForward60Days:  0.00000,
			SellingForward90Days:  0.00000,
			SellingForward120Days: 0.00000,
			SellingForward150Days: 0.00000,
			SellingForward180Days: 0.00000,
		},
		twfxr.CurrencyVND: {
			BuyingCash:            0.00098,
			BuyingSpot:            0.00000,
			BuyingForward10Days:   0.00000,
			BuyingForward30Days:   0.00000,
			BuyingForward60Days:   0.00000,
			BuyingForward90Days:   0.00000,
			BuyingForward120Days:  0.00000,
			BuyingForward150Days:  0.00000,
			BuyingForward180Days:  0.00000,
			SellingCash:           0.00139,
			SellingSpot:           0.00000,
			SellingForward10Days:  0.00000,
			SellingForward30Days:  0.00000,
			SellingForward60Days:  0.00000,
			SellingForward90Days:  0.00000,
			SellingForward120Days: 0.00000,
			SellingForward150Days: 0.00000,
			SellingForward180Days: 0.00000,
		},
		twfxr.CurrencyMYR: {
			BuyingCash:            5.65200,
			BuyingSpot:            0.00000,
			BuyingForward10Days:   0.00000,
			BuyingForward30Days:   0.00000,
			BuyingForward60Days:   0.00000,
			BuyingForward90Days:   0.00000,
			BuyingForward120Days:  0.00000,
			BuyingForward150Days:  0.00000,
			BuyingForward180Days:  0.00000,
			SellingCash:           7.13200,
			SellingSpot:           0.00000,
			SellingForward10Days:  0.00000,
			SellingForward30Days:  0.00000,
			SellingForward60Days:  0.00000,
			SellingForward90Days:  0.00000,
			SellingForward120Days: 0.00000,
			SellingForward150Days: 0.00000,
			SellingForward180Days: 0.00000,
		},
		twfxr.CurrencyCNY: {
			BuyingCash:            4.22500,
			BuyingSpot:            4.29200,
			BuyingForward10Days:   4.28080,
			BuyingForward30Days:   4.27240,
			BuyingForward60Days:   4.26080,
			BuyingForward90Days:   4.24970,
			BuyingForward120Days:  4.23800,
			BuyingForward150Days:  4.22630,
			BuyingForward180Days:  4.21470,
			SellingCash:           4.38700,
			SellingSpot:           4.35200,
			SellingForward10Days:  4.33240,
			SellingForward30Days:  4.32720,
			SellingForward60Days:  4.31820,
			SellingForward90Days:  4.30950,
			SellingForward120Days: 4.30160,
			SellingForward150Days: 4.29370,
			SellingForward180Days: 4.28580,
		},
	}

	suite.Equal(expectedCurrencies, currencies)
}

func (suite *twfxrSuite) TestGetCurrencyExchangeRate() {

	type args struct {
		ctx      context.Context
		currency twfxr.Currency
	}

	type wants struct {
		exchangeRate twfxr.CurrencyExchangeRate
		metadata     twfxr.Metadata
		err          error
	}

	type test struct {
		args  args
		wants wants
	}

	testCases := map[string]test{
		"When getting JPY exchange rate, Then it should return the corresponding results": {
			args: args{
				ctx:      context.Background(),
				currency: twfxr.CurrencyJPY,
			},
			wants: wants{
				exchangeRate: twfxr.CurrencyExchangeRate{
					BuyingCash:            0.24490,
					BuyingSpot:            0.25190,
					BuyingForward10Days:   0.25160,
					BuyingForward30Days:   0.25160,
					BuyingForward60Days:   0.25160,
					BuyingForward90Days:   0.25170,
					BuyingForward120Days:  0.25170,
					BuyingForward150Days:  0.25180,
					BuyingForward180Days:  0.25180,
					SellingCash:           0.25770,
					SellingSpot:           0.25650,
					SellingForward10Days:  0.25570,
					SellingForward30Days:  0.25580,
					SellingForward60Days:  0.25580,
					SellingForward90Days:  0.25590,
					SellingForward120Days: 0.25600,
					SellingForward150Days: 0.25620,
					SellingForward180Days: 0.25630,
				},
				metadata: twfxr.Metadata{
					QuotedAt: time.Date(2021, 8, 29, 5, 26, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
				},
			},
		},

		"When getting exchange rate of invalid currency, Then it should return an error": {
			args: args{
				ctx:      context.Background(),
				currency: "invalid",
			},
			wants: wants{err: twfxr.ErrNotFound},
		},
	}

	for name, tc := range testCases {
		suite.T().Run(name, func(t *testing.T) {
			exchangeRate, metadata, err := twfxr.GetCurrencyExchangeRate(tc.args.ctx, tc.args.currency)
			if tc.wants.err != nil {
				assert.ErrorIs(t, err, tc.wants.err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.wants.metadata, metadata)
			assert.Equal(t, tc.wants.exchangeRate, exchangeRate)
		})
	}
}

func TesttwfxrSuite(t *testing.T) {
	suite.Run(t, new(twfxrSuite))
}
