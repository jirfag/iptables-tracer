package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testPolicyNoTrace = []string{
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*mangle",
	":PREROUTING ACCEPT [6644470:1039872350]",
	":INPUT ACCEPT [6644470:1039872350]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [6093378:960619800]",
	":POSTROUTING ACCEPT [6093378:960619800]",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*raw",
	":PREROUTING ACCEPT [2502678:375978678]",
	":OUTPUT ACCEPT [2325323:325134634]",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*filter",
	":INPUT ACCEPT [38659:3282330]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [35452:1942057]",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*nat",
	":PREROUTING ACCEPT [1292886:69592056]",
	":POSTROUTING ACCEPT [71297:4698134]",
	":OUTPUT ACCEPT [71295:4698054]",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018"}

var testPolicyTraceID3 = []string{
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*mangle",
	":PREROUTING ACCEPT [6644470:1039872350]",
	":INPUT ACCEPT [6644470:1039872350]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [6093378:960619800]",
	":POSTROUTING ACCEPT [6093378:960619800]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:1\" --nflog-group 22",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:2\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:3\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:4\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:5\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*raw",
	":PREROUTING ACCEPT [2502678:375978678]",
	":OUTPUT ACCEPT [2325323:325134634]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:6\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:7\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*filter",
	":INPUT ACCEPT [38659:3282330]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [35452:1942057]",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:8\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:9\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:10\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*nat",
	":PREROUTING ACCEPT [1292886:69592056]",
	":POSTROUTING ACCEPT [71297:4698134]",
	":OUTPUT ACCEPT [71295:4698054]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:11\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:12\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:13\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018"}
var testPolicyTraceID3and4 = []string{
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*mangle",
	":PREROUTING ACCEPT [6644470:1039872350]",
	":INPUT ACCEPT [6644470:1039872350]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [6093378:960619800]",
	":POSTROUTING ACCEPT [6093378:960619800]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:1\" --nflog-group 22",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:2\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:3\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:4\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:5\" --nflog-group 22",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:1\" --nflog-group 22",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:2\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:3\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:4\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:5\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*raw",
	":PREROUTING ACCEPT [2502678:375978678]",
	":OUTPUT ACCEPT [2325323:325134634]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:6\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:7\" --nflog-group 22",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:6\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:7\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*filter",
	":INPUT ACCEPT [38659:3282330]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [35452:1942057]",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:8\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:9\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:10\" --nflog-group 22",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:8\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:9\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:10\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*nat",
	":PREROUTING ACCEPT [1292886:69592056]",
	":POSTROUTING ACCEPT [71297:4698134]",
	":OUTPUT ACCEPT [71295:4698054]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:11\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:12\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:4:13\" --nflog-group 22",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:11\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:12\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:3:13\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018"}

var testPolicyTraceID5withMark = []string{
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*mangle",
	":PREROUTING ACCEPT [6644470:1039872350]",
	":INPUT ACCEPT [6644470:1039872350]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [6093378:960619800]",
	":POSTROUTING ACCEPT [6093378:960619800]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:1\" --nflog-group 22",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:2\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:3\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:4\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:5\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*raw",
	":PREROUTING ACCEPT [2502678:375978678]",
	":OUTPUT ACCEPT [2325323:325134634]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:6\" --nflog-group 22",
	"-I PREROUTING -s 1.2.3.4 -m comment --comment \"iptr:5:limit\" -m limit --limit 60/minute --limit-burst 1 -j MARK --set-xmark 0x124/0x124",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:7\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*filter",
	":INPUT ACCEPT [38659:3282330]",
	":FORWARD ACCEPT [0:0]",
	":OUTPUT ACCEPT [35452:1942057]",
	"-I INPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:8\" --nflog-group 22",
	"-I FORWARD -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:9\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:10\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018",
	"# Generated by iptables-save v1.6.0 on Sat Nov 24 11:59:50 2018",
	"*nat",
	":PREROUTING ACCEPT [1292886:69592056]",
	":POSTROUTING ACCEPT [71297:4698134]",
	":OUTPUT ACCEPT [71295:4698054]",
	"-I PREROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:11\" --nflog-group 22",
	"-I POSTROUTING -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:12\" --nflog-group 22",
	"-I OUTPUT -s 1.2.3.4  -j NFLOG --nflog-prefix \"iptr:5:13\" --nflog-group 22",
	"COMMIT",
	"# Completed on Sat Nov 24 11:59:50 2018"}

func TestClearIptablesPolicy(t *testing.T) {
	policies := []struct {
		policy    []string
		cleanupID int
		expected  []string
	}{
		{nil, 0, nil},
		{[]string{"foobar"}, 0, []string{"foobar"}},
		{testPolicyNoTrace, 0, testPolicyNoTrace},
		{testPolicyNoTrace, 3, testPolicyNoTrace},
		{testPolicyTraceID3, 3, testPolicyNoTrace},
		{testPolicyTraceID3, 0, testPolicyNoTrace},
		{testPolicyTraceID3, 4, testPolicyTraceID3},
		{testPolicyTraceID3and4, 4, testPolicyTraceID3},
		{testPolicyTraceID3and4, 5, testPolicyTraceID3and4},
		{testPolicyTraceID3and4, 0, testPolicyNoTrace},
		{testPolicyTraceID5withMark, 5, testPolicyNoTrace},
		{testPolicyTraceID5withMark, 0, testPolicyNoTrace},
		{testPolicyTraceID5withMark, 6, testPolicyTraceID5withMark},
	}

	for _, policy := range policies {
		got := clearIptablesPolicy(policy.policy, policy.cleanupID)
		if !cmp.Equal(got, policy.expected) {
			t.Errorf("clearIptablesPolicy was incorrect, got: '%s', expected: '%s'", got, policy.expected)
		}
	}
}
