export default class Amount {
  static parse(string) {
    if (string === null) {
      return null;
    }
    const [integer, decimal] = Number(string)
      .toFixed(2)
      .split('.');
    return { integer: Number(integer), decimal: Number(decimal) };
  }

  static zero() {
    return { integer: 0, decimal: 0 };
  }

  static add(e1, e2) {
    if (!(typeof e1 === 'object')) {
      e1 = Amount.parse(e1);
    }
    if (!(typeof e2 === 'object')) {
      e2 = Amount.parse(e2);
    }

    const decimal = e1.decimal + e2.decimal;

    return {
      integer: e1.integer + e2.integer + Math.floor(decimal / 100),
      decimal: decimal % 100,
    };
  }

  static format(amount) {
    return amount == null
      ? ''
      : `${amount.integer}.${String(amount.decimal).padStart(2, '0')}`;
  }
}
