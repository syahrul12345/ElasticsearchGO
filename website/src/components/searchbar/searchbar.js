import React from 'react'
import { useState } from 'react'
import {TextField} from '@material-ui/core'

export default function Searchbar(props) {
    return(
        <h1>
            <TextField
            label="Search for flight"
            id="Search flight"
            variant="outlined"
            onChange={props.textFieldChangeHandler}
            style={{width:'60vw'}}
            />
        </h1>
    )
}