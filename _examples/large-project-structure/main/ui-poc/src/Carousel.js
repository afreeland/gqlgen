import React from 'react';
import Slider from 'react-slick';
import { Box, Typography } from '@mui/material';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

const Carousel = ({ carousel }) => {
  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    autoplay: true,
    autoplaySpeed: 3000,
  };

  return (
    <Box sx={{ maxWidth: 600, margin: '0 auto', mb: 3 }}>
      <Slider {...settings}>
        {carousel.items.map((item, index) => (
          <Box key={index} sx={{ position: 'relative', textAlign: 'center' }}>
            <img src={item.image.url} alt={item.image.alt} style={{ width: '100%', height: 'auto' }} />
            <Typography
              variant="h6"
              sx={{
                position: 'absolute',
                bottom: 10,
                left: '50%',
                transform: 'translateX(-50%)',
                color: 'white',
                backgroundColor: 'rgba(0, 0, 0, 0.5)',
                padding: '5px 10px',
                borderRadius: 1,
              }}
            >
              {item.text}
            </Typography>
          </Box>
        ))}
      </Slider>
    </Box>
  );
};

export default Carousel;