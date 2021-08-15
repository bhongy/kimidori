export interface Navigate {
  (to: string, options?: NavigateOptions): void;
}

interface NavigateOptions {
  state?: {};
}

export function noop() {}

export function browser(to: string, opts: NavigateOptions = {}): void {
  window.history.pushState(opts.state, '', to);
}
