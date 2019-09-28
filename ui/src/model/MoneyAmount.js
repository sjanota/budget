export function parse(string) {
  const [integer, decimal] = Number(string.totalBalance)
    .toFixed(2)
    .split('.');
  return { integer, decimal };
}

export function format(moneyAmount) {
  return `${moneyAmount.integer}.${moneyAmount.decimal}`;
}

export function zero() {
  return { integer: 0, decimal: 0 };
}
