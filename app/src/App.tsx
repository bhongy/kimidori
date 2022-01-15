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

// TEMPORARY: stub data from the upstream service discovery api
const pages = new Map([['/profile', 'https://localhost:8020/service/profile']]);

function MainPlaceholder(): React.ReactElement {
  const location = useLocation();
  const page = pages.get(location);
  if (!page) {
    return <NotFound />;
  }
  return (
    <iframe
      allow="clipboard-write"
      className="Main"
      frameBorder="0"
      sandbox="allow-scripts"
      src={page}
    ></iframe>
  );
}

function NotFound(): React.ReactElement {
  return <h1>404: Not Found</h1>;
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
