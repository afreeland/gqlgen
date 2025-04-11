import React from 'react';
import { Box, Typography, Chip } from '@mui/material';
import DynamicFieldSet from './DynamicFieldSet';
import Carousel from './Carousel';

const AppConnector = ({ data }) => {
  const { getAppCrowdStrike: {name, description, docUrl, logo, tags, items} } = data;

  return (
    <Box sx={{ p: 2 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
        <img src={logo.url} alt={logo.alt} width={100} style={{ marginRight: 16 }} />
        <Box>
          <Typography variant="h4">{name}</Typography>
          <Typography variant="body1">{description}</Typography>
          <Typography variant="body2" color="primary"><a href={docUrl} target="_blank" rel="noopener noreferrer">Documentation</a></Typography>
        </Box>
      </Box>

      <Box sx={{ mb: 2 }}>
        <Typography variant="h6">Tags:</Typography>
        {tags.chips.map((chip, index) => (
          <Chip key={index} label={chip} sx={{ mr: 1, mt: 1 }} />
        ))}
      </Box>

      <Box>
        {items.map((item, index) => (
          <DynamicFieldSet key={index} item={item} />
        ))}
      </Box>
    </Box>
  );
};

export default AppConnector;