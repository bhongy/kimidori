import {History} from './History';

export class NoopHistory implements History {
  listen() {
    return noop;
  }
  push = noop;
}

function noop() {}
