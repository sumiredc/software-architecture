export class ElementNotFoundError extends Error {
  status = 404

  constructor(target: string) {
    super(`Element not found. : ${target}`)
  }
}
