import { ElementNotFoundError } from "../errors/element-not-found-error";

/**
 * @param {string} selector
 * @return {E}
 * @throws {ElementNotFoundError}
 */
export function refElement<E extends Element>(selector: string): E {
  const el = document.querySelector<E>(selector);
  if (el === null) {
    throw new ElementNotFoundError(selector);
  }

  return el;
}

/**
 * @param {string} selector
 * @return {NodeListOf<E>}
 * @throws {ElementNotFoundError}
 */
export function refElements<E extends Element>(
  selector: string
): NodeListOf<E> {
  return document.querySelectorAll<E>(selector);
}

/**
 * @param {string} id
 * @return {E}
 * @throws {ElementNotFoundError}
 */
export function refElementById<E extends HTMLElement>(id: string): E {
  const el = <E>document.getElementById(id);
  if (el === null) {
    throw new ElementNotFoundError(id);
  }

  return el;
}
