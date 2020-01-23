import React from 'react'
import { useState,useEffect } from 'react'
import { Grid,TextField,Button } from '@material-ui/core';
import axios from 'axios'
import { Redirect } from 'react-router-dom';

export default function CreateAccount(props) {
    const [redirect,setRedirect] = useState({
        status:false,
        token:''
    })
    const [toLogin,changeRedirect] = useState(false)
    const [userInfo,setUserInfo] = useState({
        email:'',
        username:'',
        password:'',
        token:'',
    });
    const handleChange = (input) => event =>{
        setUserInfo({...userInfo,[input]:event.target.value})
    }
    const createAccount = () => {
        axios.post("/api/v1/createAccount",userInfo)
            .then((res) => {
                const token = res.data.account["Token"]
                setRedirect({...redirect,"token":token})
                changeRedirect(true)
            })
            .catch((err) => {
                console.log(err.response.data.message)
            })
    }
    useEffect(() => {
        // Handle redirect
        props.setState({...props.globalState,"token":redirect.token})
    },[toLogin])
    
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
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('email')}/>
            </Grid>
            <Grid item xs={12}>
                <TextField
                    label="Username"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('username')}/>
            </Grid>
            <Grid item xs={12}>
                <TextField
                    label="Password"
                    variant="outlined"
                    style={{width:'50vw',marginBlockEnd:'1vh'}}
                    onChange={handleChange('password')}/>
            </Grid>
            <Grid item xs={12}>
                <Button variant="outlined" onClick={createAccount}> Create Account </Button>
            </Grid>
            {toLogin ? 
                <Redirect to="/dashboard"/>
                : <></>
            }
        </Grid>
    )
}
