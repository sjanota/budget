import React from 'react';
import Dashboard from './Dashboard';
import Buttons from './Buttons';
import { Route } from 'react-router-dom';
import Sidebar from './Sidebar';
import Topbar from './Topbar';
import SBAdmin2 from './template/SBAdmin2';

export default function App() {
  return (
    <SBAdmin2 sidebar={Sidebar} topbar={Topbar} copyright={'Budget 2019'}>
      <Route path="/buttons" component={Buttons} />
      <Route exact path="/" component={Dashboard} />
    </SBAdmin2>
  );
}
