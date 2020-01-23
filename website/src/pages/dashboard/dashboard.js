import React from 'react';
import { useState } from 'react'
import { Grid, Button, Typography } from '@material-ui/core/'
import SearchBar from '../../components/searchbar'

function DashBoard(props) {
  const [search,setSearch] = useState('')
  const onHandleChange = (event) => {
    setSearch(event.target.value)
  }
  const startSearch = () => {
    console.log(`Searching for ${search}`)
  }
  return (
    <Grid
    container
    direction="column"
    justify="center"
    alignItems="center"
    alignContent="center"
    style={{minHeight:'60vh'}}
    >
      <Grid item xs={12}>
        <SearchBar textFieldChangeHandler={onHandleChange}/>
      </Grid>
      <Grid item xs={12}>
        <Typography variant="subtitle2"> {props.globalState.token} </Typography>
      </Grid>
      <Grid item xs={12}>
        <Button onClick={startSearch}> Search </Button>
      </Grid>
    </Grid>
  );
}

export default DashBoard;
