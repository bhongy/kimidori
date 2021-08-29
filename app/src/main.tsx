import * as React from 'react';
import * as ReactDOM from 'react-dom';
import {App} from './App';

const container = document.getElementById('react-root');
if (container) {
  ReactDOM.render(<App />, container);
}
