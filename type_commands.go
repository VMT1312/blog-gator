package main

import "errors"

type commands struct {
	callback map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil || s.cfg == nil {
		return errors.New("state or configuration not initialized")
	}

	fun, exists := c.callback[cmd.name]
	if !exists {
		return errors.New("command not found: " + cmd.name)
	}

	return fun(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.callback[name] = f
}
