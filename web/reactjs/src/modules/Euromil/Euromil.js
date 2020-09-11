// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// React
import React, {useReducer, useEffect} from 'react';

// Materials UI
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';

import axios from 'axios';

// Dev express chart
import {
    Chart,
    BarSeries,
    Title,
    ArgumentAxis,
    ValueAxis,
} from '@devexpress/dx-react-chart-material-ui';

// Code body
const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        marginTop: 20,
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary
    }
}));

const initialState = {
    loading: true,
    error: '',
    data: {
        stars: [],
        balls:[],
    }
};

const reducer = (state = initialState, action) => {
    switch(action.type){
    case 'FETCH_SUCCESS':
        return {
            ...state,
            loading: false,
            data: action.payload
        };
    case 'FETCH_ERROR':
        return {
            ...state,
            loading: false,
            error: action.payload,
        };
    default:
        return state;
    }
};

const fetcData = async (dispatch) => {
    const resp = await axios.get('/api/euromil');
    try{
        dispatch({
            type: 'FETCH_SUCCESS',
            payload: resp.data
        });
    } catch(err){
        dispatch({
            type: 'FETCH_ERROR',
            payload: resp.data
        });
    } 
};

const Euromil = () => {

    const classes = useStyles();

    const [state, dispatch] = useReducer(reducer, initialState);

    useEffect(() => {
        fetcData(dispatch);
    }, []);

    return (
        <div className={classes.root}>
             <Button onClick={ () => {fetcData(dispatch); }}>Refresh</Button>
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <Paper>
                        <Chart data={state.data.balls}>
                            <ArgumentAxis />
                            <ValueAxis max={45} />
                            <BarSeries
                                valueField="frequency"
                                argumentField="number"
                            />
                            <Title text="Ball numbers" />
                        </Chart>
                    </Paper>
                </Grid>
                <Grid item xs={12}>
                    <Paper>
                        <Chart data={state.data.stars}>
                            <ArgumentAxis />
                            <ValueAxis max={45} />
                            <BarSeries
                                valueField="frequency"
                                argumentField="number"
                            />
                            <Title text="Star numbers" />
                        </Chart>
                    </Paper>
                </Grid>
            </Grid>
        </div>
    );
};

export default Euromil;
