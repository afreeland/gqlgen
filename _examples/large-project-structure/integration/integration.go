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

// AddIndicator(ctx context.Context, input model.IndicatorInput) (*model.Indicator, error)
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
	twoColumns := int32(2)

	firstStepDesc := "Lots of additional info or tooltip could be displayed here"
	defaultValue := "dee-fault"
	typeInput := model.UIDDynamicFieldTypeInput
	typeSelect := model.UIDDynamicFieldTypeSelect
	typeStr := model.UIDefaultValueTypeString
	typeInt := model.UIDefaultValueTypeInt

	small := "10"
	large := "20"

	nnull := model.UIDynamicConditionOperatorNnull
	gt := model.UIDynamicConditionOperatorGt
	gtNum := "3"

	item1Text := "Slide 1"
	item1Img := "https://i.redd.it/h0d3v4y8xg221.jpg"

	item2Text := "Slide 2"
	item2Img := "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/i/45fb71b4-b618-45ed-b0ef-ddf151262ba5/dc3zd7q-4d0b022a-07db-444f-988e-63abfc2d2a71.png"

	app := &model.AppCrowdStrike{
		ID:          "AppCrowdStrike",
		Name:        "CrowdStrike Falcon API",
		Description: "Cybersecurity’s AI-native platform for the XDR era",
		DocURL:      &docUrl,
		Carousel: &model.UICarousel{
			Items: []*model.UICarouselItem{
				{
					Text: &item1Text,
					Image: &model.UIImage{
						URL: &item1Img,
						Alt: &item1Text,
					},
				},
				{
					Text: &item2Text,
					Image: &model.UIImage{
						URL: &item2Img,
						Alt: &item2Text,
					},
				},
			},
		},
		Logo: &model.UIImage{
			// FIXME: This should come from Fleet, no external dependencies (url)
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
						FieldName:   "thing.1",
						Description: &firstStepDesc,
						Required:    true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeStr,
							Value:     &defaultValue,
						},
						Type: &typeInput,
						Validator: model.UIRegexValidator{
							Type:    "regex",
							Message: "thing1 must be in the following format: word1",

							Pattern: "\\w\\d+",
						},
					},
					{
						Label:       "thing2",
						Description: &firstStepDesc,
						Required:    true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeStr,
							Value:     &defaultValue,
						},
						Type: &typeSelect,
						Options: []*model.UIDynamicSelectOption{
							{
								Label: "Option 1",
								Value: &model.UIDefaultValue{
									ValueType: &typeStr,
									Value:     &defaultValue,
								},
							},
							{
								Label: "Option 2",
								Value: &model.UIDefaultValue{
									ValueType: &typeStr,
									Value:     &defaultValue,
								},
							},
						},
					},
					{
						Label:     "thing3",
						FieldName: "thing.3",
						Required:  true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeInt,
							Value:     &small,
						},
						Type: &typeInput,
						Validator: model.UIRangeValidator{
							Type:    "regex",
							Message: "Number must be between 0 and 5",

							Min: 0,
							Max: 5,
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
			{
				Columns: &twoColumns,
				Label:   "Range yo",
				Fields: []*model.UIDynamicField{
					{
						Label:     "small",
						FieldName: "some.value.min",
						Required:  true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeInt,
							Value:     &small,
						},
						Type: &typeInput,
						Validator: model.UIRangeValidator{
							Type:    "regex",
							Message: "Number must be between 0 and 1000",

							Min: 0,
							Max: 1000,
						},
					},
					{
						Label:     "large",
						FieldName: "some.value.max",
						Required:  true,
						DefaultValue: &model.UIDefaultValue{
							ValueType: &typeInt,
							Value:     &large,
						},
						Type: &typeInput,
						Validator: model.UIRangeValidator{
							Type:    "regex",
							Message: "Number must be between 0 and 1000",

							Min: 0,
							Max: 1000,
						},
					},
				},
				Conditions: []*model.UIDynamicCondition{
					{
						FieldName: "thing.1",
						Operator:  &nnull,
					},
					{
						FieldName: "thing.3",
						Operator:  &gt,
						ExpectedValue: &model.UIDefaultValue{
							ValueType: &typeInt,
							Value:     &gtNum,
						},
					},
				},
			},
		},
	}
	return app, nil
}
