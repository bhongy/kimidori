import * as React from 'react';
import {History, NoopHistory} from './History';

const defaultHistory = new NoopHistory();
const HistoryContext = React.createContext<History>(defaultHistory);

export function HistoryContextProvider({
  children,
  history,
}: {
  children: React.ReactNode;
  history: History;
}): React.ReactElement {
  return (
    <HistoryContext.Provider value={history}>
      {children}
    </HistoryContext.Provider>
  );
}

export function useHistory(): History {
  return React.useContext(HistoryContext);
}
