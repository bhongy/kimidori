import * as React from 'react';
import {DashboardLayout} from './DashboardLayout';
import {Sidebar} from './Sidebar';
import './App.css';

export function App(): React.ReactElement {
  return (
    <DashboardLayout sidebar={<Sidebar />}>
      <MainPlaceholder />
    </DashboardLayout>
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
