export interface MuhohoI {
  get face(): string;
}

export class Muhoho implements MuhohoI {
  private static readonly LEFT_CHEEK = "( ";
  private static readonly RIGHT_CHEEK = ")";
  private static readonly DEFAULT_EYE = "^";
  private static readonly DEFAULT_MOUTH = "ω";

  constructor(
    private readonly eye: string = Muhoho.DEFAULT_EYE,
    private readonly mouth: string = Muhoho.DEFAULT_MOUTH
  ) {}

  get face(): string {
    return (
      Muhoho.LEFT_CHEEK + this.eye + this.mouth + this.eye + Muhoho.RIGHT_CHEEK
    );
  }
}
