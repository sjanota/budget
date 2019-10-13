export default class Day {
  constructor(year, month, day) {
    this.year = year;
    this.month = month;
    this.day = day;
  }

  static fromDate(date) {
    return new Date(date.getYear(), date.getMonth(), date.getDate());
  }

  format() {
    return `${this.year}-${String(this.month).padStart(2, '0')}-${String(
      this.day
    ).padStart(2, '0')}`;
  }
}
