import * as React from 'react';
import {DashboardLayout} from './DashboardLayout';
import {Navigation, navigate} from './Navigation';
import {Sidebar} from './Sidebar';
import './App.css';

export function App(): React.ReactElement {
  return (
    <Navigation.Provider value={navigate.browser}>
      <DashboardLayout sidebar={<Sidebar />}>
        <MainPlaceholder />
      </DashboardLayout>
    </Navigation.Provider>
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
