package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "rss-feed",
	Short: "Parse RSS feed command",
	Long:  `The rss-feed is a command that supports parsing RSS, Atom and JSON feeds.`,
	Run:   Fetch,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return CheckRequiredFlags(cmd.Flags())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// latest flag
	rootCmd.Flags().BoolP("latest", "l", false, "Only print latest item.")

	// URL flag (required)
	rootCmd.Flags().StringP("url", "u", "", "(required)RSS feed URL to fetch data.")

	rootCmd.MarkFlagRequired("url")
}

func CheckRequiredFlags(flags *pflag.FlagSet) error {
	requiredError := false
	flagName := ""

	flags.VisitAll(func(flag *pflag.Flag) {
		requiredAnnotation := flag.Annotations[cobra.BashCompOneRequiredFlag]
		if len(requiredAnnotation) == 0 {
			return
		}

		flagRequired := requiredAnnotation[0] == "true"

		if flagRequired && !flag.Changed {
			requiredError = true
			flagName = flag.Name
		}
	})

	if requiredError {
		return errors.New("Required flag `" + flagName + "` has not been set.")
	}

	return nil
}

func Fetch(cmd *cobra.Command, args []string) {
	targetUrl, err := cmd.Flags().GetString("url")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fp := gofeed.NewParser()
	feeds, err := fp.ParseURL(targetUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	latest, err := cmd.Flags().GetBool("latest")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var resp []byte
	if latest {
		resp, err = encodeJson(feeds.Items[0])
	} else {
		resp, err = encodeJson(feeds.Items)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(resp))
}

func encodeJson(i interface{}) ([]byte, error) {
	var result []byte
	switch i.(type) {
	case *gofeed.Item, []*gofeed.Item:
		result, err := json.Marshal(i)
		if err != nil {
			return result, err
		}
		return result, nil
	default:
		return result, errors.New(fmt.Sprintf("%T\n", i) + "is not allowed")
	}
}
