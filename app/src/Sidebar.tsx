import * as React from 'react';
import {BlockButton} from './Components/Button';
import './Sidebar.css';

// Sidebar is a placeholder used temporarily to check
// DashboardLayout's sidebar during the initial development
export function Sidebar(): React.ReactElement {
  return (
    <div className="Sidebar_Root">
      <Nav />
      <Footer />
    </div>
  );
}

function Nav(): React.ReactElement {
  return (
    <nav className="Sidebar_Nav" role="navigation">
      <div style={{height: '120vh'}}>
        <Menu>
          <MenuItem>Home</MenuItem>
          <MenuItem>Profile</MenuItem>
          <MenuItem>Investment</MenuItem>
          <MenuItem>Settings</MenuItem>
        </Menu>
      </div>
    </nav>
  );
}

function Footer(): React.ReactElement {
  return (
    <footer className="Sidebar_Footer">
      <MenuItem>Sidebar Footer</MenuItem>
    </footer>
  );
}

function Menu({children}: {children: React.ReactNode}): React.ReactElement {
  return <div className="Sidebar_Menu">{children}</div>;
}

function MenuItem({children}: {children: React.ReactNode}): React.ReactElement {
  return <BlockButton className="Sidebar_MenuItem">{children}</BlockButton>;
}
