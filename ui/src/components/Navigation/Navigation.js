import React from 'react';
import Nav from "react-bootstrap/Nav";
import {LinkContainer} from 'react-router-bootstrap';
import './Navigation.css';

const Link = ({to, label}) => {
  return <Nav.Item>
    <LinkContainer to={to}>
      <Nav.Link>{label}</Nav.Link>
    </LinkContainer>
  </Nav.Item>;
};

export const Navigation = () => {
  return <Nav className="flex-column Navigation" variant="pills">
    <Link to={"/expenses"} label={"Expenses"}/>
    <Link to={"/accounts"} label={"Accounts"}/>
    <Link to={"/envelopes"} label={"Envelopes"}/>
    <Link to={"/categories"} label={"Categories"}/>
  </Nav>
};