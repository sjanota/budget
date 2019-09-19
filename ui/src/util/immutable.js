export function addToList(list, element) {
  return [...list, element]
}

export function removeFromList(list, element) {
  const idx = list.indexOf(element);
  if (idx === -1) {
    return list
  }
  return [...list.slice(0, idx), ...list.slice(idx+1, list.length)]
}

export function removeFromListById(list, elementId) {
  const idx = list.findIndex(e => e.id === elementId);
  if (idx === -1) {
    return list
  }
  return [...list.slice(0, idx), ...list.slice(idx+1, list.length)]
}

export function replaceOnList(list, idx, element) {
  if (idx < 0 || idx > list.length) {
    return list
  }
  return [...list.slice(0, idx), element, ...list.slice(idx+1, list.length)]
}