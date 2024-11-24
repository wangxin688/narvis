package ipam_biz

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/ipam/schemas"
	ipam_utils "github.com/wangxin688/narvis/server/features/ipam/utils"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

type PrefixService struct{}

func NewPrefixService() *PrefixService {
	return &PrefixService{}
}

func (p *PrefixService) CreatePrefix(prefix *schemas.PrefixCreate) (string, error) {
	newPrefix := models.Prefix{
		Range:          prefix.Range,
		Gateway:        prefix.Gateway,
		VlanId:         prefix.VlanId,
		VlanName:       prefix.VlanName,
		Type:           prefix.Type,
		SiteId:         prefix.SiteId,
		OrganizationId: contextvar.OrganizationId.Get(),
		Version:        ipam_utils.CidrVersion(prefix.Range),
	}

	err := gen.Prefix.Create(&newPrefix)
	if err != nil {
		return "", err
	}
	return newPrefix.Id, nil
}

func (p *PrefixService) UpdatePrefix(g *gin.Context, id string, prefix *schemas.PrefixUpdate) error {
	dbPrefix, err := gen.Prefix.Where(gen.Prefix.Id.Eq(id), gen.Prefix.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return err
	}
	updateFields := make(map[string]*contextvar.Diff)
	if prefix.Range != nil && *prefix.Range != dbPrefix.Range {
		updateFields["range"] = &contextvar.Diff{Before: dbPrefix.Range, After: *prefix.Range}
		dbPrefix.Range = *prefix.Range
	}
	if prefix.VlanId != nil && *prefix.VlanId != *dbPrefix.VlanId {
		updateFields["vlanId"] = &contextvar.Diff{Before: dbPrefix.VlanId, After: *prefix.VlanId}
		dbPrefix.VlanId = prefix.VlanId
	}
	if prefix.VlanName != nil && *prefix.VlanName != *dbPrefix.VlanName {
		updateFields["vlanName"] = &contextvar.Diff{Before: dbPrefix.VlanName, After: *prefix.VlanName}
		dbPrefix.VlanName = prefix.VlanName
	}
	if prefix.Type != nil && *prefix.Type != dbPrefix.Type {
		updateFields["type"] = &contextvar.Diff{Before: dbPrefix.Type, After: *prefix.Type}
		dbPrefix.Type = *prefix.Type
	}
	if helpers.HasParams(g, "SiteId") && *prefix.SiteId != dbPrefix.SiteId {
		updateFields["siteId"] = &contextvar.Diff{Before: dbPrefix.SiteId, After: *prefix.SiteId}
		dbPrefix.SiteId = *prefix.SiteId
	}
	if prefix.Gateway != nil && prefix.Gateway != dbPrefix.Gateway {
		updateFields["gateway"] = &contextvar.Diff{Before: dbPrefix.Gateway, After: *prefix.Gateway}
		dbPrefix.Gateway = prefix.Gateway
	}
	if len(updateFields) == 0 {
		return nil
	}
	diffValue := make(map[string]map[string]*contextvar.Diff)
	diffValue[id] = updateFields
	contextvar.OrmDiff.Set(diffValue)
	err = gen.Prefix.UnderlyingDB().Save(dbPrefix).Error
	return err
}

func (p *PrefixService) DeletePrefix(id string) error {
	_, err := gen.Prefix.Where(gen.Prefix.Id.Eq(id), gen.Prefix.OrganizationId.Eq(contextvar.OrganizationId.Get())).Delete()
	return err
}

func (p *PrefixService) GetById(id string) (*schemas.Prefix, error) {
	prefix, err := gen.Prefix.Where(gen.Prefix.Id.Eq(id), gen.Prefix.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	_prefix := []string{prefix.Range}
	utilization, err := p.CalPrefixUsage(_prefix)
	if err != nil {
		return nil, err
	}
	return &schemas.Prefix{
		Id:          prefix.Id,
		CreatedAt:   prefix.CreatedAt,
		UpdatedAt:   prefix.UpdatedAt,
		Range:       prefix.Range,
		Version:     prefix.Version,
		Gateway:     prefix.Gateway,
		VlanId:      prefix.VlanId,
		VlanName:    prefix.VlanName,
		Type:        prefix.Type,
		SiteId:      prefix.SiteId,
		Utilization: utilization[prefix.Id],
	}, nil
}

func (p *PrefixService) ListPrefix(query *schemas.PrefixQuery) (int64, *[]*schemas.Prefix, error) {
	res := make([]*schemas.Prefix, 0)
	stmt := gen.Prefix.Where(gen.Prefix.OrganizationId.Eq(contextvar.OrganizationId.Get()))
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
	ranges := lo.Map(list, func(item *models.Prefix, _ int) string {
		return item.Range
	})
	utilization, err := p.CalPrefixUsage(ranges)
	if err != nil {
		return 0, &res, err
	}
	for _, prefix := range list {
		res = append(res, &schemas.Prefix{
			Id:          prefix.Id,
			CreatedAt:   prefix.CreatedAt,
			UpdatedAt:   prefix.UpdatedAt,
			Version:     prefix.Version,
			Gateway:     prefix.Gateway,
			Range:       prefix.Range,
			VlanId:      prefix.VlanId,
			VlanName:    prefix.VlanName,
			Type:        prefix.Type,
			SiteId:      prefix.SiteId,
			Utilization: utilization[prefix.Range],
		})
	}
	return total, &res, nil
}

func (p *PrefixService) CalPrefixUsage(prefix []string) (map[string]float64, error) {
	type prefixCount struct {
		Range string
		Count int
	}
	dbResult := make([]*prefixCount, 0)
	err := gen.IpAddress.Select(
		gen.IpAddress.Range.As("Range"),
		gen.IpAddress.Address.Count().As("Count"),
	).Where(
		gen.IpAddress.Range.In(prefix...),
		gen.IpAddress.OrganizationId.Eq(contextvar.OrganizationId.Get()),
	).Group(gen.IpAddress.Range).Scan(&dbResult)
	if err != nil {
		return nil, err
	}
	prefixUsage := make(map[string]float64)
	for _, count := range dbResult {
		if count.Count == 0 {
			prefixUsage[count.Range] = 0
			continue
		}
		rangeSize := ipam_utils.CidrSize(count.Range)
		if rangeSize == 0 {
			// ignore ipv6 mask length < 64
			prefixUsage[count.Range] = 0
			continue
		}
		prefixUsage[count.Range] = (float64(count.Count) / float64(rangeSize)) * 100
	}
	return prefixUsage, nil
}
