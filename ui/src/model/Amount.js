export default class Amount {
  static parse(string) {
    if (string === null) {
      return null;
    }
    return Math.round(Number(string) * 100);
  }

  static zero() {
    return 0;
  }

  static format(amount) {
    return amount == null ? '' : `${(amount / 100).toFixed(2)}`;
  }
}
