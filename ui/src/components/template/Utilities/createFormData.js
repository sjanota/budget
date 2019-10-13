class SimpleFormData {
  constructor($init, $process) {
    this.init = $init;
    this.process = $process || (v => v);
    this.current = null;
  }

  value() {
    if (this.current === null) {
      return null;
    }
    if (Array.isArray(this.current.value)) {
      return this.current.value.map(v => v.value());
    }
    return this.process(this.current.value);
  }

  changed() {
    return (
      (this.current === null && this.init !== null) ||
      this.current.value !== this.init
    );
  }
}

class CompositeFormData {
  constructor(model) {
    this._model = model;
    Object.keys(model).forEach(key => (this[key] = createFormData(model[key])));
  }

  changed() {
    return Object.keys(this._model).some(k => this[k].changed());
  }

  value() {
    return Object.keys(this._model).reduce((acc, key) => {
      if (!this[key].changed()) {
        return acc;
      }
      return { ...acc, [key]: this[key].value() };
    }, {});
  }
}

export function createFormData(model) {
  if (Object.prototype.hasOwnProperty.call(model, '$init')) {
    return new SimpleFormData(model.$init, model.$process);
  }
  return new CompositeFormData(model);
}
