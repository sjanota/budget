export function parse(string) {
  const [integer, decimal] = Number(string)
    .toFixed(2)
    .split('.');
  return { integer: Number(integer), decimal: Number(decimal) };
}

export function format(moneyAmount) {
  if (!(typeof moneyAmount === 'object')) {
    return moneyAmount;
  }
  return `${moneyAmount.integer}.${moneyAmount.decimal}`;
}

export function zero() {
  return { integer: 0, decimal: 0 };
}

export function add(e1, e2) {
  if (!(typeof e1 === 'object')) {
    e1 = parse(e1);
  }
  if (!(typeof e2 === 'object')) {
    e2 = parse(e2);
  }

  const decimal = e1.decimal + e2.decimal;

  return {
    integer: e1.integer + e2.integer + Math.floor(decimal / 100),
    decimal: decimal % 100,
  };
}
