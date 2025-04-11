import React from 'react';
import DynamicField from './DynamicField';
import { Box, Typography } from '@mui/material';

const DynamicFieldSet = ({ item }) => {
  return (
    <Box sx={{ mb: 2 }}>
      <Typography variant="h6">{item.label}</Typography>
      <Typography variant="body2" sx={{ mb: 1 }}>{item.description}</Typography>
      <Box
        sx={{
          display: 'grid',
          gridTemplateColumns: `repeat(${item.columns}, 1fr)`,
          gap: 2,
        }}
      >
        {item.fields.map((field, index) => (
          <DynamicField key={index} field={field} />
        ))}
      </Box>
    </Box>
  );
};

export default DynamicFieldSet;
