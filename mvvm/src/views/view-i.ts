export interface ViewSingleI<E extends Element = Element> {
  get element(): E;
}

export interface ViewMultiI<E extends Node = Node> {
  get elements(): NodeListOf<E>;
}
