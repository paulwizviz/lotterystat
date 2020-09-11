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

import React from 'react';

import PropTypes from 'prop-types';

import { ThemeProvider, createMuiTheme } from '@material-ui/core';

import { red } from '@material-ui/core/colors';

const theme = createMuiTheme({
    palette: {
        primary: red,
        type: 'light'
    },
    typography: {
        useNextVariants: true,
    },
});

const RedTheme = (props) => {
    const { children } = props;
    return (
        <ThemeProvider theme={theme}>{children}</ThemeProvider>
    );
};

RedTheme.propTypes = {
    children: PropTypes.object
};

export default RedTheme;