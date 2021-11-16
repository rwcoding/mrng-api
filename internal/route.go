package internal

import (
	_ "github.com/rwcoding/mrng/internal/gw"
	_ "github.com/rwcoding/mrng/internal/node"
	_ "github.com/rwcoding/mrng/internal/service"

	_ "github.com/rwcoding/mrng/internal/config/env"
	_ "github.com/rwcoding/mrng/internal/config/kv"
	_ "github.com/rwcoding/mrng/internal/config/log"
	_ "github.com/rwcoding/mrng/internal/config/main1"
	_ "github.com/rwcoding/mrng/internal/config/project"
	_ "github.com/rwcoding/mrng/internal/config/sync"
	_ "github.com/rwcoding/mrng/internal/config/white"
)
