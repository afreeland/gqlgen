import React from 'react';
import { TextField, Select, MenuItem, FormControl, InputLabel } from '@mui/material';

const DynamicField = ({ field }) => {
  switch (field.type) {
    case 'INPUT':
      return (
        <TextField
          label={field.label}
          defaultValue={field.defaultValue.value}
          required={field.required}
          helperText={field.description}
        />
      );
    case 'SELECT':
      return (
        <FormControl fullWidth>
          <InputLabel>{field.label}</InputLabel>
          <Select defaultValue={field.defaultValue.value}>
            {field.options && field.options.map((option, index) => (
              <MenuItem key={index} value={option.value}>
                {option.label}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      );
    default:
      return null;
  }
};

export default DynamicField;