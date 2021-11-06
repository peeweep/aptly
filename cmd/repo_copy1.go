package cmd

import (
	"fmt"
	"log"

	"github.com/aptly-dev/aptly/deb"
	"github.com/aptly-dev/aptly/query"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

func aptlyRepoCopy1CheckArgs(cmd *commander.Command, args []string) error {

	log.Println("len(args): ", len(args))

	if len(args) >= 4 {
		log.Println("args[]: ", args)
		log.Println("args[0]: ", args[0])
		log.Println("args[1]: ", args[1])
		log.Println("args[2]: ", args[2])
		log.Println("args[3]: ", args[3])
	}

	if !(len(args) == 4 && args[1] == "from" && args[2] == "package") { // nolint: goconst
		cmd.Usage()
		return commander.ErrCommandError
	} else {
		log.Println("pass")
	}

	return nil
}

func aptlyRepoCopy1Search(cmd *commander.Command, args []string) error {
	var (
		err error
		q   deb.PackageQuery
	)

	aptlyRepoCopy1CheckArgs(cmd, args)

	if len(args) == 4 {
		q, err = query.Parse(args[3])
		if err != nil {
			return fmt.Errorf("unable to search: %s", err)
		}
	} else {
		q = &deb.MatchAllQuery{}
	}

	log.Println("&deb.MatchAllQuery: ", &deb.MatchAllQuery{})
	log.Println("q: ", q)
	result := q.Query(context.CollectionFactory().PackageCollection())

	log.Println(result.Strings())
	log.Println(result.Len())

	if result.Len() == 0 {
		return fmt.Errorf("no results")
	}

	return err
}

func aptlyRepoCopy1Copy(cmd *commander.Command, args []string) error {
	var (
		err error
		q   deb.PackageQuery
	)
	aptlyRepoCopy1CheckArgs(cmd, args)

	if len(args) == 4 {
		q, err = query.Parse(args[3])
		if err != nil {
			return fmt.Errorf("unable to search: %s", err)
		}
	} else {
		q = &deb.MatchAllQuery{}
	}

	log.Println("&deb.MatchAllQuery: ", &deb.MatchAllQuery{})
	log.Println("q: ", q)
	result := q.Query(context.CollectionFactory().PackageCollection())

	log.Println(result.Strings())
	return err
}

func makeCmdRepoCopy1() *commander.Command {
	cmd := &commander.Command{
		Run:       aptlyRepoCopy1Copy,
		UsageLine: "copy1 <dst-name> from package <package-query>",
		Short:     "copy packages from packages",
		Long: `
Copy packages from local package cache

Example:

  $ aptly repo copy1 aptly-repo from package 'myapp (=0.1.12)'
  $ aptly repo copy aptly-repo from repo aptly-repo2 'myapp (=0.1.12)'
`,
		Flag: *flag.NewFlagSet("aptly-repo-copy1", flag.ExitOnError),
	}

	return cmd
}
