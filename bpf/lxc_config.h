/* Dummy configuration for test compilation */

#define LXC_MAC { .addr = { 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff } }
#define LXC_IP { .addr = { 0xbe, 0xef, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x1, 0x1, 0x65, 0x82, 0xbc } }
#define LXC_SECLABEL 0xfffff
#define LXC_SECLABEL_NB 0xfffff
#define LXC_POLICYMAP cilium_policy_foo
