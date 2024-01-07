type AsArray<T> = T extends unknown[] ? T : never;

class EventEmitter<T> {
  private events: Map<keyof T, ((...args: AsArray<T[keyof T]>) => void)[]>;

  public constructor() {
    this.events = new Map();
  }

  public on<K extends keyof T>(
    event: K,
    cb: (...args: AsArray<T[K]>) => void
  ): void {
    if (!this.events.has(event)) {
      this.events.set(event, [cb]);
    }
  }

  public off<K extends keyof T>(
    event: K,
    cb: (...args: AsArray<T[K]>) => void
  ): void {
    if (!this.events.has(event)) return;

    this.events.set(
      event,
      this.events.get(event)?.filter(fn => fn !== cb) ?? []
    );
  }

  public emit<K extends keyof T>(event: K, ...args: AsArray<T[K]>): void {
    if (!this.events.has(event)) return;

    for (let cb of this.events.get(event) ?? []) {
      cb(...args);
    }
  }
}

export { EventEmitter };
