/**
 * Created by llitfkitfk on 5/3/17.
 */
import React, {Component} from 'react';
import Drawer from 'material-ui/Drawer';
import MenuItem from "material-ui/MenuItem";

class AppNavDrawer extends Component {

    handleClose = () => this.setState({open: false});

    render() {
        const {open, handleChangeRequestNavDrawer} = this.props;
        return (
            <Drawer
                docked={false}
                width={200}
                open={open}
                onRequestChange={handleChangeRequestNavDrawer}>
                <MenuItem onTouchTap={this.handleClose}>Menu Item</MenuItem>
            </Drawer>
        );
    }
}

export default AppNavDrawer;