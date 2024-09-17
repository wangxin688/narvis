CREATE TABLE IF NOT EXISTS narvis.mac_addr
(
    `ts` DateTime64(3,'UTC') CODEC(DoubleDelta, LZ4),
    `organizationId` String
    `macAddress` FixedString(17),
    `DeviceId` LowCardinality(String),
    `ManagementIp` String CODEC(ZSTD(1)),
    `IfIndex` UInt16,
    `IfName` String,
    `VlanId` UInt16
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(ts)
ORDER BY (organizationId, toStartOfHour(ts), deviceId)
TTL ts + INTERVAL 30 DAY -- expire in 30 days
SETTINGS index_granularity = 8192;