import React, { useRef, useState, useEffect } from 'react';
import classnames from 'classnames';

export function Combobox({ allowedValues, _ref, defaultValue, className }) {
  const defaultValueObject = allowedValues.find(v => v.id === defaultValue);
  const [show, setShow] = useState(false);
  const [filter, setFilter] = useState(
    defaultValueObject ? defaultValueObject.label : ''
  );
  const [selectedIdx, setSelectedIdx] = useState(0);
  const menuRef = useRef();
  const classNames = classnames('input-group', className);
  const filtered = allowedValues.filter(v =>
    v.label.toLowerCase().includes(filter.toLowerCase())
  );
  useEffect(() => {
    _ref.current = { value: defaultValue };
    // eslint-disable-next-line
  }, []);
  function onClick(id) {
    const selectedLabel = allowedValues.find(v => v.id === id).label;
    setShow(false);
    setFilter(selectedLabel);
    _ref.current = { value: id };
  }
  function onInputChange(e) {
    const value = e.target.value;
    setShow(true);
    setFilter(value);
    setSelectedIdx(0);
    const selected = allowedValues.find(v => v.label === value);
    if (selected) {
      _ref.current = { value: selected.id };
    }
  }
  function onInputBlur() {
    setShow(false);
    const selected = allowedValues.find(v => v.label === filter);
    if (!selected) {
      setFilter('');
    }
  }
  function onKeyDown(e) {
    if (e.keyCode === 40) {
      // up
      e.preventDefault();
      if (selectedIdx >= filtered.length - 1) {
        setSelectedIdx(0);
      } else {
        setSelectedIdx(v => v + 1);
      }
    } else if (e.keyCode === 38) {
      // down
      e.preventDefault();
      if (selectedIdx <= 0) {
        setSelectedIdx(filtered.length - 1);
      } else {
        setSelectedIdx(v => v - 1);
      }
    } else if (e.keyCode === 13 && show) {
      e.preventDefault();
      onClick(filtered[selectedIdx].id);
    }
  }
  return (
    <div className={classNames}>
      <input
        className="form-control"
        value={filter}
        type="text"
        onChange={onInputChange}
        onBlur={onInputBlur}
        onKeyDown={onKeyDown}
      />
      <button
        className="btn btn-secondary dropdown-toggle dropdown-toggle-split no-arrow"
        data-toggle="dropdown"
        data-reference="parent"
        data-flip="false"
        style={{
          maxWidth: '2rem',
          borderTopLeftRadius: 0,
          borderBottomLeftRadius: 0,
          margin: -1,
        }}
      />
      <ul
        ref={menuRef}
        role="menu"
        className={classnames('dropdown-menu', { show })}
        style={{
          maxHeight: '200px',
          overflowY: 'auto',
        }}
      >
        {filtered.map((v, idx) => (
          <li
            className={`dropdown-item ${idx === selectedIdx ? 'active' : ''}`}
            onClick={() => onClick(v.id)}
            key={v.id}
          >
            {v.label}
          </li>
        ))}
      </ul>
    </div>
  );
}
