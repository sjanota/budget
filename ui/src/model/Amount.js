export default class Amount {
  static parse(string) {
    const [integer, decimal] = Number(string)
      .toFixed(2)
      .split('.');
    return { integer: Number(integer), decimal: Number(decimal) };
  }

  static format(moneyAmount) {
    if (!(typeof moneyAmount === 'object')) {
      return moneyAmount;
    }
    return `${moneyAmount.integer}.${moneyAmount.decimal}`;
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

  static Formatter(amount) {
    return amount == null
      ? ''
      : `${amount.integer}.${String(amount.decimal).padStart(2, '0')}`;
  }
}
