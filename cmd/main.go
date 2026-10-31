package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	dummydata "dummy-data-jp"
)

func main() {
	var (
		count    int
		opt      dummydata.Options
		noHeader bool
		output   string
	)

	rootCmd := &cobra.Command{
		Use:          "dummy-data-jp",
		Short:        "ダミーの氏名・住所を TSV で生成する",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opt.MinAge < 0 || opt.MaxAge > 125 || opt.MinAge > opt.MaxAge {
				return fmt.Errorf("年齢レンジは 0 <= min-age <= max-age <= 125 で指定してください")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()
			if output != "" {
				f, err := os.Create(output)
				if err != nil {
					return err
				}
				defer f.Close()
				out = f
			}
			w := csv.NewWriter(out)
			w.Comma = '\t'
			if !noHeader {
				w.Write([]string{"氏", "名", "氏カナ", "名カナ", "年齢", "性別", "電話番号", "郵便番号", "都道府県", "住所1", "住所2"})
			}
			for range count {
				p, err := dummydata.GeneratePerson(opt)
				if err != nil {
					return err
				}
				w.Write([]string{p.FamilyName, p.GivenName, p.FamilyNameKana, p.GivenNameKana, strconv.Itoa(p.Age), p.Gender, p.Tel,
					p.Area, p.Prefecture, p.Address1, p.Address2})
			}
			w.Flush()
			return w.Error()
		},
	}

	flags := rootCmd.Flags()
	flags.IntVarP(&count, "number", "n", 10, "生成件数")
	flags.BoolVar(&opt.UseAlpha, "alpha", false, "氏名を英数字(alpha)から生成する")
	flags.BoolVar(&opt.UseExtendedMJ, "extended-mj", false, "氏・住所に IVS・CJK 拡張漢字など標準フォントで表示できない MJ 文字を含める")
	flags.BoolVar(&opt.IncludeUnanswered, "unanswered", false, "性別に「未回答」を含める")
	flags.IntVar(&opt.MinAge, "min-age", 0, "年齢の最小値")
	flags.IntVar(&opt.MaxAge, "max-age", 125, "年齢の最大値")
	flags.BoolVar(&noHeader, "no-header", false, "ヘッダ行を出力しない")
	flags.StringVarP(&output, "output", "o", "", "出力先ファイル (省略時は標準出力)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
