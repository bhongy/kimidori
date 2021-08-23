import * as React from 'react';
import {DashboardLayout} from './DashboardLayout';
import {HistoryContextProvider, BrowserHistory, useHistory} from './History';
import {Sidebar} from './Sidebar';
import './App.css';

const history = new BrowserHistory();

export function App(): React.ReactElement {
  return (
    <HistoryContextProvider history={history}>
      <DashboardLayout sidebar={<Sidebar />}>
        <MainPlaceholder />
      </DashboardLayout>
    </HistoryContextProvider>
  );
}

function MainPlaceholder(): React.ReactElement {
  const location = useLocation();
  const page = location === '/profile' ? 'https://thanik.me' : 'about:blank';
  return (
    <iframe
      allow="clipboard-write"
      className="Main"
      frameBorder="0"
      sandbox=""
      src={page}
    ></iframe>
  );
}

function useLocation(): string {
  const [location, setLocation] = React.useState(window.location.pathname);
  const history = useHistory();
  React.useEffect(() => {
    const unlisten = history.listen(({to}) => setLocation(to));
    return unlisten;
  }, []);
  return location;
}
