import {Events} from './Events';
import {Action, History, Listener, Update} from './History';

export class BrowserHistory implements History {
  #events: Events<Update> = new Events();
  #globalHistory: Window['history'] = window.history;

  listen(l: Listener): () => void {
    this.#events.push(l);
    return () => this.#events.pop(l);
  }

  push(to: string, state = {}): void {
    this.#globalHistory.pushState(state, '', to);

    const update = {
      action: Action.Push,
      to,
    };
    this.#events.call(update);
  }
}
