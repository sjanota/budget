import { useState, useRef } from 'react';

function simpleFormData({ $init, $process, $default }, getRoot) {
  const process = $process || (v => v);
  const formData = { current: null };
  const _default =
    typeof $default === 'function' ? () => $default(getRoot()) : () => $default;

  formData.value = () => {
    if (formData.current === null) {
      return null;
    }
    return process(formData.current.value);
  };

  formData.changed = () => {
    return (
      (formData.current === null && $init !== null) ||
      formData.current.value !== $init
    );
  };

  formData.default = () => ($init !== null ? $init : _default());

  formData.init = () => $init;

  return formData;
}

function arrayFormData({ $model, $init }, rerender, getRoot) {
  const formData = $init.map(v => createFormData($model(v), rerender), getRoot);
  formData._originalPush = formData.push;
  formData._originalSplice = formData.splice;

  formData.value = () => {
    return formData.map(v => v.value());
  };

  formData.changed = () => {
    return formData.length !== $init.length || formData.some(v => v.changed());
  };

  formData.push = v => {
    formData._originalPush(createFormData($model(v), rerender));
    rerender();
  };

  formData.splice = (idx, n) => {
    formData._originalSplice(idx, n);
    rerender();
  };

  return formData;
}

function compositeFormData({ $includeAllValues, ...model }, rerender, getRoot) {
  const formData = Object.keys(model).reduce(
    (acc, key) => ({
      ...acc,
      [key]: createFormData(model[key], rerender, getRoot),
    }),
    {}
  );

  formData.changed = () => {
    return Object.keys(model).some(k => formData[k].changed());
  };

  formData.value = () => {
    return Object.keys(model).reduce((acc, key) => {
      if (!formData[key].changed() && !$includeAllValues) {
        return acc;
      }
      return { ...acc, [key]: formData[key].value() };
    }, {});
  };

  return formData;
}

function createFormData(model, rerender, getRoot) {
  if (Object.prototype.hasOwnProperty.call(model, '$init')) {
    if (Object.prototype.hasOwnProperty.call(model, '$model')) {
      return arrayFormData(model, rerender, getRoot);
    }
    return simpleFormData(model, getRoot);
  }
  return compositeFormData(model, rerender, getRoot);
}

export function useFormData(model) {
  const [, setValue] = useState(false);
  const rerender = () => setValue(v => !v);
  let formData,
    getRoot = () => formData;
  formData = createFormData(model, rerender, getRoot);
  const ref = useRef(formData);
  return ref.current;
}
