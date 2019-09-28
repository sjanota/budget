import React from 'react';
import Nav from 'react-bootstrap/Nav';
import { LinkContainer } from 'react-router-bootstrap';
import './Navigation.css';
import PropTypes from 'prop-types';

const Link = ({ to, label }) => {
  return (
    <Nav.Item>
      <LinkContainer to={to}>
        <Nav.Link>{label}</Nav.Link>
      </LinkContainer>
    </Nav.Item>
  );
};

Link.propTypes = {
  label: PropTypes.string,
  to: PropTypes.string,
};

export const Navigation = () => {
  return (
    <Nav className="flex-column Navigation" variant="pills">
      <Link to={'/expenses'} label={'Wydatki'} />
      <Link to={'/accounts'} label={'Konta'} />
      <Link to={'/envelopes'} label={'Koperty'} />
    </Nav>
  );
};
