//go:build !linux
// +build !linux

package unix

const (
	AF_INET                     = 0x2
	AF_INET6                    = 0xa
	AF_UNSPEC                   = 0x0
	NETLINK_ROUTE               = 0x0
	SizeofIfAddrmsg             = 0x8
	SizeofIfInfomsg             = 0x10
	SizeofNdMsg                 = 0xc
	SizeofRtMsg                 = 0xc
	SizeofRtNexthop             = 0x8
	RTM_NEWADDR                 = 0x14
	RTM_DELADDR                 = 0x15
	RTM_GETADDR                 = 0x16
	RTM_NEWLINK                 = 0x10
	RTM_DELLINK                 = 0x11
	RTM_GETLINK                 = 0x12
	RTM_SETLINK                 = 0x13
	RTM_NEWROUTE                = 0x18
	RTM_DELROUTE                = 0x19
	RTM_GETROUTE                = 0x1a
	RTM_NEWNEIGH                = 0x1c
	RTM_DELNEIGH                = 0x1d
	RTM_GETNEIGH                = 0x1e
	IFA_UNSPEC                  = 0x0
	IFA_ADDRESS                 = 0x1
	IFA_LOCAL                   = 0x2
	IFA_LABEL                   = 0x3
	IFA_BROADCAST               = 0x4
	IFA_ANYCAST                 = 0x5
	IFA_CACHEINFO               = 0x6
	IFA_MULTICAST               = 0x7
	IFA_FLAGS                   = 0x8
	IFF_UP                      = 0x1
	IFF_BROADCAST               = 0x2
	IFF_LOOPBACK                = 0x8
	IFF_POINTOPOINT             = 0x10
	IFF_MULTICAST               = 0x1000
	IFLA_UNSPEC                 = 0x0
	IFLA_ADDRESS                = 0x1
	IFLA_BROADCAST              = 0x2
	IFLA_IFNAME                 = 0x3
	IFLA_MTU                    = 0x4
	IFLA_LINK                   = 0x5
	IFLA_QDISC                  = 0x6
	IFLA_OPERSTATE              = 0x10
	IFLA_STATS                  = 0x7
	IFLA_STATS64                = 0x17
	IFLA_LINKINFO               = 0x12
	IFLA_MASTER                 = 0xa
	IFLA_INFO_KIND              = 0x1
	IFLA_INFO_SLAVE_KIND        = 0x4
	IFLA_INFO_DATA              = 0x2
	IFLA_INFO_SLAVE_DATA        = 0x5
	IFLA_XDP                    = 0x2b
	IFLA_XDP_FD                 = 0x1
	IFLA_XDP_ATTACHED           = 0x2
	IFLA_XDP_FLAGS              = 0x3
	IFLA_XDP_PROG_ID            = 0x4
	IFLA_XDP_EXPECTED_FD        = 0x8
	XDP_FLAGS_DRV_MODE          = 0x4
	XDP_FLAGS_SKB_MODE          = 0x2
	XDP_FLAGS_HW_MODE           = 0x8
	XDP_FLAGS_MODES             = 0xe
	XDP_FLAGS_MASK              = 0x1f
	XDP_FLAGS_REPLACE           = 0x10
	XDP_FLAGS_UPDATE_IF_NOEXIST = 0x1
	LWTUNNEL_ENCAP_MPLS         = 0x1
	MPLS_IPTUNNEL_DST           = 0x1
	MPLS_IPTUNNEL_TTL           = 0x2
	NDA_UNSPEC                  = 0x0
	NDA_DST                     = 0x1
	NDA_LLADDR                  = 0x2
	NDA_CACHEINFO               = 0x3
	NDA_IFINDEX                 = 0x8
	RTA_UNSPEC                  = 0x0
	RTA_DST                     = 0x1
	RTA_ENCAP                   = 0x16
	RTA_ENCAP_TYPE              = 0x15
	RTA_PREFSRC                 = 0x7
	RTA_GATEWAY                 = 0x5
	RTA_OIF                     = 0x4
	RTA_PRIORITY                = 0x6
	RTA_TABLE                   = 0xf
	RTA_MARK                    = 0x10
	RTA_EXPIRES                 = 0x17
	RTA_METRICS                 = 0x8
	RTA_MULTIPATH               = 0x9
	RTA_PREF                    = 0x14
	RTAX_ADVMSS                 = 0x8
	RTAX_FEATURES               = 0xc
	RTAX_INITCWND               = 0xb
	RTAX_INITRWND               = 0xe
	RTAX_MTU                    = 0x2
	NTF_PROXY                   = 0x8
	RTN_UNICAST                 = 0x1
	RT_TABLE_MAIN               = 0xfe
	RTPROT_BOOT                 = 0x3
	RTPROT_STATIC               = 0x4
	RT_SCOPE_UNIVERSE           = 0x0
	RT_SCOPE_HOST               = 0xfe
	RT_SCOPE_LINK               = 0xfd
)