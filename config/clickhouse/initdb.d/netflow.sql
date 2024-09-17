CREATE TABLE IF NOT EXISTS narvis.netflow
(
   `ts` DateTime('UTC'),
   `organizationId` LowCardinality(String),
   `flow.bytes` UInt32,
   `flow.client.as.label` Enum8('PRIVATE' = 0, 'PUBLIC' = 1),
   `flow.client.host.name` String,
   `flow.client.ip.addr` String, -- # TODO: ip address optimize
    --`flow.client.ip.subnet.name` LowCardinality(String),
   `flow.client.l4.port.name` LowCardinality(String),
   `flow.conversation.id` String,
   `flow.direction.name` Enum8('Ingress' = 0, 'Egress' = 1),
   `flow.dst.as.label` Enum8('PRIVATE' = 0, 'PUBLIC' = 1),
   `flow.dst.host.name` String,
   `flow.dst.ip.addr` String, -- # TODO: ip address optimize
    --`flow.dst.ip.subnet.name` LowCardinality(String),
   `flow.dst.l4.port.name` LowCardinality(String),
   `flow.export.host.name` LowCardinality(String),
   `flow.export.ip.addr` String, -- # TODO: ip address optimize
   `flow.export.version.name` LowCardinality(String),
    --`flow.in.netif.alias` String,
    --`flow.in.netif.name` LowCardinality(String),
   `flow.locality` Enum8('public' = 0, 'private' = 1),  -- # TODO: needed?
    --`flow.meter.packet_select.interval.packets` UInt32,
    --`flow.out.netif.alias` String,
    --`flow.out.netif.name` LowCardinality(String),
   `flow.packets` UInt32, -- # TODO: LowCardinality
   `flow.server.as.label` Enum8('PRIVATE' = 0, 'PUBLIC' = 1),
   `flow.server.host.name` String,
   `flow.server.ip.addr` String, -- # TODO: ip address optimize
    --`flow.server.ip.subnet.name` LowCardinality(String),
   `flow.server.l4.port.name` LowCardinality(String),
   `flow.src.host.name` String,
   `flow.src.ip.addr` String, -- # TODO: ip address optimize
    --`flow.src.ip.subnet.name` LowCardinality(String),
   `flow.src.l4.port.name` LowCardinality(String),
   `ip.version.name` Enum8('IPv4' = 0, 'IPv6' = 1),
   `l4.proto.name` LowCardinality(String),
   `flow.export.siteId` LowCardinality(String),
   `flow.export.siteCode` LowCardinality(String)
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(ts)
PRIMARY KEY (flow.export.siteId, toStartOfHour(ts), flow.client.ip.addr)
ORDER BY (flow.export.siteId, toStartOfHour(ts), flow.client.ip.addr, ts)
TTL ts + toIntervalDay(30)
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS narvis.flow_agg_5m
(
    `ts` DateTime('UTC'),
    `flow.export.siteId` LowCardinality(String),
    `flow.client.ip.addr` String,
    `flow.client.host.name` String,
    `flow.bytes` UInt64,
    `flow.direction.name` Enum8('Ingress' = 0,
 'Egress' = 1)
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(ts)
ORDER BY (flow.export.siteId,
 toStartOfHour(ts),
 flow.client.ip.addr,
 flow.client.host.name)
TTL ts + toIntervalDay(30)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS narvis.flow_agg_5m_mv TO narvis.flow_agg_5m
(
    `flow.export.siteId` LowCardinality(String),
    `ts` DateTime('UTC'),
    `flow.client.ip.addr` String,
    `flow.client.host.name` String,
    `flow.bytes` UInt64,
    `flow.direction.name` Enum8('Ingress' = 0, 'Egress' = 1)
) AS
SELECT
    `flow.export.siteId`,
    toStartOfFiveMinutes(ts) AS ts,
    `flow.client.ip.addr`,
    `flow.client.host.name`,
    sum(`flow.bytes`) AS `flow.bytes`,
    `flow.direction.name`
FROM narvis.flow
GROUP BY
    `flow.export.siteId`,
    `ts`,
    `flow.client.ip.addr`,
    `flow.client.host.name`,
    `flow.direction.name`;
