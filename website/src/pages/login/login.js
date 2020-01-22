import React from 'react'
import { Grid, TextField, Button, Typography } from '@material-ui/core'

export default function Login() {
    return(
        <Grid
        container
        direction="column"
        justify="center"
        alignItems="center"
        alignContent="center"
        style={{minHeight:'60vh'}}
        >
            <Grid item xs={12}>
                <TextField
                    label="Email"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}/>
            </Grid>
            <Grid item xs={12}>
                <TextField
                    label="Password"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}/>
            </Grid>
            <Grid item xs={12}>
                <Button variant="filled"> LOGIN </Button>
            </Grid>
            <Grid item xs={12}>
                <a href="/changePassword">
                    <Typography variant="subtitle1"> Change your password </Typography>
                </a>
            </Grid>  
        </Grid>
    )
}