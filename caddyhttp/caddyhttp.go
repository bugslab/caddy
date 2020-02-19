// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package caddyhttp

import (
	// plug in the server
	_ "caddy/caddyhttp/httpserver"

	// plug in the standard directives
	// _ "caddy/caddyhttp/basicauth"
	// _ "caddy/caddyhttp/bind"
	// _ "caddy/caddyhttp/browse"
	_ "caddy/caddyhttp/errors"
	// _ "caddy/caddyhttp/expvar"
	// _ "caddy/caddyhttp/extensions"
	// _ "caddy/caddyhttp/fastcgi"
	_ "caddy/caddyhttp/gzip"
	_ "caddy/caddyhttp/header"

	// _ "caddy/caddyhttp/index"
	// _ "caddy/caddyhttp/internalsrv"
	_ "caddy/caddyhttp/limits"
	_ "caddy/caddyhttp/log"

	// _ "caddy/caddyhttp/markdown"
	// _ "caddy/caddyhttp/mime"
	// _ "caddy/caddyhttp/pprof"
	_ "caddy/caddyhttp/proxy"
	// _ "caddy/caddyhttp/push"
	// _ "caddy/caddyhttp/redirect"
	// _ "caddy/caddyhttp/requestid"
	// _ "caddy/caddyhttp/rewrite"
	// _ "caddy/caddyhttp/root"
	// _ "caddy/caddyhttp/status"
	// _ "caddy/caddyhttp/templates"
	// _ "caddy/caddyhttp/timeouts"
	_ "caddy/caddyhttp/websocket"
	// _ "caddy/onevent"
)

