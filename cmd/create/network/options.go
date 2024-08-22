package network

import (
	opts "github.com/openshift/rosa/pkg/options/network"
	"github.com/openshift/rosa/pkg/reporter"
)

type Options struct {
	reporter *reporter.Object

	args *opts.NetworkUserOptions
}

func NewNetworkUserOptions() *opts.NetworkUserOptions {
	return &opts.NetworkUserOptions{
		Params: []string{},
	}
}

func NewNetworkOptions() *Options {
	return &Options{
		reporter: reporter.CreateReporter(),
	}
}

func (m *Options) Network() *opts.NetworkUserOptions {
	return m.args
}
