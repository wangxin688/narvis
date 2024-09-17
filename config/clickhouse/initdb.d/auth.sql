CREATE TABLE IF NOT EXISTS narvis.auth
(
    `ts` DateTime64(3, 'UTC') CODEC(DoubleDelta, LZ4),
    `organizationId` LowCardinality(String),
    `aaaProfile` LowCardinality(String),
    `apName` String,
    `authMethod` LowCardinality(String),
    `accessMethod` LowCardinality(String),
    `authServer` LowCardinality(String),
    `clientIp` String CODEC(ZSTD(1)),
    `clientMac` String CODEC(ZSTD(1)),
    `hostIp` String CODEC(ZSTD(1)),
    `role` LowCardinality(String),
    `ssid` LowCardinality(String),
    `username` String CODEC(ZSTD(1)),
    `vlan` String CODEC(ZSTD(1)),
    `meta.deviceId` LowCardinality(String),
    `meta.deviceName` String CODEC(ZSTD(1)),
    `meta.deviceRole` LowCardinality(String),
    `meta.managementIp` String CODEC(ZSTD(1)),
    `meta.manufacturer` LowCardinality(String),
    `meta.siteCode` LowCardinality(String),
    `meta.siteId` LowCardinality(String),
    `meta.siteName` LowCardinality(String),
    `meta.organizationId` LowCardinality(String)
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(ts)
ORDER BY (organizationId, toStartOfHour(ts))
TTL toDateTime(ts) + toIntervalDay(30)
SETTINGS index_granularity = 8192;