package command

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

type listJiraIDsOptions struct {
	isCommaSeparated bool
	label            string
	pullRequestID    int
	idPrefixes       []string
}

// NewListJiraIDsCommand returns definition of command list
func NewListJiraIDsCommand(cli *ManagerCli) *cobra.Command {
	opts := listJiraIDsOptions{}

	cmd := &cobra.Command{
		Use:   "list-jira-ids",
		Short: "List pull requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runListJiraIDs(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.pullRequestID, "id", "i", 0, "Pull request ID")
	flags.BoolVar(&opts.isCommaSeparated, "comma", false, "comma separated list")
	flags.StringArrayVarP(&opts.idPrefixes, "prefixes", "p", []string{}, "Comma separated list of JIRA ID prefixes")
	// flags.StringVar(&opts.label, "label", "", "Label to be applied to all issues")

	return cmd
}

func runListJiraIDs(cli *ManagerCli, opts listJiraIDsOptions) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	if len(opts.idPrefixes) == 0 {
		return fmt.Errorf("JIRA ID prefixes are not specified")
	}

	list, err := client.ListCommits(cred, repo, opts.pullRequestID)
	if err != nil {
		return err
	}

	var builder strings.Builder
	builder.WriteString(`(((`)
	for index, p := range opts.idPrefixes {
		if index == 0 {
			builder.WriteString(p)
		} else {
			builder.WriteString(fmt.Sprintf("|%s", p))
		}
	}
	builder.WriteString(`)-\d+))`)

	regex := regexp.MustCompile(builder.String())

	var ids []string
	for _, c := range list {
		matches := regex.FindAllString(c.Summary.Raw, -1)
		for _, m := range matches {
			ids = append(ids, m)
		}
	}
	distinctIDs := getDistinctNames(ids)
	sort.Strings(distinctIDs)

	for index, i := range distinctIDs {
		if opts.isCommaSeparated {
			if index == 0 {
				fmt.Printf("%s", i)
			} else {
				fmt.Printf(", %s", i)
			}
		} else {
			fmt.Println(i)
		}
	}
	return nil
}
