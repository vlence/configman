package cached

import "github.com/vlence/configman"

type CachedStore struct {
        configman.Store
        configs map[string]*configman.Config
        settings map[*configman.Config][]*configman.Setting
}
