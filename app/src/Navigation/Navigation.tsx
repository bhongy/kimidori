import * as React from 'react';
import * as navigate from './navigate';

export const Navigation = React.createContext<navigate.Navigate>(navigate.noop);
