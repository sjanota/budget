import Day from './Day';

export default class Month {
  constructor(year, month) {
    this.year = year;
    this.month = month;
  }
  static parse(string) {
    const [year, month] = string.split('-');
    return new Month(Number(year), Number(month));
  }

  firstDay() {
    return new Day(this.year, this.month, 1);
  }

  lastDay() {
    const date = new Date(this.year, this.month, 0);
    return new Day(this.year, this.month, date.getDate());
  }
}
