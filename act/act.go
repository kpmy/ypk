package act

import (
	"sync"
)

type Continue interface {
	Do(func(interface{}) (interface{}, error), ...func(error)) Continue
	Run(interface{})
}

type step struct {
	act func(interface{}) (interface{}, error)
	err func(error)
}

type cont struct {
	steps []step
}

func Seq() Continue {
	return &cont{}
}

func (c *cont) Do(fn func(interface{}) (interface{}, error), err ...func(error)) Continue {
	s := step{}
	s.act = fn
	if len(err) > 0 {
		s.err = err[0]
	}
	c.steps = append(c.steps, s)
	return c
}

func (c *cont) Run(in interface{}) {
	var next *cont
	var this *step
	if len(c.steps) > 0 {
		this = &c.steps[0]
	}
	if len(c.steps) > 1 {
		next = &cont{}
		next.steps = c.steps[1:]
	}
	if this != nil {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(this *step, next *cont) {
			if res, err := this.act(in); err == nil {
				if next != nil {
					next.Run(res)
				}
				wg.Done()
			} else if this.err != nil {
				this.err(err)
				wg.Done()
			} else {
				panic(err)
			}
		}(this, next)
		wg.Wait()
	}
}
