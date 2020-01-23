import React from 'react';
import { useState,useEffect } from 'react'
import { Grid, Button, Typography, Card, CardContent } from '@material-ui/core/'
import SearchBar from '../../components/searchbar'
import axios from 'axios'

function DashBoard(props) {
  const [search,setSearch] = useState({
    country:''
  })
  const [routes,setRoutes] = useState([])
  const onHandleChange = (event) => {
    setSearch({...search,"country":event.target.value})
  }
  const startSearch = () => {
    axios.post("/api/v1/search",search)
      .then((res) => {
        setRoutes(res.data.routeList.Routes)
      })
      .catch((err) => {
        console.log(err)
      })
  }
  useEffect(() => {

  },[routes])
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
      <Grid item xs={12}>
        <Grid 
        container 
        direction="row"
        spacing={2}
        style={{paddingLeft:'5%',paddingRight:'5%'}}>
          {routes.map((route) => {
            return(
              <Grid item xs={3}>
                <Card>
                  <CardContent>
                    <Typography variant="h3"> {route.DestinationId} </Typography>
                    <Typography variant="h6"> SGD{route.Price} </Typography>
                  </CardContent>
                </Card>
              </Grid>
            )
          })}
        </Grid>
      </Grid>
    </Grid>
  );
}

export default DashBoard;
