import React from "react";
import AppConnector from "./AppConnector";

export default {
    title: "App Connector Info",
    component: AppConnector,
};

const Template = (args) => <AppConnector {...args} />;

export const Default = Template.bind({});
Default.args = {
    "data": {
      "getAppCrowdStrike": {
        "id": "AppCrowdStrike",
        "name": "CrowdStrike Falcon API",
        "description": "Cybersecurity’s AI-native platform for the XDR era",
        "docUrl": "https://docs.corelight.com/docs/sensor/corelight-update/policyConfiguration/sources/integrations/crowdStrike.html#crowdstrike",
        "logo": {
          "url": "https://it.lbl.gov/wp-content/uploads/sites/18/2023/02/crowdstrike-logo.png",
          "alt": "CrowdStrike Falcon",
          "__typename": "UIImage"
        },
        "carousel": {
          "items": [
            {
              "image": {
                "alt": "Slide 1",
                "url": "https://i.redd.it/h0d3v4y8xg221.jpg",
                "__typename": "UIImage"
              },
              "text": "Slide 1",
              "__typename": "UICarouselItem"
            },
            {
              "image": {
                "alt": "Slide 2",
                "url": "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/i/45fb71b4-b618-45ed-b0ef-ddf151262ba5/dc3zd7q-4d0b022a-07db-444f-988e-63abfc2d2a71.png",
                "__typename": "UIImage"
              },
              "text": "Slide 2",
              "__typename": "UICarouselItem"
            }
          ],
          "__typename": "UICarousel"
        },
        "tags": {
          "chips": [
            "HOST_ENRICHMENT",
            "CVE_VULN"
          ],
          "__typename": "AppConnectorChip"
        },
        "items": [
          {
            "columns": 1,
            "label": "First Step",
            "description": "Lots of additional info or tooltip could be displayed here",
            "fields": [
              {
                "type": "INPUT",
                "label": "thing1",
                "fieldName": "thing.1",
                "description": "Lots of additional info or tooltip could be displayed here",
                "required": true,
                "defaultValue": {
                  "valueType": "STRING",
                  "value": "dee-fault"
                },
                "validator": {
                  "type": "regex",
                  "message": "thing1 must be in the following format: word1",
                  "pattern": "\\w\\d+"
                },
                "options": [],
                "__typename": "UIDynamicField"
              },
              {
                "type": "SELECT",
                "label": "thing2",
                "fieldName": "",
                "description": "Lots of additional info or tooltip could be displayed here",
                "required": true,
                "defaultValue": {
                  "valueType": "STRING",
                  "value": "dee-fault"
                },
                "validator": null,
                "options": [
                  {
                    "label": "Option 1",
                    "value": {
                      "valueType": "STRING",
                      "value": "dee-fault"
                    }
                  },
                  {
                    "label": "Option 2",
                    "value": {
                      "valueType": "STRING",
                      "value": "dee-fault"
                    }
                  }
                ],
                "__typename": "UIDynamicField"
              },
              {
                "type": "INPUT",
                "label": "thing3",
                "fieldName": "thing.3",
                "description": null,
                "required": true,
                "defaultValue": {
                  "valueType": "INT",
                  "value": "10"
                },
                "validator": {
                  "type": "regex",
                  "message": "Number must be between 0 and 5",
                  "min": 0,
                  "max": 5
                },
                "options": [],
                "__typename": "UIDynamicField"
              }
            ],
            "conditions": [],
            "__typename": "UIDynamicFieldSet"
          },
          {
            "columns": 1,
            "label": "Second Step, ahh yeah",
            "description": "Lots of additional info or tooltip could be displayed here",
            "fields": [],
            "conditions": [],
            "__typename": "UIDynamicFieldSet"
          },
          {
            "columns": 2,
            "label": "Range yo",
            "description": null,
            "fields": [
              {
                "type": "INPUT",
                "label": "small",
                "fieldName": "some.value.min",
                "description": null,
                "required": true,
                "defaultValue": {
                  "valueType": "INT",
                  "value": "10"
                },
                "validator": {
                  "type": "regex",
                  "message": "Number must be between 0 and 1000",
                  "min": 0,
                  "max": 1000
                },
                "options": [],
                "__typename": "UIDynamicField"
              },
              {
                "type": "INPUT",
                "label": "large",
                "fieldName": "some.value.max",
                "description": null,
                "required": true,
                "defaultValue": {
                  "valueType": "INT",
                  "value": "20"
                },
                "validator": {
                  "type": "regex",
                  "message": "Number must be between 0 and 1000",
                  "min": 0,
                  "max": 1000
                },
                "options": [],
                "__typename": "UIDynamicField"
              }
            ],
            "conditions": [
              {
                "fieldName": "thing.1",
                "operator": "NNULL",
                "expectedValue": null,
                "__typename": "UIDynamicCondition"
              },
              {
                "fieldName": "thing.3",
                "operator": "GT",
                "expectedValue": {
                  "valueType": "INT",
                  "value": "3"
                },
                "__typename": "UIDynamicCondition"
              }
            ],
            "__typename": "UIDynamicFieldSet"
          }
        ]
      }
    }
  }