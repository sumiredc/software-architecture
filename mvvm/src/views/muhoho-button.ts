import { refElements } from "../helpers/ref-element";
import { ViewMultiI } from "./view-i";

interface ElementDataset {
  dataset: {
    muhohoEye?: string;
    muhohoMouth?: string;
  };
}

export type MuhohoButtonElement = HTMLButtonElement & ElementDataset;

export interface MuhohoButtonI extends ViewMultiI<MuhohoButtonElement> {}

export class MuhohoButton implements MuhohoButtonI {
  static readonly _SELECTOR = "button[data-action=muhoho]";

  private readonly _elements: NodeListOf<MuhohoButtonElement>;

  constructor() {
    this._elements = refElements<MuhohoButtonElement>(MuhohoButton._SELECTOR);
  }

  get elements() {
    return this._elements;
  }
}
