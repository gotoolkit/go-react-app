/**
 * Created by llitfkitfk on 5/3/17.
 */
import React from 'react';
import {
    Route
} from 'react-router-dom';
import App from "./components/App";
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import {green100, green500, green700} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

const muiTheme = getMuiTheme({
    palette: {
        primary1Color: green500,
        primary2Color: green700,
        primary3Color: green100,
    },
}, {
    avatar: {
        borderColor: null,
    },
});

const AppRoutes = (
    <Route path="/">
        <MuiThemeProvider muiTheme={muiTheme}>
            <App/>
        </MuiThemeProvider>
    </Route>
);

export default AppRoutes;
