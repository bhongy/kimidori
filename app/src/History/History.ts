export interface Navigate {
  (to: string, options?: NavigateOptions): void;
}

export interface NavigateOptions {
  state?: {};
}

// Listener is called when history state changes
export interface Listener {
  (path: string): void;
}

export interface History {
  listen(l: Listener): void;
  unlisten(l: Listener): void;
  push: Navigate;
}

export class NoopHistory implements History {
  listen() {}
  unlisten() {}
  push() {}
}
