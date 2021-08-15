import * as React from 'react';
import {DashboardLayout} from './DashboardLayout';
import {BrowserNavigationProvider} from './Navigation';
import {Sidebar} from './Sidebar';
import './App.css';

export function App(): React.ReactElement {
  return (
    <BrowserNavigationProvider>
      <DashboardLayout sidebar={<Sidebar />}>
        <MainPlaceholder />
      </DashboardLayout>
    </BrowserNavigationProvider>
  );
}

function MainPlaceholder(): React.ReactElement {
  return (
    <section className="Main_Root">
      <div className="Main_Container">
        <h1 className="App_Title">App</h1>
      </div>
    </section>
  );
}
