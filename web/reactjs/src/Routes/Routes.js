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
import React from 'react';
import { BrowserRouter as Router, Switch, Redirect } from 'react-router-dom';


import RouteWithLayout from './RouteWithLayout';

import {
    Euromil as EuromilView,
    NotFound as NotFoundView,
    MinimalLayout,
    MainLayout as MainLayoutContainer,
} from '../modules';

const Routes = () => {
    return (
        <Router>
            <Switch>        
                <Redirect exact from="/" to="/euromil"/>
                <RouteWithLayout component={EuromilView} exact layout={MainLayoutContainer} path="/euromil"/>
                <RouteWithLayout component={NotFoundView} exact layout={MinimalLayout} path="/not-found"/>
                <Redirect to="/not-found" />
            </Switch>
        </Router>
    );
};

export default Routes;