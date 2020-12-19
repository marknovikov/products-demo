package products

import "fmt"

type optsHolder struct {
	paging  *Paging
	sorting *Sorting
}

type option func(opts *optsHolder)

type optsMethods struct{}

func Options() optsMethods {
	return optsMethods{}
}

func (so optsMethods) WithPaging(p Paging) option {
	return func(opts *optsHolder) {
		opts.paging = &p
	}
}

func (so optsMethods) WithSorting(s Sorting) (option, error) {
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("WithSorting: %s", err)
	}
	return func(opts *optsHolder) {
		opts.sorting = &s
	}, nil
}

func applyOptions(opts []option) *optsHolder {
	var hder optsHolder
	for _, o := range opts {
		o(&hder)
	}
	return &hder
}
