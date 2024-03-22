import { Muhoho } from "../models/muhoho";
import { DisplayI } from "../views/display";
import { ViewModelI } from "./view-model-i";

export interface MuhohoViewModelI extends ViewModelI {
  execute(ev: MouseEvent, eye?: string, mouth?: string): void;
}

export class MuhohoViewModel implements MuhohoViewModelI {
  constructor(private _display: DisplayI) {}

  // ビジネスロジックの呼び出し → View へ 値を渡す
  public execute(_: MouseEvent, eye?: string, mouth?: string) {
    const muhoho = new Muhoho(eye, mouth);
    this._display.value = muhoho.face;
  }
}
