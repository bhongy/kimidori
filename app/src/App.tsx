import * as React from 'react';
import {DashboardLayout} from './DashboardLayout';
import {HistoryContextProvider, BrowserHistory} from './History';
import {Sidebar} from './Sidebar';
import './App.css';

function onPathChange(path: string): void {
  console.log(`path: ${path}`);
}

export function App(): React.ReactElement {
  const history = new BrowserHistory();
  React.useEffect(() => {
    history.listen(onPathChange);
    return () => history.unlisten(onPathChange)
  })
  return (
    <HistoryContextProvider history={history}>
      <DashboardLayout sidebar={<Sidebar />}>
        <MainPlaceholder />
      </DashboardLayout>
    </HistoryContextProvider>
  );
}

function MainPlaceholder(): React.ReactElement {
  const page = 'https://thanik.me';
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
