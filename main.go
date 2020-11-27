package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/kustomize"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	envLong = templates.LongDesc(i18n.T(`
Print a set of API resources generated from instructions in a kustomization.yaml file.
The argument must be the path to the directory containing
the file, or a git repository
URL with a path suffix specifying same with respect to the
repository root.
  kubectl env somedir
	`))

	envExample = templates.Examples(i18n.T(`
# Use the current working directory
  kubectl env .
# Use some shared configuration directory
  kubectl env /home/configuration/production
# Use a URL
  kubectl env github.com/kubernetes-sigs/kustomize.git/examples/helloWorld?ref=v1.0.6
`))
)

func main() {
	var b bytes.Buffer
	out := io.Writer(&b)

	in := os.Stdin

	errOut := os.Stderr

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: errOut}
	command := kustomize.NewCmdKustomize(ioStreams)

	command.Use = "env <dir>"
	command.Long = envLong
	command.Example = envExample

	command.Flags().StringP("left-delimiter", "", "{{", "sets the left delimiters used in template parsing")
	command.Flags().StringP("right-delimiter", "", "}}", "sets the right delimiters used in template parsing")

	if err := command.Execute(); err != nil {
		fmt.Fprintf(errOut, "ERR: %v\n", err)
		os.Exit(1)
	}

	leftDelimiter, err := command.Flags().GetString("left-delimiter")
	if err != nil {
		fmt.Fprintf(errOut, "ERR: %v\n", err)
		os.Exit(1)
	} else if leftDelimiter == "" {
		fmt.Fprintf(errOut, "ERR: missing left delimiter\n")
		os.Exit(1)
	}

	rightDelimiter, err := command.Flags().GetString("right-delimiter")
	if err != nil {
		fmt.Fprintf(errOut, "ERR: %v\n", err)
		os.Exit(1)
	} else if rightDelimiter == "" {
		fmt.Fprintf(errOut, "ERR: missing right delimiter\n")
		os.Exit(1)
	}

	t, err := template.New("kustomize").Delims(leftDelimiter, rightDelimiter).Option("missingkey=error").Parse(b.String())
	if err != nil {
		fmt.Fprintf(errOut, "ERR: %v\n", err)
		os.Exit(1)
	}

	if err := t.Execute(os.Stdout, envMap()); err != nil {
		fmt.Fprintf(errOut, "\nERR: %v\n", err)
		os.Exit(1)
	}
}

func envMap() map[string]string {
	m := make(map[string]string)

	for _, v := range os.Environ() {
		split_v := strings.Split(v, "=")
		m[split_v[0]] = split_v[1]
	}

	return m
}
