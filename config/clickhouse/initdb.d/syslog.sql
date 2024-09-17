CREATE TABLE IF NOT EXISTS narvis.syslog
(
    `ts` DateTime64(3,'UTC') CODEC(DoubleDelta, LZ4),
    `facility` LowCardinality(String),
    `severity` LowCardinality(String),
    `message` String CODEC(ZSTD(5)),
    `messageToken` Array(String), -- full text tokenization, insert values tokenize('message')
    `meta.deviceId` LowCardinality(String),
    `meta.deviceName` String CODEC(ZSTD(1)),
    `meta.deviceRole` LowCardinality(String),
    `meta.managementIp` String CODEC(ZSTD(1)),
    `meta.manufacturer` LowCardinality(String),
    `meta.siteCode` LowCardinality(String),
    `meta.siteId` LowCardinality(String),
    `meta.siteName` LowCardinality(String),
    `meta.OrganizationId` LowCardinality(String),
    `OrganizationId` LowCardinality(String),
    INDEX idx_message_tokens message TYPE tokenbf_v1(30720, 3, 0) GRANULARITY 1)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(ts) -- partition by date
ORDER BY (OrganizationId, toStartOfHour(ts), meta.siteId, meta.deviceId) -- order by main fields
TTL ts + INTERVAL 30 DAY -- expire in 30 days
SETTINGS index_granularity = 8192;