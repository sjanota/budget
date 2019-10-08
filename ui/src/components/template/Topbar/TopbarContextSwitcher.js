import React from 'react';

export default function TopbarContextSwitcher({
  label,
  value,
  onChange,
  allowedValues,
}) {
  return (
    <div
      className="input-group navbar-context-switcher"
      style={{ display: 'flex', alignContent: 'stretch' }}
    >
      <div className="input-group-prepend navbar-context-switcher-label">
        <label className="input-group-text border-0">{label}</label>
      </div>
      <div
        className="input-group-append btn-group"
        style={{
          flexGrow: 1,
        }}
      >
        <span
          className="input-group-text bg-light border-0 navbar-context-display text-dark"
          aria-label={label}
        >
          {value || <span className="text-muted small">Choose {label}...</span>}
        </span>
        <button
          className="btn btn-primary dropdown-toggle dropdown-toggle-split no-arrow"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
          style={{ maxWidth: '2rem' }}
        />

        <div className="dropdown-menu">
          <span
            className="dropdown-item text-muted"
            onClick={() => onChange(null)}
          >
            Choose {label}...
          </span>
          {allowedValues.map(v => (
            <span className="dropdown-item" onClick={() => onChange(v)} key={v}>
              {v}
            </span>
          ))}
        </div>
      </div>
    </div>
  );
}
