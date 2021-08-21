interface Handler<T> {
  (v: T): void;
}

export class Events<T> {
  private handlers: Set<Handler<T>>;

  constructor() {
    this.handlers = new Set();
  }

  get length(): number {
    return this.handlers.size;
  }

  push(h: Handler<T>) {
    this.handlers.add(h);
  }

  pop(h: Handler<T>) {
    this.handlers.delete(h);
  }

  call(v: T) {
    this.handlers.forEach((h) => h(v));
  }
}
