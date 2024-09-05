package ipam_biz

import (
	"strings"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/ipam/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type PrefixService struct{}

func NewPrefixService() *PrefixService {
	return &PrefixService{}
}

func (p *PrefixService) CreatePrefix(prefix *schemas.PrefixCreate) (string, error) {
	newPrefix := models.Prefix{
		Range:          prefix.Range,
		VlanId:         prefix.VlanId,
		VlanName:       prefix.VlanName,
		Type:           prefix.Type,
		SiteId:         prefix.SiteId,
		OrganizationId: global.OrganizationId.Get(),
	}

	err := gen.Prefix.Create(&newPrefix)
	if err != nil {
		return "", err
	}
	return newPrefix.Id, nil
}

func (p *PrefixService) UpdatePrefix(id string, prefix *schemas.PrefixUpdate) error {
	updateFields := make(map[string]any)

	if prefix.Range != nil {
		updateFields["range"] = *prefix.Range
	}
	if prefix.VlanId != nil {
		updateFields["vlanId"] = *prefix.VlanId
	}
	if prefix.VlanName != nil {
		updateFields["vlanName"] = *prefix.VlanName
	}
	if prefix.Type != nil {
		updateFields["type"] = *prefix.Type
	}
	if prefix.SiteId != nil {
		updateFields["siteId"] = *prefix.SiteId
	}

	if len(updateFields) == 0 {
		return nil
	}
	_, err := gen.Prefix.Where(gen.Prefix.Id.Eq(id), gen.Prefix.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	return err
}

func (p *PrefixService) DeletePrefix(id string) error {
	_, err := gen.Prefix.Where(gen.Prefix.Id.Eq(id), gen.Prefix.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	return err
}

func (p *PrefixService) GetById(id string) (*schemas.Prefix, error) {
	prefix, err := gen.Prefix.Where(gen.Prefix.Id.Eq(id), gen.Prefix.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	return &schemas.Prefix{
		Id:       prefix.Id,
		Range:    prefix.Range,
		VlanId:   prefix.VlanId,
		VlanName: prefix.VlanName,
		Type:     prefix.Type,
		SiteId:   prefix.SiteId,
	}, nil
}

func (p *PrefixService) ListPrefix(query *schemas.PrefixQuery) (int64, *[]*schemas.Prefix, error) {
	res := make([]*schemas.Prefix, 0)
	stmt := gen.Prefix.Where(gen.Prefix.OrganizationId.Eq(global.OrganizationId.Get()))
	if query.SiteId != nil {
		stmt = stmt.Where(gen.Prefix.SiteId.Eq(*query.SiteId))
	}
	if query.Range != nil {
		stmt = stmt.Where(gen.Prefix.Range.In(*query.Range...))
	}
	if query.VlanId != nil {
		stmt = stmt.Where(gen.Prefix.VlanId.In(*query.VlanId...))
	}
	if query.VlanName != nil {
		stmt = stmt.Where(gen.Prefix.VlanName.In(*query.VlanName...))
	}
	if query.Type != nil {
		stmt = stmt.Where(gen.Prefix.Type.Eq(*query.Type))
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where(gen.Prefix.Range.Like(keyword)).Or(
			gen.Prefix.VlanName.Like(keyword))
	}

	total, err := stmt.Count()
	if err != nil || total < 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	list, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	for _, prefix := range list {
		res = append(res, &schemas.Prefix{
			Id:       prefix.Id,
			Range:    prefix.Range,
			VlanId:   prefix.VlanId,
			VlanName: prefix.VlanName,
			Type:     prefix.Type,
			SiteId:   prefix.SiteId,
		})
	}
	return total, &res, nil
}

func (p *PrefixService) CalPrefixUsage(prefix []string) (map[string]float32, error) {
	rawSql := `
	SELECT 
		p.id AS prefix_id, 
		p.range AS prefix_range, 
		COUNT(ipa.ip) AS used_ips,
		-- 计算CIDR范围内的IP总数。使用pg_catalog.broadcast() 获取CIDR的广播地址，减去网络地址
		CASE 
			WHEN family(p.range) = 4 THEN 
				(2 ^ (32 - masklen(p.range))) - 2  -- IPv4的总IP数，减去网络地址和广播地址
			WHEN family(p.range) = 6 THEN 
				(2 ^ (128 - masklen(p.range)))     -- IPv6不需要减去地址
		END AS total_ips,
		-- 计算利用率
		CASE 
			WHEN family(p.range) = 4 THEN 
				ROUND((COUNT(ipa.ip)::decimal / ((2 ^ (32 - masklen(p.range))) - 2)) * 100, 2)
			WHEN family(p.range) = 6 THEN 
				ROUND((COUNT(ipa.ip)::decimal / (2 ^ (128 - masklen(p.range))) * 100), 2)
		END AS utilization_percentage
	FROM 
		prefix p
	WHERE
		p.id IN ( ` + strings.Join(prefix, ",") + ` )
	LEFT JOIN 
		ip_address ipa ON ipa.ip << p.range  -- 使用<<操作符判断ip是否在cidr范围内
	GROUP BY 
		p.id, p.range
	ORDER BY 
		utilization_percentage DESC;
	`

	res := make(map[string]float32)
	err := gen.Prefix.UnderlyingDB().Raw(rawSql).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

