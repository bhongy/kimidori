import * as React from 'react';
import * as navigate from './navigate';

const Navigation = React.createContext<navigate.Navigate>(navigate.noop);

export function useNavigation(): navigate.Navigate {
  return React.useContext(Navigation);
}

export function BrowserNavigationProvider({
  children,
}: {
  children: React.ReactNode;
}): React.ReactElement {
  return (
    <Navigation.Provider value={navigate.browser}>
      {children}
    </Navigation.Provider>
  );
}
