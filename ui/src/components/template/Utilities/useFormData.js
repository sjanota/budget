function createFormRef({ $init, $process }) {
  const ref = { current: null };
  ref.changed = function() {
    return (
      (ref.current === null && $init !== null) || ref.current.value !== $init
    );
  };
  ref.value = function() {
    const val = ref.current === null ? ref.current : ref.current.value;
    return $process ? $process(val) : val;
  };
  ref.init = $init;
  return ref;
}

export default function useFormData(model) {
  const refs = Object.keys(model).reduce(
    (acc, key) => ({ ...acc, [key]: createFormRef(model[key]) }),
    {}
  );
  refs.changed = function() {
    return Object.values(refs).some(v => v.changed());
  };
  refs.value = function() {
    return Object.keys(model).reduce((acc, key) => {
      if (!refs[key].changed()) {
        return acc;
      }
      const raw = refs[key].value();
      const value = process[key] ? process[key](raw) : raw;
      return { ...acc, [key]: value };
    }, {});
  };
  return refs;
}
