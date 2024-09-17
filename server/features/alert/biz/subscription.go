package alert_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"gorm.io/datatypes"
)

type SubscriptionService struct{}

func NewSubscriptionService() *SubscriptionService {
	return &SubscriptionService{}
}

func (s *SubscriptionService) CreateSubscription(sub *schemas.SubscriptionCreate) (string, error) {
	newSub := &models.Subscription{
		Name:           sub.Name,
		Enabled:        sub.Enabled,
		Deduplication:  sub.Deduplication,
		SendResolved:   sub.SendResolved,
		Conditions:     condToDbCond(sub.Conditions),
		RepeatInterval: sub.RepeatInterval,
		ChannelType:    sub.ChannelType,
		ChannelConfig:  chanConfigToDbConfig(&sub.ChannelConfig),
		OrganizationId: global.OrganizationId.Get(),
		CreatedById:    global.UserId.Get(),
	}

	err := gen.Subscription.Create(newSub)
	if err != nil {
		return "", err
	}
	return newSub.Id, nil
}

func (s *SubscriptionService) UpdateSubscription(subID string, sub *schemas.SubscriptionUpdate) error {

	dbSub, err := gen.Subscription.Where(gen.Subscription.Id.Eq(subID), gen.Subscription.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return err
	}
	if sub.Name != nil {
		dbSub.Name = *sub.Name
	}
	if sub.Enabled != nil {
		dbSub.Enabled = *sub.Enabled
	}
	if sub.Deduplication != nil {
		dbSub.Deduplication = *sub.Deduplication
	}
	if sub.SendResolved != nil {
		dbSub.SendResolved = *sub.SendResolved
	}
	if sub.RepeatInterval != nil {
		dbSub.RepeatInterval = *sub.RepeatInterval
	}
	if sub.Conditions != nil {
		dbSub.Conditions = condToDbCond(*sub.Conditions)
	}
	if sub.ChannelType != nil {
		dbSub.ChannelType = *sub.ChannelType
	}
	if sub.ChannelConfig != nil {
		dbSub.ChannelConfig = chanConfigToDbConfig(sub.ChannelConfig)
	}
	err = gen.Subscription.UnderlyingDB().Save(dbSub).Error
	return err
}

func (s *SubscriptionService) DeleteSubscription(subID string) error {
	_, err := gen.Subscription.Where(
		gen.Subscription.Id.Eq(subID),
		gen.Subscription.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	return err
}

func (s *SubscriptionService) GetById(subID string) (*schemas.Subscription, error) {
	var result schemas.Subscription
	err := gen.Subscription.Where(
		gen.Subscription.Id.Eq(subID),
		gen.Subscription.OrganizationId.Eq(global.OrganizationId.Get())).Preload(
		gen.Subscription.CreatedBy,
	).Preload(gen.Subscription.UpdatedBy).Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SubscriptionService) ListSubscriptions(query *schemas.SubscriptionQuery) (int64, []*schemas.Subscription, error) {
	var result []*schemas.Subscription
	stmt := gen.Subscription.Where(gen.Subscription.OrganizationId.Eq(global.OrganizationId.Get()))
	if query.Id != nil {
		stmt = stmt.Where(gen.Subscription.Id.In(*query.Id...))
	}
	if query.Name != nil {
		stmt = stmt.Where(gen.Subscription.Name.Eq(*query.Name))
	}
	if query.Enabled != nil {
		stmt = stmt.Where(gen.Subscription.Enabled.Is(*query.Enabled))
	}
	if query.ChannelType != nil {
		stmt = stmt.Where(gen.Subscription.ChannelType.Eq(*query.ChannelType))
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where(gen.Subscription.Name.Like(keyword))
	}

	total, err := stmt.Count()
	if err != nil || total <= 0 {
		return 0, nil, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	err = stmt.Preload(gen.Subscription.CreatedBy).Preload(gen.Subscription.UpdatedBy).Scan(&result)
	if err != nil {
		return 0, nil, err
	}
	return total, result, nil
}

func chanConfigToDbConfig(chanConfig *schemas.ChannelConfig) datatypes.JSONType[models.ChannelConfig] {
	config := models.ChannelConfig{
		WebhookUrl:     chanConfig.WebhookUrl,
		WebhookHeaders: chanConfig.WebhookHeaders,
		Email:          chanConfig.Email,
	}

	return datatypes.NewJSONType(config)
}
