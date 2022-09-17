package logclient

import (
	"github.com/fsnotify/fsnotify"
)

type LogClient struct {
	*fsnotify.Watcher
}

func NewLogClient() (*LogClient, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	logClient := &LogClient{
		Watcher: watcher,
	}

	return logClient, nil
}

// TODO support watch remote file
func (c *LogClient) Watch(f string) error {
	err := c.Add(f)
	if err != nil {
		return err
	}
	return nil
}

func (c *LogClient) UnWatch(f string) error {
	err := c.Remove(f)
	if err != nil {
		return err
	}
	return nil
}

func (c *LogClient) ReplaceWatch(oldf, newf string) error {
	err := c.Watch(newf)
	if err != nil {
		return err
	}
	err = c.UnWatch(oldf)
	return err
}
