import React, {Component} from 'react';
import AppBar from 'material-ui/AppBar';
import IconButton from "material-ui/IconButton";
import AppNavDrawer from "./AppNavDrawer";
import DocumentTitle from "react-document-title";
class App extends Component {

    constructor(props) {
        super(props);
        this.state = {navDrawerOpen: false};
    }

    handleToggle = () => this.setState({navDrawerOpen: !this.state.navDrawerOpen});
    handleChangeRequestNavDrawer = (open) => this.setState({navDrawerOpen: open,});

    render() {
        return (
            <div>
                <DocumentTitle title="Go React App"/>

                <AppBar title="Go React"
                        onLeftIconButtonTouchTap={this.handleToggle}
                        iconElementRight={
                            <IconButton iconClassName="muidocs-icon-custom-github"
                                        href="https://github.com/llitfkitfk"/>
                        }
                />
                <AppNavDrawer open={this.state.navDrawerOpen}
                              handleChangeRequestNavDrawer={this.handleChangeRequestNavDrawer}/>
            </div>
        );
    }
}

export default App;
