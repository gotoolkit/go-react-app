import React from 'react';
import ReactDOM from 'react-dom';
import AppRoutes from './AppRoutes';
import injectTapEventPlugin from 'react-tap-event-plugin';
import { BrowserRouter as Router} from 'react-router-dom';

injectTapEventPlugin();


ReactDOM.render(
    <Router>
        {AppRoutes}
    </Router>,
    document.getElementById('root')
);
