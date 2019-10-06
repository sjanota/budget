import React from 'react';

export default function TopbarSearch({ onSearch }) {
  return (
    <div className="input-group">
      <input
        type="text"
        className="form-control bg-light border-0 small"
        placeholder="Search for..."
        aria-label="Search"
        aria-describedby="basic-addon2"
      />
      <div className="input-group-append">
        <button className="btn btn-primary" type="button" onClick={onSearch}>
          <i className="fas fa-search fa-sm"></i>
        </button>
      </div>
    </div>
  );
}
