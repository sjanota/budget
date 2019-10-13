function createFormRef({ $init, $process }) {
  const ref = { current: null };
  ref.changed = function() {
    return (
      (ref.current === null && $init !== null) || ref.current.value !== $init
    );
  };
  ref.value = function() {
    if (ref.current === null) {
      return null;
    }
    if (Array.isArray(ref.current.value)) {
      return ref.current.value.map(v => v.value());
    }
    return $process ? $process(ref.current.value) : ref.current.value;
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
      return { ...acc, [key]: raw };
    }, {});
  };
  return refs;
}
