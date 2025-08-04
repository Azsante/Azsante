package main

import (
	"fmt"

	"github.com/elastic/go-grok"
)

func main()	{
	g :=grok.New()
	patternDefinitions := map[string]string{
		"LOG_PREFIX": `%{SYSLOGBASE2}` ,
	}
	g.AddPatterns(patternDefinitions)
	err := g.Compile("%{LOG_PREFIX}", true)
	if err != nil {
		panic(err)
	}
	res, err := g.ParseString("Mar  6 06:31:42 ip-172-31-35-28 sshd[2424]: Failed password for backup from 65.2.161.68 port 34856 ssh2")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
