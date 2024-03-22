export interface DisplayI {
  get value(): string;
  set value(value: string);
  clear(): void;
}

export class Display implements DisplayI {
  constructor(private _value: string = "") {}

  get value(): string {
    return this._value;
  }

  set value(value: string) {
    if (this._value !== "") {
      value = `\n${value}`;
    }
    this._value += value;
  }

  clear(): void {
    this._value = "";
  }
}
