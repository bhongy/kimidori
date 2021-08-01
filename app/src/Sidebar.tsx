import * as React from 'react';
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
    <div className="Sidebar_Nav">
      <div style={{height: '120vh'}}>
        <ul>
          <li style={{width: '300px'}}>Home</li>
          <li>Profile</li>
          <li>Investment</li>
        </ul>
      </div>
    </div>
  );
}

function Footer(): React.ReactElement {
  return (
    <div className="Sidebar_Footer">
      <MenuItem>Sidebar Footer</MenuItem>
    </div>
  );
}

function MenuItem({
  children,
}: {
  children: string | React.ReactElement;
}): React.ReactElement {
  return <a className="Sidebar_MenuItem">{children}</a>;
}
