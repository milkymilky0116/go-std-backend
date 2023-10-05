package cli

import (
	"flag"
	"io"
)

type Config struct {
	IsTemplateMode bool
	User           string
	Password       string
	Dbname         string
	Port           int
}

func ParseArgs(w io.Writer, args []string) (Config, error) {
	c := Config{}
	fs := flag.NewFlagSet("Gist Application", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.BoolVar(&c.IsTemplateMode, "tmp", false, "Whether running application on Template Mode")
	fs.StringVar(&c.User, "user", "root", "Database Username")
	fs.StringVar(&c.Password, "password", "", "Database User Password")
	fs.StringVar(&c.Dbname, "dbname", "", "Database Name")
	fs.IntVar(&c.Port, "port", 3306, "MySQL port")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	return c, nil
}
