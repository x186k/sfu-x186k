module github.com/x186k/sfu1

go 1.16

require (
	github.com/caddyserver/certmagic v0.12.1-0.20210224184602-7550222c4a6a
	github.com/felixge/fgprof v0.9.1
	github.com/google/gopacket v1.1.19
	github.com/google/uuid v1.2.0 // indirect
	github.com/klauspost/cpuid v1.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/libdns/cloudflare v0.1.0
	github.com/libdns/duckdns v0.1.1
	github.com/libdns/libdns v0.2.0
	github.com/miekg/dns v1.1.40
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pion/dtls/v2 v2.0.8 // indirect
	github.com/pion/rtcp v1.2.6
	github.com/pion/rtp v1.6.2
	github.com/pion/sdp/v3 v3.0.4
	github.com/pion/webrtc/v3 v3.0.12
	github.com/pkg/profile v1.5.0
	github.com/stretchr/testify v1.7.0
	github.com/x186k/ddns5libdns v0.0.0-20210329015635-0863cd9aa3bb
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.0.0-20200513154647-78b527d18275 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

// replace github.com/libdns/duckdns => ../duckdns
//replace github.com/libdns/cloudflare => ../cloudflare

// replace github.com/x186k/dynamicdns => ../dynamicdns
//replace github.com/pion/webrtc/v3 => ../webrtc
//replace github.com/pion/webrtc/v3 v3.0.4 => github.com/cameronelliott/webrtc/v3 v3.0.5
