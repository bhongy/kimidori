import * as React from 'react';
import {BlockButton} from './Components/Button';
import {useHistory} from './History';
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
          <MenuItem href="/">Home</MenuItem>
          <MenuItem href="/profile">Profile</MenuItem>
          <MenuItem href="/investment">Investment</MenuItem>
          <MenuItem href="/settings">Settings</MenuItem>
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

function MenuItem({
  children,
  href,
}: {
  children: React.ReactNode;
  href?: string;
}): React.ReactElement {
const history = useHistory();
  const handleClick = (event: React.MouseEvent) => {
    event.preventDefault();
    href && history.push(href);
  };
  return (
    <BlockButton className="Sidebar_MenuItem" onClick={handleClick}>
      {children}
    </BlockButton>
  );
}
