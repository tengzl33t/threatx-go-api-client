/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: endpoints.go
Description: API endpoints
Author: tengzl33t
*/

package internal

func getEndpointMap() map[string][2]interface{} {
	return map[string][2]interface{}{
		"apikeys": {
			2, []string{
				"list",
				"new",
				"update",
				"revoke",
			},
		},
		"apischemas": {
			1, []string{
				"save",
				"list",
				"delete",
			},
		},
		"customers": {
			1, []string{
				"list",
				"list_all",
				"new",
				"get",
				"update",
				"delete",
				"list_api_keys",
				"new_api_key",
				"delete_api_key",
				"get_customer_config",
				"set_customer_config",
			},
		},
		"users": {
			1, []string{
				"list",
				"new",
				"get",
				"update",
				"delete",
			},
		},
		"sites": {
			2, []string{
				"list",
				"new",
				"get",
				"delete",
				"update",
				"unset",
			},
		},
		"sitegroups": {
			1, []string{
				"list",
				"save",
				"delete",
			},
		},
		"templates": {
			1, []string{
				"set",
				"get",
				"delete",
			},
		},
		"sensors": {
			1, []string{
				"list",
				"tags",
			},
		},
		"services": {
			1, []string{
				"list",
			},
		},
		"entities": {
			1, []string{
				"list",
				"show",
				"state_changes",
				"risk_changes",
				"notes",
				"new_note",
				"reset",
				"block_entity",
				"blacklist_entity",
				"whitelist_entity",
				"watch_entity",
				"list_most_risky",
				"count",
			},
		},
		"metrics": {
			1, []string{
				"request_stats_by_hour",
				"request_stats_by_minute",
				"match_stats_by_hour",
				"block_stats_by_endpoint",
				"entity_stats_by_entity_by_quarter_hour",
				"rules_matched_by_ip_by_quarter_hour",
				"request_stats_by_endpoint",
				"threat_stats_by_endpoint",
				"threat_stats_by_hour",
				"threat_stats_by_quarter_hour",
				"threat_stats_by_site",
				"status_codes_by_site",
				"request_stats_hourly_by_site",
				"request_stats_hourly_by_endpoint",
			},
		},
		"subscriptions": {
			1, []string{
				"save",
				"delete",
				"list",
				"enable",
				"disable",
			},
		},
		"globaltags": {
			1, []string{
				"new",
				"list",
			},
		},
		"actortags": {
			1, []string{
				"new",
				"list",
				"delete",
			},
		},
		"features": {
			1, []string{
				"list",
				"query",
				"save",
				"delete",
			},
		},
		"channels": {
			1, []string{
				"new",
				"list",
				"update",
			},
		},
		"globalsettings": {
			1, []string{
				"get",
			},
		},
		"dnsinfo": {
			1, []string{
				"list",
			},
		},
		"logs": {
			1, []string{
				"events",
				"entities",
				"blocks",
				"actions",
				"matches",
				"rule_hits",
				"sysinfo",
				"audit_log",
			},
		},
		"logsv2": {
			2, []string{
				"block_events",
				"match_events",
				"audit_events",
			},
		},
		"lists": {
			1, []string{
				"list_blacklist",
				"list_blocklist",
				"list_whitelist",
				"list_ignorelist",
				"new_blacklist",
				"new_blocklist",
				"new_whitelist",
				"new_ignorelist",
				"bulk_new_blacklist",
				"bulk_new_blocklist",
				"bulk_new_whitelist",
				"bulk_new_ignorelist",
				"get_blacklist",
				"get_blocklist",
				"get_whitelist",
				"get_ignorelist",
				"delete_blacklist",
				"delete_blocklist",
				"delete_whitelist",
				"delete_ignorelist",
				"bulk_delete_blacklist",
				"bulk_delete_blocklist",
				"bulk_delete_whitelist",
				"bulk_delete_ignorelist",
				"ip_to_link",
			},
		},
		"rules": {
			1, []string{
				"list_customer_rules",
				"list_whitelist_rules",
				"list_profiler_rules",
				"list_common_rules",
				"new_customer_rule",
				"new_whitelist_rule",
				"new_common_rule",
				"update_customer_rule",
				"update_whitelist_rule",
				"update_profiler_rule",
				"update_common_rule",
				"get_customer_rule",
				"get_whitelist_rule",
				"get_profiler_rule",
				"get_common_rule",
				"delete_customer_rule",
				"delete_whitelist_rule",
				"delete_profiler_rule",
				"delete_common_rule",
				"validate_rule",
			},
		},
	}
}
