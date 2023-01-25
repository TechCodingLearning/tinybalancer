package balancer

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestBounded_Add .
func TestBounded_Add(t *testing.T) {
	expect, err := Build(BoundedBalancer, []string{
		"192.168.1.1:1015",
		"192.168.1.1:1016",
		"192.168.1.1:1017",
		"192.168.1.1:1018",
	})
	assert.Equal(t, err, nil)

	cases := []struct {
		name   string
		lb     Balancer
		args   []string
		expect Balancer
	}{
		{
			"test-1",
			NewBounded(nil),
			[]string{
				"192.168.1.1:1015",
				"192.168.1.1:1016",
				"192.168.1.1:1017",
				"192.168.1.1:1018",
			},
			expect,
		},
	}
	//bounded := NewBounded(nil)
	//bounded.Add("192.168.1.1:1015")
	//bounded.Add("192.168.1.1:1016")
	//bounded.Add("192.168.1.1:1017")
	//bounded.Add("192.168.1.1:1018")
	//assert.Equal(t, true, reflect.DeepEqual(expect, bounded))
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, arg := range c.args {
				c.lb.Add(arg)
			}
			assert.Equal(t, true, reflect.DeepEqual(c.expect, c.lb))
		})
	}
}

// TestBounded_Remove .
func TestBounded_Remove(t *testing.T) {
	expect, err := Build(BoundedBalancer, []string{
		"192.168.1.1:1015",
		"192.168.1.1:1016",
	})
	assert.Equal(t, err, nil)
	//bounded := NewBounded([]string{
	//	"192.168.1.1:1015",
	//	"192.168.1.1:1016",
	//	"192.168.1.1:1017",
	//})
	//bounded.Remove("192.168.1.1:1017")
	//assert.Equal(t, true, reflect.DeepEqual(expect, bounded))

	cases := []struct {
		name   string
		lb     Balancer
		args   string
		expect Balancer
	}{
		{
			"test-1",
			NewBounded([]string{
				"192.168.1.1:1015",
				"192.168.1.1:1016",
				"192.168.1.1:1017",
			}),
			"192.168.1.1:1017",
			expect,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.lb.Remove(c.args)
			assert.Equal(t, true, reflect.DeepEqual(c.expect, c.lb))
		})
	}
}

// TestBounded_Balance .
func TestBounded_Balance(t *testing.T) {
	expect, _ := Build(BoundedBalancer, []string{
		"192.168.1.1:1015",
		"192.168.1.1:1016",
		"192.168.1.1:1017",
		"192.168.1.1:1018",
	})

	cases := []struct {
		name     string
		lb       Balancer
		inc_args []string
		des_args []string
		args     string
		expect   string
	}{
		{
			"test-1",
			expect,
			[]string{},
			[]string{},
			"172.166.2.44",
			"192.168.1.1:1017",
		},
	}

	//expect.Inc("192.168.1.1:1015")
	//expect.Inc("192.168.1.1:1015")
	//expect.Inc("NIL")
	//expect.Done("192.168.1.1:1015")
	//expect.Done("NIL")
	//host, _ := expect.Balance("172.166.2.44")
	//assert.Equal(t, "192.168.1.1:1017", host)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, arg := range c.inc_args {
				c.lb.Inc(arg)
			}

			for _, arg := range c.des_args {
				c.lb.Done(arg)
			}
			host, _ := c.lb.Balance(c.args)
			assert.Equal(t, "192.168.1.1:1017", host)
		})
	}
}
