/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"net"
	"net/url"
	"strings"

	"github.com/gravitational/trace"
)

// IsDatabaseEndpoint returns true if provided endpoint is a valid database
// endpoint.
func IsDatabaseEndpoint(endpoint string) bool {
	return strings.Contains(endpoint, DatabaseEndpointSuffix)
}

// IsCacheForRedisEndpoint returns true if provided endpoint is a valid Azure
// Cache for Redis endpoint.
func IsCacheForRedisEndpoint(endpoint string) bool {
	return IsRedisEndpoint(endpoint) || IsRedisEnterpriseEndpoint(endpoint)
}

// IsRedisEndpoint returns true if provided endpoint is a valid Redis
// (non-Enterprise tier) endpoint.
func IsRedisEndpoint(endpoint string) bool {
	return strings.Contains(endpoint, RedisEndpointSuffix)
}

// IsRedisEnterpriseEndpoint returns true if provided endpoint is a valid Redis
// Enterprise endpoint.
func IsRedisEnterpriseEndpoint(endpoint string) bool {
	return strings.Contains(endpoint, RedisEnterpriseEndpointSuffix)
}

// ParseDatabaseEndpoint extracts database server name from Azure endpoint.
func ParseDatabaseEndpoint(endpoint string) (name string, err error) {
	host, _, err := net.SplitHostPort(endpoint)
	if err != nil {
		return "", trace.Wrap(err)
	}
	// Azure endpoint looks like this:
	// name.mysql.database.azure.com
	parts := strings.Split(host, ".")
	if !strings.HasSuffix(host, DatabaseEndpointSuffix) || len(parts) != 5 {
		return "", trace.BadParameter("failed to parse %v as Azure endpoint", endpoint)
	}
	return parts[0], nil
}

// ParseCacheForRedisEndpoint extracts database server name from Azure Cache
// for Redis endpoint.
func ParseCacheForRedisEndpoint(endpoint string) (name string, err error) {
	// Note that the Redis URI may contain schema and parameters.
	host, err := GetHostFromRedisURI(endpoint)
	if err != nil {
		return "", trace.Wrap(err)
	}

	switch {
	// Redis (non-Enterprise) endpoint looks like this:
	// name.redis.cache.windows.net
	case strings.HasSuffix(host, RedisEndpointSuffix):
		return strings.TrimSuffix(host, RedisEndpointSuffix), nil

	// Redis Enterprise endpoint looks like this:
	// name.region.redisenterprise.cache.azure.net
	case strings.HasSuffix(host, RedisEnterpriseEndpointSuffix):
		name, _, ok := strings.Cut(strings.TrimSuffix(host, RedisEnterpriseEndpointSuffix), ".")
		if !ok {
			return "", trace.BadParameter("failed to parse %v as Azure Cache endpoint", endpoint)
		}
		return name, nil

	default:
		return "", trace.BadParameter("failed to parse %v as Azure Cache endpoint", endpoint)
	}
}

// GetHostFromRedisURI extracts host name from a Redis URI. The URI may start
// with "redis://", "rediss://", or without. The URI may also have parameters
// like "?mode=cluster".
func GetHostFromRedisURI(uri string) (string, error) {
	// Add a temporary schema to make a valid URL for url.Parse if schema is
	// not found.
	if !strings.Contains(uri, "://") {
		uri = "schema://" + uri
	}

	parsed, err := url.Parse(uri)
	if err != nil {
		return "", trace.Wrap(err)
	}
	return parsed.Hostname(), nil
}

const (
	// DatabaseEndpointSuffix is the Azure database endpoint suffix. Used for
	// MySQL, PostgresSQL, etc.
	DatabaseEndpointSuffix = ".database.azure.com"

	// RedisEndpointSuffix is the endpoint suffix for Redis.
	RedisEndpointSuffix = ".redis.cache.windows.net"

	// RedisEnterpriseEndpointSuffix is the endpoint suffix for Redis Enterprise.
	RedisEnterpriseEndpointSuffix = ".redisenterprise.cache.azure.net"
)
