import React from 'react';
import SidebarCollapsibleLink from './SidebarCollapsibleLink';
import SidebarSimpleLink from './SidebarSimpleLink';
import PropTypes from 'prop-types';

export function SidebarGroup({ group }) {
  return (
    <>
      <hr className="sidebar-divider my-1" />
      {group.name && <div className="sidebar-heading">{group.name}</div>}
      {group.entries.map(entry =>
        entry.to !== undefined ? (
          <SidebarSimpleLink key={entry.name} {...entry} />
        ) : (
          <SidebarCollapsibleLink
            key={entry.name}
            parent="accordionSidebar"
            {...entry}
          />
        )
      )}
    </>
  );
}

SidebarGroup.propTypes = {
  group: PropTypes.shape({
    name: PropTypes.string,
    entries: PropTypes.arrayOf(
      PropTypes.oneOfType([
        PropTypes.shape(SidebarCollapsibleLink.propTypes),
        PropTypes.shape(SidebarSimpleLink.propTypes),
      ])
    ).isRequired,
  }),
};
