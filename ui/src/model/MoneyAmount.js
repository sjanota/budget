export function parse(string) {
  const [integer, decimal] = Number(string.totalBalance)
    .toFixed(2)
    .split('.');
  return { integer, decimal };
}

export function format(moneyAmount) {
  return `${moneyAmount.integer}.${moneyAmount.decimal}`;
}
