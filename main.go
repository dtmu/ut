package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "ut"
	app.Usage = "This is cli tool to convert unix time to utc time or vice versa."
	app.Version = "0.0.1"
	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) error {
	args := c.Args().Slice()
	if len(args) == 0 || strings.Trim(c.Args().Get(0), " ") == "" {
		fmt.Println("Please check usage of this command by executing 'ut -h'")
		return nil
	}
	arg := strings.Join(args, " ")
	if strings.Contains(arg, "/") || strings.Contains(arg, "-") {
		t, e := time.Parse("2006-01-02 15:04:05.000", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006/01/02 15:04:05.000", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006-01-02 15:04:05", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006/01/02 15:04:05", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006-01-02 15:04", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006/01/02 15:04", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006-01-02 15", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006/01/02 15", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006-01-02", arg)
		if e == nil {
			fmt.Println(t.UnixNano() / int64(time.Millisecond))
			return nil
		}
		t, e = time.Parse("2006/01/02", arg)
		if e != nil {
			fmt.Println(arg, "cannot be convert to unix time.")
			return e
		}
		fmt.Println(t.UnixNano() / int64(time.Millisecond))
		return nil
	}

	// unix time to UTC time.
	ut, e := strconv.Atoi(arg)
	if e != nil {
		fmt.Println(arg, "cannot be convert to UTC time.")
		return e
	}
	fmt.Println(time.Unix(0, int64(ut)*int64(time.Millisecond)).In(time.UTC))
	return nil
}
