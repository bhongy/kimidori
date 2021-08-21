import {Events} from './Events';
import {History, Listener, Navigate, NavigateOptions} from './History';

export class BrowserHistory implements History {
  #events: Events<string> = new Events();

  listen(l: Listener): void {
    this.#events.push(l);
  }

  unlisten(l: Listener): void {
    this.#events.pop(l);
  }

  push(to: string, opts: NavigateOptions = {}): void {
    window.history.pushState(opts.state, '', to);
    this.#events.call(to);
  }
}
