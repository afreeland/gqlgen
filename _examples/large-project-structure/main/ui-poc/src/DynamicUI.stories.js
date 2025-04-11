import React from 'react';
import DynamicFieldSet from './DynamicFieldSet';

export default {
  title: 'Dynamic UI',
  component: DynamicFieldSet,
};

const Template = (args) => <DynamicFieldSet {...args} />;

export const Default = Template.bind({});
Default.args = {
  item: {
    label: 'First Step',
    description: 'Lots of additional info or tooltip could be displayed here',
    columns: 1,
    fields: [
      {
        type: 'INPUT',
        label: 'thing1',
        fieldName: 'thing.1',
        description: 'Lots of additional info or tooltip could be displayed here',
        required: true,
        defaultValue: {
          valueType: 'STRING',
          value: 'dee-fault'
        },
        validator: {
          type: 'regex',
          message: 'thing1 must be in the following format: word1',
          pattern: '\\w\\d+'
        }
      },
      {
        type: 'SELECT',
        label: 'thing2',
        fieldName: '',
        description: 'Lots of additional info or tooltip could be displayed here',
        required: true,
        defaultValue: {
          valueType: 'STRING',
          value: 'dee-fault'
        },
        options: [
          { label: 'Option 1', value: 'option1' },
          { label: 'Option 2', value: 'option2' },
        ]
      }
    ]
  }
};