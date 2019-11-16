export default class Amount {
  static parse(string) {
    if (string === null || string === '') {
      return null;
    }
    return Math.round(Number(string) * 100);
  }

  static zero() {
    return 0;
  }

  static format(amount) {
    if (amount === null) {
      return null;
    }
    var parts = (amount / 100).toFixed(2).split('.');
    parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ' ');
    return parts.join('.');
  }
}
