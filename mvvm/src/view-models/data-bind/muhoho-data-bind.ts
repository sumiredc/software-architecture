import {
  MuhohoButton,
  MuhohoButtonElement,
  MuhohoButtonI,
} from "../../views/muhoho-button";
import { Display } from "../../views/display";
import { DataBindI } from "./data-bind-i";
import { MuhohoViewModel, MuhohoViewModelI } from "../muhoho-view-model";

export class MuhohoDataBind implements DataBindI {
  static new() {
    const muhohoButton = new MuhohoButton();
    const display = new Display();
    const viewModel = new MuhohoViewModel(display);

    return new MuhohoDataBind(viewModel, muhohoButton);
  }

  private constructor(
    private readonly _viewModel: MuhohoViewModelI,
    private readonly _muhohoButton: MuhohoButtonI
  ) {}

  bind() {
    this._muhohoButton.elements.forEach((el) => this._attachEvent(el));
  }

  private _attachEvent(el: MuhohoButtonElement) {
    el.addEventListener("click", (ev: MouseEvent) =>
      this._viewModel.execute(ev, el.dataset.muhohoEye, el.dataset.muhohoMouth)
    );
  }
}
