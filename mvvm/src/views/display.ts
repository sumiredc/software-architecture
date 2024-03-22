import { refElementById } from "../helpers/ref-element";
import { ViewSingleI } from "./view-i";

export interface DisplayI extends ViewSingleI<HTMLDivElement> {
  set value(v: string);
}

export class Display implements DisplayI {
  private static readonly _ID = "display";

  private readonly _element: HTMLDivElement;

  constructor() {
    this._element = refElementById(Display._ID);
  }

  get element() {
    return this._element;
  }

  set value(v: string) {
    if (this._element.innerText !== "") {
      v = `\n${v}`;
    }
    this._element.innerText += v;
  }
}
