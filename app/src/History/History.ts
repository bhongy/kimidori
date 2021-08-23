export enum Action {
  /**
   * A PUSH indicates a new entry being added to the history stack, such as when
   * a link is clicked and a new page loads. When this happens, all subsequent
   * entries in the stack are lost.
   */
  Push = 'PUSH',
}

export interface Update {
  action: Action;
  to: string;
}

// Listener is called when history state changes
export interface Listener {
  (update: Update): void;
}

export interface History {
  /**
   * listen adds `listener` to be called when the history state changes.
   *
   * @param listener - A function that will be called when the location changes
   * @returns unlisten - A function that may be used to stop listening
   */
  listen(l: Listener): () => void;

  /**
   * Pushes a new location onto the history stack, increasing its length by one.
   * If there were any entries in the stack after the current one, they are
   * lost.
   *
   * @param to - The new URL
   * @param state - Data to associate with the new location
   */
  push(to: string, state?: {}): void;
}
