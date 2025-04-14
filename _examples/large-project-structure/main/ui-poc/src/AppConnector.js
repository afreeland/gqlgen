import React from "react";
import { Box, Grid, Typography, Chip, Link } from "@mui/material";
import Carousel from "./Carousel";
import DynamicFieldSet from "./DynamicFieldSet";

const AppConnector = ({ data }) => {
    const {
        getAppCrowdStrike: {
            name,
            description,
            docUrl,
            logo,
            tags,
            items,
            carousel,
        },
    } = data;

    return (
        <Box sx={{ p: 2 }}>
            <Grid container spacing={3}>
                {/* Left Section: Name, Description, Logo */}
                <Grid size={5}>
                    <Box sx={{ mb: 2 }}>
                        <Grid container>
                            <Grid item xs={5}>
                                <img
                                    src={logo.url}
                                    alt={logo.alt}
                                    width="100%"
                                    style={{
                                        maxHeight: "150px",
                                        objectFit: "contain",
                                        float: "left",
                                    }}
                                />
                            </Grid>
                            <Grid item xs={7}>
                                <Typography variant="h4" component="h1">
                                    {name}
                                </Typography>
                                <Typography
                                    variant="body1"
                                    color="textSecondary"
                                    sx={{ mb: 1 }}
                                >
                                    {description}
                                </Typography>
                                <Typography variant="body2">
                                    Documentation:{" "}
                                    <Link
                                        href={docUrl}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        color="primary"
                                    >
                                        {docUrl}
                                    </Link>
                                </Typography>
                            </Grid>
                        </Grid>
                        <Box sx={{ mt: 2 }}>
                            <Typography variant="h6">Tags:</Typography>
                            <Box
                                sx={{
                                    display: "flex",
                                    flexWrap: "wrap",
                                    gap: 1,
                                }}
                            >
                                {tags.chips.map((chip, index) => (
                                    <Chip
                                        key={index}
                                        label={chip}
                                        variant="outlined"
                                    />
                                ))}
                            </Box>
                        </Box>
                    </Box>
                </Grid>

                {/* Right Section: Carousel */}
                <Grid item size={7}>
                    <Carousel carousel={carousel} />
                </Grid>
            </Grid>

            {/* Dynamic Fields Section */}
            <Box sx={{ mt: 4 }}>
                {items.map((item, index) => (
                    <DynamicFieldSet key={index} item={item} />
                ))}
            </Box>
        </Box>
    );
};

export default AppConnector;
