package integration

import (
	"context"
	"errors"
	"fmt"

	priv "github.com/99designs/gqlgen/_examples/large-project-structure/main/private/graph"
	"github.com/99designs/gqlgen/_examples/large-project-structure/main/public/graph/model"
	shared "github.com/99designs/gqlgen/_examples/large-project-structure/shared/graph"
)

type Resolver struct {
	Logger shared.Logger
}

// Implement the Tezz method that is managed by another team
func (r *Resolver) Tezz(ctx context.Context) (*model.Test, error) {
	r.Logger.Info("Tezz was hit!")
	// Can do whatever logic is needed...
	return &model.Test{ID: "external-1"}, nil
}

func (r *Resolver) GetYaSome(ctx context.Context, input *model.CustomInput) ([]*model.CustomZeekIntel, error) {
	r.Logger.Info("Better GetYaSome")
	intels := []*model.CustomZeekIntel{}

	if input.Error != nil && *input.Error {
		return intels, errors.New("error as requested")
	}

	if input.Limit != nil {
		count := int(*input.Limit)
		for i := 0; i < count; i++ {
			czi := &model.CustomZeekIntel{
				ID:         fmt.Sprintf("%d", i),
				Name:       fmt.Sprintf("external-%d", i),
				ExtraField: "let other teams resolve",
			}
			intels = append(intels, czi)
		}
	}

	return intels, nil
}

func (r *Resolver) AddIndicator(ctx context.Context, input model.IndicatorInput) (*model.Indicator, error) {
	r.Logger.Infof("Oh, looking to add indicator", input)

	q := priv.QResolver{}
	c, err := q.GetConnector(ctx, 47)
	if err != nil {
		return nil, errors.New("unable to get ur connector, my bad bruh")
	}
	fmt.Println("got a connector for your..")
	fmt.Printf("%+v\n", c)
	return &model.Indicator{
		ID:            "1234",
		Indicator:     input.Indicator,
		IndicatorType: input.IndicatorType,
		MetaSource:    input.MetaSource,
	}, nil
}

func (r *Resolver) GetAppCrowdStrike(ctx context.Context) (*model.AppCrowdStrike, error) {

	docUrl := "https://docs.corelight.com/docs/sensor/corelight-update/policyConfiguration/sources/integrations/crowdStrike.html#crowdstrike"
	logoUrl := "https://it.lbl.gov/wp-content/uploads/sites/18/2023/02/crowdstrike-logo.png"
	logoAlt := "CrowdStrike Falcon"

	numColumns := int32(1)
	firstStepDesc := "Lots of additional info or tooltip could be displayed here"
	defaultValue := "dee-fault"
	typeInput := model.UIDDynamicFieldTypeInput
	typeSelect := model.UIDDynamicFieldTypeSelect

	app := &model.AppCrowdStrike{
		ID:          "AppCrowdStrike",
		Name:        "CrowdStrike Falcon API",
		Description: "Cybersecurity’s AI-native platform for the XDR era",
		DocURL:      &docUrl,
		Logo: &model.UIImage{
			// FIXME: This should come from Fleet, no external dependencies
			URL: &logoUrl,
			Alt: &logoAlt,
		},
		Tags: &model.AppConnectorChip{
			Chips: []model.AppConnectorChipEnum{
				model.AppConnectorChipEnumHostEnrichment,
				model.AppConnectorChipEnumCveVuln,
			},
		},
		Items: []*model.UIDynamicFieldSet{
			{
				Columns:     &numColumns,
				Label:       "First Step",
				Description: &firstStepDesc,
				Fields: []*model.UIDynamicField{
					{
						Label:       "thing1",
						Description: &firstStepDesc,
						Required:    true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeInput,
							Value:     &defaultValue,
						},
					},
					{
						Label:       "thing2",
						Description: &firstStepDesc,
						Required:    true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeSelect,
							Value:     &defaultValue,
						},
					},
				},
				Conditions: []*model.UIDynamicCondition{},
			},
			{
				Columns:     &numColumns,
				Label:       "Second Step, ahh yeah",
				Description: &firstStepDesc,
				Fields:      []*model.UIDynamicField{},
				Conditions:  []*model.UIDynamicCondition{},
			},
		},
	}
	return app, nil
}
