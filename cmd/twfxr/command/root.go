package command

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mkfsn/twfxr"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Persistent Flags
var (
	output string
)

var (
	rootCmd = &cobra.Command{
		Use:   "twfxr",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			results, _, err := twfxr.GetCurrencyExchangeRates(context.Background())
			if err != nil {
				log.Printf("error: %s\n", err)
				return
			}

			switch strings.ToLower(output) {
			case "":
				toTableRow := func(exchangeRate twfxr.CurrencyExchangeRate) []string {
					toValue := func(f float64) string {
						if f == 0 {
							return "-"
						}
						return fmt.Sprintf("%f", f)
					}

					return []string{
						exchangeRate.Currency,
						toValue(exchangeRate.BuyingCash),
						toValue(exchangeRate.BuyingSpot),
						toValue(exchangeRate.SellingCash),
						toValue(exchangeRate.SellingSpot),
					}
				}

				data := [][]string{
					toTableRow(results[twfxr.CurrencyUSD]),
					toTableRow(results[twfxr.CurrencyHKD]),
					toTableRow(results[twfxr.CurrencyGBP]),
					toTableRow(results[twfxr.CurrencyAUD]),
					toTableRow(results[twfxr.CurrencyCAD]),
					toTableRow(results[twfxr.CurrencySGD]),
					toTableRow(results[twfxr.CurrencyCHF]),
					toTableRow(results[twfxr.CurrencyJPY]),
					toTableRow(results[twfxr.CurrencyZAR]),
					toTableRow(results[twfxr.CurrencySEK]),
					toTableRow(results[twfxr.CurrencyNZD]),
					toTableRow(results[twfxr.CurrencyTHB]),
					toTableRow(results[twfxr.CurrencyPHP]),
					toTableRow(results[twfxr.CurrencyIDR]),
					toTableRow(results[twfxr.CurrencyEUR]),
					toTableRow(results[twfxr.CurrencyKRW]),
					toTableRow(results[twfxr.CurrencyVND]),
					toTableRow(results[twfxr.CurrencyMYR]),
					toTableRow(results[twfxr.CurrencyCNY]),
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"外幣", "本行買入:現金", "本行買入:即期", "本行賣出:現金", "本行賣出:即期"})
				table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
				table.SetCenterSeparator("|")
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				table.AppendBulk(data)
				table.Render()

			default:
				_, _ = fmt.Fprintf(cmd.OutOrStderr(), "unsupported output %s\n", output)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output")
}

func Execute() error {
	return rootCmd.Execute()
}
