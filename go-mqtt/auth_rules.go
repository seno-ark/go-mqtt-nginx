package main

import "github.com/mochi-mqtt/server/v2/hooks/auth"

var AuthRules = &auth.Ledger{
	Auth: auth.AuthRules{ // Auth disallows all by default
		{Username: "peach", Password: "peach1", Allow: true},
		{Username: "melon", Password: "melon1", Allow: true},
		{Remote: "127.0.0.1:*", Allow: true},
		{Remote: "localhost:*", Allow: true},
	},
	ACL: auth.ACLRules{ // ACL allows all by default
		{Remote: "127.0.0.1:*"}, // local superuser allow all
		{
			// user melon can read and write to their own topic
			Username: "peach", Filters: auth.Filters{
				"device/peach/#":  auth.ReadWrite,
				"device/+/status": auth.ReadOnly,
				// "updates/#": auth.WriteOnly, // can write to updates, but can't read updates from others
			},
		},
		{
			// user melon can read and write to their own topic
			Username: "melon", Filters: auth.Filters{
				"device/melon/#":  auth.ReadWrite,
				"device/+/status": auth.ReadOnly,
				// "updates/#": auth.WriteOnly, // can write to updates, but can't read updates from others
			},
		},
		{
			// Otherwise, no clients have publishing permissions
			Filters: auth.Filters{
				"#":         auth.ReadOnly,
				"updates/#": auth.Deny,
			},
		},
	},
}
